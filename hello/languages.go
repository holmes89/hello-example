package hello

import (
	"encoding/json"
	"strings"
)

// LanguageService provides the ability to find languages and hello in that language
type LanguageService struct {
	languageMap map[string]string
	languages   []string
}

// NewLanguageService creates the service by loading prefconfigured file
func NewLanguageService() *LanguageService {
	type helloStruct struct {
		Language string `json:"language"`
		Word     string `json:"hello"`
	}
	var hellos []helloStruct
	if err := json.Unmarshal([]byte(helloMap), &hellos); err != nil {
		panic(err)
	}
	languageMap := make(map[string]string)
	var languages []string
	for _, h := range hellos {
		language := strings.ToLower(h.Language)
		languageMap[language] = h.Word
		languages = append(languages, language)
	}
	return &LanguageService{
		languageMap: languageMap,
		languages:   languages,
	}
}

// ListLanguages returns all kesy in the language map
func (s *LanguageService) ListLanguages() []string {
	return s.languages
}

// GetHellos will return all languages if no filters are provided.
func (s *LanguageService) GetHellos(languages []string) map[string]string {
	//TODO lower languages
	if len(languages) == 0 {
		return s.languageMap
	}

	m := make(map[string]string)
	for _, l := range languages {
		m[l] = s.languageMap[l]
	}
	return m
}

const helloMap = `
[{
    "language": "English",
    "hello": "Welcome!"
  },
  {
    "language": "Afrikaans",
    "hello": "hallo"
  },
  {
    "language": "Albanian",
    "hello": "Përshëndetje"
  },
  {
    "language": "Amharic",
    "hello": "ሰላም"
  },
  {
    "language": "Arabic",
    "hello": "مرحبا"
  },
  {
    "language": "Armenian",
    "hello": "Բարեւ"
  },
  {
    "language": "Azerbaijani",
    "hello": "Salam"
  },
  {
    "language": "Basque",
    "hello": "Kaixo"
  },
  {
    "language": "Belarusian",
    "hello": "добры дзень"
  },
  {
    "language": "Bengali",
    "hello": "হ্যালো"
  },
  {
    "language": "Bosnian",
    "hello": "zdravo"
  },
  {
    "language": "Bulgarian",
    "hello": "Здравейте"
  },
  {
    "language": "Catalan",
    "hello": "Hola"
  },
  {
    "language": "Cebuano",
    "hello": "Hello"
  },
  {
    "language": "Chichewa",
    "hello": "Moni"
  },
  {
    "language": "Chinese (Simplified)",
    "hello": "您好"
  },
  {
    "language": "Chinese (Traditional)",
    "hello": "您好"
  },
  {
    "language": "Corsican",
    "hello": "Bonghjornu"
  },
  {
    "language": "Croatian",
    "hello": "zdravo"
  },
  {
    "language": "Czech",
    "hello": "Ahoj"
  },
  {
    "language": "Danish",
    "hello": "Hej"
  },
  {
    "language": "Dutch",
    "hello": "Hallo"
  },
  {
    "language": "English",
    "hello": "Hello"
  },
  {
    "language": "Esperanto",
    "hello": "Saluton"
  },
  {
    "language": "Estonian",
    "hello": "Tere"
  },
  {
    "language": "Filipino",
    "hello": "Hello"
  },
  {
    "language": "Finnish",
    "hello": "Hei"
  },
  {
    "language": "French",
    "hello": "Bonjour"
  },
  {
    "language": "Frisian",
    "hello": "Hello"
  },
  {
    "language": "Galician",
    "hello": "Ola"
  },
  {
    "language": "Georgian",
    "hello": "გამარჯობა"
  },
  {
    "language": "German",
    "hello": "Hallo"
  },
  {
    "language": "Greek",
    "hello": "Γεια σας"
  },
  {
    "language": "Gujarati",
    "hello": "હેલો"
  },
  {
    "language": "Haitian Creole",
    "hello": "Bonjou"
  },
  {
    "language": "Hausa",
    "hello": "Sannu"
  },
  {
    "language": "Hawaiian",
    "hello": "Alohaʻoe"
  },
  {
    "language": "Hebrew",
    "hello": "שלום"
  },
  {
    "language": "Hindi",
    "hello": "नमस्ते"
  },
  {
    "language": "Hmong",
    "hello": "Nyob zoo"
  },
  {
    "language": "Hungarian",
    "hello": "Helló"
  },
  {
    "language": "Icelandic",
    "hello": "Halló"
  },
  {
    "language": "Igbo",
    "hello": "Ndewo"
  },
  {
    "language": "Indonesian",
    "hello": "Halo"
  },
  {
    "language": "Irish",
    "hello": "Dia duit"
  },
  {
    "language": "Italian",
    "hello": "Ciao"
  },
  {
    "language": "Japanese",
    "hello": "こんにちは"
  },
  {
    "language": "Javanese",
    "hello": "Hello"
  },
  {
    "language": "Kannada",
    "hello": "ಹಲೋ"
  },
  {
    "language": "Kazakh",
    "hello": "Сәлем"
  },
  {
    "language": "Khmer",
    "hello": "ជំរាបសួរ"
  },
  {
    "language": "Korean",
    "hello": "안녕하세요."
  },
  {
    "language": "Kurdish (Kurmanji)",
    "hello": "Hello"
  },
  {
    "language": "Kyrgyz",
    "hello": "салам"
  },
  {
    "language": "Lao",
    "hello": "ສະບາຍດີ"
  },
  {
    "language": "Latin",
    "hello": "salve"
  },
  {
    "language": "Latvian",
    "hello": "Labdien!"
  },
  {
    "language": "Lithuanian",
    "hello": "Sveiki"
  },
  {
    "language": "Luxembourgish",
    "hello": "Moien"
  },
  {
    "language": "Macedonian",
    "hello": "Здраво"
  },
  {
    "language": "Malagasy",
    "hello": "Hello"
  },
  {
    "language": "Malay",
    "hello": "Hello"
  },
  {
    "language": "Malayalam",
    "hello": "ഹലോ"
  },
  {
    "language": "Maltese",
    "hello": "Hello"
  },
  {
    "language": "Maori",
    "hello": "Hiha"
  },
  {
    "language": "Marathi",
    "hello": "हॅलो"
  },
  {
    "language": "Mongolian",
    "hello": "Сайн байна уу"
  },
  {
    "language": "Myanmar (Burmese)",
    "hello": "မင်္ဂလာပါ"
  },
  {
    "language": "Nepali",
    "hello": "नमस्ते"
  },
  {
    "language": "Norwegian",
    "hello": "Hallo"
  },
  {
    "language": "Pashto",
    "hello": "سلام"
  },
  {
    "language": "Persian",
    "hello": "سلام"
  },
  {
    "language": "Polish",
    "hello": "Cześć"
  },
  {
    "language": "Portuguese",
    "hello": "Olá"
  },
  {
    "language": "Punjabi",
    "hello": "ਹੈਲੋ"
  },
  {
    "language": "Romanian",
    "hello": "Alo"
  },
  {
    "language": "Russian",
    "hello": "привет"
  },
  {
    "language": "Samoan",
    "hello": "Talofa"
  },
  {
    "language": "Scots Gaelic",
    "hello": "Hello"
  },
  {
    "language": "Serbian",
    "hello": "Здраво"
  },
  {
    "language": "Sesotho",
    "hello": "Hello"
  },
  {
    "language": "Shona",
    "hello": "Hello"
  },
  {
    "language": "Sindhi",
    "hello": "هيلو"
  },
  {
    "language": "Sinhala",
    "hello": "හෙලෝ"
  },
  {
    "language": "Slovak",
    "hello": "ahoj"
  },
  {
    "language": "Slovenian",
    "hello": "Pozdravljeni"
  },
  {
    "language": "Somali",
    "hello": "Hello"
  },
  {
    "language": "Spanish",
    "hello": "Hola"
  },
  {
    "language": "Sundanese",
    "hello": "halo"
  },
  {
    "language": "Swahili",
    "hello": "Sawa"
  },
  {
    "language": "Swedish",
    "hello": "Hallå"
  },
  {
    "language": "Tajik",
    "hello": "Салом"
  },
  {
    "language": "Tamil",
    "hello": "ஹலோ"
  },
  {
    "language": "Telugu",
    "hello": "హలో"
  },
  {
    "language": "Thai",
    "hello": "สวัสดี"
  },
  {
    "language": "Turkish",
    "hello": "Merhaba"
  },
  {
    "language": "Ukranian",
    "hello": "Здрастуйте"
  },
  {
    "language": "Urdu",
    "hello": "ہیلو"
  },
  {
    "language": "Uzbek",
    "hello": "Salom"
  },
  {
    "language": "Vietnamese",
    "hello": "Xin chào"
  },
  {
    "language": "Welsh",
    "hello": "Helo"
  },
  {
    "language": "Xhosa",
    "hello": "Sawubona"
  },
  {
    "language": "Yiddish",
    "hello": "העלא"
  },
  {
    "language": "Yoruba",
    "hello": "Kaabo"
  },
  {
    "language": "Zulu",
    "hello": "Sawubona"
  }
]
`
