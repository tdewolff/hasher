package main

// uses github.com/tdewolff/hasher
//go:generate hasher -type=Hash -file=data_test.go

//import "github.com/dgryski/go-metro"

// Hash defines perfect hashes for a predefined list of strings
type Hash uint32

// Identifiers for the hashes associated with the text in the comments.
const (
	A                Hash = 0x1     // a
	Abbr             Hash = 0x4     // abbr
	Accept           Hash = 0x3206  // accept
	Accept_Charset   Hash = 0x320e  // accept-charset
	Accesskey        Hash = 0x4409  // accesskey
	Acronym          Hash = 0xbb07  // acronym
	Action           Hash = 0x2b906 // action
	Address          Hash = 0x67607 // address
	Align            Hash = 0x1605  // align
	Alink            Hash = 0xd205  // alink
	Allowfullscreen  Hash = 0x23c0f // allowfullscreen
	Alt              Hash = 0xeb03  // alt
	Annotation       Hash = 0x2060a // annotation
	AnnotationXml    Hash = 0x2060d // annotationXml
	Applet           Hash = 0x16106 // applet
	Area             Hash = 0x38604 // area
	Article          Hash = 0x40707 // article
	Aside            Hash = 0x8305  // aside
	Async            Hash = 0xf705  // async
	Audio            Hash = 0x11305 // audio
	Autocomplete     Hash = 0x14a0c // autocomplete
	Autofocus        Hash = 0x15609 // autofocus
	Autoplay         Hash = 0x16b08 // autoplay
	Axis             Hash = 0x17304 // axis
	B                Hash = 0x101   // b
	Background       Hash = 0x1e0a  // background
	Base             Hash = 0x44d04 // base
	Basefont         Hash = 0x44d08 // basefont
	Bdi              Hash = 0xcb03  // bdi
	Bdo              Hash = 0x18a03 // bdo
	Bgcolor          Hash = 0x19d07 // bgcolor
	Bgsound          Hash = 0x1a407 // bgsound
	Big              Hash = 0x1ac03 // big
	Blink            Hash = 0x1af05 // blink
	Blockquote       Hash = 0x1b40a // blockquote
	Body             Hash = 0x4004  // body
	Border           Hash = 0x33806 // border
	Br               Hash = 0x202   // br
	Button           Hash = 0x1be06 // button
	Canvas           Hash = 0x7f06  // canvas
	Caption          Hash = 0x27e07 // caption
	Center           Hash = 0x62306 // center
	Challenge        Hash = 0x1eb09 // challenge
	Charset          Hash = 0x3907  // charset
	Checked          Hash = 0x3ad07 // checked
	Cite             Hash = 0xfb04  // cite
	Class            Hash = 0x1c905 // class
	Classid          Hash = 0x1c907 // classid
	Clear            Hash = 0x40b05 // clear
	Code             Hash = 0x1dc04 // code
	Codebase         Hash = 0x44908 // codebase
	Codetype         Hash = 0x1dc08 // codetype
	Col              Hash = 0x19f03 // col
	Colgroup         Hash = 0x1f408 // colgroup
	Color            Hash = 0x19f05 // color
	Cols             Hash = 0x20104 // cols
	Colspan          Hash = 0x20107 // colspan
	Command          Hash = 0x21307 // command
	Compact          Hash = 0x21a07 // compact
	Content          Hash = 0x58107 // content
	Contenteditable  Hash = 0x5810f // contenteditable
	Contextmenu      Hash = 0x3b60b // contextmenu
	Controls         Hash = 0x22908 // controls
	Coords           Hash = 0x23506 // coords
	Crossorigin      Hash = 0x25a0b // crossorigin
	Data             Hash = 0x4a604 // data
	Datalist         Hash = 0x4a608 // datalist
	Datetime         Hash = 0x2e908 // datetime
	Dd               Hash = 0x31602 // dd
	Declare          Hash = 0x8607  // declare
	Default          Hash = 0x5407  // default
	DefaultChecked   Hash = 0x4ea0e // defaultChecked
	DefaultMuted     Hash = 0x54b0c // defaultMuted
	DefaultSelected  Hash = 0x540f  // defaultSelected
	Defer            Hash = 0x6205  // defer
	Del              Hash = 0x7203  // del
	Desc             Hash = 0x7c04  // desc
	Details          Hash = 0x9207  // details
	Dfn              Hash = 0xab03  // dfn
	Dialog           Hash = 0xcc06  // dialog
	Dir              Hash = 0xd903  // dir
	Dirname          Hash = 0xd907  // dirname
	Disabled         Hash = 0x10108 // disabled
	Div              Hash = 0x10803 // div
	Dl               Hash = 0x1aa02 // dl
	Download         Hash = 0x47f08 // download
	Draggable        Hash = 0x1cf09 // draggable
	Dropzone         Hash = 0x41208 // dropzone
	Dt               Hash = 0x5ff02 // dt
	Em               Hash = 0x6e02  // em
	Embed            Hash = 0x6e05  // embed
	Enabled          Hash = 0x4e07  // enabled
	Enctype          Hash = 0x2ce07 // enctype
	Face             Hash = 0x62104 // face
	Fieldset         Hash = 0x26b08 // fieldset
	Figcaption       Hash = 0x27b0a // figcaption
	Figure           Hash = 0x28f06 // figure
	Font             Hash = 0x45104 // font
	Footer           Hash = 0xee06  // footer
	For              Hash = 0x29b03 // for
	ForeignObject    Hash = 0x29b0d // foreignObject
	Foreignobject    Hash = 0x2a80d // foreignobject
	Form             Hash = 0x2b504 // form
	Formaction       Hash = 0x2b50a // formaction
	Formenctype      Hash = 0x2ca0b // formenctype
	Formmethod       Hash = 0x2d50a // formmethod
	Formnovalidate   Hash = 0x2df0e // formnovalidate
	Formtarget       Hash = 0x2f40a // formtarget
	Frame            Hash = 0xa305  // frame
	Frameborder      Hash = 0x3330b // frameborder
	Frameset         Hash = 0xa308  // frameset
	H1               Hash = 0x19b02 // h1
	H2               Hash = 0x32402 // h2
	H3               Hash = 0x34902 // h3
	H4               Hash = 0x37f02 // h4
	H5               Hash = 0x60102 // h5
	H6               Hash = 0x2fe02 // h6
	Head             Hash = 0x36b04 // head
	Header           Hash = 0x36b06 // header
	Headers          Hash = 0x36b07 // headers
	Height           Hash = 0x30006 // height
	Hgroup           Hash = 0x30806 // hgroup
	Hidden           Hash = 0x31406 // hidden
	High             Hash = 0x32104 // high
	Hr               Hash = 0xaf02  // hr
	Href             Hash = 0xaf04  // href
	Hreflang         Hash = 0xaf08  // hreflang
	Html             Hash = 0x30404 // html
	Http_Equiv       Hash = 0x3260a // http-equiv
	I                Hash = 0x601   // i
	Icon             Hash = 0x58004 // icon
	Id               Hash = 0x8502  // id
	Iframe           Hash = 0x33206 // iframe
	Image            Hash = 0x33e05 // image
	Img              Hash = 0x34303 // img
	Inert            Hash = 0x53605 // inert
	Input            Hash = 0x46c05 // input
	Ins              Hash = 0x26303 // ins
	Isindex          Hash = 0x17507 // isindex
	Ismap            Hash = 0x34b05 // ismap
	Itemid           Hash = 0xfc06  // itemid
	Itemprop         Hash = 0x56e08 // itemprop
	Itemref          Hash = 0x61b07 // itemref
	Itemscope        Hash = 0x35609 // itemscope
	Itemtype         Hash = 0x36008 // itemtype
	Kbd              Hash = 0xca03  // kbd
	Keygen           Hash = 0x4a06  // keygen
	Keytype          Hash = 0x5b007 // keytype
	Kind             Hash = 0xd604  // kind
	Label            Hash = 0x7405  // label
	Lang             Hash = 0xb304  // lang
	Language         Hash = 0xb308  // language
	Legend           Hash = 0x1d606 // legend
	Li               Hash = 0x1702  // li
	Link             Hash = 0xd304  // link
	List             Hash = 0x4aa04 // list
	Listing          Hash = 0x4aa07 // listing
	Longdesc         Hash = 0x7808  // longdesc
	Loop             Hash = 0x11e04 // loop
	Low              Hash = 0x23e03 // low
	Main             Hash = 0x1004  // main
	Malignmark       Hash = 0xc10a  // malignmark
	Manifest         Hash = 0x65708 // manifest
	Map              Hash = 0x16003 // map
	Mark             Hash = 0xc704  // mark
	Marquee          Hash = 0x69907 // marquee
	Math             Hash = 0x36804 // math
	Max              Hash = 0x37703 // max
	Maxlength        Hash = 0x37709 // maxlength
	Media            Hash = 0xde05  // media
	Mediagroup       Hash = 0xde0a  // mediagroup
	Menu             Hash = 0x3bd04 // menu
	Meta             Hash = 0x4b904 // meta
	Meter            Hash = 0x2ef05 // meter
	Method           Hash = 0x2d906 // method
	Mglyph           Hash = 0x34406 // mglyph
	Mi               Hash = 0x2c02  // mi
	Min              Hash = 0x2c03  // min
	Mn               Hash = 0x2e202 // mn
	Mo               Hash = 0x4dd02 // mo
	Ms               Hash = 0x35902 // ms
	Mtext            Hash = 0x38105 // mtext
	Multiple         Hash = 0x38f08 // multiple
	Muted            Hash = 0x39705 // muted
	Name             Hash = 0xdc04  // name
	Nav              Hash = 0x1303  // nav
	Nobr             Hash = 0x1a04  // nobr
	Noembed          Hash = 0x6c07  // noembed
	Noframes         Hash = 0xa108  // noframes
	Nohref           Hash = 0xad06  // nohref
	Noresize         Hash = 0x24a08 // noresize
	Noscript         Hash = 0x31908 // noscript
	Noshade          Hash = 0x4e507 // noshade
	Novalidate       Hash = 0x2e30a // novalidate
	Nowrap           Hash = 0x57706 // nowrap
	Object           Hash = 0x2af06 // object
	Ol               Hash = 0x12d02 // ol
	Onabort          Hash = 0x1c207 // onabort
	Onafterprint     Hash = 0x2830c // onafterprint
	Onbeforeprint    Hash = 0x2bd0d // onbeforeprint
	Onbeforeunload   Hash = 0x66a0e // onbeforeunload
	Onblur           Hash = 0x14406 // onblur
	Oncancel         Hash = 0x11708 // oncancel
	Oncanplay        Hash = 0x18c09 // oncanplay
	Oncanplaythrough Hash = 0x18c10 // oncanplaythrough
	Onchange         Hash = 0x42808 // onchange
	Onclick          Hash = 0x6a407 // onclick
	Onclose          Hash = 0x39c07 // onclose
	Oncontextmenu    Hash = 0x3b40d // oncontextmenu
	Oncuechange      Hash = 0x3c10b // oncuechange
	Ondblclick       Hash = 0x3cc0a // ondblclick
	Ondrag           Hash = 0x3d606 // ondrag
	Ondragend        Hash = 0x3d609 // ondragend
	Ondragenter      Hash = 0x3df0b // ondragenter
	Ondragleave      Hash = 0x3ea0b // ondragleave
	Ondragover       Hash = 0x3f50a // ondragover
	Ondragstart      Hash = 0x3ff0b // ondragstart
	Ondrop           Hash = 0x41006 // ondrop
	Ondurationchange Hash = 0x42010 // ondurationchange
	Onemptied        Hash = 0x41709 // onemptied
	Onended          Hash = 0x43007 // onended
	Onerror          Hash = 0x43707 // onerror
	Onfocus          Hash = 0x43e07 // onfocus
	Onhashchange     Hash = 0x45e0c // onhashchange
	Oninput          Hash = 0x46a07 // oninput
	Oninvalid        Hash = 0x47109 // oninvalid
	Onkeydown        Hash = 0x47a09 // onkeydown
	Onkeypress       Hash = 0x4870a // onkeypress
	Onkeyup          Hash = 0x49707 // onkeyup
	Onload           Hash = 0x49e06 // onload
	Onloadeddata     Hash = 0x49e0c // onloadeddata
	Onloadedmetadata Hash = 0x4b110 // onloadedmetadata
	Onloadstart      Hash = 0x4c70b // onloadstart
	Onmessage        Hash = 0x4d209 // onmessage
	Onmousedown      Hash = 0x4db0b // onmousedown
	Onmousemove      Hash = 0x4f80b // onmousemove
	Onmouseout       Hash = 0x5030a // onmouseout
	Onmouseover      Hash = 0x5100b // onmouseover
	Onmouseup        Hash = 0x51b09 // onmouseup
	Onmousewheel     Hash = 0x5240c // onmousewheel
	Onoffline        Hash = 0x53009 // onoffline
	Ononline         Hash = 0x53b08 // ononline
	Onpagehide       Hash = 0x5430a // onpagehide
	Onpageshow       Hash = 0x5570a // onpageshow
	Onpause          Hash = 0x56307 // onpause
	Onplay           Hash = 0x59006 // onplay
	Onplaying        Hash = 0x59009 // onplaying
	Onpopstate       Hash = 0x5990a // onpopstate
	Onprogress       Hash = 0x5a30a // onprogress
	Onratechange     Hash = 0x5b70c // onratechange
	Onreset          Hash = 0x5c307 // onreset
	Onresize         Hash = 0x5ca08 // onresize
	Onscroll         Hash = 0x5d208 // onscroll
	Onseeked         Hash = 0x5dd08 // onseeked
	Onseeking        Hash = 0x5e509 // onseeking
	Onselect         Hash = 0x5ee08 // onselect
	Onshow           Hash = 0x5f806 // onshow
	Onstalled        Hash = 0x60309 // onstalled
	Onstorage        Hash = 0x60c09 // onstorage
	Onsubmit         Hash = 0x61508 // onsubmit
	Onsuspend        Hash = 0x62909 // onsuspend
	Ontimeupdate     Hash = 0x4520c // ontimeupdate
	Onunload         Hash = 0x63208 // onunload
	Onvolumechange   Hash = 0x63a0e // onvolumechange
	Onwaiting        Hash = 0x64809 // onwaiting
	Open             Hash = 0x57404 // open
	Optgroup         Hash = 0x12008 // optgroup
	Optimum          Hash = 0x65107 // optimum
	Option           Hash = 0x66606 // option
	Output           Hash = 0x50a06 // output
	P                Hash = 0xc01   // p
	Param            Hash = 0xc05   // param
	Pattern          Hash = 0x9b07  // pattern
	Pauseonexit      Hash = 0x5650b // pauseonexit
	Ping             Hash = 0xe704  // ping
	Placeholder      Hash = 0x1270b // placeholder
	Plaintext        Hash = 0x17d09 // plaintext
	Poster           Hash = 0x1fb06 // poster
	Pre              Hash = 0x30d03 // pre
	Preload          Hash = 0x30d07 // preload
	Profile          Hash = 0x34f07 // profile
	Progress         Hash = 0x5a508 // progress
	Prompt           Hash = 0x66006 // prompt
	Public           Hash = 0x57c06 // public
	Q                Hash = 0x8d01  // q
	Radiogroup       Hash = 0x30a   // radiogroup
	Rb               Hash = 0x1d02  // rb
	Readonly         Hash = 0x38708 // readonly
	Rel              Hash = 0x30e03 // rel
	Required         Hash = 0x8b08  // required
	Rev              Hash = 0x29303 // rev
	Reversed         Hash = 0x29308 // reversed
	Rows             Hash = 0x6604  // rows
	Rowspan          Hash = 0x6607  // rowspan
	Rp               Hash = 0x28902 // rp
	Rt               Hash = 0x1c702 // rt
	Rtc              Hash = 0x1c703 // rtc
	Ruby             Hash = 0xf304  // ruby
	Rules            Hash = 0x13105 // rules
	S                Hash = 0x3d01  // s
	Samp             Hash = 0x9804  // samp
	Sandbox          Hash = 0x13507 // sandbox
	Scope            Hash = 0x35a05 // scope
	Scoped           Hash = 0x35a06 // scoped
	Script           Hash = 0x31b06 // script
	Scrolling        Hash = 0x5d409 // scrolling
	Seamless         Hash = 0x3a108 // seamless
	Section          Hash = 0x13f07 // section
	Select           Hash = 0x5f006 // select
	Selected         Hash = 0x5f008 // selected
	Shape            Hash = 0x23005 // shape
	Size             Hash = 0x24e04 // size
	Sizes            Hash = 0x24e05 // sizes
	Small            Hash = 0x23a05 // small
	Sortable         Hash = 0x25208 // sortable
	Source           Hash = 0x26506 // source
	Spacer           Hash = 0x37106 // spacer
	Span             Hash = 0x6904  // span
	Spellcheck       Hash = 0x3a80a // spellcheck
	Src              Hash = 0x44403 // src
	Srcdoc           Hash = 0x44406 // srcdoc
	Srclang          Hash = 0x49007 // srclang
	Start            Hash = 0x40505 // start
	Step             Hash = 0x65d04 // step
	Strike           Hash = 0x5ac06 // strike
	Strong           Hash = 0x67c06 // strong
	Style            Hash = 0x68205 // style
	Sub              Hash = 0x61703 // sub
	Summary          Hash = 0x68707 // summary
	Sup              Hash = 0x68e03 // sup
	Svg              Hash = 0x69103 // svg
	System           Hash = 0x69406 // system
	Tabindex         Hash = 0x4bf08 // tabindex
	Target           Hash = 0x2f806 // target
	Tbody            Hash = 0x3f05  // tbody
	Td               Hash = 0xaa02  // td
	Text             Hash = 0x18204 // text
	Textarea         Hash = 0x38208 // textarea
	Tfoot            Hash = 0xed05  // tfoot
	Th               Hash = 0x19502 // th
	Thead            Hash = 0x36a05 // thead
	Time             Hash = 0x2ed04 // time
	Title            Hash = 0x16605 // title
	Tr               Hash = 0x18502 // tr
	Track            Hash = 0x18505 // track
	Translate        Hash = 0x22009 // translate
	Truespeed        Hash = 0x27209 // truespeed
	Tt               Hash = 0x9d02  // tt
	Type             Hash = 0x10f04 // type
	Typemustmatch    Hash = 0x1e00d // typemustmatch
	U                Hash = 0xb01   // u
	Ul               Hash = 0x5802  // ul
	Undeterminate    Hash = 0x250d  // undeterminate
	Usemap           Hash = 0x15d06 // usemap
	Valign           Hash = 0x1506  // valign
	Value            Hash = 0x10a05 // value
	Valuetype        Hash = 0x10a09 // valuetype
	Var              Hash = 0x32f03 // var
	Video            Hash = 0x6a005 // video
	Visible          Hash = 0x6ab07 // visible
	Vlink            Hash = 0x6b205 // vlink
	Wbr              Hash = 0x56003 // wbr
	Width            Hash = 0x5fd05 // width
	Wrap             Hash = 0x57904 // wrap
	Xmlns            Hash = 0x13b05 // xmlns
	Xmp              Hash = 0x17b03 // xmp
)

var HashMap = map[string]Hash{
	"a":                A,
	"abbr":             Abbr,
	"accept":           Accept,
	"accept-charset":   Accept_Charset,
	"accesskey":        Accesskey,
	"acronym":          Acronym,
	"action":           Action,
	"address":          Address,
	"align":            Align,
	"alink":            Alink,
	"allowfullscreen":  Allowfullscreen,
	"alt":              Alt,
	"annotation":       Annotation,
	"annotationXml":    AnnotationXml,
	"applet":           Applet,
	"area":             Area,
	"article":          Article,
	"aside":            Aside,
	"async":            Async,
	"audio":            Audio,
	"autocomplete":     Autocomplete,
	"autofocus":        Autofocus,
	"autoplay":         Autoplay,
	"axis":             Axis,
	"b":                B,
	"background":       Background,
	"base":             Base,
	"basefont":         Basefont,
	"bdi":              Bdi,
	"bdo":              Bdo,
	"bgcolor":          Bgcolor,
	"bgsound":          Bgsound,
	"big":              Big,
	"blink":            Blink,
	"blockquote":       Blockquote,
	"body":             Body,
	"border":           Border,
	"br":               Br,
	"button":           Button,
	"canvas":           Canvas,
	"caption":          Caption,
	"center":           Center,
	"challenge":        Challenge,
	"charset":          Charset,
	"checked":          Checked,
	"cite":             Cite,
	"class":            Class,
	"classid":          Classid,
	"clear":            Clear,
	"code":             Code,
	"codebase":         Codebase,
	"codetype":         Codetype,
	"col":              Col,
	"colgroup":         Colgroup,
	"color":            Color,
	"cols":             Cols,
	"colspan":          Colspan,
	"command":          Command,
	"compact":          Compact,
	"content":          Content,
	"contenteditable":  Contenteditable,
	"contextmenu":      Contextmenu,
	"controls":         Controls,
	"coords":           Coords,
	"crossorigin":      Crossorigin,
	"data":             Data,
	"datalist":         Datalist,
	"datetime":         Datetime,
	"dd":               Dd,
	"declare":          Declare,
	"default":          Default,
	"defaultChecked":   DefaultChecked,
	"defaultMuted":     DefaultMuted,
	"defaultSelected":  DefaultSelected,
	"defer":            Defer,
	"del":              Del,
	"desc":             Desc,
	"details":          Details,
	"dfn":              Dfn,
	"dialog":           Dialog,
	"dir":              Dir,
	"dirname":          Dirname,
	"disabled":         Disabled,
	"div":              Div,
	"dl":               Dl,
	"download":         Download,
	"draggable":        Draggable,
	"dropzone":         Dropzone,
	"dt":               Dt,
	"em":               Em,
	"embed":            Embed,
	"enabled":          Enabled,
	"enctype":          Enctype,
	"face":             Face,
	"fieldset":         Fieldset,
	"figcaption":       Figcaption,
	"figure":           Figure,
	"font":             Font,
	"footer":           Footer,
	"for":              For,
	"foreignObject":    ForeignObject,
	"foreignobject":    Foreignobject,
	"form":             Form,
	"formaction":       Formaction,
	"formenctype":      Formenctype,
	"formmethod":       Formmethod,
	"formnovalidate":   Formnovalidate,
	"formtarget":       Formtarget,
	"frame":            Frame,
	"frameborder":      Frameborder,
	"frameset":         Frameset,
	"h1":               H1,
	"h2":               H2,
	"h3":               H3,
	"h4":               H4,
	"h5":               H5,
	"h6":               H6,
	"head":             Head,
	"header":           Header,
	"headers":          Headers,
	"height":           Height,
	"hgroup":           Hgroup,
	"hidden":           Hidden,
	"high":             High,
	"hr":               Hr,
	"href":             Href,
	"hreflang":         Hreflang,
	"html":             Html,
	"http-equiv":       Http_Equiv,
	"i":                I,
	"icon":             Icon,
	"id":               Id,
	"iframe":           Iframe,
	"image":            Image,
	"img":              Img,
	"inert":            Inert,
	"input":            Input,
	"ins":              Ins,
	"isindex":          Isindex,
	"ismap":            Ismap,
	"itemid":           Itemid,
	"itemprop":         Itemprop,
	"itemref":          Itemref,
	"itemscope":        Itemscope,
	"itemtype":         Itemtype,
	"kbd":              Kbd,
	"keygen":           Keygen,
	"keytype":          Keytype,
	"kind":             Kind,
	"label":            Label,
	"lang":             Lang,
	"language":         Language,
	"legend":           Legend,
	"li":               Li,
	"link":             Link,
	"list":             List,
	"listing":          Listing,
	"longdesc":         Longdesc,
	"loop":             Loop,
	"low":              Low,
	"main":             Main,
	"malignmark":       Malignmark,
	"manifest":         Manifest,
	"map":              Map,
	"mark":             Mark,
	"marquee":          Marquee,
	"math":             Math,
	"max":              Max,
	"maxlength":        Maxlength,
	"media":            Media,
	"mediagroup":       Mediagroup,
	"menu":             Menu,
	"meta":             Meta,
	"meter":            Meter,
	"method":           Method,
	"mglyph":           Mglyph,
	"mi":               Mi,
	"min":              Min,
	"mn":               Mn,
	"mo":               Mo,
	"ms":               Ms,
	"mtext":            Mtext,
	"multiple":         Multiple,
	"muted":            Muted,
	"name":             Name,
	"nav":              Nav,
	"nobr":             Nobr,
	"noembed":          Noembed,
	"noframes":         Noframes,
	"nohref":           Nohref,
	"noresize":         Noresize,
	"noscript":         Noscript,
	"noshade":          Noshade,
	"novalidate":       Novalidate,
	"nowrap":           Nowrap,
	"object":           Object,
	"ol":               Ol,
	"onabort":          Onabort,
	"onafterprint":     Onafterprint,
	"onbeforeprint":    Onbeforeprint,
	"onbeforeunload":   Onbeforeunload,
	"onblur":           Onblur,
	"oncancel":         Oncancel,
	"oncanplay":        Oncanplay,
	"oncanplaythrough": Oncanplaythrough,
	"onchange":         Onchange,
	"onclick":          Onclick,
	"onclose":          Onclose,
	"oncontextmenu":    Oncontextmenu,
	"oncuechange":      Oncuechange,
	"ondblclick":       Ondblclick,
	"ondrag":           Ondrag,
	"ondragend":        Ondragend,
	"ondragenter":      Ondragenter,
	"ondragleave":      Ondragleave,
	"ondragover":       Ondragover,
	"ondragstart":      Ondragstart,
	"ondrop":           Ondrop,
	"ondurationchange": Ondurationchange,
	"onemptied":        Onemptied,
	"onended":          Onended,
	"onerror":          Onerror,
	"onfocus":          Onfocus,
	"onhashchange":     Onhashchange,
	"oninput":          Oninput,
	"oninvalid":        Oninvalid,
	"onkeydown":        Onkeydown,
	"onkeypress":       Onkeypress,
	"onkeyup":          Onkeyup,
	"onload":           Onload,
	"onloadeddata":     Onloadeddata,
	"onloadedmetadata": Onloadedmetadata,
	"onloadstart":      Onloadstart,
	"onmessage":        Onmessage,
	"onmousedown":      Onmousedown,
	"onmousemove":      Onmousemove,
	"onmouseout":       Onmouseout,
	"onmouseover":      Onmouseover,
	"onmouseup":        Onmouseup,
	"onmousewheel":     Onmousewheel,
	"onoffline":        Onoffline,
	"ononline":         Ononline,
	"onpagehide":       Onpagehide,
	"onpageshow":       Onpageshow,
	"onpause":          Onpause,
	"onplay":           Onplay,
	"onplaying":        Onplaying,
	"onpopstate":       Onpopstate,
	"onprogress":       Onprogress,
	"onratechange":     Onratechange,
	"onreset":          Onreset,
	"onresize":         Onresize,
	"onscroll":         Onscroll,
	"onseeked":         Onseeked,
	"onseeking":        Onseeking,
	"onselect":         Onselect,
	"onshow":           Onshow,
	"onstalled":        Onstalled,
	"onstorage":        Onstorage,
	"onsubmit":         Onsubmit,
	"onsuspend":        Onsuspend,
	"ontimeupdate":     Ontimeupdate,
	"onunload":         Onunload,
	"onvolumechange":   Onvolumechange,
	"onwaiting":        Onwaiting,
	"open":             Open,
	"optgroup":         Optgroup,
	"optimum":          Optimum,
	"option":           Option,
	"output":           Output,
	"p":                P,
	"param":            Param,
	"pattern":          Pattern,
	"pauseonexit":      Pauseonexit,
	"ping":             Ping,
	"placeholder":      Placeholder,
	"plaintext":        Plaintext,
	"poster":           Poster,
	"pre":              Pre,
	"preload":          Preload,
	"profile":          Profile,
	"progress":         Progress,
	"prompt":           Prompt,
	"public":           Public,
	"q":                Q,
	"radiogroup":       Radiogroup,
	"rb":               Rb,
	"readonly":         Readonly,
	"rel":              Rel,
	"required":         Required,
	"rev":              Rev,
	"reversed":         Reversed,
	"rows":             Rows,
	"rowspan":          Rowspan,
	"rp":               Rp,
	"rt":               Rt,
	"rtc":              Rtc,
	"ruby":             Ruby,
	"rules":            Rules,
	"s":                S,
	"samp":             Samp,
	"sandbox":          Sandbox,
	"scope":            Scope,
	"scoped":           Scoped,
	"script":           Script,
	"scrolling":        Scrolling,
	"seamless":         Seamless,
	"section":          Section,
	"select":           Select,
	"selected":         Selected,
	"shape":            Shape,
	"size":             Size,
	"sizes":            Sizes,
	"small":            Small,
	"sortable":         Sortable,
	"source":           Source,
	"spacer":           Spacer,
	"span":             Span,
	"spellcheck":       Spellcheck,
	"src":              Src,
	"srcdoc":           Srcdoc,
	"srclang":          Srclang,
	"start":            Start,
	"step":             Step,
	"strike":           Strike,
	"strong":           Strong,
	"style":            Style,
	"sub":              Sub,
	"summary":          Summary,
	"sup":              Sup,
	"svg":              Svg,
	"system":           System,
	"tabindex":         Tabindex,
	"target":           Target,
	"tbody":            Tbody,
	"td":               Td,
	"text":             Text,
	"textarea":         Textarea,
	"tfoot":            Tfoot,
	"th":               Th,
	"thead":            Thead,
	"time":             Time,
	"title":            Title,
	"tr":               Tr,
	"track":            Track,
	"translate":        Translate,
	"truespeed":        Truespeed,
	"tt":               Tt,
	"type":             Type,
	"typemustmatch":    Typemustmatch,
	"u":                U,
	"ul":               Ul,
	"undeterminate":    Undeterminate,
	"usemap":           Usemap,
	"valign":           Valign,
	"value":            Value,
	"valuetype":        Valuetype,
	"var":              Var,
	"video":            Video,
	"visible":          Visible,
	"vlink":            Vlink,
	"wbr":              Wbr,
	"width":            Width,
	"wrap":             Wrap,
	"xmlns":            Xmlns,
	"xmp":              Xmp,
}

// String returns the text associated with the hash.
func (i Hash) String() string {
	return string(i.Bytes())
}

// Bytes returns the text associated with the hash.
func (i Hash) Bytes() []byte {
	start := uint32(i >> 8)
	n := uint32(i & 0xff)
	if start+n > uint32(len(_Hash_text)) {
		return []byte{}
	}
	return _Hash_text[start : start+n]
}

// ToHash returns a hash Hash for a given []byte. Hash is a uint32 that is associated with the text in []byte. It returns zero if no match found.
func ToHash(s []byte) Hash {
	if 3 < len(s) {
		return HashMap[string(s)]
	}
	h := uint32(_Hash_hash0)
	for i := 0; i < len(s); i++ {
		h ^= uint32(s[i])
		h *= 16777619
	}
	if i := _Hash_table[h&uint32(len(_Hash_table)-1)]; int(i&0xff) == len(s) {
		t := _Hash_text[i>>8 : i>>8+i&0xff]
		for i := 0; i < len(s); i++ {
			if t[i] != s[i] {
				goto NEXT
			}
		}
		return i
	}
NEXT:
	if i := _Hash_table[(h>>16)&uint32(len(_Hash_table)-1)]; int(i&0xff) == len(s) {
		t := _Hash_text[i>>8 : i>>8+i&0xff]
		for i := 0; i < len(s); i++ {
			if t[i] != s[i] {
				return 0
			}
		}
		return i
	}
	return 0
}

const _Hash_hash0 = 0x5334b67c
const _Hash_maxLen = 16

var _Hash_text = []byte("" +
	"abbradiogrouparamainavalignobrbackgroundeterminateaccept-cha" +
	"rsetbodyaccesskeygenabledefaultSelectedeferowspanoembedelabe" +
	"longdescanvasideclarequiredetailsampatternoframesetdfnohrefl" +
	"anguageacronymalignmarkbdialogalinkindirnamediagroupingaltfo" +
	"oterubyasyncitemidisabledivaluetypeaudioncancelooptgrouplace" +
	"holderulesandboxmlnsectionblurautocompleteautofocusemappleti" +
	"tleautoplayaxisindexmplaintextrackbdoncanplaythrough1bgcolor" +
	"bgsoundlbigblinkblockquotebuttonabortclassidraggablegendcode" +
	"typemustmatchallengecolgroupostercolspannotationXmlcommandco" +
	"mpactranslatecontrolshapecoordsmallowfullscreenoresizesortab" +
	"lecrossoriginsourcefieldsetruespeedfigcaptionafterprintfigur" +
	"eversedforeignObjectforeignobjectformactionbeforeprintformen" +
	"ctypeformmethodformnovalidatetimeterformtargeth6heightmlhgro" +
	"upreloadhiddenoscripthigh2http-equivariframeborderimageimgly" +
	"ph3ismaprofileitemscopeditemtypematheaderspacermaxlength4mte" +
	"xtareadonlymultiplemutedoncloseamlesspellcheckedoncontextmen" +
	"uoncuechangeondblclickondragendondragenterondragleaveondrago" +
	"verondragstarticlearondropzonemptiedondurationchangeonendedo" +
	"nerroronfocusrcdocodebasefontimeupdateonhashchangeoninputoni" +
	"nvalidonkeydownloadonkeypressrclangonkeyuponloadeddatalistin" +
	"gonloadedmetadatabindexonloadstartonmessageonmousedownoshade" +
	"faultCheckedonmousemoveonmouseoutputonmouseoveronmouseuponmo" +
	"usewheelonofflinertononlineonpagehidefaultMutedonpageshowbro" +
	"npauseonexitempropenowrapublicontenteditableonplayingonpopst" +
	"ateonprogresstrikeytypeonratechangeonresetonresizeonscrollin" +
	"gonseekedonseekingonselectedonshowidth5onstalledonstorageons" +
	"ubmitemrefacenteronsuspendonunloadonvolumechangeonwaitingopt" +
	"imumanifestepromptoptionbeforeunloaddresstrongstylesummarysu" +
	"psvgsystemarqueevideonclickvisiblevlink")

var _Hash_table = [1 << 9]Hash{
	0x0:   0x2ca0b, // formenctype
	0x1:   0x2d50a, // formmethod
	0x2:   0x3c10b, // oncuechange
	0x3:   0x3d606, // ondrag
	0x6:   0x5ac06, // strike
	0x7:   0x6a005, // video
	0x9:   0x58107, // content
	0xa:   0x4e07,  // enabled
	0xb:   0x57706, // nowrap
	0xc:   0xd304,  // link
	0xe:   0x28902, // rp
	0xf:   0x2830c, // onafterprint
	0x10:  0x16106, // applet
	0x11:  0xed05,  // tfoot
	0x12:  0x4ea0e, // defaultChecked
	0x13:  0x3330b, // frameborder
	0x14:  0xee06,  // footer
	0x15:  0x5f008, // selected
	0x16:  0x49007, // srclang
	0x18:  0x5100b, // onmouseover
	0x19:  0x1dc04, // code
	0x1b:  0x47109, // oninvalid
	0x1c:  0x62104, // face
	0x1e:  0x3b60b, // contextmenu
	0x1f:  0xa308,  // frameset
	0x21:  0x54b0c, // defaultMuted
	0x22:  0x19f05, // color
	0x23:  0x59006, // onplay
	0x25:  0x2ef05, // meter
	0x26:  0x60c09, // onstorage
	0x27:  0x38708, // readonly
	0x29:  0x34f07, // profile
	0x2a:  0x5570a, // onpageshow
	0x2b:  0xb01,   // u
	0x2c:  0x31908, // noscript
	0x2d:  0x65708, // manifest
	0x2e:  0x1be06, // button
	0x2f:  0x2e908, // datetime
	0x30:  0x46c05, // input
	0x31:  0x5407,  // default
	0x32:  0x1dc08, // codetype
	0x33:  0x2a80d, // foreignobject
	0x34:  0x69907, // marquee
	0x36:  0x19d07, // bgcolor
	0x37:  0x19b02, // h1
	0x39:  0x1e0a,  // background
	0x3b:  0x2f40a, // formtarget
	0x41:  0x2f806, // target
	0x43:  0x23a05, // small
	0x44:  0x44908, // codebase
	0x45:  0x53605, // inert
	0x47:  0x38105, // mtext
	0x48:  0x6607,  // rowspan
	0x49:  0x2bd0d, // onbeforeprint
	0x4a:  0x53b08, // ononline
	0x4c:  0x28f06, // figure
	0x4d:  0x4b110, // onloadedmetadata
	0x4e:  0xbb07,  // acronym
	0x50:  0x38f08, // multiple
	0x51:  0x320e,  // accept-charset
	0x52:  0x24e05, // sizes
	0x53:  0x29b0d, // foreignObject
	0x55:  0x2e30a, // novalidate
	0x56:  0x5430a, // onpagehide
	0x57:  0x2e202, // mn
	0x58:  0x37f02, // h4
	0x5a:  0x1c702, // rt
	0x5b:  0xd205,  // alink
	0x5e:  0x66006, // prompt
	0x5f:  0x12d02, // ol
	0x61:  0x5ca08, // onresize
	0x64:  0x68707, // summary
	0x65:  0x5990a, // onpopstate
	0x66:  0x38604, // area
	0x68:  0x64809, // onwaiting
	0x6b:  0xdc04,  // name
	0x6c:  0x23506, // coords
	0x6d:  0x34303, // img
	0x6e:  0x65d04, // step
	0x6f:  0x5e509, // onseeking
	0x70:  0x32104, // high
	0x71:  0x49707, // onkeyup
	0x72:  0x5f006, // select
	0x73:  0x18505, // track
	0x74:  0x34b05, // ismap
	0x77:  0x8d01,  // q
	0x78:  0x47a09, // onkeydown
	0x79:  0x33e05, // image
	0x7a:  0x2b504, // form
	0x7b:  0x60309, // onstalled
	0x7d:  0x42808, // onchange
	0x7e:  0x1af05, // blink
	0x7f:  0xeb03,  // alt
	0x80:  0xf705,  // async
	0x82:  0x1702,  // li
	0x84:  0x2c02,  // mi
	0x85:  0xfc06,  // itemid
	0x86:  0x11305, // audio
	0x87:  0x31b06, // script
	0x8b:  0x44406, // srcdoc
	0x8e:  0xc704,  // mark
	0x8f:  0x18a03, // bdo
	0x91:  0x4f80b, // onmousemove
	0x93:  0x3bd04, // menu
	0x94:  0x45104, // font
	0x95:  0x16b08, // autoplay
	0x96:  0x6b205, // vlink
	0x98:  0x6e02,  // em
	0x9b:  0x1f408, // colgroup
	0x9c:  0x57404, // open
	0x9d:  0x1d606, // legend
	0x9e:  0x4c70b, // onloadstart
	0xa2:  0x22009, // translate
	0xa3:  0x6e05,  // embed
	0xa4:  0x1c905, // class
	0xa7:  0x36b06, // header
	0xa9:  0x49e06, // onload
	0xaa:  0x36a05, // thead
	0xab:  0x5d409, // scrolling
	0xac:  0xc05,   // param
	0xae:  0x9b07,  // pattern
	0xaf:  0x9207,  // details
	0xb1:  0x57c06, // public
	0xb3:  0x4db0b, // onmousedown
	0xb4:  0x16003, // map
	0xb6:  0x25a0b, // crossorigin
	0xb7:  0x1506,  // valign
	0xb9:  0x1c207, // onabort
	0xba:  0x66606, // option
	0xbb:  0x26506, // source
	0xbc:  0x6205,  // defer
	0xbd:  0x1eb09, // challenge
	0xbf:  0x10a05, // value
	0xc0:  0x23c0f, // allowfullscreen
	0xc1:  0xca03,  // kbd
	0xc2:  0x2060d, // annotationXml
	0xc3:  0x5b70c, // onratechange
	0xc4:  0x4dd02, // mo
	0xc6:  0x3a80a, // spellcheck
	0xc7:  0x2c03,  // min
	0xc8:  0x49e0c, // onloadeddata
	0xc9:  0x40b05, // clear
	0xca:  0x42010, // ondurationchange
	0xcb:  0x1a04,  // nobr
	0xcd:  0x27209, // truespeed
	0xcf:  0x30806, // hgroup
	0xd0:  0x40505, // start
	0xd3:  0x41208, // dropzone
	0xd5:  0x7405,  // label
	0xd8:  0xde0a,  // mediagroup
	0xd9:  0x14406, // onblur
	0xdb:  0x27e07, // caption
	0xdd:  0x7c04,  // desc
	0xde:  0x13b05, // xmlns
	0xdf:  0x30006, // height
	0xe0:  0x21307, // command
	0xe2:  0x5650b, // pauseonexit
	0xe3:  0x67c06, // strong
	0xe4:  0x43707, // onerror
	0xe5:  0x61508, // onsubmit
	0xe6:  0xb308,  // language
	0xe7:  0x47f08, // download
	0xe9:  0x51b09, // onmouseup
	0xec:  0x2ce07, // enctype
	0xed:  0x5ee08, // onselect
	0xee:  0x2af06, // object
	0xef:  0x17d09, // plaintext
	0xf0:  0x3cc0a, // ondblclick
	0xf1:  0x18c10, // oncanplaythrough
	0xf2:  0xd903,  // dir
	0xf3:  0x38208, // textarea
	0xf4:  0xe704,  // ping
	0xf5:  0x2d906, // method
	0xf6:  0x22908, // controls
	0xf7:  0x37106, // spacer
	0xf8:  0x69103, // svg
	0xf9:  0x30404, // html
	0xfa:  0x3d01,  // s
	0xfc:  0xcc06,  // dialog
	0xfe:  0x1e00d, // typemustmatch
	0xff:  0x3ad07, // checked
	0x101: 0x1fb06, // poster
	0x102: 0x3260a, // http-equiv
	0x103: 0x44403, // src
	0x104: 0x10108, // disabled
	0x105: 0x36b07, // headers
	0x106: 0x5a30a, // onprogress
	0x107: 0x26b08, // fieldset
	0x108: 0x32f03, // var
	0x10a: 0xa305,  // frame
	0x10b: 0x36008, // itemtype
	0x10c: 0x3f50a, // ondragover
	0x10d: 0x15609, // autofocus
	0x10f: 0x601,   // i
	0x110: 0x35902, // ms
	0x111: 0x44d04, // base
	0x113: 0x35a05, // scope
	0x114: 0x3206,  // accept
	0x115: 0x56e08, // itemprop
	0x117: 0xfb04,  // cite
	0x118: 0x3907,  // charset
	0x119: 0x16605, // title
	0x11a: 0x5b007, // keytype
	0x11b: 0x18204, // text
	0x11c: 0x65107, // optimum
	0x11e: 0x36b04, // head
	0x121: 0x21a07, // compact
	0x123: 0x62909, // onsuspend
	0x124: 0x4aa04, // list
	0x125: 0x4520c, // ontimeupdate
	0x126: 0x62306, // center
	0x127: 0x31406, // hidden
	0x129: 0x35609, // itemscope
	0x12c: 0x1aa02, // dl
	0x12d: 0x13f07, // section
	0x12e: 0x11708, // oncancel
	0x12f: 0x6a407, // onclick
	0x130: 0xde05,  // media
	0x131: 0x50a06, // output
	0x132: 0x4a608, // datalist
	0x133: 0x5240c, // onmousewheel
	0x134: 0x44d08, // basefont
	0x135: 0x37709, // maxlength
	0x136: 0x6ab07, // visible
	0x137: 0x2df0e, // formnovalidate
	0x139: 0x17b03, // xmp
	0x13a: 0x101,   // b
	0x13b: 0x46a07, // oninput
	0x13c: 0xf304,  // ruby
	0x13d: 0x1270b, // placeholder
	0x13e: 0x4aa07, // listing
	0x140: 0x26303, // ins
	0x141: 0x61b07, // itemref
	0x144: 0x540f,  // defaultSelected
	0x146: 0x3ea0b, // ondragleave
	0x147: 0x1b40a, // blockquote
	0x148: 0x57904, // wrap
	0x14a: 0x1ac03, // big
	0x14b: 0x30e03, // rel
	0x14c: 0x41006, // ondrop
	0x14e: 0x69406, // system
	0x14f: 0x30a,   // radiogroup
	0x150: 0xb304,  // lang
	0x152: 0x56003, // wbr
	0x153: 0x3b40d, // oncontextmenu
	0x155: 0x250d,  // undeterminate
	0x157: 0x20104, // cols
	0x158: 0x13507, // sandbox
	0x159: 0x1303,  // nav
	0x15a: 0x37703, // max
	0x15b: 0x7808,  // longdesc
	0x15c: 0x5fd05, // width
	0x15d: 0x34902, // h3
	0x15e: 0x1a407, // bgsound
	0x161: 0x10a09, // valuetype
	0x162: 0x68205, // style
	0x164: 0x3f05,  // tbody
	0x165: 0x40707, // article
	0x169: 0xcb03,  // bdi
	0x16a: 0x67607, // address
	0x16b: 0x23005, // shape
	0x16c: 0x2b906, // action
	0x16e: 0x18502, // tr
	0x16f: 0xaa02,  // td
	0x170: 0x3d609, // ondragend
	0x171: 0x5802,  // ul
	0x172: 0x33806, // border
	0x174: 0x4a06,  // keygen
	0x175: 0x4004,  // body
	0x177: 0x1cf09, // draggable
	0x178: 0x2b50a, // formaction
	0x17b: 0x34406, // mglyph
	0x17d: 0x1d02,  // rb
	0x17e: 0x2fe02, // h6
	0x17f: 0x41709, // onemptied
	0x180: 0x5c307, // onreset
	0x181: 0x1004,  // main
	0x182: 0x11e04, // loop
	0x183: 0x4870a, // onkeypress
	0x184: 0x9d02,  // tt
	0x186: 0x20107, // colspan
	0x188: 0x36804, // math
	0x189: 0x1605,  // align
	0x18a: 0xa108,  // noframes
	0x18b: 0xaf02,  // hr
	0x18c: 0xc10a,  // malignmark
	0x18e: 0x23e03, // low
	0x18f: 0x8502,  // id
	0x190: 0x6604,  // rows
	0x191: 0x29303, // rev
	0x192: 0x63208, // onunload
	0x193: 0x39705, // muted
	0x194: 0x35a06, // scoped
	0x195: 0x31602, // dd
	0x196: 0x5ff02, // dt
	0x197: 0x66a0e, // onbeforeunload
	0x199: 0x2060a, // annotation
	0x19a: 0x29308, // reversed
	0x19c: 0x10f04, // type
	0x19d: 0x56307, // onpause
	0x19e: 0xd604,  // kind
	0x19f: 0x4a604, // data
	0x1a0: 0x4e507, // noshade
	0x1a3: 0x13105, // rules
	0x1a4: 0x12008, // optgroup
	0x1a5: 0x202,   // br
	0x1a7: 0x1,     // a
	0x1a8: 0x5030a, // onmouseout
	0x1aa: 0x53009, // onoffline
	0x1ab: 0x63a0e, // onvolumechange
	0x1ae: 0x61703, // sub
	0x1b3: 0x29b03, // for
	0x1b5: 0x8b08,  // required
	0x1b6: 0x5a508, // progress
	0x1b7: 0x15d06, // usemap
	0x1b8: 0x7f06,  // canvas
	0x1b9: 0x58004, // icon
	0x1bb: 0x1c703, // rtc
	0x1bc: 0x8305,  // aside
	0x1bd: 0x2ed04, // time
	0x1be: 0x3ff0b, // ondragstart
	0x1c0: 0x27b0a, // figcaption
	0x1c1: 0xaf04,  // href
	0x1c2: 0x33206, // iframe
	0x1c3: 0x18c09, // oncanplay
	0x1c4: 0x6904,  // span
	0x1c5: 0x30d03, // pre
	0x1c6: 0x6c07,  // noembed
	0x1c8: 0x5dd08, // onseeked
	0x1c9: 0x4b904, // meta
	0x1ca: 0x32402, // h2
	0x1cb: 0x3a108, // seamless
	0x1cc: 0xab03,  // dfn
	0x1cd: 0x17304, // axis
	0x1cf: 0x3df0b, // ondragenter
	0x1d0: 0x19502, // th
	0x1d1: 0x45e0c, // onhashchange
	0x1d2: 0x8607,  // declare
	0x1d3: 0x43e07, // onfocus
	0x1d5: 0x24e04, // size
	0x1d8: 0x14a0c, // autocomplete
	0x1d9: 0xaf08,  // hreflang
	0x1da: 0x9804,  // samp
	0x1de: 0x19f03, // col
	0x1df: 0x10803, // div
	0x1e0: 0x25208, // sortable
	0x1e1: 0x7203,  // del
	0x1e3: 0x39c07, // onclose
	0x1e6: 0xd907,  // dirname
	0x1e8: 0x1c907, // classid
	0x1e9: 0x30d07, // preload
	0x1ea: 0x4bf08, // tabindex
	0x1eb: 0x60102, // h5
	0x1ec: 0x5d208, // onscroll
	0x1ed: 0x5810f, // contenteditable
	0x1ee: 0x4d209, // onmessage
	0x1ef: 0x4,     // abbr
	0x1f0: 0x17507, // isindex
	0x1f1: 0x68e03, // sup
	0x1f3: 0x24a08, // noresize
	0x1f5: 0x59009, // onplaying
	0x1f6: 0x4409,  // accesskey
	0x1fa: 0xc01,   // p
	0x1fb: 0x43007, // onended
	0x1fc: 0x5f806, // onshow
	0x1fe: 0xad06,  // nohref
}
