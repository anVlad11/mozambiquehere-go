package models

type NewsLanguage string

// Languages are taken from Apex Legends official
// news page language selector and presented in saved order
const (
	NewsLanguageEnglishUSA       = NewsLanguage("en-us")
	NewsLanguageEnglishUK        = NewsLanguage("en-uk")
	NewsLanguageEnglishAustralia = NewsLanguage("en-au")
	NewsLanguageGermanGermany    = NewsLanguage("de-de")
	NewsLanguageSpanishSpain     = NewsLanguage("es-es")
	NewsLanguageFrenchFrance     = NewsLanguage("fr-fr")
	NewsLanguageItalianItaly     = NewsLanguage("it-it")
	NewsLanguagePolishPoland     = NewsLanguage("pl-pl")
	NewsLanguageRussianRussia    = NewsLanguage("ru-ru")
	NewsLanguagePortugueseBrazil = NewsLanguage("pt-br")
	NewsLanguageSpanishMexico    = NewsLanguage("es-mx")
	NewsLanguageJapaneseJapan    = NewsLanguage("ja-jp")
	NewsLanguageKoreanKorea      = NewsLanguage("ko-kr")
	NewsLanguageChineseTaiwan    = NewsLanguage("zh-tw")
)
