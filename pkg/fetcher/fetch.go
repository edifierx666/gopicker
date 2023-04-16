package fetcher

import (
  "fmt"
  "log"
  "net/url"
  "regexp"
  "strconv"
  "strings"

  "github.com/PuerkitoBio/goquery"
  "github.com/go-resty/resty/v2"
  "gopicker/cfg"
)

type Fetcher struct {
}

func makeURL(cfg2 *cfg.Cfg) string {
  u := &url.URL{
    Scheme: "https",
    Host:   "pkg.go.dev",
    Path:   "search",
  }

  q := u.Query()
  q.Set("q", cfg2.Name)
  q.Set("limit", strconv.Itoa(cfg2.Limit))
  u.RawQuery = q.Encode()

  return u.String()
}

var client = resty.New()

type Result struct {
  HTML        string `json:"HTML"`
  Name        string `json:"name"`
  Link        string `json:"link"`
  PKGURL      string `json:"pkgurl"`
  License     string `json:"license"`
  LicenseLink string `json:"licenseceLink"`
  Description string `json:"description"`
  Snippet     string `json:"snippet"`
}

func Parse(html string) ([]*Result, error) {
  doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
  if err != nil {
    return nil, fmt.Errorf("failed to parse the content: %w", err)
  }

  var (
    results []*Result
  )
  doc.Find("div.SearchSnippet").Each(func(i int, div *goquery.Selection) {
    result := &Result{}
    result.HTML = html

    div.Find("h2").Each(func(i int, h2 *goquery.Selection) {
      result.Name = clean(h2.Text())
      result.Link = clean(h2.Find("a").AttrOr("href", ""))
      result.PKGURL = pkgUrl(result.Link)
    })

    if desc := div.Find("p[data-test-id=snippet-synopsis]"); desc != nil {
      result.Description = clean(desc.Text())
    }

    if license := div.Find("span[data-test-id=snippet-license] a"); license != nil {
      result.LicenseLink = clean(license.AttrOr("href", ""))
      result.License = clean(license.Text())
    }

    div.Find("div.SearchSnippet-infoLabel").Each(func(i int, label *goquery.Selection) {
      result.Snippet = clean(label.Text())
    })

    results = append(results, result)
  })

  return results, nil
}

func ParseResult(cfg *cfg.Cfg) ([]*Result, error) {
  r := client.R()
  get, err := r.Get(makeURL(cfg))
  if err != nil {
    log.Println(err)
    return nil, err
  }
  return Parse(get.String())
}

func clean(str string) string {
  s := strings.ReplaceAll(str, "\n", "")
  re := regexp.MustCompile(` +`)
  s = re.ReplaceAllString(s, " ")

  return strings.TrimSpace(s)
}
func pkgUrl(link string) string {
  return strings.Replace(link, "/", "", 1)
}
