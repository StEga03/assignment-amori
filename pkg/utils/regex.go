package utils

import "regexp"

// The \s character class in JavaScript regex typically matches:
// - Space character (U+0020)
// - Tab character (U+0009)
// - Line feed or new line character (U+000A)
// - Carriage return character (U+000D)
// - Form feed character (U+000C)
// - Vertical tab character (U+000B)
// - No-break space (U+00A0)
// - Ogham space mark (U+1680)
// - En quad (U+2000)
// - Em quad (U+2001)
// - En space (U+2002)
// - Em space (U+2003)
// - Three-per-em space (U+2004)
// - Four-per-em space (U+2005)
// - Six-per-em space (U+2006)
// - Figure space (U+2007)
// - Punctuation space (U+2008)
// - Thin space (U+2009)
// - Hair space (U+200A)
// - Line separator (U+2028)
// - Paragraph separator (U+2029)
// - Narrow no-break space (U+202F)
// - Medium mathematical space (U+205F)
// - Ideographic space (U+3000)
//
// The inclusion of these characters helps ensure that JavaScript regex
// can effectively handle text containing various types of whitespace,
// including those used in different languages and typographic
// conventions.
//
// However, In Go (Golang), the behavior of the \s character class in
// regular expressions can differ from that in JavaScript. Go's regexp
// package implements regular expressions compatible with Perl, but it
// does not necessarily include all Unicode white space characters in
// the \s character class by default.
//
// The \s character class in Go regex typically matches:
// - Space character (U+0020)
// - Tab character (U+0009)
// - Line feed or new line character (U+000A)
// - Carriage return character (U+000D)
// - Form feed character (U+000C)
// - Vertical tab character (U+000B)
//
// This set is more limited compared to the broader range of white space
// characters recognized by JavaScript's \s character class. Notably,
// it does not include the Narrow No-Break Space (U+202F) by default.
// To make sure we do correct match, we use [\s\u202F] instead of \s.

var (
	RegexDate             = regexp.MustCompile("^(?:\u200E|\u200F)*\\[?(\\d{1,4}[-/.][\\s\u202F]?\\d{1,4}[-/.][\\s\u202F]?\\d{1,4})[,.]?[\\s\u202F]\\D*?(\\d{1,2}[.:]\\d{1,2}(?:[.:]\\d{1,2})?)(?:[\\s\u202F]([ap]\\.?[\\s\u202F]?m\\.?))?\\]?(?:[\\s\u202F]-|:)?[\\s\u202F]")
	RegexAuthorAndMessage = regexp.MustCompile("(.+?):[\\s\u202F](?:\u200E|\u200F)*([\\S\\s]*)")
	RegexMessage          = regexp.MustCompile("(?:\u200E|\u200F)*([\\s\\S]+)")
	RegexParserRegular    = regexp.MustCompile("(?i)" + RegexDate.String() + RegexAuthorAndMessage.String())
	RegexParserSystem     = regexp.MustCompile("(?i)" + RegexDate.String() + RegexMessage.String())
	RegexSplitTime        = regexp.MustCompile(`[:.]`)
	RegexSplitDate        = regexp.MustCompile(`[-/.] ?`)
	RegexNonAPM           = regexp.MustCompile(`(?i)[^apm]`)
	RegexNewlines         = regexp.MustCompile(`\r\n|\r|\n`)
)
