package text

import (
	"os"
	"strings"
)

func DefaultStopwords() map[string]struct{} {
	return map[string]struct{}{
		"a": {}, "about": {}, "above": {}, "after": {}, "again": {}, "against": {}, "all": {},
		"am": {}, "an": {}, "and": {}, "any": {}, "are": {}, "as": {}, "at": {},
		"be": {}, "because": {}, "been": {}, "before": {}, "being": {}, "below": {},
		"between": {}, "both": {}, "but": {}, "by": {}, "can": {}, "cannot": {}, "could": {},
		"did": {}, "do": {}, "does": {}, "doing": {}, "down": {}, "during": {}, "each": {},
		"few": {}, "for": {}, "from": {}, "further": {}, "had": {}, "has": {}, "have": {},
		"having": {}, "he": {}, "her": {}, "here": {}, "hers": {}, "herself": {}, "him": {},
		"himself": {}, "his": {}, "how": {}, "i": {}, "if": {}, "in": {}, "into": {},
		"is": {}, "it": {}, "its": {}, "itself": {}, "just": {}, "me": {}, "more": {},
		"most": {}, "must": {}, "my": {}, "myself": {}, "no": {}, "nor": {}, "not": {},
		"now": {}, "of": {}, "off": {}, "on": {}, "once": {}, "only": {}, "or": {},
		"other": {}, "our": {}, "ours": {}, "ourselves": {}, "out": {}, "over": {}, "own": {},
		"same": {}, "she": {}, "should": {}, "so": {}, "some": {}, "such": {}, "than": {},
		"that": {}, "the": {}, "their": {}, "theirs": {}, "them": {}, "themselves": {},
		"then": {}, "there": {}, "these": {}, "they": {}, "this": {}, "those": {}, "through": {},
		"to": {}, "too": {}, "under": {}, "until": {}, "up": {}, "very": {}, "was": {},
		"we": {}, "were": {}, "what": {}, "when": {}, "where": {}, "which": {}, "while": {},
		"who": {}, "whom": {}, "why": {}, "will": {}, "with": {}, "would": {}, "you": {},
		"your": {}, "yours": {}, "yourself": {}, "yourselves": {},
	}
}

func GermanStopwords() map[string]struct{} {
	m := make(map[string]struct{})
	words := []string{
		"aber", "alle", "allem", "allen", "aller", "alles", "als", "also", "am", "an",
		"ander", "andere", "anderem", "anderen", "anderer", "anderes", "anders",
		"auf", "auch", "aus", "bei", "bin", "bis", "bist", "da", "damit", "dann",
		"das", "dass", "dein", "deine", "deinem", "deinen", "deiner", "dem",
		"den", "denn", "der", "des", "dessen", "dich", "die", "dies", "diese",
		"dieselbe", "dieselben", "diesem", "diesen", "dieser", "dieses", "dir",
		"doch", "dort", "durch", "ein", "eine", "einem", "einen", "einer",
		"eines", "einig", "einige", "einigem", "einiger", "einiges", "einmal",
		"er", "es", "etwas", "euch", "euer", "eure", "eurem", "euren", "eurer",
		"für", "gegen", "gewesen", "gut", "hab", "habe", "haben", "hat",
		"hatte", "hatten", "hier", "hin", "hinter", "ich", "ihm", "ihn", "ihnen",
		"ihr", "ihre", "ihrem", "ihren", "ihrer", "im", "in", "indem", "ins",
		"ist", "jede", "jedem", "jeden", "jeder", "jedes", "jene", "jenem",
		"jenen", "jener", "jenes", "jetzt", "kann", "kein", "keine", "keinem",
		"keinen", "keiner", "können", "könnte", "machen", "mag", "magst",
		"man", "manche", "manchem", "manchen", "mancher", "manches", "mein",
		"meine", "meinem", "meinen", "meiner", "mir", "mit", "muss", "musste",
		"nach", "nicht", "nichts", "noch", "nun", "nur", "ob", "oder",
		"ohne", "sehr", "sein", "seine", "seinem", "seinen", "seiner",
		"selbst", "sich", "sie", "sind", "so", "sobald", "sogar", "sonst",
		"über", "um", "und", "uns", "unser", "unsere", "unserem", "unseren",
		"unter", "viel", "vom", "von", "vor", "während", "war", "waren",
		"warst", "was", "weil", "welch", "welche", "welchem", "welchen",
		"welcher", "welches", "wenn", "werde", "werden", "wie", "wieder",
		"will", "wir", "wird", "wirst", "wo", "wollen", "wollte", "würde",
		"würden", "zu", "zum", "zur", "zwar", "zwischen",
	}
	for _, w := range words {
		m[w] = struct{}{}
	}
	return m
}

func FrenchStopwords() map[string]struct{} {
	m := make(map[string]struct{})
	words := []string{
		"à", "alors", "après", "avec", "autre", "autres", "avant", "bien", "bon",
		"cette", "ceux", "chaque", "chez", "choisir", "chose", "comme",
		"comment", "contre", "croire", "dans", "donc", "dont", "du", "elle",
		"elles", "encore", "enfin", "ensemble", "ensuite", "entre", "être",
		"eu", "fait", "gens", "grand", "gross", "hors", "ici", "il", "ils",
		"jamais", "je", "juste", "la", "le", "les", "leur", "leurs",
		"lieu", "lors", "lorsque", "lui", "mais", "même", "mes", "moins",
		"moment", "mon", "monsieur", "moyen", "naître", "naturel", "ne", "ni",
		"niveau", "nom", "non", "nos", "notre", "nous", "nouveau",
		"nouveaux", "oui", "ou", "par", "paraître", "parce", "parler",
		"parmi", "pas", "passer", "passé", "pendant", "penser", "perdre",
		"permet", "personne", "petit", "peu", "peut", "peuvent", "pire",
		"plus", "plutôt", "point", "porter", "poser", "possible", "pour",
		"pouv", "pouvoir", "près", "premier", "première", "prendre",
		"proche", "propos", "puis", "que", "quel", "quelle", "quelles",
		"quels", "qui", "quand", "quarante", "quatre", "rappeler",
		"recherche", "regarder", "relever", "remplacer", "rentrer",
		"retour", "rien", "rire", "risque", "roi", "rôle", "rond",
		"rose", "rouge", "route", "sa", "saisir", "sans", "savoir",
		"se", "second", "seconde", "selon", "semaine", "sembler",
		"sens", "sent", "sentence", "sentir", "serait", "série",
		"serious", "serveur", "service", "ses", "seul", "seule",
		"seulement", "si", "sien", "site", "situ", "situation",
		"smart", "social", "société", "soi", "soin", "soir", "soit",
		"soldat", "soleil", "solid", "somme", "son", "sont", "sort",
		"sou", "souhait", "sous", "sourire", "stat", "statue",
		"statut", "store", "sub", "sucre", "sud", "suer", "suivre",
		"sujet", "super", "suppose", "sur", "sure", "surveiller",
		"system", "système", "ta", "tâche", "taille", "take",
		"talon", "tapis", "tel", "telle", "temps", "tendre",
		"tenir", "term", "terminer", "terra", "terrain",
		"tester", "texte", "tien", "tim", "tire", "titre",
		"toi", "tomber", "ton", "top", "total", "toucher",
		"tour", "toute", "toutes", "trace", "train", "traiter",
		"tranquille", "travaille", "travailler", "travailleur",
		"travers", "tres", "tribu", "trimestre", "triste",
		"trouver", "tu", "turbo", "turn", "twist",
		"type", "types", "un", "une", "uniques", "uns",
		"usage", "util", "utiliser", "va", "vague",
		"vaincre", "valoir", "valeur", "vamp", "vase", "vaut",
		"vent", "ver", "verger", "vers", "verser", "vert",
		"veut", "vez", "via", "vie", "vieux", "vif", "vingt",
		"virer", "vis", "viser", "vision", "vite", "voici",
		"voila", "voir", "voisin", "voiture", "vol",
		"voler", "volet", "volonté", "volume", "vont",
		"vos", "votre", "vouloir", "voyage", "vrai",
		"vraiment", "vue", "vues", "wait", "walk",
		"wan", "war", "ward", "warm", "wash", "waste",
		"watch", "water", "wave", "ways", "web",
		"weekend", "weird", "welcome", "well", "west",
		"what", "when", "where", "which", "while",
		"white", "who", "whom", "why", "wide", "wife",
		"wild", "will", "win", "wind", "window",
		"wine", "wing", "winner", "winter", "wire",
		"wise", "wish", "within", "without", "witness",
		"woman", "wonder", "wood", "wooden", "word",
		"wore", "work", "worker", "works", "world",
		"worm", "worn", "wort", "would",
		"wound", "wrist", "write", "writer", "writing",
		"wrong", "wrote", "yard", "yeah", "year",
		"yellow", "yell", "yes", "yet", "yield",
		"young", "your", "yours", "yourself",
		"yourselves", "youth", "zones",
	}
	for _, w := range words {
		m[w] = struct{}{}
	}
	return m
}

func SpanishStopwords() map[string]struct{} {
	m := make(map[string]struct{})
	words := []string{
		"a", "al", "algo", "algunas", "algunos", "allí", "allá", "ante", "aquí",
		"así", "aunque", "bajo", "bien", "caber", "cada", "casi", "con",
		"contra", "cosa", "creer", "cual", "cuando", "de", "del", "dentro",
		"desde", "donde", "durante", "e", "el", "ella", "ellas", "ello",
		"ellos", "entonces", "entre", "ese", "eso", "esos", "esta", "está",
		"están", "este", "esto", "estos", "estar", "fue", "gustar", "haber",
		"hacer", "hacia", "han", "hasta", "hay", "en", "era", "eran", "es",
		"está", "esto", "estos", "excepto", "favor", "fueron", "gran", "ha",
		"había", "hacer", "hay", "hora", "hoy", "idea", "ir", "la", "las",
		"le", "les", "lo", "los", "luego", "lugar", "más", "mayor", "me",
		"medio", "mejor", "menos", "mientras", "mis", "mismo", "mucho", "muchos",
		"muy", "nacer", "nada", "nadie", "ni", "ninguno", "ninguna", "no",
		"nos", "nosotros", "nuestro", "nuestra", "o", "ocupar", "ojalá",
		"otro", "otra", "otros", "otras", "pagar", "para", "parecer", "parte",
		"pasado", "pedir", "pequeño", "perder", "pero", "pese", "poder",
		"por", "porque", "posible", "preciso", "preferir", "primera",
		"primer", "primero", "principal", "propio", "próximo", "prxima",
		"pues", "q", "quedar", "querer", "quien", "quier", "según",
		"ser", "si", "siempre", "significar", "sin", "sino", "siguiente",
		"sobre", "solamente", "solo", "sólo", "son", "somos", "sr", "su",
		"sustantivo", "también", "tcp", "tener", "tiempo", "tienen",
		"tipo", "toda", "todas", "todavía", "todo", "todos", "tomar",
		"total", "trabajar", "traer", "tras", "tratar", "través", "tres",
		"tu", "tú", "tuvo", "u", "último", "un", "una", "uno", "unos",
		"us", "usted", "ustedes", "utilizar", "va", "valor", "vamos",
		"van", "varios", "ve", "ver", "vez", "via", "vista", "vivir",
		"volar", "volver", "voto", "voy", "verdad", "verdadero", "vez",
		"viaje", "vino", "visto", "vote", "y", "ya", "yo", "zona",
	}
	for _, w := range words {
		m[w] = struct{}{}
	}
	return m
}

func LoadStopwords(path string, language string) (map[string]struct{}, error) {
	if path != "" {
		data, err := os.ReadFile(path)
		if err != nil {
			return nil, err
		}
		words := make(map[string]struct{})
		for _, w := range strings.Fields(string(data)) {
			words[strings.ToLower(w)] = struct{}{}
		}
		return words, nil
	}

	switch language {
	case "de":
		return GermanStopwords(), nil
	case "fr":
		return FrenchStopwords(), nil
	case "es":
		return SpanishStopwords(), nil
	default:
		return DefaultStopwords(), nil
	}
}