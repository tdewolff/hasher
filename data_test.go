package main

// uses github.com/tdewolff/hasher
//go:generate hasher -type=Hash -file=data_test.go

import "github.com/dgryski/go-metro"

// Hash defines perfect hashes for a predefined list of strings
type Hash int32

// Unique hash definitions to be used instead of strings
const (
	A                Hash = 0x0   // a
	Abbr             Hash = 0x1   // abbr
	Accept           Hash = 0x2   // accept
	Accept_Charset   Hash = 0x3   // accept-charset
	Accesskey        Hash = 0x4   // accesskey
	Acronym          Hash = 0x5   // acronym
	Action           Hash = 0x6   // action
	Address          Hash = 0x7   // address
	Align            Hash = 0x8   // align
	Alink            Hash = 0x9   // alink
	Allowfullscreen  Hash = 0xa   // allowfullscreen
	Alt              Hash = 0xb   // alt
	Annotation       Hash = 0xc   // annotation
	AnnotationXml    Hash = 0xd   // annotationXml
	Applet           Hash = 0xe   // applet
	Area             Hash = 0xf   // area
	Article          Hash = 0x10  // article
	Aside            Hash = 0x11  // aside
	Async            Hash = 0x12  // async
	Audio            Hash = 0x13  // audio
	Autocomplete     Hash = 0x14  // autocomplete
	Autofocus        Hash = 0x15  // autofocus
	Autoplay         Hash = 0x16  // autoplay
	Axis             Hash = 0x17  // axis
	B                Hash = 0x18  // b
	Background       Hash = 0x19  // background
	Base             Hash = 0x1a  // base
	Basefont         Hash = 0x1b  // basefont
	Bdi              Hash = 0x1c  // bdi
	Bdo              Hash = 0x1d  // bdo
	Bgcolor          Hash = 0x1e  // bgcolor
	Bgsound          Hash = 0x1f  // bgsound
	Big              Hash = 0x20  // big
	Blink            Hash = 0x21  // blink
	Blockquote       Hash = 0x22  // blockquote
	Body             Hash = 0x23  // body
	Border           Hash = 0x24  // border
	Br               Hash = 0x25  // br
	Button           Hash = 0x26  // button
	Canvas           Hash = 0x27  // canvas
	Caption          Hash = 0x28  // caption
	Center           Hash = 0x29  // center
	Challenge        Hash = 0x2a  // challenge
	Charset          Hash = 0x2b  // charset
	Checked          Hash = 0x2c  // checked
	Cite             Hash = 0x2d  // cite
	Class            Hash = 0x2e  // class
	Classid          Hash = 0x2f  // classid
	Clear            Hash = 0x30  // clear
	Code             Hash = 0x31  // code
	Codebase         Hash = 0x32  // codebase
	Codetype         Hash = 0x33  // codetype
	Col              Hash = 0x34  // col
	Colgroup         Hash = 0x35  // colgroup
	Color            Hash = 0x36  // color
	Cols             Hash = 0x37  // cols
	Colspan          Hash = 0x38  // colspan
	Command          Hash = 0x39  // command
	Compact          Hash = 0x3a  // compact
	Content          Hash = 0x3b  // content
	Contenteditable  Hash = 0x3c  // contenteditable
	Contextmenu      Hash = 0x3d  // contextmenu
	Controls         Hash = 0x3e  // controls
	Coords           Hash = 0x3f  // coords
	Crossorigin      Hash = 0x40  // crossorigin
	Data             Hash = 0x41  // data
	Datalist         Hash = 0x42  // datalist
	Datetime         Hash = 0x43  // datetime
	Dd               Hash = 0x44  // dd
	Declare          Hash = 0x45  // declare
	Default          Hash = 0x46  // default
	DefaultChecked   Hash = 0x47  // defaultChecked
	DefaultMuted     Hash = 0x48  // defaultMuted
	DefaultSelected  Hash = 0x49  // defaultSelected
	Defer            Hash = 0x4a  // defer
	Del              Hash = 0x4b  // del
	Desc             Hash = 0x4c  // desc
	Details          Hash = 0x4d  // details
	Dfn              Hash = 0x4e  // dfn
	Dialog           Hash = 0x4f  // dialog
	Dir              Hash = 0x50  // dir
	Dirname          Hash = 0x51  // dirname
	Disabled         Hash = 0x52  // disabled
	Div              Hash = 0x53  // div
	Dl               Hash = 0x54  // dl
	Download         Hash = 0x55  // download
	Draggable        Hash = 0x56  // draggable
	Dropzone         Hash = 0x57  // dropzone
	Dt               Hash = 0x58  // dt
	Em               Hash = 0x59  // em
	Embed            Hash = 0x5a  // embed
	Enabled          Hash = 0x5b  // enabled
	Enctype          Hash = 0x5c  // enctype
	Face             Hash = 0x5d  // face
	Fieldset         Hash = 0x5e  // fieldset
	Figcaption       Hash = 0x5f  // figcaption
	Figure           Hash = 0x60  // figure
	Font             Hash = 0x61  // font
	Footer           Hash = 0x62  // footer
	For              Hash = 0x63  // for
	ForeignObject    Hash = 0x64  // foreignObject
	Foreignobject    Hash = 0x65  // foreignobject
	Form             Hash = 0x66  // form
	Formaction       Hash = 0x67  // formaction
	Formenctype      Hash = 0x68  // formenctype
	Formmethod       Hash = 0x69  // formmethod
	Formnovalidate   Hash = 0x6a  // formnovalidate
	Formtarget       Hash = 0x6b  // formtarget
	Frame            Hash = 0x6c  // frame
	Frameborder      Hash = 0x6d  // frameborder
	Frameset         Hash = 0x6e  // frameset
	H1               Hash = 0x6f  // h1
	H2               Hash = 0x70  // h2
	H3               Hash = 0x71  // h3
	H4               Hash = 0x72  // h4
	H5               Hash = 0x73  // h5
	H6               Hash = 0x74  // h6
	Head             Hash = 0x75  // head
	Header           Hash = 0x76  // header
	Headers          Hash = 0x77  // headers
	Height           Hash = 0x78  // height
	Hgroup           Hash = 0x79  // hgroup
	Hidden           Hash = 0x7a  // hidden
	High             Hash = 0x7b  // high
	Hr               Hash = 0x7c  // hr
	Href             Hash = 0x7d  // href
	Hreflang         Hash = 0x7e  // hreflang
	Html             Hash = 0x7f  // html
	Http_Equiv       Hash = 0x80  // http-equiv
	I                Hash = 0x81  // i
	Icon             Hash = 0x82  // icon
	Id               Hash = 0x83  // id
	Iframe           Hash = 0x84  // iframe
	Image            Hash = 0x85  // image
	Img              Hash = 0x86  // img
	Inert            Hash = 0x87  // inert
	Input            Hash = 0x88  // input
	Ins              Hash = 0x89  // ins
	Isindex          Hash = 0x8a  // isindex
	Ismap            Hash = 0x8b  // ismap
	Itemid           Hash = 0x8c  // itemid
	Itemprop         Hash = 0x8d  // itemprop
	Itemref          Hash = 0x8e  // itemref
	Itemscope        Hash = 0x8f  // itemscope
	Itemtype         Hash = 0x90  // itemtype
	Kbd              Hash = 0x91  // kbd
	Keygen           Hash = 0x92  // keygen
	Keytype          Hash = 0x93  // keytype
	Kind             Hash = 0x94  // kind
	Label            Hash = 0x95  // label
	Lang             Hash = 0x96  // lang
	Language         Hash = 0x97  // language
	Legend           Hash = 0x98  // legend
	Li               Hash = 0x99  // li
	Link             Hash = 0x9a  // link
	List             Hash = 0x9b  // list
	Listing          Hash = 0x9c  // listing
	Longdesc         Hash = 0x9d  // longdesc
	Loop             Hash = 0x9e  // loop
	Low              Hash = 0x9f  // low
	Main             Hash = 0xa0  // main
	Malignmark       Hash = 0xa1  // malignmark
	Manifest         Hash = 0xa2  // manifest
	Map              Hash = 0xa3  // map
	Mark             Hash = 0xa4  // mark
	Marquee          Hash = 0xa5  // marquee
	Math             Hash = 0xa6  // math
	Max              Hash = 0xa7  // max
	Maxlength        Hash = 0xa8  // maxlength
	Media            Hash = 0xa9  // media
	Mediagroup       Hash = 0xaa  // mediagroup
	Menu             Hash = 0xab  // menu
	Meta             Hash = 0xac  // meta
	Meter            Hash = 0xad  // meter
	Method           Hash = 0xae  // method
	Mglyph           Hash = 0xaf  // mglyph
	Mi               Hash = 0xb0  // mi
	Min              Hash = 0xb1  // min
	Mn               Hash = 0xb2  // mn
	Mo               Hash = 0xb3  // mo
	Ms               Hash = 0xb4  // ms
	Mtext            Hash = 0xb5  // mtext
	Multiple         Hash = 0xb6  // multiple
	Muted            Hash = 0xb7  // muted
	Name             Hash = 0xb8  // name
	Nav              Hash = 0xb9  // nav
	Nobr             Hash = 0xba  // nobr
	Noembed          Hash = 0xbb  // noembed
	Noframes         Hash = 0xbc  // noframes
	Nohref           Hash = 0xbd  // nohref
	Noresize         Hash = 0xbe  // noresize
	Noscript         Hash = 0xbf  // noscript
	Noshade          Hash = 0xc0  // noshade
	Novalidate       Hash = 0xc1  // novalidate
	Nowrap           Hash = 0xc2  // nowrap
	Object           Hash = 0xc3  // object
	Ol               Hash = 0xc4  // ol
	Onabort          Hash = 0xc5  // onabort
	Onafterprint     Hash = 0xc6  // onafterprint
	Onbeforeprint    Hash = 0xc7  // onbeforeprint
	Onbeforeunload   Hash = 0xc8  // onbeforeunload
	Onblur           Hash = 0xc9  // onblur
	Oncancel         Hash = 0xca  // oncancel
	Oncanplay        Hash = 0xcb  // oncanplay
	Oncanplaythrough Hash = 0xcc  // oncanplaythrough
	Onchange         Hash = 0xcd  // onchange
	Onclick          Hash = 0xce  // onclick
	Onclose          Hash = 0xcf  // onclose
	Oncontextmenu    Hash = 0xd0  // oncontextmenu
	Oncuechange      Hash = 0xd1  // oncuechange
	Ondblclick       Hash = 0xd2  // ondblclick
	Ondrag           Hash = 0xd3  // ondrag
	Ondragend        Hash = 0xd4  // ondragend
	Ondragenter      Hash = 0xd5  // ondragenter
	Ondragleave      Hash = 0xd6  // ondragleave
	Ondragover       Hash = 0xd7  // ondragover
	Ondragstart      Hash = 0xd8  // ondragstart
	Ondrop           Hash = 0xd9  // ondrop
	Ondurationchange Hash = 0xda  // ondurationchange
	Onemptied        Hash = 0xdb  // onemptied
	Onended          Hash = 0xdc  // onended
	Onerror          Hash = 0xdd  // onerror
	Onfocus          Hash = 0xde  // onfocus
	Onhashchange     Hash = 0xdf  // onhashchange
	Oninput          Hash = 0xe0  // oninput
	Oninvalid        Hash = 0xe1  // oninvalid
	Onkeydown        Hash = 0xe2  // onkeydown
	Onkeypress       Hash = 0xe3  // onkeypress
	Onkeyup          Hash = 0xe4  // onkeyup
	Onload           Hash = 0xe5  // onload
	Onloadeddata     Hash = 0xe6  // onloadeddata
	Onloadedmetadata Hash = 0xe7  // onloadedmetadata
	Onloadstart      Hash = 0xe8  // onloadstart
	Onmessage        Hash = 0xe9  // onmessage
	Onmousedown      Hash = 0xea  // onmousedown
	Onmousemove      Hash = 0xeb  // onmousemove
	Onmouseout       Hash = 0xec  // onmouseout
	Onmouseover      Hash = 0xed  // onmouseover
	Onmouseup        Hash = 0xee  // onmouseup
	Onmousewheel     Hash = 0xef  // onmousewheel
	Onoffline        Hash = 0xf0  // onoffline
	Ononline         Hash = 0xf1  // ononline
	Onpagehide       Hash = 0xf2  // onpagehide
	Onpageshow       Hash = 0xf3  // onpageshow
	Onpause          Hash = 0xf4  // onpause
	Onplay           Hash = 0xf5  // onplay
	Onplaying        Hash = 0xf6  // onplaying
	Onpopstate       Hash = 0xf7  // onpopstate
	Onprogress       Hash = 0xf8  // onprogress
	Onratechange     Hash = 0xf9  // onratechange
	Onreset          Hash = 0xfa  // onreset
	Onresize         Hash = 0xfb  // onresize
	Onscroll         Hash = 0xfc  // onscroll
	Onseeked         Hash = 0xfd  // onseeked
	Onseeking        Hash = 0xfe  // onseeking
	Onselect         Hash = 0xff  // onselect
	Onshow           Hash = 0x100 // onshow
	Onstalled        Hash = 0x101 // onstalled
	Onstorage        Hash = 0x102 // onstorage
	Onsubmit         Hash = 0x103 // onsubmit
	Onsuspend        Hash = 0x104 // onsuspend
	Ontimeupdate     Hash = 0x105 // ontimeupdate
	Onunload         Hash = 0x106 // onunload
	Onvolumechange   Hash = 0x107 // onvolumechange
	Onwaiting        Hash = 0x108 // onwaiting
	Open             Hash = 0x109 // open
	Optgroup         Hash = 0x10a // optgroup
	Optimum          Hash = 0x10b // optimum
	Option           Hash = 0x10c // option
	Output           Hash = 0x10d // output
	P                Hash = 0x10e // p
	Param            Hash = 0x10f // param
	Pattern          Hash = 0x110 // pattern
	Pauseonexit      Hash = 0x111 // pauseonexit
	Ping             Hash = 0x112 // ping
	Placeholder      Hash = 0x113 // placeholder
	Plaintext        Hash = 0x114 // plaintext
	Poster           Hash = 0x115 // poster
	Pre              Hash = 0x116 // pre
	Preload          Hash = 0x117 // preload
	Profile          Hash = 0x118 // profile
	Progress         Hash = 0x119 // progress
	Prompt           Hash = 0x11a // prompt
	Public           Hash = 0x11b // public
	Q                Hash = 0x11c // q
	Radiogroup       Hash = 0x11d // radiogroup
	Rb               Hash = 0x11e // rb
	Readonly         Hash = 0x11f // readonly
	Rel              Hash = 0x120 // rel
	Required         Hash = 0x121 // required
	Rev              Hash = 0x122 // rev
	Reversed         Hash = 0x123 // reversed
	Rows             Hash = 0x124 // rows
	Rowspan          Hash = 0x125 // rowspan
	Rp               Hash = 0x126 // rp
	Rt               Hash = 0x127 // rt
	Rtc              Hash = 0x128 // rtc
	Ruby             Hash = 0x129 // ruby
	Rules            Hash = 0x12a // rules
	S                Hash = 0x12b // s
	Samp             Hash = 0x12c // samp
	Sandbox          Hash = 0x12d // sandbox
	Scope            Hash = 0x12e // scope
	Scoped           Hash = 0x12f // scoped
	Script           Hash = 0x130 // script
	Scrolling        Hash = 0x131 // scrolling
	Seamless         Hash = 0x132 // seamless
	Section          Hash = 0x133 // section
	Select           Hash = 0x134 // select
	Selected         Hash = 0x135 // selected
	Shape            Hash = 0x136 // shape
	Size             Hash = 0x137 // size
	Sizes            Hash = 0x138 // sizes
	Small            Hash = 0x139 // small
	Sortable         Hash = 0x13a // sortable
	Source           Hash = 0x13b // source
	Spacer           Hash = 0x13c // spacer
	Span             Hash = 0x13d // span
	Spellcheck       Hash = 0x13e // spellcheck
	Src              Hash = 0x13f // src
	Srcdoc           Hash = 0x140 // srcdoc
	Srclang          Hash = 0x141 // srclang
	Start            Hash = 0x142 // start
	Step             Hash = 0x143 // step
	Strike           Hash = 0x144 // strike
	Strong           Hash = 0x145 // strong
	Style            Hash = 0x146 // style
	Sub              Hash = 0x147 // sub
	Summary          Hash = 0x148 // summary
	Sup              Hash = 0x149 // sup
	Svg              Hash = 0x14a // svg
	System           Hash = 0x14b // system
	Tabindex         Hash = 0x14c // tabindex
	Target           Hash = 0x14d // target
	Tbody            Hash = 0x14e // tbody
	Td               Hash = 0x14f // td
	Text             Hash = 0x150 // text
	Textarea         Hash = 0x151 // textarea
	Tfoot            Hash = 0x152 // tfoot
	Th               Hash = 0x153 // th
	Thead            Hash = 0x154 // thead
	Time             Hash = 0x155 // time
	Title            Hash = 0x156 // title
	Tr               Hash = 0x157 // tr
	Track            Hash = 0x158 // track
	Translate        Hash = 0x159 // translate
	Truespeed        Hash = 0x15a // truespeed
	Tt               Hash = 0x15b // tt
	Type             Hash = 0x15c // type
	Typemustmatch    Hash = 0x15d // typemustmatch
	U                Hash = 0x15e // u
	Ul               Hash = 0x15f // ul
	Undeterminate    Hash = 0x160 // undeterminate
	Usemap           Hash = 0x161 // usemap
	Valign           Hash = 0x162 // valign
	Value            Hash = 0x163 // value
	Valuetype        Hash = 0x164 // valuetype
	Var              Hash = 0x165 // var
	Video            Hash = 0x166 // video
	Visible          Hash = 0x167 // visible
	Vlink            Hash = 0x168 // vlink
	Wbr              Hash = 0x169 // wbr
	Width            Hash = 0x16a // width
	Wrap             Hash = 0x16b // wrap
	Xmlns            Hash = 0x16c // xmlns
	Xmp              Hash = 0x16d // xmp
)

var HashMap = map[Hash][]byte{
	A:                []byte("a"),
	Abbr:             []byte("abbr"),
	Accept:           []byte("accept"),
	Accept_Charset:   []byte("accept-charset"),
	Accesskey:        []byte("accesskey"),
	Acronym:          []byte("acronym"),
	Action:           []byte("action"),
	Address:          []byte("address"),
	Align:            []byte("align"),
	Alink:            []byte("alink"),
	Allowfullscreen:  []byte("allowfullscreen"),
	Alt:              []byte("alt"),
	Annotation:       []byte("annotation"),
	AnnotationXml:    []byte("annotationXml"),
	Applet:           []byte("applet"),
	Area:             []byte("area"),
	Article:          []byte("article"),
	Aside:            []byte("aside"),
	Async:            []byte("async"),
	Audio:            []byte("audio"),
	Autocomplete:     []byte("autocomplete"),
	Autofocus:        []byte("autofocus"),
	Autoplay:         []byte("autoplay"),
	Axis:             []byte("axis"),
	B:                []byte("b"),
	Background:       []byte("background"),
	Base:             []byte("base"),
	Basefont:         []byte("basefont"),
	Bdi:              []byte("bdi"),
	Bdo:              []byte("bdo"),
	Bgcolor:          []byte("bgcolor"),
	Bgsound:          []byte("bgsound"),
	Big:              []byte("big"),
	Blink:            []byte("blink"),
	Blockquote:       []byte("blockquote"),
	Body:             []byte("body"),
	Border:           []byte("border"),
	Br:               []byte("br"),
	Button:           []byte("button"),
	Canvas:           []byte("canvas"),
	Caption:          []byte("caption"),
	Center:           []byte("center"),
	Challenge:        []byte("challenge"),
	Charset:          []byte("charset"),
	Checked:          []byte("checked"),
	Cite:             []byte("cite"),
	Class:            []byte("class"),
	Classid:          []byte("classid"),
	Clear:            []byte("clear"),
	Code:             []byte("code"),
	Codebase:         []byte("codebase"),
	Codetype:         []byte("codetype"),
	Col:              []byte("col"),
	Colgroup:         []byte("colgroup"),
	Color:            []byte("color"),
	Cols:             []byte("cols"),
	Colspan:          []byte("colspan"),
	Command:          []byte("command"),
	Compact:          []byte("compact"),
	Content:          []byte("content"),
	Contenteditable:  []byte("contenteditable"),
	Contextmenu:      []byte("contextmenu"),
	Controls:         []byte("controls"),
	Coords:           []byte("coords"),
	Crossorigin:      []byte("crossorigin"),
	Data:             []byte("data"),
	Datalist:         []byte("datalist"),
	Datetime:         []byte("datetime"),
	Dd:               []byte("dd"),
	Declare:          []byte("declare"),
	Default:          []byte("default"),
	DefaultChecked:   []byte("defaultChecked"),
	DefaultMuted:     []byte("defaultMuted"),
	DefaultSelected:  []byte("defaultSelected"),
	Defer:            []byte("defer"),
	Del:              []byte("del"),
	Desc:             []byte("desc"),
	Details:          []byte("details"),
	Dfn:              []byte("dfn"),
	Dialog:           []byte("dialog"),
	Dir:              []byte("dir"),
	Dirname:          []byte("dirname"),
	Disabled:         []byte("disabled"),
	Div:              []byte("div"),
	Dl:               []byte("dl"),
	Download:         []byte("download"),
	Draggable:        []byte("draggable"),
	Dropzone:         []byte("dropzone"),
	Dt:               []byte("dt"),
	Em:               []byte("em"),
	Embed:            []byte("embed"),
	Enabled:          []byte("enabled"),
	Enctype:          []byte("enctype"),
	Face:             []byte("face"),
	Fieldset:         []byte("fieldset"),
	Figcaption:       []byte("figcaption"),
	Figure:           []byte("figure"),
	Font:             []byte("font"),
	Footer:           []byte("footer"),
	For:              []byte("for"),
	ForeignObject:    []byte("foreignObject"),
	Foreignobject:    []byte("foreignobject"),
	Form:             []byte("form"),
	Formaction:       []byte("formaction"),
	Formenctype:      []byte("formenctype"),
	Formmethod:       []byte("formmethod"),
	Formnovalidate:   []byte("formnovalidate"),
	Formtarget:       []byte("formtarget"),
	Frame:            []byte("frame"),
	Frameborder:      []byte("frameborder"),
	Frameset:         []byte("frameset"),
	H1:               []byte("h1"),
	H2:               []byte("h2"),
	H3:               []byte("h3"),
	H4:               []byte("h4"),
	H5:               []byte("h5"),
	H6:               []byte("h6"),
	Head:             []byte("head"),
	Header:           []byte("header"),
	Headers:          []byte("headers"),
	Height:           []byte("height"),
	Hgroup:           []byte("hgroup"),
	Hidden:           []byte("hidden"),
	High:             []byte("high"),
	Hr:               []byte("hr"),
	Href:             []byte("href"),
	Hreflang:         []byte("hreflang"),
	Html:             []byte("html"),
	Http_Equiv:       []byte("http-equiv"),
	I:                []byte("i"),
	Icon:             []byte("icon"),
	Id:               []byte("id"),
	Iframe:           []byte("iframe"),
	Image:            []byte("image"),
	Img:              []byte("img"),
	Inert:            []byte("inert"),
	Input:            []byte("input"),
	Ins:              []byte("ins"),
	Isindex:          []byte("isindex"),
	Ismap:            []byte("ismap"),
	Itemid:           []byte("itemid"),
	Itemprop:         []byte("itemprop"),
	Itemref:          []byte("itemref"),
	Itemscope:        []byte("itemscope"),
	Itemtype:         []byte("itemtype"),
	Kbd:              []byte("kbd"),
	Keygen:           []byte("keygen"),
	Keytype:          []byte("keytype"),
	Kind:             []byte("kind"),
	Label:            []byte("label"),
	Lang:             []byte("lang"),
	Language:         []byte("language"),
	Legend:           []byte("legend"),
	Li:               []byte("li"),
	Link:             []byte("link"),
	List:             []byte("list"),
	Listing:          []byte("listing"),
	Longdesc:         []byte("longdesc"),
	Loop:             []byte("loop"),
	Low:              []byte("low"),
	Main:             []byte("main"),
	Malignmark:       []byte("malignmark"),
	Manifest:         []byte("manifest"),
	Map:              []byte("map"),
	Mark:             []byte("mark"),
	Marquee:          []byte("marquee"),
	Math:             []byte("math"),
	Max:              []byte("max"),
	Maxlength:        []byte("maxlength"),
	Media:            []byte("media"),
	Mediagroup:       []byte("mediagroup"),
	Menu:             []byte("menu"),
	Meta:             []byte("meta"),
	Meter:            []byte("meter"),
	Method:           []byte("method"),
	Mglyph:           []byte("mglyph"),
	Mi:               []byte("mi"),
	Min:              []byte("min"),
	Mn:               []byte("mn"),
	Mo:               []byte("mo"),
	Ms:               []byte("ms"),
	Mtext:            []byte("mtext"),
	Multiple:         []byte("multiple"),
	Muted:            []byte("muted"),
	Name:             []byte("name"),
	Nav:              []byte("nav"),
	Nobr:             []byte("nobr"),
	Noembed:          []byte("noembed"),
	Noframes:         []byte("noframes"),
	Nohref:           []byte("nohref"),
	Noresize:         []byte("noresize"),
	Noscript:         []byte("noscript"),
	Noshade:          []byte("noshade"),
	Novalidate:       []byte("novalidate"),
	Nowrap:           []byte("nowrap"),
	Object:           []byte("object"),
	Ol:               []byte("ol"),
	Onabort:          []byte("onabort"),
	Onafterprint:     []byte("onafterprint"),
	Onbeforeprint:    []byte("onbeforeprint"),
	Onbeforeunload:   []byte("onbeforeunload"),
	Onblur:           []byte("onblur"),
	Oncancel:         []byte("oncancel"),
	Oncanplay:        []byte("oncanplay"),
	Oncanplaythrough: []byte("oncanplaythrough"),
	Onchange:         []byte("onchange"),
	Onclick:          []byte("onclick"),
	Onclose:          []byte("onclose"),
	Oncontextmenu:    []byte("oncontextmenu"),
	Oncuechange:      []byte("oncuechange"),
	Ondblclick:       []byte("ondblclick"),
	Ondrag:           []byte("ondrag"),
	Ondragend:        []byte("ondragend"),
	Ondragenter:      []byte("ondragenter"),
	Ondragleave:      []byte("ondragleave"),
	Ondragover:       []byte("ondragover"),
	Ondragstart:      []byte("ondragstart"),
	Ondrop:           []byte("ondrop"),
	Ondurationchange: []byte("ondurationchange"),
	Onemptied:        []byte("onemptied"),
	Onended:          []byte("onended"),
	Onerror:          []byte("onerror"),
	Onfocus:          []byte("onfocus"),
	Onhashchange:     []byte("onhashchange"),
	Oninput:          []byte("oninput"),
	Oninvalid:        []byte("oninvalid"),
	Onkeydown:        []byte("onkeydown"),
	Onkeypress:       []byte("onkeypress"),
	Onkeyup:          []byte("onkeyup"),
	Onload:           []byte("onload"),
	Onloadeddata:     []byte("onloadeddata"),
	Onloadedmetadata: []byte("onloadedmetadata"),
	Onloadstart:      []byte("onloadstart"),
	Onmessage:        []byte("onmessage"),
	Onmousedown:      []byte("onmousedown"),
	Onmousemove:      []byte("onmousemove"),
	Onmouseout:       []byte("onmouseout"),
	Onmouseover:      []byte("onmouseover"),
	Onmouseup:        []byte("onmouseup"),
	Onmousewheel:     []byte("onmousewheel"),
	Onoffline:        []byte("onoffline"),
	Ononline:         []byte("ononline"),
	Onpagehide:       []byte("onpagehide"),
	Onpageshow:       []byte("onpageshow"),
	Onpause:          []byte("onpause"),
	Onplay:           []byte("onplay"),
	Onplaying:        []byte("onplaying"),
	Onpopstate:       []byte("onpopstate"),
	Onprogress:       []byte("onprogress"),
	Onratechange:     []byte("onratechange"),
	Onreset:          []byte("onreset"),
	Onresize:         []byte("onresize"),
	Onscroll:         []byte("onscroll"),
	Onseeked:         []byte("onseeked"),
	Onseeking:        []byte("onseeking"),
	Onselect:         []byte("onselect"),
	Onshow:           []byte("onshow"),
	Onstalled:        []byte("onstalled"),
	Onstorage:        []byte("onstorage"),
	Onsubmit:         []byte("onsubmit"),
	Onsuspend:        []byte("onsuspend"),
	Ontimeupdate:     []byte("ontimeupdate"),
	Onunload:         []byte("onunload"),
	Onvolumechange:   []byte("onvolumechange"),
	Onwaiting:        []byte("onwaiting"),
	Open:             []byte("open"),
	Optgroup:         []byte("optgroup"),
	Optimum:          []byte("optimum"),
	Option:           []byte("option"),
	Output:           []byte("output"),
	P:                []byte("p"),
	Param:            []byte("param"),
	Pattern:          []byte("pattern"),
	Pauseonexit:      []byte("pauseonexit"),
	Ping:             []byte("ping"),
	Placeholder:      []byte("placeholder"),
	Plaintext:        []byte("plaintext"),
	Poster:           []byte("poster"),
	Pre:              []byte("pre"),
	Preload:          []byte("preload"),
	Profile:          []byte("profile"),
	Progress:         []byte("progress"),
	Prompt:           []byte("prompt"),
	Public:           []byte("public"),
	Q:                []byte("q"),
	Radiogroup:       []byte("radiogroup"),
	Rb:               []byte("rb"),
	Readonly:         []byte("readonly"),
	Rel:              []byte("rel"),
	Required:         []byte("required"),
	Rev:              []byte("rev"),
	Reversed:         []byte("reversed"),
	Rows:             []byte("rows"),
	Rowspan:          []byte("rowspan"),
	Rp:               []byte("rp"),
	Rt:               []byte("rt"),
	Rtc:              []byte("rtc"),
	Ruby:             []byte("ruby"),
	Rules:            []byte("rules"),
	S:                []byte("s"),
	Samp:             []byte("samp"),
	Sandbox:          []byte("sandbox"),
	Scope:            []byte("scope"),
	Scoped:           []byte("scoped"),
	Script:           []byte("script"),
	Scrolling:        []byte("scrolling"),
	Seamless:         []byte("seamless"),
	Section:          []byte("section"),
	Select:           []byte("select"),
	Selected:         []byte("selected"),
	Shape:            []byte("shape"),
	Size:             []byte("size"),
	Sizes:            []byte("sizes"),
	Small:            []byte("small"),
	Sortable:         []byte("sortable"),
	Source:           []byte("source"),
	Spacer:           []byte("spacer"),
	Span:             []byte("span"),
	Spellcheck:       []byte("spellcheck"),
	Src:              []byte("src"),
	Srcdoc:           []byte("srcdoc"),
	Srclang:          []byte("srclang"),
	Start:            []byte("start"),
	Step:             []byte("step"),
	Strike:           []byte("strike"),
	Strong:           []byte("strong"),
	Style:            []byte("style"),
	Sub:              []byte("sub"),
	Summary:          []byte("summary"),
	Sup:              []byte("sup"),
	Svg:              []byte("svg"),
	System:           []byte("system"),
	Tabindex:         []byte("tabindex"),
	Target:           []byte("target"),
	Tbody:            []byte("tbody"),
	Td:               []byte("td"),
	Text:             []byte("text"),
	Textarea:         []byte("textarea"),
	Tfoot:            []byte("tfoot"),
	Th:               []byte("th"),
	Thead:            []byte("thead"),
	Time:             []byte("time"),
	Title:            []byte("title"),
	Tr:               []byte("tr"),
	Track:            []byte("track"),
	Translate:        []byte("translate"),
	Truespeed:        []byte("truespeed"),
	Tt:               []byte("tt"),
	Type:             []byte("type"),
	Typemustmatch:    []byte("typemustmatch"),
	U:                []byte("u"),
	Ul:               []byte("ul"),
	Undeterminate:    []byte("undeterminate"),
	Usemap:           []byte("usemap"),
	Valign:           []byte("valign"),
	Value:            []byte("value"),
	Valuetype:        []byte("valuetype"),
	Var:              []byte("var"),
	Video:            []byte("video"),
	Visible:          []byte("visible"),
	Vlink:            []byte("vlink"),
	Wbr:              []byte("wbr"),
	Width:            []byte("width"),
	Wrap:             []byte("wrap"),
	Xmlns:            []byte("xmlns"),
	Xmp:              []byte("xmp"),
}

// String returns the hash' name.
func (i Hash) String() string {
	return string(HashMap[i])
}

// Bytes returns the hash' name.
func (i Hash) Bytes() []byte {
	return HashMap[i]
}

// ToHash returns the hash whose name is s. It returns zero if there is no
// such hash. It is case sensitive.
func ToHash(k []byte) Hash {
	size := uint64(len(_Hash_values))
	hash := metro.Hash64(k, 0)
	i := hash & (size - 1)
	seed := _Hash_seeds[i]
	if seed < 0 {
		return _Hash_values[-seed-1]
	}

	i = xorshiftMult64(uint64(seed)+hash) & (size - 1)
	return _Hash_values[i]
}

func xorshiftMult64(x uint64) uint64 {
	x ^= x >> 12 // a
	x ^= x << 25 // b
	x ^= x >> 27 // c
	return x * 2685821657736338717
}

var _Hash_values = []Hash{286, 44, 57, 150, 77, 148, 218, 278, 66, 113, 332, 102, 149, 293, 243, 261, 285, 201, 78, 251, 185, 223, 105, 255, 141, 245, 24, 59, 297, 91, 316, 350, 6, 23, 206, 82, 63, 336, 162, 76, 124, 197, 196, 194, 246, 260, 20, 281, 38, 152, 22, 312, 211, 271, 252, 49, 351, 54, 97, 298, 296, 69, 89, 168, 85, 80, 3, 19, 346, 311, 233, 327, 79, 195, 235, 309, 230, 21, 0, 317, 56, 155, 177, 173, 280, 330, 232, 365, 31, 193, 349, 11, 300, 90, 64, 70, 108, 138, 62, 67, 17, 176, 345, 220, 103, 359, 283, 140, 265, 170, 237, 263, 35, 131, 88, 165, 272, 358, 32, 302, 337, 187, 96, 199, 244, 221, 208, 307, 174, 360, 95, 190, 289, 292, 181, 169, 27, 60, 315, 15, 268, 248, 42, 106, 274, 219, 135, 348, 340, 180, 189, 83, 249, 357, 326, 240, 153, 341, 93, 160, 34, 200, 321, 192, 212, 33, 347, 291, 74, 161, 116, 51, 98, 334, 269, 43, 328, 333, 104, 186, 238, 264, 322, 266, 198, 331, 81, 184, 118, 99, 68, 353, 342, 290, 125, 7, 259, 4, 143, 175, 92, 299, 275, 306, 355, 343, 26, 132, 270, 361, 356, 267, 204, 213, 224, 273, 154, 50, 126, 145, 171, 236, 256, 128, 254, 253, 58, 207, 178, 250, 310, 262, 217, 209, 101, 229, 320, 139, 40, 157, 188, 36, 39, 247, 231, 277, 48, 142, 127, 279, 109, 46, 28, 134, 100, 10, 210, 362, 117, 314, 242, 16, 114, 305, 301, 276, 338, 354, 303, 107, 87, 146, 203, 364, 133, 363, 352, 129, 45, 163, 215, 183, 14, 115, 287, 282, 304, 65, 0, 222, 0, 61, 0, 0, 0, 0, 294, 0, 0, 0, 30, 0, 239, 123, 0, 0, 0, 0, 227, 0, 0, 0, 137, 234, 0, 0, 0, 0, 110, 0, 0, 0, 0, 0, 119, 156, 0, 2, 0, 0, 0, 0, 0, 0, 319, 41, 0, 0, 324, 0, 0, 0, 0, 0, 0, 0, 53, 130, 0, 0, 122, 0, 0, 0, 0, 0, 308, 111, 226, 0, 318, 241, 0, 94, 0, 0, 86, 0, 29, 0, 47, 313, 0, 0, 172, 0, 0, 257, 0, 284, 0, 0, 0, 0, 166, 0, 0, 0, 0, 0, 191, 0, 0, 0, 295, 0, 159, 0, 216, 0, 335, 0, 0, 0, 52, 0, 75, 0, 5, 228, 0, 0, 0, 167, 225, 158, 0, 9, 0, 0, 0, 0, 136, 55, 0, 0, 0, 0, 0, 151, 1, 0, 329, 0, 0, 0, 8, 164, 147, 13, 288, 0, 0, 0, 144, 0, 0, 0, 0, 121, 0, 0, 202, 84, 0, 0, 0, 0, 0, 179, 0, 0, 0, 112, 0, 0, 120, 205, 0, 0, 71, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 18, 12, 0, 25, 0, 0, 72, 73, 37, 0, 0, 0, 344, 0, 323, 0, 258, 0, 0, 0, 0, 0, 0, 339, 0, 214, 0, 182, 325, 0, 0}

var _Hash_seeds = []int32{-67, 0, 0, 1, 1, 0, 0, 0, 0, -266, 0, 1, -254, -251, 0, -244, -239, -231, 0, 0, -223, 0, 0, -215, 0, 0, 0, 0, -209, 0, -194, 0, -1, -188, 0, -174, -158, -157, -154, 1, 3, 0, -146, 0, 0, -126, 0, 0, 1, 1, 1, -110, -91, 0, 0, 0, 1, 0, 0, 4, -78, 1, -75, -68, 0, -59, 6, -58, -56, 1, 0, -46, 0, 0, -35, -27, 7, 0, 0, 0, -14, -10, 0, -3, 0, 0, -9, 0, 0, -12, 0, 0, 0, 0, -19, 0, -25, 0, -76, -28, -29, -31, -32, 0, 0, 0, 0, 0, 0, -47, 0, 0, -52, 0, 0, 0, 1, 2, 0, 0, 0, 0, 0, 0, 0, -70, 0, -72, 0, 0, -127, -77, 2, 0, 5, -83, 0, 0, 0, 0, 3, 0, -92, 0, 0, 0, -99, 0, 0, -106, 0, -109, 0, 0, -115, -117, 0, 0, -122, 0, -125, 0, 0, 0, 0, 0, -131, -133, -134, 0, -138, -139, -141, -144, 0, 0, -149, 0, 0, -153, 0, 0, 0, -159, -161, 0, 0, 0, 0, 0, -172, -173, -284, 0, 0, 0, -181, 0, -184, 0, 1, 0, -190, 0, 0, -193, 0, -195, -197, -198, 2, -202, -206, 0, -208, 0, 1, 0, 0, 1, 1, -219, 0, -222, 0, 0, 0, 1, -229, 1, -233, -237, 0, 0, -242, 1, 0, 0, 0, 0, 0, -255, -259, 0, 1, 0, 1, 3, -272, -273, -274, 0, 0, 2, 0, 0, -170, 0, 0, 0, 0, 1, -287, -286, 0, 0, 0, -285, 0, -281, -280, 1, 0, -276, 0, -264, 0, -260, 0, 1, 1, 0, -248, 0, 0, 0, 0, -247, -245, 0, 5, -238, 0, 0, 0, -227, 0, -226, 1, -221, -214, 1, 3, 0, -207, 3, 0, -192, 1, 0, -189, 0, -185, 0, -183, -180, 0, 0, -178, -176, 1, -169, -167, -165, -164, 0, 0, 1, 1, 0, 1, 0, -151, 2, 0, 0, 0, 0, -147, 1, 0, 0, -135, 1, 1, 0, -128, 0, 0, 1, 2, 0, 0, 1, 0, -120, -119, 0, -113, 0, 2, 0, 0, 0, -107, 0, -103, 0, 0, 0, -101, -95, -94, -93, 1, 1, 0, -87, 0, -85, -84, 0, 1, -288, 0, 2, 0, -73, 0, -71, -69, 1, 1, 0, -65, 0, 0, -64, 0, -62, -60, -57, 1, 0, 0, 0, 8, -51, 0, -50, 0, 0, 0, 1, 0, -43, -42, 0, 7, 0, 0, 1, 0, 0, 2, -26, 0, 0, 4, 2, -16, -15, 0, 0, 0, 0, 0, 0, 1, -11, 2, 0, 2, -5, 0, 0, -8, 0, 0, 0, -18, 0, 0, -21, 0, -38, -41, 0, -53, -79, -81, 2, 0, -89, 0, -90, 0, 0, 0, 0, -123, 1, -129, -130, 0, 0, -150, 1, 0, 0, -191, -201, 0, 0, 2, 0, -210, 1, 0, -212, -217, 0, 1, 2, -228, 1, -263, 0, 0, 0, 0, 0, 1, 0, -268, 0, -277, -279, 1, 0, 1, 3}
