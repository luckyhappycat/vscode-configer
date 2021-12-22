package cpp

// launch.json
type Launch struct {
	Version        string          `json:"version"`
	Configurations []Configuration `json:"configurations"`
}
type Configuration struct {
	Name            string   `json:"name"`
	Type            string   `json:"type"`
	Request         string   `json:"request"`
	Program         string   `json:"program"`
	Args            []string `json:"args"`
	StopAtEntry     bool     `json:"stopAtEntry"`
	Cwd             string   `json:"cwd"`
	Environment     []string `json:"environment"`
	ExternalConsole bool     `json:"externalConsole"`
	MIMode          string   `json:"MIMode"`
	PreLaunchTask   string   `json:"preLaunchTask"`
	PostDebugTask   string   `json:"postDebugTask"`
}

// settings.json
type Settings struct {
	CmakeConfigureOnOpen        bool              `json:"cmake.configureOnOpen"`
	FilesAssociations           FilesAssociations `json:"files.associations"`
	CCppClangFormatSortIncludes bool              `json:"C_Cpp.clang_format_sortIncludes"`
}
type FilesAssociations struct {
	Tcc             string `json:"*.tcc"`
	Streambuf       string `json:"streambuf"`
	String          string `json:"string"`
	Array           string `json:"array"`
	Cctype          string `json:"cctype"`
	Clocale         string `json:"clocale"`
	Cmath           string `json:"cmath"`
	Cstdarg         string `json:"cstdarg"`
	Cstdint         string `json:"cstdint"`
	Cstdio          string `json:"cstdio"`
	Cstdlib         string `json:"cstdlib"`
	Cstring         string `json:"cstring"`
	Ctime           string `json:"ctime"`
	Cwchar          string `json:"cwchar"`
	Cwctype         string `json:"cwctype"`
	Deque           string `json:"deque"`
	List            string `json:"list"`
	UnorderedMap    string `json:"unordered_map"`
	Vector          string `json:"vector"`
	Exception       string `json:"exception"`
	Fstream         string `json:"fstream"`
	Functional      string `json:"functional"`
	InitializerList string `json:"initializer_list"`
	Iosfwd          string `json:"iosfwd"`
	Iostream        string `json:"iostream"`
	Istream         string `json:"istream"`
	Limits          string `json:"limits"`
	New             string `json:"new"`
	Ostream         string `json:"ostream"`
	Numeric         string `json:"numeric"`
	Sstream         string `json:"sstream"`
	Stdexcept       string `json:"stdexcept"`
	Cinttypes       string `json:"cinttypes"`
	Tuple           string `json:"tuple"`
	TypeTraits      string `json:"type_traits"`
	Utility         string `json:"utility"`
	Typeinfo        string `json:"typeinfo"`
}

// tasks.json
type Tasks struct {
	Tasks   []Task `json:"tasks"`
	Version string `json:"version"`
}
type Options struct {
	Cwd string `json:"cwd"`
}
type Group struct {
	Kind      string `json:"kind"`
	IsDefault bool   `json:"isDefault"`
}
type Presentation struct {
	Echo             bool   `json:"echo"`
	Reveal           string `json:"reveal"`
	Focus            bool   `json:"focus"`
	Panel            string `json:"panel"`
	Clear            bool   `json:"clear"`
	ShowReuseMessage bool   `json:"showReuseMessage"`
}
type Task struct {
	Type           string       `json:"type"`
	Label          string       `json:"label"`
	Command        string       `json:"command"`
	Args           []string     `json:"args,omitempty"`
	Options        Options      `json:"options"`
	ProblemMatcher []string     `json:"problemMatcher,omitempty"`
	Group          Group        `json:"group"`
	Presentation   Presentation `json:"presentation,omitempty"`
}
