package goLanguageDetection

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type testpair struct {
	language string
	text     string
}

var tests = []testpair{
	{"danish", "Quizdeltagerne spiste jordbær med fløde, mens cirkusklovnen Walther spillede på xylofon"},
	{"dutch", "Sexy qua lijf, doch bang voor 't zwempak Pa's wijze lynx bezag vroom het fikse aquaduct "},
	{"english", "Jackdaws love my big sphinx of quartz The quick brown fox jumps over the lazy dog Grumpy wizards make toxic brew for the evil queen and Jack"},
	{"french", "Portez ce whisky au vieux juge blond qui fume. Le cœur déçu mais l'âme plutôt naïve, Louÿs rêva de crapaüter en canoë au delà des îles, près du mälströn où brûlent les novæ"},
	{"german", "Sylvia wagt quick den Jux bei Pforzheim Zwölf Boxkämpfer jagten Victor quer über den großen Sylter Deich Franz jagt im komplett verwahrlosten Taxi quer durch Bayern"},
	{"italian", "Ma la volpe col suo balzo ha raggiunto il quieto Fido Quel vituperabile xenofobo zelante assaggia il whisky ed esclama: alleluja!"},
	{"portuguese", "Um pequeno jabuti xereta viu dez cegonhas felizes. Zebras caolhas de Java querem mandar fax para gigante em New York"},
	{"spanish", "El cadáver de Wamba, rey godo de España, fue exhumado y trasladado en una caja de zinc que pesó un kilo"},
	{"swedish", "Flygande bäckasiner söka hwila på mjuka tuvor. Yxskaftbud, ge vår wczonmö iqhjälp"},
}

func TestDetectText(t *testing.T) {
	var detect Detect = *New()
	for _, pair := range tests {
		result, _ := detect.Text(pair.text)
		assert.Equal(t, pair.language, result)
	}
}
