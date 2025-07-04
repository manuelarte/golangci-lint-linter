// Code generated from JSON Schema using quicktype. DO NOT EDIT.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    golangci, err := UnmarshalGolangci(bytes)
//    bytes, err = golangci.Marshal()

package golangci

import "encoding/json"

func UnmarshalGolangci(data []byte) (Golangci, error) {
	var r Golangci
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Golangci) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type Golangci struct {
	Formatters *Formatters `json:"formatters,omitempty" yaml:"formatters,omitempty" toml:"formatters,omitempty"`
	Issues     *Issues     `json:"issues,omitempty" yaml:"issues,omitempty" toml:"issues,omitempty"`
	Linters    *Linters    `json:"linters,omitempty" yaml:"linters,omitempty" toml:"linters,omitempty"`
	// Output configuration options.
	Output *Output `json:"output,omitempty" yaml:"output,omitempty" toml:"output,omitempty"`
	// Options for analysis running,
	Run      *Run           `json:"run,omitempty" yaml:"run,omitempty" toml:"run,omitempty"`
	Severity *SeverityClass `json:"severity,omitempty" yaml:"severity,omitempty" toml:"severity,omitempty"`
	Version  string         `json:"version" yaml:"version" toml:"version"`
}

type Formatters struct {
	// List of enabled formatters.
	Enable     []FormatterNames      `json:"enable,omitempty" yaml:"enable,omitempty" toml:"enable,omitempty"`
	Exclusions *FormattersExclusions `json:"exclusions,omitempty" yaml:"exclusions,omitempty" toml:"exclusions,omitempty"`
	Settings   *FormattersSettings   `json:"settings,omitempty" yaml:"settings,omitempty" toml:"settings,omitempty"`
}

type FormattersExclusions struct {
	Generated  *Generated `json:"generated,omitempty" yaml:"generated,omitempty" toml:"generated,omitempty"`
	Paths      []string   `json:"paths,omitempty" yaml:"paths,omitempty" toml:"paths,omitempty"`
	WarnUnused *bool      `json:"warn-unused,omitempty" yaml:"warn-unused,omitempty" toml:"warn-unused,omitempty"`
}

type FormattersSettings struct {
	Gci       *GciSettings       `json:"gci,omitempty" yaml:"gci,omitempty" toml:"gci,omitempty"`
	Gofmt     *GofmtSettings     `json:"gofmt,omitempty" yaml:"gofmt,omitempty" toml:"gofmt,omitempty"`
	Gofumpt   *GofumptSettings   `json:"gofumpt,omitempty" yaml:"gofumpt,omitempty" toml:"gofumpt,omitempty"`
	Goimports *GoimportsSettings `json:"goimports,omitempty" yaml:"goimports,omitempty" toml:"goimports,omitempty"`
	Golines   *GolinesSettings   `json:"golines,omitempty" yaml:"golines,omitempty" toml:"golines,omitempty"`
}

type GciSettings struct {
	// Enable custom order of sections.
	CustomOrder *bool `json:"custom-order,omitempty" yaml:"custom-order,omitempty" toml:"custom-order,omitempty"`
	// Checks that no inline Comments are present.
	NoInlineComments *bool `json:"no-inline-comments,omitempty" yaml:"no-inline-comments,omitempty" toml:"no-inline-comments,omitempty"`
	// Drops lexical ordering for custom sections.
	NoLexOrder *bool `json:"no-lex-order,omitempty" yaml:"no-lex-order,omitempty" toml:"no-lex-order,omitempty"`
	// Checks that no prefix Comments(comment lines above an import) are present.
	NoPrefixComments *bool `json:"no-prefix-comments,omitempty" yaml:"no-prefix-comments,omitempty" toml:"no-prefix-comments,omitempty"`
	// Section configuration to compare against.
	Sections []string `json:"sections,omitempty" yaml:"sections,omitempty" toml:"sections,omitempty"`
}

type GofmtSettings struct {
	// Apply the rewrite rules to the source before reformatting.
	RewriteRules []RewriteRule `json:"rewrite-rules,omitempty" yaml:"rewrite-rules,omitempty" toml:"rewrite-rules,omitempty"`
	// Simplify code.
	Simplify *bool `json:"simplify,omitempty" yaml:"simplify,omitempty" toml:"simplify,omitempty"`
}

type RewriteRule struct {
	Pattern     *string `json:"pattern,omitempty" yaml:"pattern,omitempty" toml:"pattern,omitempty"`
	Replacement *string `json:"replacement,omitempty" yaml:"replacement,omitempty" toml:"replacement,omitempty"`
}

type GofumptSettings struct {
	// Choose whether or not to use the extra rules that are disabled by default.
	ExtraRules *bool `json:"extra-rules,omitempty" yaml:"extra-rules,omitempty" toml:"extra-rules,omitempty"`
	// Module path which contains the source code being formatted.
	ModulePath *string `json:"module-path,omitempty" yaml:"module-path,omitempty" toml:"module-path,omitempty"`
}

type GoimportsSettings struct {
	// Put imports beginning with prefix after 3rd-party packages. It is a list of prefixes.
	LocalPrefixes []string `json:"local-prefixes,omitempty" yaml:"local-prefixes,omitempty" toml:"local-prefixes,omitempty"`
}

type GolinesSettings struct {
	ChainSplitDots  *bool  `json:"chain-split-dots,omitempty" yaml:"chain-split-dots,omitempty" toml:"chain-split-dots,omitempty"`
	MaxLen          *int64 `json:"max-len,omitempty" yaml:"max-len,omitempty" toml:"max-len,omitempty"`
	ReformatTags    *bool  `json:"reformat-tags,omitempty" yaml:"reformat-tags,omitempty" toml:"reformat-tags,omitempty"`
	ShortenComments *bool  `json:"shorten-comments,omitempty" yaml:"shorten-comments,omitempty" toml:"shorten-comments,omitempty"`
	TabLen          *int64 `json:"tab-len,omitempty" yaml:"tab-len,omitempty" toml:"tab-len,omitempty"`
}

type Issues struct {
	// Fix found issues (if it's supported by the linter).
	Fix *bool `json:"fix,omitempty" yaml:"fix,omitempty" toml:"fix,omitempty"`
	// Maximum issues count per one linter. Set to 0 to disable.
	MaxIssuesPerLinter *int64 `json:"max-issues-per-linter,omitempty" yaml:"max-issues-per-linter,omitempty" toml:"max-issues-per-linter,omitempty"`
	// Maximum count of issues with the same text. Set to 0 to disable.
	MaxSameIssues *int64 `json:"max-same-issues,omitempty" yaml:"max-same-issues,omitempty" toml:"max-same-issues,omitempty"`
	// Show only new issues: if there are unstaged changes or untracked files, only those
	// changes are analyzed, else only changes in HEAD~ are analyzed.
	New *bool `json:"new,omitempty" yaml:"new,omitempty" toml:"new,omitempty"`
	// Show only new issues created after the best common ancestor (merge-base against HEAD).
	NewFromMergeBase *string `json:"new-from-merge-base,omitempty" yaml:"new-from-merge-base,omitempty" toml:"new-from-merge-base,omitempty"`
	// Show only new issues created in git patch with this file path.
	NewFromPatch *string `json:"new-from-patch,omitempty" yaml:"new-from-patch,omitempty" toml:"new-from-patch,omitempty"`
	// Show only new issues created after this git revision.
	NewFromRev *string `json:"new-from-rev,omitempty" yaml:"new-from-rev,omitempty" toml:"new-from-rev,omitempty"`
	// Make issues output unique by line.
	UniqByLine *bool `json:"uniq-by-line,omitempty" yaml:"uniq-by-line,omitempty" toml:"uniq-by-line,omitempty"`
	// Show issues in any part of update files (requires new-from-rev or new-from-patch).
	WholeFiles *bool `json:"whole-files,omitempty" yaml:"whole-files,omitempty" toml:"whole-files,omitempty"`
}

type Linters struct {
	Default *LintersDefault `json:"default,omitempty" yaml:"default,omitempty" toml:"default,omitempty"`
	// List of disabled linters.
	Disable []string `json:"disable,omitempty" yaml:"disable,omitempty" toml:"disable,omitempty"`
	// List of enabled linters.
	Enable     []string           `json:"enable,omitempty" yaml:"enable,omitempty" toml:"enable,omitempty"`
	Exclusions *LintersExclusions `json:"exclusions,omitempty" yaml:"exclusions,omitempty" toml:"exclusions,omitempty"`
	// All available settings of specific linters.
	Settings *LintersSettings `json:"settings,omitempty" yaml:"settings,omitempty" toml:"settings,omitempty"`
}

type LintersExclusions struct {
	Generated   *Generated       `json:"generated,omitempty" yaml:"generated,omitempty" toml:"generated,omitempty"`
	Paths       []string         `json:"paths,omitempty" yaml:"paths,omitempty" toml:"paths,omitempty"`
	PathsExcept []string         `json:"paths-except,omitempty" yaml:"paths-except,omitempty" toml:"paths-except,omitempty"`
	Presets     []Preset         `json:"presets,omitempty" yaml:"presets,omitempty" toml:"presets,omitempty"`
	Rules       []ExclusionsRule `json:"rules,omitempty" yaml:"rules,omitempty" toml:"rules,omitempty"`
	WarnUnused  *bool            `json:"warn-unused,omitempty" yaml:"warn-unused,omitempty" toml:"warn-unused,omitempty"`
}

type ExclusionsRule struct {
	Linters    []string `json:"linters,omitempty" yaml:"linters,omitempty" toml:"linters,omitempty"`
	Path       *string  `json:"path,omitempty" yaml:"path,omitempty" toml:"path,omitempty"`
	PathExcept *string  `json:"path-except,omitempty" yaml:"path-except,omitempty" toml:"path-except,omitempty"`
	Source     *string  `json:"source,omitempty" yaml:"source,omitempty" toml:"source,omitempty"`
	Text       *string  `json:"text,omitempty" yaml:"text,omitempty" toml:"text,omitempty"`
}

// All available settings of specific linters.
type LintersSettings struct {
	Asasalint                *AsasalintSettings                `json:"asasalint,omitempty" yaml:"asasalint,omitempty" toml:"asasalint,omitempty"`
	Bidichk                  *BidichkSettings                  `json:"bidichk,omitempty" yaml:"bidichk,omitempty" toml:"bidichk,omitempty"`
	Copyloopvar              *CopyloopvarSettings              `json:"copyloopvar,omitempty" yaml:"copyloopvar,omitempty" toml:"copyloopvar,omitempty"`
	Custom                   map[string]interface{}            `json:"custom,omitempty" yaml:"custom,omitempty" toml:"custom,omitempty"`
	Cyclop                   *CyclopSettings                   `json:"cyclop,omitempty" yaml:"cyclop,omitempty" toml:"cyclop,omitempty"`
	Decorder                 *DecorderSettings                 `json:"decorder,omitempty" yaml:"decorder,omitempty" toml:"decorder,omitempty"`
	Depguard                 *DepguardSettings                 `json:"depguard,omitempty" yaml:"depguard,omitempty" toml:"depguard,omitempty"`
	Dogsled                  *DogsledSettings                  `json:"dogsled,omitempty" yaml:"dogsled,omitempty" toml:"dogsled,omitempty"`
	Dupl                     *DuplSettings                     `json:"dupl,omitempty" yaml:"dupl,omitempty" toml:"dupl,omitempty"`
	Dupword                  *DupwordSettings                  `json:"dupword,omitempty" yaml:"dupword,omitempty" toml:"dupword,omitempty"`
	Embeddedstructfieldcheck *EmbeddedstructfieldcheckSettings `json:"embeddedstructfieldcheck,omitempty" yaml:"embeddedstructfieldcheck,omitempty" toml:"embeddedstructfieldcheck,omitempty"`
	Errcheck                 *ErrcheckSettings                 `json:"errcheck,omitempty" yaml:"errcheck,omitempty" toml:"errcheck,omitempty"`
	Errchkjson               *ErrchkjsonSettings               `json:"errchkjson,omitempty" yaml:"errchkjson,omitempty" toml:"errchkjson,omitempty"`
	Errorlint                *ErrorlintSettings                `json:"errorlint,omitempty" yaml:"errorlint,omitempty" toml:"errorlint,omitempty"`
	Exhaustive               *ExhaustiveSettings               `json:"exhaustive,omitempty" yaml:"exhaustive,omitempty" toml:"exhaustive,omitempty"`
	Exhaustruct              *ExhaustructSettings              `json:"exhaustruct,omitempty" yaml:"exhaustruct,omitempty" toml:"exhaustruct,omitempty"`
	Fatcontext               *FatcontextSettings               `json:"fatcontext,omitempty" yaml:"fatcontext,omitempty" toml:"fatcontext,omitempty"`
	Forbidigo                *ForbidigoSettings                `json:"forbidigo,omitempty" yaml:"forbidigo,omitempty" toml:"forbidigo,omitempty"`
	Funcorder                *FuncorderSettings                `json:"funcorder,omitempty" yaml:"funcorder,omitempty" toml:"funcorder,omitempty"`
	Funlen                   *FunlenSettings                   `json:"funlen,omitempty" yaml:"funlen,omitempty" toml:"funlen,omitempty"`
	Ginkgolinter             *GinkgolinterSettings             `json:"ginkgolinter,omitempty" yaml:"ginkgolinter,omitempty" toml:"ginkgolinter,omitempty"`
	Gochecksumtype           *GochecksumtypeSettings           `json:"gochecksumtype,omitempty" yaml:"gochecksumtype,omitempty" toml:"gochecksumtype,omitempty"`
	Gocognit                 *GocognitSettings                 `json:"gocognit,omitempty" yaml:"gocognit,omitempty" toml:"gocognit,omitempty"`
	Goconst                  *GoconstSettings                  `json:"goconst,omitempty" yaml:"goconst,omitempty" toml:"goconst,omitempty"`
	Gocritic                 *GocriticClass                    `json:"gocritic,omitempty" yaml:"gocritic,omitempty" toml:"gocritic,omitempty"`
	Gocyclo                  *GocycloSettings                  `json:"gocyclo,omitempty" yaml:"gocyclo,omitempty" toml:"gocyclo,omitempty"`
	Godot                    *GodotSettings                    `json:"godot,omitempty" yaml:"godot,omitempty" toml:"godot,omitempty"`
	Godox                    *GodoxSettings                    `json:"godox,omitempty" yaml:"godox,omitempty" toml:"godox,omitempty"`
	Goheader                 *GoheaderSettings                 `json:"goheader,omitempty" yaml:"goheader,omitempty" toml:"goheader,omitempty"`
	Gomoddirectives          *GomoddirectivesSettings          `json:"gomoddirectives,omitempty" yaml:"gomoddirectives,omitempty" toml:"gomoddirectives,omitempty"`
	Gomodguard               *GomodguardSettings               `json:"gomodguard,omitempty" yaml:"gomodguard,omitempty" toml:"gomodguard,omitempty"`
	Gosec                    *GosecSettings                    `json:"gosec,omitempty" yaml:"gosec,omitempty" toml:"gosec,omitempty"`
	Gosmopolitan             *GosmopolitanSettings             `json:"gosmopolitan,omitempty" yaml:"gosmopolitan,omitempty" toml:"gosmopolitan,omitempty"`
	Govet                    *GovetSettings                    `json:"govet,omitempty" yaml:"govet,omitempty" toml:"govet,omitempty"`
	Grouper                  *GrouperSettings                  `json:"grouper,omitempty" yaml:"grouper,omitempty" toml:"grouper,omitempty"`
	Iface                    *IfaceClass                       `json:"iface,omitempty" yaml:"iface,omitempty" toml:"iface,omitempty"`
	Importas                 *ImportasSettings                 `json:"importas,omitempty" yaml:"importas,omitempty" toml:"importas,omitempty"`
	Inamedparam              *InamedparamSettings              `json:"inamedparam,omitempty" yaml:"inamedparam,omitempty" toml:"inamedparam,omitempty"`
	Interfacebloat           *InterfacebloatSettings           `json:"interfacebloat,omitempty" yaml:"interfacebloat,omitempty" toml:"interfacebloat,omitempty"`
	Ireturn                  *IreturnSettings                  `json:"ireturn,omitempty" yaml:"ireturn,omitempty" toml:"ireturn,omitempty"`
	Lll                      *LllSettings                      `json:"lll,omitempty" yaml:"lll,omitempty" toml:"lll,omitempty"`
	Loggercheck              *LoggercheckSettings              `json:"loggercheck,omitempty" yaml:"loggercheck,omitempty" toml:"loggercheck,omitempty"`
	Maintidx                 *MaintidxSettings                 `json:"maintidx,omitempty" yaml:"maintidx,omitempty" toml:"maintidx,omitempty"`
	Makezero                 *MakezeroSettings                 `json:"makezero,omitempty" yaml:"makezero,omitempty" toml:"makezero,omitempty"`
	Misspell                 *MisspellSettings                 `json:"misspell,omitempty" yaml:"misspell,omitempty" toml:"misspell,omitempty"`
	Mnd                      *MndSettings                      `json:"mnd,omitempty" yaml:"mnd,omitempty" toml:"mnd,omitempty"`
	Musttag                  *MusttagSettings                  `json:"musttag,omitempty" yaml:"musttag,omitempty" toml:"musttag,omitempty"`
	Nakedret                 *NakedretSettings                 `json:"nakedret,omitempty" yaml:"nakedret,omitempty" toml:"nakedret,omitempty"`
	Nestif                   *NestifSettings                   `json:"nestif,omitempty" yaml:"nestif,omitempty" toml:"nestif,omitempty"`
	Nilnil                   *NilnilSettings                   `json:"nilnil,omitempty" yaml:"nilnil,omitempty" toml:"nilnil,omitempty"`
	Nlreturn                 *NlreturnSettings                 `json:"nlreturn,omitempty" yaml:"nlreturn,omitempty" toml:"nlreturn,omitempty"`
	Nolintlint               *NolintlintSettings               `json:"nolintlint,omitempty" yaml:"nolintlint,omitempty" toml:"nolintlint,omitempty"`
	Nonamedreturns           *NonamedreturnsSettings           `json:"nonamedreturns,omitempty" yaml:"nonamedreturns,omitempty" toml:"nonamedreturns,omitempty"`
	Paralleltest             *ParalleltestSettings             `json:"paralleltest,omitempty" yaml:"paralleltest,omitempty" toml:"paralleltest,omitempty"`
	Perfsprint               *PerfsprintSettings               `json:"perfsprint,omitempty" yaml:"perfsprint,omitempty" toml:"perfsprint,omitempty"`
	Prealloc                 *PreallocSettings                 `json:"prealloc,omitempty" yaml:"prealloc,omitempty" toml:"prealloc,omitempty"`
	Predeclared              *PredeclaredSettings              `json:"predeclared,omitempty" yaml:"predeclared,omitempty" toml:"predeclared,omitempty"`
	Promlinter               *PromlinterSettings               `json:"promlinter,omitempty" yaml:"promlinter,omitempty" toml:"promlinter,omitempty"`
	Protogetter              *ProtogetterSettings              `json:"protogetter,omitempty" yaml:"protogetter,omitempty" toml:"protogetter,omitempty"`
	Reassign                 *ReassignSettings                 `json:"reassign,omitempty" yaml:"reassign,omitempty" toml:"reassign,omitempty"`
	Recvcheck                *RecvcheckSettings                `json:"recvcheck,omitempty" yaml:"recvcheck,omitempty" toml:"recvcheck,omitempty"`
	Revive                   *ReviveSettings                   `json:"revive,omitempty" yaml:"revive,omitempty" toml:"revive,omitempty"`
	Rowserrcheck             *RowserrcheckSettings             `json:"rowserrcheck,omitempty" yaml:"rowserrcheck,omitempty" toml:"rowserrcheck,omitempty"`
	Sloglint                 *SloglintSettings                 `json:"sloglint,omitempty" yaml:"sloglint,omitempty" toml:"sloglint,omitempty"`
	Spancheck                *SpancheckSettings                `json:"spancheck,omitempty" yaml:"spancheck,omitempty" toml:"spancheck,omitempty"`
	Staticcheck              *StaticcheckSettings              `json:"staticcheck,omitempty" yaml:"staticcheck,omitempty" toml:"staticcheck,omitempty"`
	Tagalign                 *TagalignSettings                 `json:"tagalign,omitempty" yaml:"tagalign,omitempty" toml:"tagalign,omitempty"`
	Tagliatelle              *TagliatelleSettings              `json:"tagliatelle,omitempty" yaml:"tagliatelle,omitempty" toml:"tagliatelle,omitempty"`
	Testifylint              *TestifylintSettings              `json:"testifylint,omitempty" yaml:"testifylint,omitempty" toml:"testifylint,omitempty"`
	Testpackage              *TestpackageSettings              `json:"testpackage,omitempty" yaml:"testpackage,omitempty" toml:"testpackage,omitempty"`
	Thelper                  *ThelperSettings                  `json:"thelper,omitempty" yaml:"thelper,omitempty" toml:"thelper,omitempty"`
	Unconvert                *UnconvertSettings                `json:"unconvert,omitempty" yaml:"unconvert,omitempty" toml:"unconvert,omitempty"`
	Unparam                  *UnparamSettings                  `json:"unparam,omitempty" yaml:"unparam,omitempty" toml:"unparam,omitempty"`
	Unused                   *UnusedSettings                   `json:"unused,omitempty" yaml:"unused,omitempty" toml:"unused,omitempty"`
	Usestdlibvars            *UsestdlibvarsSettings            `json:"usestdlibvars,omitempty" yaml:"usestdlibvars,omitempty" toml:"usestdlibvars,omitempty"`
	Usetesting               *UsetestingSettings               `json:"usetesting,omitempty" yaml:"usetesting,omitempty" toml:"usetesting,omitempty"`
	Varnamelen               *VarnamelenSettings               `json:"varnamelen,omitempty" yaml:"varnamelen,omitempty" toml:"varnamelen,omitempty"`
	Whitespace               *WhitespaceSettings               `json:"whitespace,omitempty" yaml:"whitespace,omitempty" toml:"whitespace,omitempty"`
	Wrapcheck                *WrapcheckSettings                `json:"wrapcheck,omitempty" yaml:"wrapcheck,omitempty" toml:"wrapcheck,omitempty"`
	Wsl                      *WslSettings                      `json:"wsl,omitempty" yaml:"wsl,omitempty" toml:"wsl,omitempty"`
	WslV5                    *WslSettingsV5                    `json:"wsl_v5,omitempty" yaml:"wsl_v5,omitempty" toml:"wsl_v5,omitempty"`
}

type AsasalintSettings struct {
	// To specify a set of function names to exclude.
	Exclude []string `json:"exclude,omitempty" yaml:"exclude,omitempty" toml:"exclude,omitempty"`
	// To enable/disable the asasalint builtin exclusions of function names.
	UseBuiltinExclusions *bool `json:"use-builtin-exclusions,omitempty" yaml:"use-builtin-exclusions,omitempty" toml:"use-builtin-exclusions,omitempty"`
}

type BidichkSettings struct {
	// Disallow: FIRST-STRONG-ISOLATE
	FirstStrongIsolate *bool `json:"first-strong-isolate,omitempty" yaml:"first-strong-isolate,omitempty" toml:"first-strong-isolate,omitempty"`
	// Disallow: LEFT-TO-RIGHT-EMBEDDING
	LeftToRightEmbedding *bool `json:"left-to-right-embedding,omitempty" yaml:"left-to-right-embedding,omitempty" toml:"left-to-right-embedding,omitempty"`
	// Disallow: LEFT-TO-RIGHT-ISOLATE
	LeftToRightIsolate *bool `json:"left-to-right-isolate,omitempty" yaml:"left-to-right-isolate,omitempty" toml:"left-to-right-isolate,omitempty"`
	// Disallow: LEFT-TO-RIGHT-OVERRIDE
	LeftToRightOverride *bool `json:"left-to-right-override,omitempty" yaml:"left-to-right-override,omitempty" toml:"left-to-right-override,omitempty"`
	// Disallow: POP-DIRECTIONAL-FORMATTING
	PopDirectionalFormatting *bool `json:"pop-directional-formatting,omitempty" yaml:"pop-directional-formatting,omitempty" toml:"pop-directional-formatting,omitempty"`
	// Disallow: POP-DIRECTIONAL-ISOLATE
	PopDirectionalIsolate *bool `json:"pop-directional-isolate,omitempty" yaml:"pop-directional-isolate,omitempty" toml:"pop-directional-isolate,omitempty"`
	// Disallow: RIGHT-TO-LEFT-EMBEDDING
	RightToLeftEmbedding *bool `json:"right-to-left-embedding,omitempty" yaml:"right-to-left-embedding,omitempty" toml:"right-to-left-embedding,omitempty"`
	// Disallow: RIGHT-TO-LEFT-ISOLATE
	RightToLeftIsolate *bool `json:"right-to-left-isolate,omitempty" yaml:"right-to-left-isolate,omitempty" toml:"right-to-left-isolate,omitempty"`
	// Disallow: RIGHT-TO-LEFT-OVERRIDE
	RightToLeftOverride *bool `json:"right-to-left-override,omitempty" yaml:"right-to-left-override,omitempty" toml:"right-to-left-override,omitempty"`
}

type CopyloopvarSettings struct {
	CheckAlias *bool `json:"check-alias,omitempty" yaml:"check-alias,omitempty" toml:"check-alias,omitempty"`
}

type CyclopSettings struct {
	// Max complexity the function can have
	MaxComplexity *int64 `json:"max-complexity,omitempty" yaml:"max-complexity,omitempty" toml:"max-complexity,omitempty"`
	// Max average complexity in package
	PackageAverage *float64 `json:"package-average,omitempty" yaml:"package-average,omitempty" toml:"package-average,omitempty"`
}

type DecorderSettings struct {
	DECOrder []DECOrder `json:"dec-order,omitempty" yaml:"dec-order,omitempty" toml:"dec-order,omitempty"`
	// Const declarations will be ignored for dec num check
	DisableConstDECNumCheck *bool `json:"disable-const-dec-num-check,omitempty" yaml:"disable-const-dec-num-check,omitempty" toml:"disable-const-dec-num-check,omitempty"`
	// Multiple global type, const and var declarations are allowed
	DisableDECNumCheck *bool `json:"disable-dec-num-check,omitempty" yaml:"disable-dec-num-check,omitempty" toml:"disable-dec-num-check,omitempty"`
	// Order of declarations is not checked
	DisableDECOrderCheck *bool `json:"disable-dec-order-check,omitempty" yaml:"disable-dec-order-check,omitempty" toml:"disable-dec-order-check,omitempty"`
	// Allow init func to be anywhere in file
	DisableInitFuncFirstCheck *bool `json:"disable-init-func-first-check,omitempty" yaml:"disable-init-func-first-check,omitempty" toml:"disable-init-func-first-check,omitempty"`
	// Type declarations will be ignored for dec num check
	DisableTypeDECNumCheck *bool `json:"disable-type-dec-num-check,omitempty" yaml:"disable-type-dec-num-check,omitempty" toml:"disable-type-dec-num-check,omitempty"`
	// Var declarations will be ignored for dec num check
	DisableVarDECNumCheck *bool `json:"disable-var-dec-num-check,omitempty" yaml:"disable-var-dec-num-check,omitempty" toml:"disable-var-dec-num-check,omitempty"`
	// Underscore vars (vars with "_" as the name) will be ignored at all checks
	IgnoreUnderscoreVars *bool `json:"ignore-underscore-vars,omitempty" yaml:"ignore-underscore-vars,omitempty" toml:"ignore-underscore-vars,omitempty"`
}

type DepguardSettings struct {
	// Rules to apply.
	Rules *Rules `json:"rules,omitempty" yaml:"rules,omitempty" toml:"rules,omitempty"`
}

// Rules to apply.
type Rules struct {
}

type DogsledSettings struct {
	// Check assignments with too many blank identifiers.
	MaxBlankIdentifiers *int64 `json:"max-blank-identifiers,omitempty" yaml:"max-blank-identifiers,omitempty" toml:"max-blank-identifiers,omitempty"`
}

type DuplSettings struct {
	// Tokens count to trigger issue.
	Threshold *int64 `json:"threshold,omitempty" yaml:"threshold,omitempty" toml:"threshold,omitempty"`
}

type DupwordSettings struct {
	// Keywords used to ignore detection.
	Ignore []string `json:"ignore,omitempty" yaml:"ignore,omitempty" toml:"ignore,omitempty"`
	// Keywords for detecting duplicate words. If this list is not empty, only the words defined
	// in this list will be detected.
	Keywords []string `json:"keywords,omitempty" yaml:"keywords,omitempty" toml:"keywords,omitempty"`
}

type EmbeddedstructfieldcheckSettings struct {
	// Checks that sync.Mutex and sync.RWMutex are not used as embedded fields.
	ForbidMutex *bool `json:"forbid-mutex,omitempty" yaml:"forbid-mutex,omitempty" toml:"forbid-mutex,omitempty"`
}

type ErrcheckSettings struct {
	// Report about assignment of errors to blank identifier
	CheckBlank *bool `json:"check-blank,omitempty" yaml:"check-blank,omitempty" toml:"check-blank,omitempty"`
	// Report about not checking errors in type assertions, i.e.: `a := b.(MyStruct)`
	CheckTypeAssertions *bool `json:"check-type-assertions,omitempty" yaml:"check-type-assertions,omitempty" toml:"check-type-assertions,omitempty"`
	// To disable the errcheck built-in exclude list
	DisableDefaultExclusions *bool `json:"disable-default-exclusions,omitempty" yaml:"disable-default-exclusions,omitempty" toml:"disable-default-exclusions,omitempty"`
	// List of functions to exclude from checking, where each entry is a single function to
	// exclude
	ExcludeFunctions []string `json:"exclude-functions,omitempty" yaml:"exclude-functions,omitempty" toml:"exclude-functions,omitempty"`
	// Display function signature instead of selector
	Verbose *bool `json:"verbose,omitempty" yaml:"verbose,omitempty" toml:"verbose,omitempty"`
}

type ErrchkjsonSettings struct {
	CheckErrorFreeEncoding *bool `json:"check-error-free-encoding,omitempty" yaml:"check-error-free-encoding,omitempty" toml:"check-error-free-encoding,omitempty"`
	// Issue on struct that doesn't have exported fields.
	ReportNoExported *bool `json:"report-no-exported,omitempty" yaml:"report-no-exported,omitempty" toml:"report-no-exported,omitempty"`
}

type ErrorlintSettings struct {
	AllowedErrors         []AllowedError          `json:"allowed-errors,omitempty" yaml:"allowed-errors,omitempty" toml:"allowed-errors,omitempty"`
	AllowedErrorsWildcard []AllowedErrorsWildcard `json:"allowed-errors-wildcard,omitempty" yaml:"allowed-errors-wildcard,omitempty" toml:"allowed-errors-wildcard,omitempty"`
	// Check for plain type assertions and type switches.
	Asserts *bool `json:"asserts,omitempty" yaml:"asserts,omitempty" toml:"asserts,omitempty"`
	// Check for plain error comparisons
	Comparison *bool `json:"comparison,omitempty" yaml:"comparison,omitempty" toml:"comparison,omitempty"`
	// Check whether fmt.Errorf uses the %w verb for formatting errors
	Errorf *bool `json:"errorf,omitempty" yaml:"errorf,omitempty" toml:"errorf,omitempty"`
	// Permit more than 1 %w verb, valid per Go 1.20
	ErrorfMulti *bool `json:"errorf-multi,omitempty" yaml:"errorf-multi,omitempty" toml:"errorf-multi,omitempty"`
}

type AllowedError struct {
	Err *string `json:"err,omitempty" yaml:"err,omitempty" toml:"err,omitempty"`
	Fun *string `json:"fun,omitempty" yaml:"fun,omitempty" toml:"fun,omitempty"`
}

type AllowedErrorsWildcard struct {
	Err *string `json:"err,omitempty" yaml:"err,omitempty" toml:"err,omitempty"`
	Fun *string `json:"fun,omitempty" yaml:"fun,omitempty" toml:"fun,omitempty"`
}

type ExhaustiveSettings struct {
	// Program elements to check for exhaustiveness.
	Check []string `json:"check,omitempty" yaml:"check,omitempty" toml:"check,omitempty"`
	// Switch statement requires default case even if exhaustive.
	DefaultCaseRequired *bool `json:"default-case-required,omitempty" yaml:"default-case-required,omitempty" toml:"default-case-required,omitempty"`
	// Presence of `default` case in switch statements satisfies exhaustiveness, even if all
	// enum members are not listed.
	DefaultSignifiesExhaustive *bool `json:"default-signifies-exhaustive,omitempty" yaml:"default-signifies-exhaustive,omitempty" toml:"default-signifies-exhaustive,omitempty"`
	// Only run exhaustive check on map literals with "//exhaustive:enforce" comment.
	ExplicitExhaustiveMap *bool `json:"explicit-exhaustive-map,omitempty" yaml:"explicit-exhaustive-map,omitempty" toml:"explicit-exhaustive-map,omitempty"`
	// Only run exhaustive check on switches with "//exhaustive:enforce" comment.
	ExplicitExhaustiveSwitch *bool `json:"explicit-exhaustive-switch,omitempty" yaml:"explicit-exhaustive-switch,omitempty" toml:"explicit-exhaustive-switch,omitempty"`
	// Enum members matching `regex` do not have to be listed in switch statements to satisfy
	// exhaustiveness
	IgnoreEnumMembers *string `json:"ignore-enum-members,omitempty" yaml:"ignore-enum-members,omitempty" toml:"ignore-enum-members,omitempty"`
	// Enum types matching the supplied regex do not have to be listed in switch statements to
	// satisfy exhaustiveness.
	IgnoreEnumTypes *string `json:"ignore-enum-types,omitempty" yaml:"ignore-enum-types,omitempty" toml:"ignore-enum-types,omitempty"`
	// Consider enums only in package scopes, not in inner scopes.
	PackageScopeOnly *bool `json:"package-scope-only,omitempty" yaml:"package-scope-only,omitempty" toml:"package-scope-only,omitempty"`
}

type ExhaustructSettings struct {
	// List of regular expressions to exclude struct packages and names from check.
	Exclude []string `json:"exclude,omitempty" yaml:"exclude,omitempty" toml:"exclude,omitempty"`
	// List of regular expressions to match struct packages and names.
	Include []string `json:"include,omitempty" yaml:"include,omitempty" toml:"include,omitempty"`
}

type FatcontextSettings struct {
	// Check for potential fat contexts in struct pointers.
	CheckStructPointers *bool `json:"check-struct-pointers,omitempty" yaml:"check-struct-pointers,omitempty" toml:"check-struct-pointers,omitempty"`
}

type ForbidigoSettings struct {
	// Instead of matching the literal source code, use type information to replace expressions
	// with strings that contain the package name and (for methods and fields) the type name.
	AnalyzeTypes *bool `json:"analyze-types,omitempty" yaml:"analyze-types,omitempty" toml:"analyze-types,omitempty"`
	// Exclude code in godoc examples.
	ExcludeGodocExamples *bool `json:"exclude-godoc-examples,omitempty" yaml:"exclude-godoc-examples,omitempty" toml:"exclude-godoc-examples,omitempty"`
	// List of identifiers to forbid (written using `regexp`)
	Forbid []Forbid `json:"forbid,omitempty" yaml:"forbid,omitempty" toml:"forbid,omitempty"`
}

type Forbid struct {
	// Message
	Msg *string `json:"msg,omitempty" yaml:"msg,omitempty" toml:"msg,omitempty"`
	// Pattern
	Pattern *string `json:"pattern,omitempty" yaml:"pattern,omitempty" toml:"pattern,omitempty"`
	// Package
	Pkg *string `json:"pkg,omitempty" yaml:"pkg,omitempty" toml:"pkg,omitempty"`
}

type FuncorderSettings struct {
	// Checks if the constructors and/or structure methods are sorted alphabetically.
	Alphabetical *bool `json:"alphabetical,omitempty" yaml:"alphabetical,omitempty" toml:"alphabetical,omitempty"`
	// Checks that constructors are placed after the structure declaration.
	Constructor *bool `json:"constructor,omitempty" yaml:"constructor,omitempty" toml:"constructor,omitempty"`
	// Checks if the exported methods of a structure are placed before the non-exported ones.
	StructMethod *bool `json:"struct-method,omitempty" yaml:"struct-method,omitempty" toml:"struct-method,omitempty"`
}

type FunlenSettings struct {
	// Ignore comments when counting lines.
	IgnoreComments *bool `json:"ignore-comments,omitempty" yaml:"ignore-comments,omitempty" toml:"ignore-comments,omitempty"`
	// Limit lines number per function.
	Lines *int64 `json:"lines,omitempty" yaml:"lines,omitempty" toml:"lines,omitempty"`
	// Limit statements number per function.
	Statements *int64 `json:"statements,omitempty" yaml:"statements,omitempty" toml:"statements,omitempty"`
}

type GinkgolinterSettings struct {
	// Don't trigger warnings for HaveLen(0).
	AllowHavelenZero *bool `json:"allow-havelen-zero,omitempty" yaml:"allow-havelen-zero,omitempty" toml:"allow-havelen-zero,omitempty"`
	// Trigger warning for ginkgo focus containers like FDescribe, FContext, FWhen or FIt.
	ForbidFocusContainer *bool `json:"forbid-focus-container,omitempty" yaml:"forbid-focus-container,omitempty" toml:"forbid-focus-container,omitempty"`
	// Trigger a warning for variable assignments in ginkgo containers like `Describe`,
	// `Context` and `When`, instead of in `BeforeEach()`.
	ForbidSpecPollution *bool `json:"forbid-spec-pollution,omitempty" yaml:"forbid-spec-pollution,omitempty" toml:"forbid-spec-pollution,omitempty"`
	// Force using `Expect` with `To`, `ToNot` or `NotTo`
	ForceExpectTo *bool `json:"force-expect-to,omitempty" yaml:"force-expect-to,omitempty" toml:"force-expect-to,omitempty"`
	// Force using the Succeed matcher for error functions, and the HaveOccurred matcher for
	// non-function error values.
	ForceSucceed *bool `json:"force-succeed,omitempty" yaml:"force-succeed,omitempty" toml:"force-succeed,omitempty"`
	// Suppress the function all in async assertion warning.
	SuppressAsyncAssertion *bool `json:"suppress-async-assertion,omitempty" yaml:"suppress-async-assertion,omitempty" toml:"suppress-async-assertion,omitempty"`
	// Suppress the wrong comparison assertion warning.
	SuppressCompareAssertion *bool `json:"suppress-compare-assertion,omitempty" yaml:"suppress-compare-assertion,omitempty" toml:"suppress-compare-assertion,omitempty"`
	// Suppress the wrong error assertion warning.
	SuppressErrAssertion *bool `json:"suppress-err-assertion,omitempty" yaml:"suppress-err-assertion,omitempty" toml:"suppress-err-assertion,omitempty"`
	// Suppress the wrong length assertion warning.
	SuppressLenAssertion *bool `json:"suppress-len-assertion,omitempty" yaml:"suppress-len-assertion,omitempty" toml:"suppress-len-assertion,omitempty"`
	// Suppress the wrong nil assertion warning.
	SuppressNilAssertion *bool `json:"suppress-nil-assertion,omitempty" yaml:"suppress-nil-assertion,omitempty" toml:"suppress-nil-assertion,omitempty"`
	// Suppress warning for comparing values from different types, like int32 and uint32.
	SuppressTypeCompareAssertion *bool `json:"suppress-type-compare-assertion,omitempty" yaml:"suppress-type-compare-assertion,omitempty" toml:"suppress-type-compare-assertion,omitempty"`
	// Best effort validation of async intervals (timeout and polling).
	ValidateAsyncIntervals *bool `json:"validate-async-intervals,omitempty" yaml:"validate-async-intervals,omitempty" toml:"validate-async-intervals,omitempty"`
}

type GochecksumtypeSettings struct {
	// Presence of `default` case in switch statements satisfies exhaustiveness, if all members
	// are not listed.
	DefaultSignifiesExhaustive *bool `json:"default-signifies-exhaustive,omitempty" yaml:"default-signifies-exhaustive,omitempty" toml:"default-signifies-exhaustive,omitempty"`
	// Include shared interfaces in the exhaustiviness check.
	IncludeSharedInterfaces *bool `json:"include-shared-interfaces,omitempty" yaml:"include-shared-interfaces,omitempty" toml:"include-shared-interfaces,omitempty"`
}

type GocognitSettings struct {
	// Minimal code complexity to report (we recommend 10-20).
	MinComplexity *int64 `json:"min-complexity,omitempty" yaml:"min-complexity,omitempty" toml:"min-complexity,omitempty"`
}

type GoconstSettings struct {
	// Evaluates of constant expressions like Prefix + "suffix"
	EvalConstExpressions *bool `json:"eval-const-expressions,omitempty" yaml:"eval-const-expressions,omitempty" toml:"eval-const-expressions,omitempty"`
	// Detects constants with identical values
	FindDuplicates *bool `json:"find-duplicates,omitempty" yaml:"find-duplicates,omitempty" toml:"find-duplicates,omitempty"`
	// Ignore when constant is not used as function argument
	IgnoreCalls *bool `json:"ignore-calls,omitempty" yaml:"ignore-calls,omitempty" toml:"ignore-calls,omitempty"`
	// Exclude strings matching the given regular expression
	IgnoreStringValues []string `json:"ignore-string-values,omitempty" yaml:"ignore-string-values,omitempty" toml:"ignore-string-values,omitempty"`
	// Look for existing constants matching the values
	MatchConstant *bool `json:"match-constant,omitempty" yaml:"match-constant,omitempty" toml:"match-constant,omitempty"`
	// Maximum value, only works with `numbers`
	Max *int64 `json:"max,omitempty" yaml:"max,omitempty" toml:"max,omitempty"`
	// Minimum value, only works with `numbers`
	Min *int64 `json:"min,omitempty" yaml:"min,omitempty" toml:"min,omitempty"`
	// Minimum length of string constant.
	MinLen *int64 `json:"min-len,omitempty" yaml:"min-len,omitempty" toml:"min-len,omitempty"`
	// Minimum occurrences count to trigger.
	MinOccurrences *int64 `json:"min-occurrences,omitempty" yaml:"min-occurrences,omitempty" toml:"min-occurrences,omitempty"`
	// Search also for duplicated numbers.
	Numbers *bool `json:"numbers,omitempty" yaml:"numbers,omitempty" toml:"numbers,omitempty"`
}

type GocriticClass struct {
	DisableAll *bool `json:"disable-all,omitempty" yaml:"disable-all,omitempty" toml:"disable-all,omitempty"`
	// Which checks should be disabled.
	DisabledChecks []GocriticChecks `json:"disabled-checks,omitempty" yaml:"disabled-checks,omitempty" toml:"disabled-checks,omitempty"`
	// Disable multiple checks by tags, run `GL_DEBUG=gocritic golangci-lint run` to see all
	// tags and checks.
	DisabledTags []GocriticTags `json:"disabled-tags,omitempty" yaml:"disabled-tags,omitempty" toml:"disabled-tags,omitempty"`
	EnableAll    *bool          `json:"enable-all,omitempty" yaml:"enable-all,omitempty" toml:"enable-all,omitempty"`
	// Which checks should be enabled. By default, a list of stable checks is used. To see it,
	// run `GL_DEBUG=gocritic golangci-lint run`.
	EnabledChecks []GocriticChecks `json:"enabled-checks,omitempty" yaml:"enabled-checks,omitempty" toml:"enabled-checks,omitempty"`
	// Enable multiple checks by tags, run `GL_DEBUG=gocritic golangci-lint run` to see all tags
	// and checks.
	EnabledTags []GocriticTags `json:"enabled-tags,omitempty" yaml:"enabled-tags,omitempty" toml:"enabled-tags,omitempty"`
	// Settings passed to gocritic. Properties must be valid and enabled check names.
	Settings *GocriticSettings `json:"settings,omitempty" yaml:"settings,omitempty" toml:"settings,omitempty"`
}

// Settings passed to gocritic. Properties must be valid and enabled check names.
type GocriticSettings struct {
	CaptLocal             *CaptLocalClass             `json:"captLocal,omitempty" yaml:"captLocal,omitempty" toml:"captLocal,omitempty"`
	CommentedOutCode      *CommentedOutCodeClass      `json:"commentedOutCode,omitempty" yaml:"commentedOutCode,omitempty" toml:"commentedOutCode,omitempty"`
	Elseif                *ElseifClass                `json:"elseif,omitempty" yaml:"elseif,omitempty" toml:"elseif,omitempty"`
	HugeParam             *HugeParamClass             `json:"hugeParam,omitempty" yaml:"hugeParam,omitempty" toml:"hugeParam,omitempty"`
	IfElseChain           *IfElseChainClass           `json:"ifElseChain,omitempty" yaml:"ifElseChain,omitempty" toml:"ifElseChain,omitempty"`
	NestingReduce         *NestingReduceClass         `json:"nestingReduce,omitempty" yaml:"nestingReduce,omitempty" toml:"nestingReduce,omitempty"`
	RangeExprCopy         *RangeExprCopyClass         `json:"rangeExprCopy,omitempty" yaml:"rangeExprCopy,omitempty" toml:"rangeExprCopy,omitempty"`
	RangeValCopy          *RangeValCopyClass          `json:"rangeValCopy,omitempty" yaml:"rangeValCopy,omitempty" toml:"rangeValCopy,omitempty"`
	Ruleguard             *RuleguardClass             `json:"ruleguard,omitempty" yaml:"ruleguard,omitempty" toml:"ruleguard,omitempty"`
	TooManyResultsChecker *TooManyResultsCheckerClass `json:"tooManyResultsChecker,omitempty" yaml:"tooManyResultsChecker,omitempty" toml:"tooManyResultsChecker,omitempty"`
	TruncateCmp           *TruncateCmpClass           `json:"truncateCmp,omitempty" yaml:"truncateCmp,omitempty" toml:"truncateCmp,omitempty"`
	Underef               *UnderefClass               `json:"underef,omitempty" yaml:"underef,omitempty" toml:"underef,omitempty"`
	UnnamedResult         *UnnamedResultClass         `json:"unnamedResult,omitempty" yaml:"unnamedResult,omitempty" toml:"unnamedResult,omitempty"`
}

type CaptLocalClass struct {
	ParamsOnly *bool `json:"paramsOnly,omitempty" yaml:"paramsOnly,omitempty" toml:"paramsOnly,omitempty"`
}

type CommentedOutCodeClass struct {
	MinLength *float64 `json:"minLength,omitempty" yaml:"minLength,omitempty" toml:"minLength,omitempty"`
}

type ElseifClass struct {
	SkipBalanced *bool `json:"skipBalanced,omitempty" yaml:"skipBalanced,omitempty" toml:"skipBalanced,omitempty"`
}

type HugeParamClass struct {
	SizeThreshold *float64 `json:"sizeThreshold,omitempty" yaml:"sizeThreshold,omitempty" toml:"sizeThreshold,omitempty"`
}

type IfElseChainClass struct {
	MinThreshold *float64 `json:"minThreshold,omitempty" yaml:"minThreshold,omitempty" toml:"minThreshold,omitempty"`
}

type NestingReduceClass struct {
	BodyWidth *float64 `json:"bodyWidth,omitempty" yaml:"bodyWidth,omitempty" toml:"bodyWidth,omitempty"`
}

type RangeExprCopyClass struct {
	SizeThreshold *float64 `json:"sizeThreshold,omitempty" yaml:"sizeThreshold,omitempty" toml:"sizeThreshold,omitempty"`
	SkipTestFuncs *bool    `json:"skipTestFuncs,omitempty" yaml:"skipTestFuncs,omitempty" toml:"skipTestFuncs,omitempty"`
}

type RangeValCopyClass struct {
	SizeThreshold *float64 `json:"sizeThreshold,omitempty" yaml:"sizeThreshold,omitempty" toml:"sizeThreshold,omitempty"`
	SkipTestFuncs *bool    `json:"skipTestFuncs,omitempty" yaml:"skipTestFuncs,omitempty" toml:"skipTestFuncs,omitempty"`
}

type RuleguardClass struct {
	Debug   *string `json:"debug,omitempty" yaml:"debug,omitempty" toml:"debug,omitempty"`
	Disable *string `json:"disable,omitempty" yaml:"disable,omitempty" toml:"disable,omitempty"`
	Enable  *string `json:"enable,omitempty" yaml:"enable,omitempty" toml:"enable,omitempty"`
	FailOn  *string `json:"failOn,omitempty" yaml:"failOn,omitempty" toml:"failOn,omitempty"`
	Rules   *string `json:"rules,omitempty" yaml:"rules,omitempty" toml:"rules,omitempty"`
}

type TooManyResultsCheckerClass struct {
	MaxResults *float64 `json:"maxResults,omitempty" yaml:"maxResults,omitempty" toml:"maxResults,omitempty"`
}

type TruncateCmpClass struct {
	SkipArchDependent *bool `json:"skipArchDependent,omitempty" yaml:"skipArchDependent,omitempty" toml:"skipArchDependent,omitempty"`
}

type UnderefClass struct {
	SkipRecvDeref *bool `json:"skipRecvDeref,omitempty" yaml:"skipRecvDeref,omitempty" toml:"skipRecvDeref,omitempty"`
}

type UnnamedResultClass struct {
	CheckExported *bool `json:"checkExported,omitempty" yaml:"checkExported,omitempty" toml:"checkExported,omitempty"`
}

type GocycloSettings struct {
	// Minimum code complexity to report (we recommend 10-20).
	MinComplexity *int64 `json:"min-complexity,omitempty" yaml:"min-complexity,omitempty" toml:"min-complexity,omitempty"`
}

type GodotSettings struct {
	// Check that each sentence starts with a capital letter.
	Capital *bool `json:"capital,omitempty" yaml:"capital,omitempty" toml:"capital,omitempty"`
	// DEPRECATED: Check all top-level comments, not only declarations.
	CheckAll *bool `json:"check-all,omitempty" yaml:"check-all,omitempty" toml:"check-all,omitempty"`
	// List of regexps for excluding particular comment lines from check.
	Exclude []string `json:"exclude,omitempty" yaml:"exclude,omitempty" toml:"exclude,omitempty"`
	// Check that each sentence ends with a period.
	Period *bool `json:"period,omitempty" yaml:"period,omitempty" toml:"period,omitempty"`
	// Comments to be checked.
	Scope *ScopeEnum `json:"scope,omitempty" yaml:"scope,omitempty" toml:"scope,omitempty"`
}

type GodoxSettings struct {
	// Report any comments starting with one of these keywords. This is useful for TODO or FIXME
	// comments that might be left in the code accidentally and should be resolved before
	// merging.
	Keywords []string `json:"keywords,omitempty" yaml:"keywords,omitempty" toml:"keywords,omitempty"`
}

type GoheaderSettings struct {
	// Template to put on top of every file.
	Template *string `json:"template,omitempty" yaml:"template,omitempty" toml:"template,omitempty"`
	// Path to the file containing the template source.
	TemplatePath *string `json:"template-path,omitempty" yaml:"template-path,omitempty" toml:"template-path,omitempty"`
	Values       *Values `json:"values,omitempty" yaml:"values,omitempty" toml:"values,omitempty"`
}

type Values struct {
	// Constants to use in the template.
	Const *ConstClass `json:"const,omitempty" yaml:"const,omitempty" toml:"const,omitempty"`
	// Regular expressions to use in your template.
	Regexp *RegexpClass `json:"regexp,omitempty" yaml:"regexp,omitempty" toml:"regexp,omitempty"`
}

// Constants to use in the template.
type ConstClass struct {
}

// Regular expressions to use in your template.
type RegexpClass struct {
}

type GomoddirectivesSettings struct {
	// Forbid the use of the `exclude` directives.
	ExcludeForbidden *bool `json:"exclude-forbidden,omitempty" yaml:"exclude-forbidden,omitempty" toml:"exclude-forbidden,omitempty"`
	// Forbid the use of the `godebug` directive.
	GoDebugForbidden *bool `json:"go-debug-forbidden,omitempty" yaml:"go-debug-forbidden,omitempty" toml:"go-debug-forbidden,omitempty"`
	// Defines a pattern to validate `go` minimum version directive.
	GoVersionPattern *string `json:"go-version-pattern,omitempty" yaml:"go-version-pattern,omitempty" toml:"go-version-pattern,omitempty"`
	// Forbid the use of the `ignore` directives. (>= go1.25)
	IgnoreForbidden *bool `json:"ignore-forbidden,omitempty" yaml:"ignore-forbidden,omitempty" toml:"ignore-forbidden,omitempty"`
	// List of allowed `replace` directives.
	ReplaceAllowList []string `json:"replace-allow-list,omitempty" yaml:"replace-allow-list,omitempty" toml:"replace-allow-list,omitempty"`
	// Allow local `replace` directives.
	ReplaceLocal *bool `json:"replace-local,omitempty" yaml:"replace-local,omitempty" toml:"replace-local,omitempty"`
	// Allow to not explain why the version has been retracted in the `retract` directives.
	RetractAllowNoExplanation *bool `json:"retract-allow-no-explanation,omitempty" yaml:"retract-allow-no-explanation,omitempty" toml:"retract-allow-no-explanation,omitempty"`
	// Forbid the use of the `tool` directives.
	ToolForbidden *bool `json:"tool-forbidden,omitempty" yaml:"tool-forbidden,omitempty" toml:"tool-forbidden,omitempty"`
	// Forbid the use of the `toolchain` directive.
	ToolchainForbidden *bool `json:"toolchain-forbidden,omitempty" yaml:"toolchain-forbidden,omitempty" toml:"toolchain-forbidden,omitempty"`
	// Defines a pattern to validate `toolchain` directive.
	ToolchainPattern *string `json:"toolchain-pattern,omitempty" yaml:"toolchain-pattern,omitempty" toml:"toolchain-pattern,omitempty"`
}

type GomodguardSettings struct {
	Allowed *Allowed `json:"allowed,omitempty" yaml:"allowed,omitempty" toml:"allowed,omitempty"`
	Blocked *Blocked `json:"blocked,omitempty" yaml:"blocked,omitempty" toml:"blocked,omitempty"`
}

type Allowed struct {
	// List of allowed module domains.
	Domains []string `json:"domains,omitempty" yaml:"domains,omitempty" toml:"domains,omitempty"`
	// List of allowed modules.
	Modules []string `json:"modules,omitempty" yaml:"modules,omitempty" toml:"modules,omitempty"`
}

type Blocked struct {
	// Raise lint issues if loading local path with replace directive
	LocalReplaceDirectives *bool `json:"local-replace-directives,omitempty" yaml:"local-replace-directives,omitempty" toml:"local-replace-directives,omitempty"`
	// List of blocked modules.
	Modules []Module `json:"modules,omitempty" yaml:"modules,omitempty" toml:"modules,omitempty"`
	// List of blocked module version constraints.
	Versions []map[string]interface{} `json:"versions,omitempty" yaml:"versions,omitempty" toml:"versions,omitempty"`
}

type Module struct {
}

type GosecSettings struct {
	// Concurrency value
	Concurrency *int64 `json:"concurrency,omitempty" yaml:"concurrency,omitempty" toml:"concurrency,omitempty"`
	// Filter out the issues with a lower confidence than the given value
	Confidence *Confidence `json:"confidence,omitempty" yaml:"confidence,omitempty" toml:"confidence,omitempty"`
	// To specify the configuration of rules
	Config map[string]interface{} `json:"config,omitempty" yaml:"config,omitempty" toml:"config,omitempty"`
	// To specify a set of rules to explicitly exclude
	Excludes []GosecRules `json:"excludes,omitempty" yaml:"excludes,omitempty" toml:"excludes,omitempty"`
	// To select a subset of rules to run
	Includes []GosecRules `json:"includes,omitempty" yaml:"includes,omitempty" toml:"includes,omitempty"`
	// Filter out the issues with a lower severity than the given value
	Severity *Confidence `json:"severity,omitempty" yaml:"severity,omitempty" toml:"severity,omitempty"`
}

type GosmopolitanSettings struct {
	// Allow and ignore `time.Local` usages.
	AllowTimeLocal *bool `json:"allow-time-local,omitempty" yaml:"allow-time-local,omitempty" toml:"allow-time-local,omitempty"`
	// List of fully qualified names in the `full/pkg/path.name` form, to act as "i18n escape
	// hatches".
	EscapeHatches []string `json:"escape-hatches,omitempty" yaml:"escape-hatches,omitempty" toml:"escape-hatches,omitempty"`
	// List of Unicode scripts to watch for any usage in string literals.
	WatchForScripts []string `json:"watch-for-scripts,omitempty" yaml:"watch-for-scripts,omitempty" toml:"watch-for-scripts,omitempty"`
}

type GovetSettings struct {
	// Disable analyzers by name.
	Disable []GovetAnalyzers `json:"disable,omitempty" yaml:"disable,omitempty" toml:"disable,omitempty"`
	// Disable all analyzers.
	DisableAll *bool `json:"disable-all,omitempty" yaml:"disable-all,omitempty" toml:"disable-all,omitempty"`
	// Enable analyzers by name.
	Enable []GovetAnalyzers `json:"enable,omitempty" yaml:"enable,omitempty" toml:"enable,omitempty"`
	// Enable all analyzers.
	EnableAll *bool `json:"enable-all,omitempty" yaml:"enable-all,omitempty" toml:"enable-all,omitempty"`
	// Settings per analyzer. Map of analyzer name to specific settings.
	// Run `go tool vet help` to find out more.
	Settings map[string]interface{} `json:"settings,omitempty" yaml:"settings,omitempty" toml:"settings,omitempty"`
}

type GrouperSettings struct {
	ConstRequireGrouping      *bool `json:"const-require-grouping,omitempty" yaml:"const-require-grouping,omitempty" toml:"const-require-grouping,omitempty"`
	ConstRequireSingleConst   *bool `json:"const-require-single-const,omitempty" yaml:"const-require-single-const,omitempty" toml:"const-require-single-const,omitempty"`
	ImportRequireGrouping     *bool `json:"import-require-grouping,omitempty" yaml:"import-require-grouping,omitempty" toml:"import-require-grouping,omitempty"`
	ImportRequireSingleImport *bool `json:"import-require-single-import,omitempty" yaml:"import-require-single-import,omitempty" toml:"import-require-single-import,omitempty"`
	TypeRequireGrouping       *bool `json:"type-require-grouping,omitempty" yaml:"type-require-grouping,omitempty" toml:"type-require-grouping,omitempty"`
	TypeRequireSingleType     *bool `json:"type-require-single-type,omitempty" yaml:"type-require-single-type,omitempty" toml:"type-require-single-type,omitempty"`
	VarRequireGrouping        *bool `json:"var-require-grouping,omitempty" yaml:"var-require-grouping,omitempty" toml:"var-require-grouping,omitempty"`
	VarRequireSingleVar       *bool `json:"var-require-single-var,omitempty" yaml:"var-require-single-var,omitempty" toml:"var-require-single-var,omitempty"`
}

type IfaceClass struct {
	// Enable analyzers by name.
	Enable   []IfaceAnalyzers `json:"enable,omitempty" yaml:"enable,omitempty" toml:"enable,omitempty"`
	Settings *IfaceSettings   `json:"settings,omitempty" yaml:"settings,omitempty" toml:"settings,omitempty"`
}

type IfaceSettings struct {
	Unused *UnusedClass `json:"unused,omitempty" yaml:"unused,omitempty" toml:"unused,omitempty"`
}

type UnusedClass struct {
	Exclude []string `json:"exclude,omitempty" yaml:"exclude,omitempty" toml:"exclude,omitempty"`
}

type ImportasSettings struct {
	// List of aliases
	Alias []Alias `json:"alias,omitempty" yaml:"alias,omitempty" toml:"alias,omitempty"`
	// Do not allow non-required aliases.
	NoExtraAliases *bool `json:"no-extra-aliases,omitempty" yaml:"no-extra-aliases,omitempty" toml:"no-extra-aliases,omitempty"`
	// Do not allow unaliased imports of aliased packages.
	NoUnaliased *bool `json:"no-unaliased,omitempty" yaml:"no-unaliased,omitempty" toml:"no-unaliased,omitempty"`
}

type Alias struct {
	// Package alias e.g. autoscalingv1alpha1
	Alias string `json:"alias" yaml:"alias" toml:"alias"`
	// Package path e.g. knative.dev/serving/pkg/apis/autoscaling/v1alpha1
	Pkg string `json:"pkg" yaml:"pkg" toml:"pkg"`
}

type InamedparamSettings struct {
	// Skips check for interface methods with only a single parameter.
	SkipSingleParam *bool `json:"skip-single-param,omitempty" yaml:"skip-single-param,omitempty" toml:"skip-single-param,omitempty"`
}

type InterfacebloatSettings struct {
	// The maximum number of methods allowed for an interface.
	Max *int64 `json:"max,omitempty" yaml:"max,omitempty" toml:"max,omitempty"`
}

// Use either `reject` or `allow` properties for interfaces matching.
type IreturnSettings struct {
	Allow  []string `json:"allow,omitempty" yaml:"allow,omitempty" toml:"allow,omitempty"`
	Reject []string `json:"reject,omitempty" yaml:"reject,omitempty" toml:"reject,omitempty"`
}

type LllSettings struct {
	// Maximum allowed line length, lines longer will be reported.
	LineLength *int64 `json:"line-length,omitempty" yaml:"line-length,omitempty" toml:"line-length,omitempty"`
	// Width of "\t" in spaces.
	TabWidth *int64 `json:"tab-width,omitempty" yaml:"tab-width,omitempty" toml:"tab-width,omitempty"`
}

type LoggercheckSettings struct {
	// Allow check for the github.com/go-kit/log library.
	Kitlog *bool `json:"kitlog,omitempty" yaml:"kitlog,omitempty" toml:"kitlog,omitempty"`
	// Allow check for the k8s.io/klog/v2 library.
	Klog *bool `json:"klog,omitempty" yaml:"klog,omitempty" toml:"klog,omitempty"`
	// Allow check for the github.com/go-logr/logr library.
	Logr *bool `json:"logr,omitempty" yaml:"logr,omitempty" toml:"logr,omitempty"`
	// Require printf-like format specifier (%s, %d for example) not present.
	NoPrintfLike *bool `json:"no-printf-like,omitempty" yaml:"no-printf-like,omitempty" toml:"no-printf-like,omitempty"`
	// Require all logging keys to be inlined constant strings.
	RequireStringKey *bool `json:"require-string-key,omitempty" yaml:"require-string-key,omitempty" toml:"require-string-key,omitempty"`
	// List of custom rules to check against, where each rule is a single logger pattern, useful
	// for wrapped loggers.
	Rules []string `json:"rules,omitempty" yaml:"rules,omitempty" toml:"rules,omitempty"`
	// Allow check for the log/slog library.
	Slog *bool `json:"slog,omitempty" yaml:"slog,omitempty" toml:"slog,omitempty"`
	// Allow check for the "sugar logger" from go.uber.org/zap library.
	Zap *bool `json:"zap,omitempty" yaml:"zap,omitempty" toml:"zap,omitempty"`
}

// Maintainability index
// https://docs.microsoft.com/en-us/visualstudio/code-quality/code-metrics-maintainability-index-range-and-meaning?view=vs-2022
type MaintidxSettings struct {
	// Minimum accatpable maintainability index level (see
	// https://docs.microsoft.com/en-us/visualstudio/code-quality/code-metrics-maintainability-index-range-and-meaning?view=vs-2022)
	Under *float64 `json:"under,omitempty" yaml:"under,omitempty" toml:"under,omitempty"`
}

type MakezeroSettings struct {
	// Allow only slices initialized with a length of zero.
	Always *bool `json:"always,omitempty" yaml:"always,omitempty" toml:"always,omitempty"`
}

// Correct spellings using locale preferences for US or UK. Default is to use a neutral
// variety of English.
type MisspellSettings struct {
	// Extra word corrections.
	ExtraWords []ExtraWord `json:"extra-words,omitempty" yaml:"extra-words,omitempty" toml:"extra-words,omitempty"`
	// List of rules to ignore.
	IgnoreRules []string `json:"ignore-rules,omitempty" yaml:"ignore-rules,omitempty" toml:"ignore-rules,omitempty"`
	Locale      *Locale  `json:"locale,omitempty" yaml:"locale,omitempty" toml:"locale,omitempty"`
	// Mode of the analysis.
	Mode *MisspellMode `json:"mode,omitempty" yaml:"mode,omitempty" toml:"mode,omitempty"`
}

type ExtraWord struct {
	Correction *string `json:"correction,omitempty" yaml:"correction,omitempty" toml:"correction,omitempty"`
	Typo       *string `json:"typo,omitempty" yaml:"typo,omitempty" toml:"typo,omitempty"`
}

type MndSettings struct {
	// The list of enabled checks, see https://github.com/tommy-muehle/go-mnd/#checks for
	// description.
	Checks []MndCheck `json:"checks,omitempty" yaml:"checks,omitempty" toml:"checks,omitempty"`
	// List of file patterns to exclude from analysis.
	IgnoredFiles []string `json:"ignored-files,omitempty" yaml:"ignored-files,omitempty" toml:"ignored-files,omitempty"`
	// Comma-separated list of function patterns to exclude from the analysis.
	IgnoredFunctions []string `json:"ignored-functions,omitempty" yaml:"ignored-functions,omitempty" toml:"ignored-functions,omitempty"`
	// List of numbers to exclude from analysis.
	IgnoredNumbers []string `json:"ignored-numbers,omitempty" yaml:"ignored-numbers,omitempty" toml:"ignored-numbers,omitempty"`
}

type MusttagSettings struct {
	Functions []Function `json:"functions,omitempty" yaml:"functions,omitempty" toml:"functions,omitempty"`
}

type Function struct {
	ArgPos *int64  `json:"arg-pos,omitempty" yaml:"arg-pos,omitempty" toml:"arg-pos,omitempty"`
	Name   *string `json:"name,omitempty" yaml:"name,omitempty" toml:"name,omitempty"`
	Tag    *string `json:"tag,omitempty" yaml:"tag,omitempty" toml:"tag,omitempty"`
}

type NakedretSettings struct {
	// Report if a function has more lines of code than this value and it has naked returns.
	MaxFuncLines *int64 `json:"max-func-lines,omitempty" yaml:"max-func-lines,omitempty" toml:"max-func-lines,omitempty"`
}

type NestifSettings struct {
	// Minimum complexity of "if" statements to report.
	MinComplexity *int64 `json:"min-complexity,omitempty" yaml:"min-complexity,omitempty" toml:"min-complexity,omitempty"`
}

type NilnilSettings struct {
	// List of return types to check.
	CheckedTypes []CheckedType `json:"checked-types,omitempty" yaml:"checked-types,omitempty" toml:"checked-types,omitempty"`
	// In addition, detect opposite situation (simultaneous return of non-nil error and valid
	// value).
	DetectOpposite *bool `json:"detect-opposite,omitempty" yaml:"detect-opposite,omitempty" toml:"detect-opposite,omitempty"`
	// To check functions with only two return values.
	OnlyTwo *bool `json:"only-two,omitempty" yaml:"only-two,omitempty" toml:"only-two,omitempty"`
}

type NlreturnSettings struct {
	// set block size that is still ok
	BlockSize *float64 `json:"block-size,omitempty" yaml:"block-size,omitempty" toml:"block-size,omitempty"`
}

type NolintlintSettings struct {
	// Exclude these linters from requiring an explanation.
	AllowNoExplanation []string `json:"allow-no-explanation,omitempty" yaml:"allow-no-explanation,omitempty" toml:"allow-no-explanation,omitempty"`
	// Enable to ensure that nolint directives are all used.
	AllowUnused *bool `json:"allow-unused,omitempty" yaml:"allow-unused,omitempty" toml:"allow-unused,omitempty"`
	// Enable to require an explanation of nonzero length after each nolint directive.
	RequireExplanation *bool `json:"require-explanation,omitempty" yaml:"require-explanation,omitempty" toml:"require-explanation,omitempty"`
	// Enable to require nolint directives to mention the specific linter being suppressed.
	RequireSpecific *bool `json:"require-specific,omitempty" yaml:"require-specific,omitempty" toml:"require-specific,omitempty"`
}

type NonamedreturnsSettings struct {
	// Report named error if it is assigned inside defer.
	ReportErrorInDefer *bool `json:"report-error-in-defer,omitempty" yaml:"report-error-in-defer,omitempty" toml:"report-error-in-defer,omitempty"`
}

type ParalleltestSettings struct {
	// Ignore missing calls to `t.Parallel()` and only report incorrect uses of it.
	IgnoreMissing *bool `json:"ignore-missing,omitempty" yaml:"ignore-missing,omitempty" toml:"ignore-missing,omitempty"`
	// Ignore missing calls to `t.Parallel()` in subtests. Top-level tests are still required to
	// have `t.Parallel`, but subtests are allowed to skip it.
	IgnoreMissingSubtests *bool `json:"ignore-missing-subtests,omitempty" yaml:"ignore-missing-subtests,omitempty" toml:"ignore-missing-subtests,omitempty"`
}

type PerfsprintSettings struct {
	// Enable/disable optimization of bool formatting.
	BoolFormat *bool `json:"bool-format,omitempty" yaml:"bool-format,omitempty" toml:"bool-format,omitempty"`
	// Optimizes into `err.Error()` even if it is only equivalent for non-nil errors.
	ErrError *bool `json:"err-error,omitempty" yaml:"err-error,omitempty" toml:"err-error,omitempty"`
	// Enable/disable optimization of error formatting.
	ErrorFormat *bool `json:"error-format,omitempty" yaml:"error-format,omitempty" toml:"error-format,omitempty"`
	// Optimizes `fmt.Errorf`.
	Errorf *bool `json:"errorf,omitempty" yaml:"errorf,omitempty" toml:"errorf,omitempty"`
	// Enable/disable optimization of hex formatting.
	HexFormat *bool `json:"hex-format,omitempty" yaml:"hex-format,omitempty" toml:"hex-format,omitempty"`
	// Optimizes even if it requires an int or uint type cast.
	IntConversion *bool `json:"int-conversion,omitempty" yaml:"int-conversion,omitempty" toml:"int-conversion,omitempty"`
	// Enable/disable optimization of integer formatting.
	IntegerFormat *bool `json:"integer-format,omitempty" yaml:"integer-format,omitempty" toml:"integer-format,omitempty"`
	// Optimizes `fmt.Sprintf` with only one argument.
	Sprintf1 *bool `json:"sprintf1,omitempty" yaml:"sprintf1,omitempty" toml:"sprintf1,omitempty"`
	// Optimizes into strings concatenation.
	Strconcat *bool `json:"strconcat,omitempty" yaml:"strconcat,omitempty" toml:"strconcat,omitempty"`
	// Enable/disable optimization of string formatting.
	StringFormat *bool `json:"string-format,omitempty" yaml:"string-format,omitempty" toml:"string-format,omitempty"`
}

// We do not recommend using this linter before doing performance profiling.
// For most programs usage of `prealloc` will be premature optimization.
type PreallocSettings struct {
	// Report preallocation suggestions on for loops.
	ForLoops *bool `json:"for-loops,omitempty" yaml:"for-loops,omitempty" toml:"for-loops,omitempty"`
	// Report preallocation suggestions on range loops.
	RangeLoops *bool `json:"range-loops,omitempty" yaml:"range-loops,omitempty" toml:"range-loops,omitempty"`
	// Report preallocation suggestions only on simple loops that have no
	// returns/breaks/continues/gotos in them.
	Simple *bool `json:"simple,omitempty" yaml:"simple,omitempty" toml:"simple,omitempty"`
}

type PredeclaredSettings struct {
	// List of predeclared identifiers to not report on.
	Ignore []string `json:"ignore,omitempty" yaml:"ignore,omitempty" toml:"ignore,omitempty"`
	// Include method names and field names in checks.
	QualifiedName *bool `json:"qualified-name,omitempty" yaml:"qualified-name,omitempty" toml:"qualified-name,omitempty"`
}

type PromlinterSettings struct {
	DisabledLinters []DisabledLinter `json:"disabled-linters,omitempty" yaml:"disabled-linters,omitempty" toml:"disabled-linters,omitempty"`
	Strict          interface{}      `json:"strict" yaml:"strict" toml:"strict"`
}

type ProtogetterSettings struct {
	// Skip first argument of append function.
	ReplaceFirstArgInAppend *bool `json:"replace-first-arg-in-append,omitempty" yaml:"replace-first-arg-in-append,omitempty" toml:"replace-first-arg-in-append,omitempty"`
	// Skip any generated files from the checking.
	SkipAnyGenerated *bool    `json:"skip-any-generated,omitempty" yaml:"skip-any-generated,omitempty" toml:"skip-any-generated,omitempty"`
	SkipFiles        []string `json:"skip-files,omitempty" yaml:"skip-files,omitempty" toml:"skip-files,omitempty"`
	SkipGeneratedBy  []string `json:"skip-generated-by,omitempty" yaml:"skip-generated-by,omitempty" toml:"skip-generated-by,omitempty"`
}

type ReassignSettings struct {
	Patterns []string `json:"patterns,omitempty" yaml:"patterns,omitempty" toml:"patterns,omitempty"`
}

type RecvcheckSettings struct {
	// Disables the built-in method exclusions.
	DisableBuiltin *bool `json:"disable-builtin,omitempty" yaml:"disable-builtin,omitempty" toml:"disable-builtin,omitempty"`
	// User-defined method exclusions.
	Exclusions []string `json:"exclusions,omitempty" yaml:"exclusions,omitempty" toml:"exclusions,omitempty"`
}

type ReviveSettings struct {
	Confidence     *float64           `json:"confidence,omitempty" yaml:"confidence,omitempty" toml:"confidence,omitempty"`
	Directives     []DirectiveElement `json:"directives,omitempty" yaml:"directives,omitempty" toml:"directives,omitempty"`
	EnableAllRules *bool              `json:"enable-all-rules,omitempty" yaml:"enable-all-rules,omitempty" toml:"enable-all-rules,omitempty"`
	MaxOpenFiles   *int64             `json:"max-open-files,omitempty" yaml:"max-open-files,omitempty" toml:"max-open-files,omitempty"`
	Rules          []Rule             `json:"rules,omitempty" yaml:"rules,omitempty" toml:"rules,omitempty"`
	Severity       *SeverityEnum      `json:"severity,omitempty" yaml:"severity,omitempty" toml:"severity,omitempty"`
}

type DirectiveElement struct {
	Arguments []interface{} `json:"arguments,omitempty" yaml:"arguments,omitempty" toml:"arguments,omitempty"`
	Exclude   []string      `json:"exclude,omitempty" yaml:"exclude,omitempty" toml:"exclude,omitempty"`
	Name      *Name         `json:"name,omitempty" yaml:"name,omitempty" toml:"name,omitempty"`
	Severity  *SeverityEnum `json:"severity,omitempty" yaml:"severity,omitempty" toml:"severity,omitempty"`
}

type Rule struct {
	Arguments []interface{} `json:"arguments,omitempty" yaml:"arguments,omitempty" toml:"arguments,omitempty"`
	Disabled  *bool         `json:"disabled,omitempty" yaml:"disabled,omitempty" toml:"disabled,omitempty"`
	Exclude   []string      `json:"exclude,omitempty" yaml:"exclude,omitempty" toml:"exclude,omitempty"`
	Name      ReviveRules   `json:"name" yaml:"name" toml:"name"`
	Severity  *SeverityEnum `json:"severity,omitempty" yaml:"severity,omitempty" toml:"severity,omitempty"`
}

type RowserrcheckSettings struct {
	Packages []string `json:"packages,omitempty" yaml:"packages,omitempty" toml:"packages,omitempty"`
}

type SloglintSettings struct {
	// Enforce putting arguments on separate lines.
	ArgsOnSepLines *bool `json:"args-on-sep-lines,omitempty" yaml:"args-on-sep-lines,omitempty" toml:"args-on-sep-lines,omitempty"`
	// Enforce using attributes only (incompatible with kv-only).
	AttrOnly *bool `json:"attr-only,omitempty" yaml:"attr-only,omitempty" toml:"attr-only,omitempty"`
	// Enforce using methods that accept a context.
	Context *ContextEnum `json:"context,omitempty" yaml:"context,omitempty" toml:"context,omitempty"`
	// Enforce not using specific keys.
	ForbiddenKeys []string `json:"forbidden-keys,omitempty" yaml:"forbidden-keys,omitempty" toml:"forbidden-keys,omitempty"`
	// Enforce a single key naming convention.
	KeyNamingCase *KeyNamingCase `json:"key-naming-case,omitempty" yaml:"key-naming-case,omitempty" toml:"key-naming-case,omitempty"`
	// Enforce using key-value pairs only (incompatible with attr-only).
	KvOnly *bool `json:"kv-only,omitempty" yaml:"kv-only,omitempty" toml:"kv-only,omitempty"`
	// Enforce message style.
	MsgStyle *MsgStyleEnum `json:"msg-style,omitempty" yaml:"msg-style,omitempty" toml:"msg-style,omitempty"`
	// Enforce not using global loggers.
	NoGlobal *NoGlobalEnum `json:"no-global,omitempty" yaml:"no-global,omitempty" toml:"no-global,omitempty"`
	// Enforce not mixing key-value pairs and attributes.
	NoMixedArgs *bool `json:"no-mixed-args,omitempty" yaml:"no-mixed-args,omitempty" toml:"no-mixed-args,omitempty"`
	// Enforce using constants instead of raw keys.
	NoRawKeys *bool `json:"no-raw-keys,omitempty" yaml:"no-raw-keys,omitempty" toml:"no-raw-keys,omitempty"`
	// Enforce using static values for log messages.
	StaticMsg *bool `json:"static-msg,omitempty" yaml:"static-msg,omitempty" toml:"static-msg,omitempty"`
}

type SpancheckSettings struct {
	// Checks to enable.
	Checks []SpancheckCheck `json:"checks,omitempty" yaml:"checks,omitempty" toml:"checks,omitempty"`
	// A list of regexes for additional function signatures that create spans.
	ExtraStartSpanSignatures []string `json:"extra-start-span-signatures,omitempty" yaml:"extra-start-span-signatures,omitempty" toml:"extra-start-span-signatures,omitempty"`
	// A list of regexes for function signatures that silence `record-error` and `set-status`
	// reports if found in the call path to a returned error.
	IgnoreCheckSignatures []string `json:"ignore-check-signatures,omitempty" yaml:"ignore-check-signatures,omitempty" toml:"ignore-check-signatures,omitempty"`
}

type StaticcheckSettings struct {
	Checks []string `json:"checks,omitempty" yaml:"checks,omitempty" toml:"checks,omitempty"`
	// By default, ST1001 forbids all uses of dot imports in non-test packages. This setting
	// allows setting a whitelist of import paths that can be dot-imported anywhere.
	DotImportWhitelist []string `json:"dot-import-whitelist,omitempty" yaml:"dot-import-whitelist,omitempty" toml:"dot-import-whitelist,omitempty"`
	// ST1013 recommends using constants from the net/http package instead of hard-coding
	// numeric HTTP status codes. This setting specifies a list of numeric status codes that
	// this check does not complain about.
	HTTPStatusCodeWhitelist []HTTPStatusCodeWhitelist `json:"http-status-code-whitelist,omitempty" yaml:"http-status-code-whitelist,omitempty" toml:"http-status-code-whitelist,omitempty"`
	// ST1003 check, among other things, for the correct capitalization of initialisms. The set
	// of known initialisms can be configured with this option.
	Initialisms []string `json:"initialisms,omitempty" yaml:"initialisms,omitempty" toml:"initialisms,omitempty"`
}

type TagalignSettings struct {
	// Align and sort can be used together or separately.
	Align *bool `json:"align,omitempty" yaml:"align,omitempty" toml:"align,omitempty"`
	// Specify the order of tags, the other tags will be sorted by name.
	Order []string `json:"order,omitempty" yaml:"order,omitempty" toml:"order,omitempty"`
	// Whether enable tags sort.
	Sort *bool `json:"sort,omitempty" yaml:"sort,omitempty" toml:"sort,omitempty"`
	// Whether enable strict style.
	Strict *bool `json:"strict,omitempty" yaml:"strict,omitempty" toml:"strict,omitempty"`
}

type TagliatelleSettings struct {
	Case *CaseClass `json:"case,omitempty" yaml:"case,omitempty" toml:"case,omitempty"`
}

type CaseClass struct {
	// Defines the association between tag name and case.
	ExtendedRules map[string]interface{} `json:"extended-rules,omitempty" yaml:"extended-rules,omitempty" toml:"extended-rules,omitempty"`
	// The field names to ignore.
	IgnoredFields []string `json:"ignored-fields,omitempty" yaml:"ignored-fields,omitempty" toml:"ignored-fields,omitempty"`
	// Overrides the default/root configuration.
	Overrides []Override             `json:"overrides,omitempty" yaml:"overrides,omitempty" toml:"overrides,omitempty"`
	Rules     map[string]interface{} `json:"rules,omitempty" yaml:"rules,omitempty" toml:"rules,omitempty"`
	// Use the struct field name to check the name of the struct tag.
	UseFieldName *bool `json:"use-field-name,omitempty" yaml:"use-field-name,omitempty" toml:"use-field-name,omitempty"`
}

type Override struct {
	// Defines the association between tag name and case.
	ExtendedRules map[string]interface{} `json:"extended-rules,omitempty" yaml:"extended-rules,omitempty" toml:"extended-rules,omitempty"`
	// Ignore the package (takes precedence over all other configurations).
	Ignore *bool `json:"ignore,omitempty" yaml:"ignore,omitempty" toml:"ignore,omitempty"`
	// The field names to ignore.
	IgnoredFields []string `json:"ignored-fields,omitempty" yaml:"ignored-fields,omitempty" toml:"ignored-fields,omitempty"`
	// A package path.
	Pkg   string                 `json:"pkg" yaml:"pkg" toml:"pkg"`
	Rules map[string]interface{} `json:"rules,omitempty" yaml:"rules,omitempty" toml:"rules,omitempty"`
	// Use the struct field name to check the name of the struct tag.
	UseFieldName *bool `json:"use-field-name,omitempty" yaml:"use-field-name,omitempty" toml:"use-field-name,omitempty"`
}

type TestifylintSettings struct {
	BoolCompare *BoolCompareClass `json:"bool-compare,omitempty" yaml:"bool-compare,omitempty" toml:"bool-compare,omitempty"`
	// Disable specific checkers.
	Disable []Able `json:"disable,omitempty" yaml:"disable,omitempty" toml:"disable,omitempty"`
	// Disable all checkers.
	DisableAll *bool `json:"disable-all,omitempty" yaml:"disable-all,omitempty" toml:"disable-all,omitempty"`
	// Enable specific checkers.
	Enable []Able `json:"enable,omitempty" yaml:"enable,omitempty" toml:"enable,omitempty"`
	// Enable all checkers.
	EnableAll            *bool                      `json:"enable-all,omitempty" yaml:"enable-all,omitempty" toml:"enable-all,omitempty"`
	ExpectedActual       *ExpectedActualClass       `json:"expected-actual,omitempty" yaml:"expected-actual,omitempty" toml:"expected-actual,omitempty"`
	Formatter            *FormatterClass            `json:"formatter,omitempty" yaml:"formatter,omitempty" toml:"formatter,omitempty"`
	GoRequire            *GoRequireClass            `json:"go-require,omitempty" yaml:"go-require,omitempty" toml:"go-require,omitempty"`
	RequireError         *RequireErrorClass         `json:"require-error,omitempty" yaml:"require-error,omitempty" toml:"require-error,omitempty"`
	SuiteExtraAssertCall *SuiteExtraAssertCallClass `json:"suite-extra-assert-call,omitempty" yaml:"suite-extra-assert-call,omitempty" toml:"suite-extra-assert-call,omitempty"`
}

type BoolCompareClass struct {
	// To ignore user defined types (over builtin bool).
	IgnoreCustomTypes *bool `json:"ignore-custom-types,omitempty" yaml:"ignore-custom-types,omitempty" toml:"ignore-custom-types,omitempty"`
}

type ExpectedActualClass struct {
	// Regexp for expected variable name.
	Pattern *string `json:"pattern,omitempty" yaml:"pattern,omitempty" toml:"pattern,omitempty"`
}

type FormatterClass struct {
	// To enable go vet's printf checks.
	CheckFormatString *bool `json:"check-format-string,omitempty" yaml:"check-format-string,omitempty" toml:"check-format-string,omitempty"`
	// To require f-assertions (e.g. assert.Equalf) if format string is used, even if there are
	// no variable-length variables.
	RequireFFuncs *bool `json:"require-f-funcs,omitempty" yaml:"require-f-funcs,omitempty" toml:"require-f-funcs,omitempty"`
	// To require that the first element of msgAndArgs (msg) has a string type.
	RequireStringMsg *bool `json:"require-string-msg,omitempty" yaml:"require-string-msg,omitempty" toml:"require-string-msg,omitempty"`
}

type GoRequireClass struct {
	// To ignore HTTP handlers (like http.HandlerFunc).
	IgnoreHTTPHandlers *bool `json:"ignore-http-handlers,omitempty" yaml:"ignore-http-handlers,omitempty" toml:"ignore-http-handlers,omitempty"`
}

type RequireErrorClass struct {
	// Regexp for assertions to analyze. If defined, then only matched error assertions will be
	// reported.
	FnPattern *string `json:"fn-pattern,omitempty" yaml:"fn-pattern,omitempty" toml:"fn-pattern,omitempty"`
}

type SuiteExtraAssertCallClass struct {
	// To require or remove extra Assert() call?
	Mode *SuiteExtraAssertCallMode `json:"mode,omitempty" yaml:"mode,omitempty" toml:"mode,omitempty"`
}

type TestpackageSettings struct {
	// List of packages that don't end with _test that tests are allowed to be in.
	AllowPackages []string `json:"allow-packages,omitempty" yaml:"allow-packages,omitempty" toml:"allow-packages,omitempty"`
	// Files with names matching this regular expression are skipped.
	SkipRegexp *string `json:"skip-regexp,omitempty" yaml:"skip-regexp,omitempty" toml:"skip-regexp,omitempty"`
}

type ThelperSettings struct {
	Benchmark *Benchmark `json:"benchmark,omitempty" yaml:"benchmark,omitempty" toml:"benchmark,omitempty"`
	Fuzz      *Fuzz      `json:"fuzz,omitempty" yaml:"fuzz,omitempty" toml:"fuzz,omitempty"`
	TB        *TB        `json:"tb,omitempty" yaml:"tb,omitempty" toml:"tb,omitempty"`
	Test      *Test      `json:"test,omitempty" yaml:"test,omitempty" toml:"test,omitempty"`
}

type Benchmark struct {
	// Check if `b.Helper()` begins helper function.
	Begin *bool `json:"begin,omitempty" yaml:"begin,omitempty" toml:"begin,omitempty"`
	// Check if *testing.B is first param of helper function.
	First *bool `json:"first,omitempty" yaml:"first,omitempty" toml:"first,omitempty"`
	// Check if *testing.B param has b name.
	Name *bool `json:"name,omitempty" yaml:"name,omitempty" toml:"name,omitempty"`
}

type Fuzz struct {
	// Check if `f.Helper()` begins helper function.
	Begin *bool `json:"begin,omitempty" yaml:"begin,omitempty" toml:"begin,omitempty"`
	// Check if *testing.F is first param of helper function.
	First *bool `json:"first,omitempty" yaml:"first,omitempty" toml:"first,omitempty"`
	// Check if *testing.F param has f name.
	Name *bool `json:"name,omitempty" yaml:"name,omitempty" toml:"name,omitempty"`
}

type TB struct {
	// Check if `tb.Helper()` begins helper function.
	Begin *bool `json:"begin,omitempty" yaml:"begin,omitempty" toml:"begin,omitempty"`
	// Check if *testing.TB is first param of helper function.
	First *bool `json:"first,omitempty" yaml:"first,omitempty" toml:"first,omitempty"`
	// Check if *testing.TB param has tb name.
	Name *bool `json:"name,omitempty" yaml:"name,omitempty" toml:"name,omitempty"`
}

type Test struct {
	// Check if `t.Helper()` begins helper function.
	Begin *bool `json:"begin,omitempty" yaml:"begin,omitempty" toml:"begin,omitempty"`
	// Check if *testing.T is first param of helper function.
	First *bool `json:"first,omitempty" yaml:"first,omitempty" toml:"first,omitempty"`
	// Check if *testing.T param has t name.
	Name *bool `json:"name,omitempty" yaml:"name,omitempty" toml:"name,omitempty"`
}

type UnconvertSettings struct {
	FastMath *bool `json:"fast-math,omitempty" yaml:"fast-math,omitempty" toml:"fast-math,omitempty"`
	Safe     *bool `json:"safe,omitempty" yaml:"safe,omitempty" toml:"safe,omitempty"`
}

type UnparamSettings struct {
	// Inspect exported functions. Set to true if no external program/library imports your
	// code.
	//
	// WARNING: if you enable this setting, unparam will report a lot of false-positives in text
	// editors:
	// if it's called for subdir of a project it can't find external interfaces. All text editor
	// integrations
	// with golangci-lint call it on a directory with the changed file.
	CheckExported *bool `json:"check-exported,omitempty" yaml:"check-exported,omitempty" toml:"check-exported,omitempty"`
}

type UnusedSettings struct {
	ExportedFieldsAreUsed  *bool `json:"exported-fields-are-used,omitempty" yaml:"exported-fields-are-used,omitempty" toml:"exported-fields-are-used,omitempty"`
	FieldWritesAreUses     *bool `json:"field-writes-are-uses,omitempty" yaml:"field-writes-are-uses,omitempty" toml:"field-writes-are-uses,omitempty"`
	GeneratedIsUsed        *bool `json:"generated-is-used,omitempty" yaml:"generated-is-used,omitempty" toml:"generated-is-used,omitempty"`
	LocalVariablesAreUsed  *bool `json:"local-variables-are-used,omitempty" yaml:"local-variables-are-used,omitempty" toml:"local-variables-are-used,omitempty"`
	ParametersAreUsed      *bool `json:"parameters-are-used,omitempty" yaml:"parameters-are-used,omitempty" toml:"parameters-are-used,omitempty"`
	PostStatementsAreReads *bool `json:"post-statements-are-reads,omitempty" yaml:"post-statements-are-reads,omitempty" toml:"post-statements-are-reads,omitempty"`
}

type UsestdlibvarsSettings struct {
	// Suggest the use of constant.Kind.String().
	ConstantKind *bool `json:"constant-kind,omitempty" yaml:"constant-kind,omitempty" toml:"constant-kind,omitempty"`
	// Suggest the use of crypto.Hash.String().
	CryptoHash *bool `json:"crypto-hash,omitempty" yaml:"crypto-hash,omitempty" toml:"crypto-hash,omitempty"`
	// Suggest the use of rpc.DefaultXXPath.
	DefaultRPCPath *bool `json:"default-rpc-path,omitempty" yaml:"default-rpc-path,omitempty" toml:"default-rpc-path,omitempty"`
	// Suggest the use of http.MethodXX.
	HTTPMethod *bool `json:"http-method,omitempty" yaml:"http-method,omitempty" toml:"http-method,omitempty"`
	// Suggest the use of http.StatusXX.
	HTTPStatusCode *bool `json:"http-status-code,omitempty" yaml:"http-status-code,omitempty" toml:"http-status-code,omitempty"`
	// Suggest the use of sql.LevelXX.String().
	SQLIsolationLevel *bool `json:"sql-isolation-level,omitempty" yaml:"sql-isolation-level,omitempty" toml:"sql-isolation-level,omitempty"`
	// Suggest the use of time.Month in time.Date.
	TimeDateMonth *bool `json:"time-date-month,omitempty" yaml:"time-date-month,omitempty" toml:"time-date-month,omitempty"`
	// Suggest the use of time.Layout.
	TimeLayout *bool `json:"time-layout,omitempty" yaml:"time-layout,omitempty" toml:"time-layout,omitempty"`
	// Suggest the use of time.Month.String().
	TimeMonth *bool `json:"time-month,omitempty" yaml:"time-month,omitempty" toml:"time-month,omitempty"`
	// Suggest the use of time.Weekday.String().
	TimeWeekday *bool `json:"time-weekday,omitempty" yaml:"time-weekday,omitempty" toml:"time-weekday,omitempty"`
	// Suggest the use of tls.SignatureScheme.String().
	TLSSignatureScheme *bool `json:"tls-signature-scheme,omitempty" yaml:"tls-signature-scheme,omitempty" toml:"tls-signature-scheme,omitempty"`
}

type UsetestingSettings struct {
	ContextBackground *bool `json:"context-background,omitempty" yaml:"context-background,omitempty" toml:"context-background,omitempty"`
	ContextTodo       *bool `json:"context-todo,omitempty" yaml:"context-todo,omitempty" toml:"context-todo,omitempty"`
	OSChdir           *bool `json:"os-chdir,omitempty" yaml:"os-chdir,omitempty" toml:"os-chdir,omitempty"`
	OSCreateTemp      *bool `json:"os-create-temp,omitempty" yaml:"os-create-temp,omitempty" toml:"os-create-temp,omitempty"`
	OSMkdirTemp       *bool `json:"os-mkdir-temp,omitempty" yaml:"os-mkdir-temp,omitempty" toml:"os-mkdir-temp,omitempty"`
	OSSetenv          *bool `json:"os-setenv,omitempty" yaml:"os-setenv,omitempty" toml:"os-setenv,omitempty"`
	OSTempDir         *bool `json:"os-temp-dir,omitempty" yaml:"os-temp-dir,omitempty" toml:"os-temp-dir,omitempty"`
}

type VarnamelenSettings struct {
	// Check method receiver names.
	CheckReceiver *bool `json:"check-receiver,omitempty" yaml:"check-receiver,omitempty" toml:"check-receiver,omitempty"`
	// Check named return values.
	CheckReturn *bool `json:"check-return,omitempty" yaml:"check-return,omitempty" toml:"check-return,omitempty"`
	// Check type parameters.
	CheckTypeParam *bool `json:"check-type-param,omitempty" yaml:"check-type-param,omitempty" toml:"check-type-param,omitempty"`
	// Ignore `ok` variables that hold the bool return value of a channel receive.
	IgnoreChanRecvOk *bool `json:"ignore-chan-recv-ok,omitempty" yaml:"ignore-chan-recv-ok,omitempty" toml:"ignore-chan-recv-ok,omitempty"`
	// Optional list of variable declarations that should be ignored completely.
	IgnoreDecls []string `json:"ignore-decls,omitempty" yaml:"ignore-decls,omitempty" toml:"ignore-decls,omitempty"`
	// Ignore `ok` variables that hold the bool return value of a map index.
	IgnoreMapIndexOk *bool `json:"ignore-map-index-ok,omitempty" yaml:"ignore-map-index-ok,omitempty" toml:"ignore-map-index-ok,omitempty"`
	// Optional list of variable names that should be ignored completely.
	IgnoreNames []string `json:"ignore-names,omitempty" yaml:"ignore-names,omitempty" toml:"ignore-names,omitempty"`
	// Ignore `ok` variables that hold the bool return value of a type assertion
	IgnoreTypeAssertOk *bool `json:"ignore-type-assert-ok,omitempty" yaml:"ignore-type-assert-ok,omitempty" toml:"ignore-type-assert-ok,omitempty"`
	// Variables used in at most this N-many lines will be ignored.
	MaxDistance *int64 `json:"max-distance,omitempty" yaml:"max-distance,omitempty" toml:"max-distance,omitempty"`
	// The minimum length of a variable's name that is considered `long`.
	MinNameLength *int64 `json:"min-name-length,omitempty" yaml:"min-name-length,omitempty" toml:"min-name-length,omitempty"`
}

type WhitespaceSettings struct {
	// Enforces newlines (or comments) after every multi-line function signature
	MultiFunc *bool `json:"multi-func,omitempty" yaml:"multi-func,omitempty" toml:"multi-func,omitempty"`
	// Enforces newlines (or comments) after every multi-line if statement
	MultiIf *bool `json:"multi-if,omitempty" yaml:"multi-if,omitempty" toml:"multi-if,omitempty"`
}

type WrapcheckSettings struct {
	// An array of strings specifying additional substrings of signatures to ignore.
	ExtraIgnoreSigs []string `json:"extra-ignore-sigs,omitempty" yaml:"extra-ignore-sigs,omitempty" toml:"extra-ignore-sigs,omitempty"`
	// An array of glob patterns which, if matched to an underlying interface name, will ignore
	// unwrapped errors returned from a function whose call is defined on the given interface.
	IgnoreInterfaceRegexps []string `json:"ignore-interface-regexps,omitempty" yaml:"ignore-interface-regexps,omitempty" toml:"ignore-interface-regexps,omitempty"`
	// An array of glob patterns which, if any match the package of the function returning the
	// error, will skip wrapcheck analysis for this error.
	IgnorePackageGlobs []string `json:"ignore-package-globs,omitempty" yaml:"ignore-package-globs,omitempty" toml:"ignore-package-globs,omitempty"`
	// An array of strings which specify regular expressions of signatures to ignore.
	IgnoreSigRegexps []string `json:"ignore-sig-regexps,omitempty" yaml:"ignore-sig-regexps,omitempty" toml:"ignore-sig-regexps,omitempty"`
	// An array of strings which specify substrings of signatures to ignore.
	IgnoreSigs []string `json:"ignore-sigs,omitempty" yaml:"ignore-sigs,omitempty" toml:"ignore-sigs,omitempty"`
	// Determines whether wrapcheck should report errors returned from inside the package.
	ReportInternalErrors *bool `json:"report-internal-errors,omitempty" yaml:"report-internal-errors,omitempty" toml:"report-internal-errors,omitempty"`
}

type WslSettings struct {
	// Controls if you may cuddle assignments and anything without needing an empty line between
	// them.
	AllowAssignAndAnything *bool `json:"allow-assign-and-anything,omitempty" yaml:"allow-assign-and-anything,omitempty" toml:"allow-assign-and-anything,omitempty"`
	// Allow calls and assignments to be cuddled as long as the lines have any matching
	// variables, fields or types.
	AllowAssignAndCall *bool `json:"allow-assign-and-call,omitempty" yaml:"allow-assign-and-call,omitempty" toml:"allow-assign-and-call,omitempty"`
	// Allow declarations (var) to be cuddled.
	AllowCuddleDeclarations *bool `json:"allow-cuddle-declarations,omitempty" yaml:"allow-cuddle-declarations,omitempty" toml:"allow-cuddle-declarations,omitempty"`
	// Allow cuddling with any block as long as the variable is used somewhere in the block
	AllowCuddleUsedInBlock *bool `json:"allow-cuddle-used-in-block,omitempty" yaml:"allow-cuddle-used-in-block,omitempty" toml:"allow-cuddle-used-in-block,omitempty"`
	// A list of call idents that everything can be cuddled with.
	AllowCuddleWithCalls []string `json:"allow-cuddle-with-calls,omitempty" yaml:"allow-cuddle-with-calls,omitempty" toml:"allow-cuddle-with-calls,omitempty"`
	// AllowCuddleWithRHS is a list of right hand side variables that is allowed to be cuddled
	// with anything.
	AllowCuddleWithRhs []string `json:"allow-cuddle-with-rhs,omitempty" yaml:"allow-cuddle-with-rhs,omitempty" toml:"allow-cuddle-with-rhs,omitempty"`
	// Allow multiline assignments to be cuddled.
	AllowMultilineAssign *bool `json:"allow-multiline-assign,omitempty" yaml:"allow-multiline-assign,omitempty" toml:"allow-multiline-assign,omitempty"`
	// Allow leading comments to be separated with empty lines.
	AllowSeparatedLeadingComment *bool `json:"allow-separated-leading-comment,omitempty" yaml:"allow-separated-leading-comment,omitempty" toml:"allow-separated-leading-comment,omitempty"`
	// Allow trailing comments in ending of blocks.
	AllowTrailingComment *bool `json:"allow-trailing-comment,omitempty" yaml:"allow-trailing-comment,omitempty" toml:"allow-trailing-comment,omitempty"`
	// When force-err-cuddling is enabled this is a list of names used for error variables to
	// check for in the conditional.
	ErrorVariableNames []string `json:"error-variable-names,omitempty" yaml:"error-variable-names,omitempty" toml:"error-variable-names,omitempty"`
	// Force newlines in end of case at this limit (0 = never).
	ForceCaseTrailingWhitespace *int64 `json:"force-case-trailing-whitespace,omitempty" yaml:"force-case-trailing-whitespace,omitempty" toml:"force-case-trailing-whitespace,omitempty"`
	// Causes an error when an If statement that checks an error variable doesn't cuddle with
	// the assignment of that variable.
	ForceErrCuddling *bool `json:"force-err-cuddling,omitempty" yaml:"force-err-cuddling,omitempty" toml:"force-err-cuddling,omitempty"`
	// Causes an error if a short declaration (:=) cuddles with anything other than another
	// short declaration.
	ForceShortDeclCuddling *bool `json:"force-short-decl-cuddling,omitempty" yaml:"force-short-decl-cuddling,omitempty" toml:"force-short-decl-cuddling,omitempty"`
	// If true, append is only allowed to be cuddled if appending value is matching variables,
	// fields or types on line above.
	StrictAppend *bool `json:"strict-append,omitempty" yaml:"strict-append,omitempty" toml:"strict-append,omitempty"`
}

type WslSettingsV5 struct {
	AllowFirstInBlock *bool         `json:"allow-first-in-block,omitempty" yaml:"allow-first-in-block,omitempty" toml:"allow-first-in-block,omitempty"`
	AllowWholeBlock   *bool         `json:"allow-whole-block,omitempty" yaml:"allow-whole-block,omitempty" toml:"allow-whole-block,omitempty"`
	BranchMaxLines    *int64        `json:"branch-max-lines,omitempty" yaml:"branch-max-lines,omitempty" toml:"branch-max-lines,omitempty"`
	CaseMaxLines      *int64        `json:"case-max-lines,omitempty" yaml:"case-max-lines,omitempty" toml:"case-max-lines,omitempty"`
	Default           *WslV5Default `json:"default,omitempty" yaml:"default,omitempty" toml:"default,omitempty"`
	Disable           []WslChecks   `json:"disable,omitempty" yaml:"disable,omitempty" toml:"disable,omitempty"`
	Enable            []WslChecks   `json:"enable,omitempty" yaml:"enable,omitempty" toml:"enable,omitempty"`
}

// Output configuration options.
type Output struct {
	// Output formats to use.
	Formats  *Formats `json:"formats,omitempty" yaml:"formats,omitempty" toml:"formats,omitempty"`
	PathMode *string  `json:"path-mode,omitempty" yaml:"path-mode,omitempty" toml:"path-mode,omitempty"`
	// Add a prefix to the output file references.
	PathPrefix *string `json:"path-prefix,omitempty" yaml:"path-prefix,omitempty" toml:"path-prefix,omitempty"`
	// Show statistics per linter.
	ShowStats *bool       `json:"show-stats,omitempty" yaml:"show-stats,omitempty" toml:"show-stats,omitempty"`
	SortOrder []SortOrder `json:"sort-order,omitempty" yaml:"sort-order,omitempty" toml:"sort-order,omitempty"`
}

// Output formats to use.
type Formats struct {
	Checkstyle  *SimpleFormat `json:"checkstyle,omitempty" yaml:"checkstyle,omitempty" toml:"checkstyle,omitempty"`
	CodeClimate *SimpleFormat `json:"code-climate,omitempty" yaml:"code-climate,omitempty" toml:"code-climate,omitempty"`
	HTML        *SimpleFormat `json:"html,omitempty" yaml:"html,omitempty" toml:"html,omitempty"`
	JSON        *SimpleFormat `json:"json,omitempty" yaml:"json,omitempty" toml:"json,omitempty"`
	JunitXML    *JunitXML     `json:"junit-xml,omitempty" yaml:"junit-xml,omitempty" toml:"junit-xml,omitempty"`
	Sarif       *SimpleFormat `json:"sarif,omitempty" yaml:"sarif,omitempty" toml:"sarif,omitempty"`
	Tab         *Tab          `json:"tab,omitempty" yaml:"tab,omitempty" toml:"tab,omitempty"`
	Teamcity    *SimpleFormat `json:"teamcity,omitempty" yaml:"teamcity,omitempty" toml:"teamcity,omitempty"`
	Text        *Text         `json:"text,omitempty" yaml:"text,omitempty" toml:"text,omitempty"`
}

type SimpleFormat struct {
	Path *string `json:"path,omitempty" yaml:"path,omitempty" toml:"path,omitempty"`
}

type JunitXML struct {
	Extended *bool   `json:"extended,omitempty" yaml:"extended,omitempty" toml:"extended,omitempty"`
	Path     *string `json:"path,omitempty" yaml:"path,omitempty" toml:"path,omitempty"`
}

type Tab struct {
	Colors          *bool   `json:"colors,omitempty" yaml:"colors,omitempty" toml:"colors,omitempty"`
	Path            *string `json:"path,omitempty" yaml:"path,omitempty" toml:"path,omitempty"`
	PrintLinterName *bool   `json:"print-linter-name,omitempty" yaml:"print-linter-name,omitempty" toml:"print-linter-name,omitempty"`
}

type Text struct {
	Colors           *bool   `json:"colors,omitempty" yaml:"colors,omitempty" toml:"colors,omitempty"`
	Path             *string `json:"path,omitempty" yaml:"path,omitempty" toml:"path,omitempty"`
	PrintIssuedLines *bool   `json:"print-issued-lines,omitempty" yaml:"print-issued-lines,omitempty" toml:"print-issued-lines,omitempty"`
	PrintLinterName  *bool   `json:"print-linter-name,omitempty" yaml:"print-linter-name,omitempty" toml:"print-linter-name,omitempty"`
}

// Options for analysis running,
type Run struct {
	// Allow multiple parallel golangci-lint instances running. If disabled, golangci-lint
	// acquires file lock on start.
	AllowParallelRunners *bool `json:"allow-parallel-runners,omitempty" yaml:"allow-parallel-runners,omitempty" toml:"allow-parallel-runners,omitempty"`
	// Allow multiple golangci-lint instances running, but serialize them around a lock.
	AllowSerialRunners *bool `json:"allow-serial-runners,omitempty" yaml:"allow-serial-runners,omitempty" toml:"allow-serial-runners,omitempty"`
	// List of build tags to pass to all linters.
	BuildTags []string `json:"build-tags,omitempty" yaml:"build-tags,omitempty" toml:"build-tags,omitempty"`
	// Number of concurrent runners. Defaults to the number of available CPU cores.
	Concurrency *int64 `json:"concurrency,omitempty" yaml:"concurrency,omitempty" toml:"concurrency,omitempty"`
	// Targeted Go version.
	Go *string `json:"go,omitempty" yaml:"go,omitempty" toml:"go,omitempty"`
	// Exit code when at least one issue was found.
	IssuesExitCode *int64 `json:"issues-exit-code,omitempty" yaml:"issues-exit-code,omitempty" toml:"issues-exit-code,omitempty"`
	// Option to pass to "go list -mod={option}".
	// See "go help modules" for more information.
	ModulesDownloadMode *ModulesDownloadMode `json:"modules-download-mode,omitempty" yaml:"modules-download-mode,omitempty" toml:"modules-download-mode,omitempty"`
	// The mode used to evaluate relative paths.
	RelativePathMode *RelativePathModeEnum `json:"relative-path-mode,omitempty" yaml:"relative-path-mode,omitempty" toml:"relative-path-mode,omitempty"`
	// Enable inclusion of test files.
	Tests *bool `json:"tests,omitempty" yaml:"tests,omitempty" toml:"tests,omitempty"`
	// Timeout for the analysis.
	Timeout *string `json:"timeout,omitempty" yaml:"timeout,omitempty" toml:"timeout,omitempty"`
}

type SeverityClass struct {
	// Set the default severity for issues. If severity rules are defined and the issues do not
	// match or no severity is provided to the rule this will be the default severity applied.
	// Severities should match the supported severity names of the selected out format.
	Default string `json:"default" yaml:"default" toml:"default"`
	// When a list of severity rules are provided, severity information will be added to lint
	// issues. Severity rules have the same filtering capability as exclude rules except you are
	// allowed to specify one matcher per severity rule.
	// Only affects out formats that support setting severity information.
	Rules []SeverityRule `json:"rules,omitempty" yaml:"rules,omitempty" toml:"rules,omitempty"`
}

type SeverityRule struct {
	Linters    []string `json:"linters,omitempty" yaml:"linters,omitempty" toml:"linters,omitempty"`
	Path       *string  `json:"path,omitempty" yaml:"path,omitempty" toml:"path,omitempty"`
	PathExcept *string  `json:"path-except,omitempty" yaml:"path-except,omitempty" toml:"path-except,omitempty"`
	Severity   string   `json:"severity" yaml:"severity" toml:"severity"`
	Source     *string  `json:"source,omitempty" yaml:"source,omitempty" toml:"source,omitempty"`
	Text       *string  `json:"text,omitempty" yaml:"text,omitempty" toml:"text,omitempty"`
}

// Usable formatter names.
type FormatterNames string

const (
	Gci       FormatterNames = "gci"
	Gofmt     FormatterNames = "gofmt"
	Gofumpt   FormatterNames = "gofumpt"
	Goimports FormatterNames = "goimports"
	Golines   FormatterNames = "golines"
	Swaggo    FormatterNames = "swaggo"
)

type Generated string

const (
	Disable Generated = "disable"
	Lax     Generated = "lax"
	Strict  Generated = "strict"
)

type LintersDefault string

const (
	Fast       LintersDefault = "fast"
	PurpleAll  LintersDefault = "all"
	PurpleNone LintersDefault = "none"
	Standard   LintersDefault = "standard"
)

type Preset string

const (
	Comments             Preset = "comments"
	CommonFalsePositives Preset = "common-false-positives"
	Legacy               Preset = "legacy"
	StdErrorHandling     Preset = "std-error-handling"
)

type DECOrder string

const (
	Const        DECOrder = "const"
	DECOrderFunc DECOrder = "func"
	Type         DECOrder = "type"
	Var          DECOrder = "var"
)

type GocriticChecks string

const (
	AppendAssign             GocriticChecks = "appendAssign"
	AppendCombine            GocriticChecks = "appendCombine"
	ArgOrder                 GocriticChecks = "argOrder"
	AssignOp                 GocriticChecks = "assignOp"
	BadCall                  GocriticChecks = "badCall"
	BadCond                  GocriticChecks = "badCond"
	BadLock                  GocriticChecks = "badLock"
	BadRegexp                GocriticChecks = "badRegexp"
	BadSorting               GocriticChecks = "badSorting"
	BadSyncOnceFunc          GocriticChecks = "badSyncOnceFunc"
	BoolExprSimplify         GocriticChecks = "boolExprSimplify"
	BuiltinShadow            GocriticChecks = "builtinShadow"
	BuiltinShadowDecl        GocriticChecks = "builtinShadowDecl"
	CaptLocal                GocriticChecks = "captLocal"
	CaseOrder                GocriticChecks = "caseOrder"
	CodegenComment           GocriticChecks = "codegenComment"
	CommentFormatting        GocriticChecks = "commentFormatting"
	CommentedOutCode         GocriticChecks = "commentedOutCode"
	CommentedOutImport       GocriticChecks = "commentedOutImport"
	DefaultCaseOrder         GocriticChecks = "defaultCaseOrder"
	DeferInLoop              GocriticChecks = "deferInLoop"
	DeferUnlambda            GocriticChecks = "deferUnlambda"
	DeprecatedComment        GocriticChecks = "deprecatedComment"
	DocStub                  GocriticChecks = "docStub"
	DupArg                   GocriticChecks = "dupArg"
	DupBranchBody            GocriticChecks = "dupBranchBody"
	DupCase                  GocriticChecks = "dupCase"
	DupImport                GocriticChecks = "dupImport"
	DupSubExpr               GocriticChecks = "dupSubExpr"
	DynamicFmtString         GocriticChecks = "dynamicFmtString"
	Elseif                   GocriticChecks = "elseif"
	EmptyDecl                GocriticChecks = "emptyDecl"
	EmptyFallthrough         GocriticChecks = "emptyFallthrough"
	EmptyStringTest          GocriticChecks = "emptyStringTest"
	EqualFold                GocriticChecks = "equalFold"
	EvalOrder                GocriticChecks = "evalOrder"
	ExitAfterDefer           GocriticChecks = "exitAfterDefer"
	ExposedSyncMutex         GocriticChecks = "exposedSyncMutex"
	ExternalErrorReassign    GocriticChecks = "externalErrorReassign"
	FilepathJoin             GocriticChecks = "filepathJoin"
	FlagDeref                GocriticChecks = "flagDeref"
	FlagName                 GocriticChecks = "flagName"
	HTTPNoBody               GocriticChecks = "httpNoBody"
	HexLiteral               GocriticChecks = "hexLiteral"
	HugeParam                GocriticChecks = "hugeParam"
	IfElseChain              GocriticChecks = "ifElseChain"
	ImportShadow             GocriticChecks = "importShadow"
	IndexAlloc               GocriticChecks = "indexAlloc"
	InitClause               GocriticChecks = "initClause"
	IoutilDeprecated         GocriticChecks = "ioutilDeprecated"
	MapKey                   GocriticChecks = "mapKey"
	MethodExprCall           GocriticChecks = "methodExprCall"
	NestingReduce            GocriticChecks = "nestingReduce"
	NewDeref                 GocriticChecks = "newDeref"
	NilValReturn             GocriticChecks = "nilValReturn"
	OctalLiteral             GocriticChecks = "octalLiteral"
	OffBy1                   GocriticChecks = "offBy1"
	ParamTypeCombine         GocriticChecks = "paramTypeCombine"
	PreferDecodeRune         GocriticChecks = "preferDecodeRune"
	PreferFilepathJoin       GocriticChecks = "preferFilepathJoin"
	PreferFprint             GocriticChecks = "preferFprint"
	PreferStringWriter       GocriticChecks = "preferStringWriter"
	PreferWriteByte          GocriticChecks = "preferWriteByte"
	PtrToRefParam            GocriticChecks = "ptrToRefParam"
	RangeAppendAll           GocriticChecks = "rangeAppendAll"
	RangeExprCopy            GocriticChecks = "rangeExprCopy"
	RangeValCopy             GocriticChecks = "rangeValCopy"
	RedundantSprint          GocriticChecks = "redundantSprint"
	RegexpMust               GocriticChecks = "regexpMust"
	RegexpPattern            GocriticChecks = "regexpPattern"
	RegexpSimplify           GocriticChecks = "regexpSimplify"
	ReturnAfterHTTPError     GocriticChecks = "returnAfterHttpError"
	Ruleguard                GocriticChecks = "ruleguard"
	SQLQuery                 GocriticChecks = "sqlQuery"
	SingleCaseSwitch         GocriticChecks = "singleCaseSwitch"
	SliceClear               GocriticChecks = "sliceClear"
	SloppyLen                GocriticChecks = "sloppyLen"
	SloppyReassign           GocriticChecks = "sloppyReassign"
	SloppyTypeAssert         GocriticChecks = "sloppyTypeAssert"
	SortSlice                GocriticChecks = "sortSlice"
	SprintfQuotedString      GocriticChecks = "sprintfQuotedString"
	StringConcatSimplify     GocriticChecks = "stringConcatSimplify"
	StringXbytes             GocriticChecks = "stringXbytes"
	StringsCompare           GocriticChecks = "stringsCompare"
	SwitchTrue               GocriticChecks = "switchTrue"
	SyncMapLoadAndDelete     GocriticChecks = "syncMapLoadAndDelete"
	TimeExprSimplify         GocriticChecks = "timeExprSimplify"
	TodoCommentWithoutDetail GocriticChecks = "todoCommentWithoutDetail"
	TooManyResultsChecker    GocriticChecks = "tooManyResultsChecker"
	TruncateCmp              GocriticChecks = "truncateCmp"
	TypeAssertChain          GocriticChecks = "typeAssertChain"
	TypeDefFirst             GocriticChecks = "typeDefFirst"
	TypeSwitchVar            GocriticChecks = "typeSwitchVar"
	TypeUnparen              GocriticChecks = "typeUnparen"
	UncheckedInlineErr       GocriticChecks = "uncheckedInlineErr"
	Underef                  GocriticChecks = "underef"
	UnlabelStmt              GocriticChecks = "unlabelStmt"
	Unlambda                 GocriticChecks = "unlambda"
	UnnamedResult            GocriticChecks = "unnamedResult"
	UnnecessaryBlock         GocriticChecks = "unnecessaryBlock"
	UnnecessaryDefer         GocriticChecks = "unnecessaryDefer"
	Unslice                  GocriticChecks = "unslice"
	ValSwap                  GocriticChecks = "valSwap"
	WeakCond                 GocriticChecks = "weakCond"
	WhyNoLint                GocriticChecks = "whyNoLint"
	WrapperFunc              GocriticChecks = "wrapperFunc"
	YodaStyleExpr            GocriticChecks = "yodaStyleExpr"
)

type GocriticTags string

const (
	Diagnostic   GocriticTags = "diagnostic"
	Experimental GocriticTags = "experimental"
	Opinionated  GocriticTags = "opinionated"
	Performance  GocriticTags = "performance"
	Security     GocriticTags = "security"
	Style        GocriticTags = "style"
)

// Comments to be checked.
type ScopeEnum string

const (
	Declarations ScopeEnum = "declarations"
	ScopeAll     ScopeEnum = "all"
	Toplevel     ScopeEnum = "toplevel"
)

// Filter out the issues with a lower confidence than the given value
//
// Filter out the issues with a lower severity than the given value
type Confidence string

const (
	High   Confidence = "high"
	Low    Confidence = "low"
	Medium Confidence = "medium"
)

type GosecRules string

const (
	G101 GosecRules = "G101"
	G102 GosecRules = "G102"
	G103 GosecRules = "G103"
	G104 GosecRules = "G104"
	G106 GosecRules = "G106"
	G107 GosecRules = "G107"
	G108 GosecRules = "G108"
	G109 GosecRules = "G109"
	G110 GosecRules = "G110"
	G111 GosecRules = "G111"
	G112 GosecRules = "G112"
	G114 GosecRules = "G114"
	G115 GosecRules = "G115"
	G201 GosecRules = "G201"
	G202 GosecRules = "G202"
	G203 GosecRules = "G203"
	G204 GosecRules = "G204"
	G301 GosecRules = "G301"
	G302 GosecRules = "G302"
	G303 GosecRules = "G303"
	G304 GosecRules = "G304"
	G305 GosecRules = "G305"
	G306 GosecRules = "G306"
	G307 GosecRules = "G307"
	G401 GosecRules = "G401"
	G402 GosecRules = "G402"
	G403 GosecRules = "G403"
	G404 GosecRules = "G404"
	G405 GosecRules = "G405"
	G406 GosecRules = "G406"
	G501 GosecRules = "G501"
	G502 GosecRules = "G502"
	G503 GosecRules = "G503"
	G504 GosecRules = "G504"
	G505 GosecRules = "G505"
	G506 GosecRules = "G506"
	G507 GosecRules = "G507"
	G601 GosecRules = "G601"
	G602 GosecRules = "G602"
)

type GovetAnalyzers string

const (
	Appends                 GovetAnalyzers = "appends"
	Asmdecl                 GovetAnalyzers = "asmdecl"
	Atomicalign             GovetAnalyzers = "atomicalign"
	Bools                   GovetAnalyzers = "bools"
	Buildtag                GovetAnalyzers = "buildtag"
	Cgocall                 GovetAnalyzers = "cgocall"
	Composites              GovetAnalyzers = "composites"
	Copylocks               GovetAnalyzers = "copylocks"
	Deepequalerrors         GovetAnalyzers = "deepequalerrors"
	Defers                  GovetAnalyzers = "defers"
	Directive               GovetAnalyzers = "directive"
	Errorsas                GovetAnalyzers = "errorsas"
	Fieldalignment          GovetAnalyzers = "fieldalignment"
	Findcall                GovetAnalyzers = "findcall"
	Framepointer            GovetAnalyzers = "framepointer"
	GovetAnalyzersAssign    GovetAnalyzers = "assign"
	GovetAnalyzersAtomic    GovetAnalyzers = "atomic"
	GovetAnalyzersUnsafeptr GovetAnalyzers = "unsafeptr"
	Hostport                GovetAnalyzers = "hostport"
	Httpmux                 GovetAnalyzers = "httpmux"
	Httpresponse            GovetAnalyzers = "httpresponse"
	Ifaceassert             GovetAnalyzers = "ifaceassert"
	Loopclosure             GovetAnalyzers = "loopclosure"
	Lostcancel              GovetAnalyzers = "lostcancel"
	Nilfunc                 GovetAnalyzers = "nilfunc"
	Nilness                 GovetAnalyzers = "nilness"
	Printf                  GovetAnalyzers = "printf"
	Reflectvaluecompare     GovetAnalyzers = "reflectvaluecompare"
	Shadow                  GovetAnalyzers = "shadow"
	Shift                   GovetAnalyzers = "shift"
	Sigchanyzer             GovetAnalyzers = "sigchanyzer"
	Slog                    GovetAnalyzers = "slog"
	Sortslice               GovetAnalyzers = "sortslice"
	Stdmethods              GovetAnalyzers = "stdmethods"
	Stdversion              GovetAnalyzers = "stdversion"
	Stringintconv           GovetAnalyzers = "stringintconv"
	Structtag               GovetAnalyzers = "structtag"
	Testinggoroutine        GovetAnalyzers = "testinggoroutine"
	Tests                   GovetAnalyzers = "tests"
	Timeformat              GovetAnalyzers = "timeformat"
	Unmarshal               GovetAnalyzers = "unmarshal"
	Unreachable             GovetAnalyzers = "unreachable"
	Unusedresult            GovetAnalyzers = "unusedresult"
	Unusedwrite             GovetAnalyzers = "unusedwrite"
	Waitgroup               GovetAnalyzers = "waitgroup"
)

type IfaceAnalyzers string

const (
	Identical  IfaceAnalyzers = "identical"
	Opaque     IfaceAnalyzers = "opaque"
	Unexported IfaceAnalyzers = "unexported"
	Unused     IfaceAnalyzers = "unused"
)

type Locale string

const (
	Uk Locale = "UK"
	Us Locale = "US"
)

// Mode of the analysis.
type MisspellMode string

const (
	Mode        MisspellMode = ""
	ModeDefault MisspellMode = "default"
	Restricted  MisspellMode = "restricted"
)

type MndCheck string

const (
	Argument    MndCheck = "argument"
	Case        MndCheck = "case"
	CheckAssign MndCheck = "assign"
	CheckReturn MndCheck = "return"
	Condition   MndCheck = "condition"
	Operation   MndCheck = "operation"
)

type CheckedType string

const (
	Chan                 CheckedType = "chan"
	CheckedTypeFunc      CheckedType = "func"
	CheckedTypeUnsafeptr CheckedType = "unsafeptr"
	Iface                CheckedType = "iface"
	Map                  CheckedType = "map"
	Ptr                  CheckedType = "ptr"
	Uintptr              CheckedType = "uintptr"
)

type DisabledLinter string

const (
	CamelCase                DisabledLinter = "CamelCase"
	Counter                  DisabledLinter = "Counter"
	Help                     DisabledLinter = "Help"
	HistogramSummaryReserved DisabledLinter = "HistogramSummaryReserved"
	MetricTypeInName         DisabledLinter = "MetricTypeInName"
	MetricUnits              DisabledLinter = "MetricUnits"
	ReservedChars            DisabledLinter = "ReservedChars"
	UnitAbbreviations        DisabledLinter = "UnitAbbreviations"
)

type Name string

const (
	SpecifyDisableReason Name = "specify-disable-reason"
)

type SeverityEnum string

const (
	Error   SeverityEnum = "error"
	Warning SeverityEnum = "warning"
)

type ReviveRules string

const (
	AddConstant                 ReviveRules = "add-constant"
	ArgumentLimit               ReviveRules = "argument-limit"
	BannedCharacters            ReviveRules = "banned-characters"
	BareReturn                  ReviveRules = "bare-return"
	BlankImports                ReviveRules = "blank-imports"
	BoolLiteralInExpr           ReviveRules = "bool-literal-in-expr"
	CallToGc                    ReviveRules = "call-to-gc"
	CognitiveComplexity         ReviveRules = "cognitive-complexity"
	CommentSpacings             ReviveRules = "comment-spacings"
	CommentsDensity             ReviveRules = "comments-density"
	ConfusingNaming             ReviveRules = "confusing-naming"
	ConfusingResults            ReviveRules = "confusing-results"
	ConstantLogicalExpr         ReviveRules = "constant-logical-expr"
	ContextAsArgument           ReviveRules = "context-as-argument"
	ContextKeysType             ReviveRules = "context-keys-type"
	Cyclomatic                  ReviveRules = "cyclomatic"
	Datarace                    ReviveRules = "datarace"
	DeepExit                    ReviveRules = "deep-exit"
	DotImports                  ReviveRules = "dot-imports"
	DuplicatedImports           ReviveRules = "duplicated-imports"
	EarlyReturn                 ReviveRules = "early-return"
	EmptyBlock                  ReviveRules = "empty-block"
	EmptyLines                  ReviveRules = "empty-lines"
	EnforceMapStyle             ReviveRules = "enforce-map-style"
	EnforceRepeatedArgTypeStyle ReviveRules = "enforce-repeated-arg-type-style"
	EnforceSliceStyle           ReviveRules = "enforce-slice-style"
	ErrorNaming                 ReviveRules = "error-naming"
	ErrorReturn                 ReviveRules = "error-return"
	ErrorStrings                ReviveRules = "error-strings"
	Errorf                      ReviveRules = "errorf"
	Exported                    ReviveRules = "exported"
	FileHeader                  ReviveRules = "file-header"
	FileLengthLimit             ReviveRules = "file-length-limit"
	FilenameFormat              ReviveRules = "filename-format"
	FlagParameter               ReviveRules = "flag-parameter"
	FunctionLength              ReviveRules = "function-length"
	FunctionResultLimit         ReviveRules = "function-result-limit"
	GetReturn                   ReviveRules = "get-return"
	IdenticalBranches           ReviveRules = "identical-branches"
	IfReturn                    ReviveRules = "if-return"
	ImportAliasNaming           ReviveRules = "import-alias-naming"
	ImportShadowing             ReviveRules = "import-shadowing"
	ImportsBlocklist            ReviveRules = "imports-blocklist"
	IncrementDecrement          ReviveRules = "increment-decrement"
	IndentErrorFlow             ReviveRules = "indent-error-flow"
	LineLengthLimit             ReviveRules = "line-length-limit"
	MaxControlNesting           ReviveRules = "max-control-nesting"
	MaxPublicStructs            ReviveRules = "max-public-structs"
	ModifiesParameter           ReviveRules = "modifies-parameter"
	ModifiesValueReceiver       ReviveRules = "modifies-value-receiver"
	NestedStructs               ReviveRules = "nested-structs"
	OptimizeOperandsOrder       ReviveRules = "optimize-operands-order"
	PackageComments             ReviveRules = "package-comments"
	RangeValAddress             ReviveRules = "range-val-address"
	RangeValInClosure           ReviveRules = "range-val-in-closure"
	ReceiverNaming              ReviveRules = "receiver-naming"
	RedefinesBuiltinID          ReviveRules = "redefines-builtin-id"
	RedundantBuildTag           ReviveRules = "redundant-build-tag"
	RedundantImportAlias        ReviveRules = "redundant-import-alias"
	RedundantTestMainExit       ReviveRules = "redundant-test-main-exit"
	ReviveRulesAtomic           ReviveRules = "atomic"
	ReviveRulesDefer            ReviveRules = "defer"
	ReviveRulesRange            ReviveRules = "range"
	StringFormat                ReviveRules = "string-format"
	StringOfInt                 ReviveRules = "string-of-int"
	StructTag                   ReviveRules = "struct-tag"
	SuperfluousElse             ReviveRules = "superfluous-else"
	TimeDate                    ReviveRules = "time-date"
	TimeEqual                   ReviveRules = "time-equal"
	TimeNaming                  ReviveRules = "time-naming"
	UncheckedTypeAssertion      ReviveRules = "unchecked-type-assertion"
	UnconditionalRecursion      ReviveRules = "unconditional-recursion"
	UnexportedNaming            ReviveRules = "unexported-naming"
	UnexportedReturn            ReviveRules = "unexported-return"
	UnhandledError              ReviveRules = "unhandled-error"
	UnnecessaryFormat           ReviveRules = "unnecessary-format"
	UnnecessaryStmt             ReviveRules = "unnecessary-stmt"
	UnreachableCode             ReviveRules = "unreachable-code"
	UnusedParameter             ReviveRules = "unused-parameter"
	UnusedReceiver              ReviveRules = "unused-receiver"
	UseAny                      ReviveRules = "use-any"
	UseErrorsNew                ReviveRules = "use-errors-new"
	UseFmtPrint                 ReviveRules = "use-fmt-print"
	UselessBreak                ReviveRules = "useless-break"
	VarDeclaration              ReviveRules = "var-declaration"
	VarNaming                   ReviveRules = "var-naming"
	WaitgroupByValue            ReviveRules = "waitgroup-by-value"
)

// Enforce using methods that accept a context.
type ContextEnum string

const (
	Context    ContextEnum = ""
	ContextAll ContextEnum = "all"
	Scope      ContextEnum = "scope"
)

// Enforce a single key naming convention.
type KeyNamingCase string

const (
	Camel  KeyNamingCase = "camel"
	Kebab  KeyNamingCase = "kebab"
	Pascal KeyNamingCase = "pascal"
	Snake  KeyNamingCase = "snake"
)

// Enforce message style.
type MsgStyleEnum string

const (
	Capitalized MsgStyleEnum = "capitalized"
	Lowercased  MsgStyleEnum = "lowercased"
	MsgStyle    MsgStyleEnum = ""
)

// Enforce not using global loggers.
type NoGlobalEnum string

const (
	NoGlobal        NoGlobalEnum = ""
	NoGlobalAll     NoGlobalEnum = "all"
	NoGlobalDefault NoGlobalEnum = "default"
)

type SpancheckCheck string

const (
	End         SpancheckCheck = "end"
	RecordError SpancheckCheck = "record-error"
	SetStatus   SpancheckCheck = "set-status"
)

type HTTPStatusCodeWhitelist string

const (
	The100 HTTPStatusCodeWhitelist = "100"
	The101 HTTPStatusCodeWhitelist = "101"
	The102 HTTPStatusCodeWhitelist = "102"
	The103 HTTPStatusCodeWhitelist = "103"
	The200 HTTPStatusCodeWhitelist = "200"
	The201 HTTPStatusCodeWhitelist = "201"
	The202 HTTPStatusCodeWhitelist = "202"
	The203 HTTPStatusCodeWhitelist = "203"
	The204 HTTPStatusCodeWhitelist = "204"
	The205 HTTPStatusCodeWhitelist = "205"
	The206 HTTPStatusCodeWhitelist = "206"
	The207 HTTPStatusCodeWhitelist = "207"
	The208 HTTPStatusCodeWhitelist = "208"
	The226 HTTPStatusCodeWhitelist = "226"
	The300 HTTPStatusCodeWhitelist = "300"
	The301 HTTPStatusCodeWhitelist = "301"
	The302 HTTPStatusCodeWhitelist = "302"
	The303 HTTPStatusCodeWhitelist = "303"
	The304 HTTPStatusCodeWhitelist = "304"
	The305 HTTPStatusCodeWhitelist = "305"
	The306 HTTPStatusCodeWhitelist = "306"
	The307 HTTPStatusCodeWhitelist = "307"
	The308 HTTPStatusCodeWhitelist = "308"
	The400 HTTPStatusCodeWhitelist = "400"
	The401 HTTPStatusCodeWhitelist = "401"
	The402 HTTPStatusCodeWhitelist = "402"
	The403 HTTPStatusCodeWhitelist = "403"
	The404 HTTPStatusCodeWhitelist = "404"
	The405 HTTPStatusCodeWhitelist = "405"
	The406 HTTPStatusCodeWhitelist = "406"
	The407 HTTPStatusCodeWhitelist = "407"
	The408 HTTPStatusCodeWhitelist = "408"
	The409 HTTPStatusCodeWhitelist = "409"
	The410 HTTPStatusCodeWhitelist = "410"
	The411 HTTPStatusCodeWhitelist = "411"
	The412 HTTPStatusCodeWhitelist = "412"
	The413 HTTPStatusCodeWhitelist = "413"
	The414 HTTPStatusCodeWhitelist = "414"
	The415 HTTPStatusCodeWhitelist = "415"
	The416 HTTPStatusCodeWhitelist = "416"
	The417 HTTPStatusCodeWhitelist = "417"
	The418 HTTPStatusCodeWhitelist = "418"
	The421 HTTPStatusCodeWhitelist = "421"
	The422 HTTPStatusCodeWhitelist = "422"
	The423 HTTPStatusCodeWhitelist = "423"
	The424 HTTPStatusCodeWhitelist = "424"
	The425 HTTPStatusCodeWhitelist = "425"
	The426 HTTPStatusCodeWhitelist = "426"
	The428 HTTPStatusCodeWhitelist = "428"
	The429 HTTPStatusCodeWhitelist = "429"
	The431 HTTPStatusCodeWhitelist = "431"
	The451 HTTPStatusCodeWhitelist = "451"
	The500 HTTPStatusCodeWhitelist = "500"
	The501 HTTPStatusCodeWhitelist = "501"
	The502 HTTPStatusCodeWhitelist = "502"
	The503 HTTPStatusCodeWhitelist = "503"
	The504 HTTPStatusCodeWhitelist = "504"
	The505 HTTPStatusCodeWhitelist = "505"
	The506 HTTPStatusCodeWhitelist = "506"
	The507 HTTPStatusCodeWhitelist = "507"
	The508 HTTPStatusCodeWhitelist = "508"
	The510 HTTPStatusCodeWhitelist = "510"
	The511 HTTPStatusCodeWhitelist = "511"
)

type Able string

const (
	BlankImport          Able = "blank-import"
	BoolCompare          Able = "bool-compare"
	Compares             Able = "compares"
	Contains             Able = "contains"
	Empty                Able = "empty"
	EncodedCompare       Able = "encoded-compare"
	EqualValues          Able = "equal-values"
	ErrorIsAs            Able = "error-is-as"
	ErrorNil             Able = "error-nil"
	ExpectedActual       Able = "expected-actual"
	FloatCompare         Able = "float-compare"
	Formatter            Able = "formatter"
	GoRequire            Able = "go-require"
	Len                  Able = "len"
	NegativePositive     Able = "negative-positive"
	NilCompare           Able = "nil-compare"
	Regexp               Able = "regexp"
	RequireError         Able = "require-error"
	SuiteBrokenParallel  Able = "suite-broken-parallel"
	SuiteDontUsePkg      Able = "suite-dont-use-pkg"
	SuiteExtraAssertCall Able = "suite-extra-assert-call"
	SuiteMethodSignature Able = "suite-method-signature"
	SuiteSubtestRun      Able = "suite-subtest-run"
	SuiteThelper         Able = "suite-thelper"
	UselessAssert        Able = "useless-assert"
)

// To require or remove extra Assert() call?
type SuiteExtraAssertCallMode string

const (
	Remove  SuiteExtraAssertCallMode = "remove"
	Require SuiteExtraAssertCallMode = "require"
)

type WslV5Default string

const (
	Default        WslV5Default = ""
	DefaultDefault WslV5Default = "default"
	FluffyAll      WslV5Default = "all"
	FluffyNone     WslV5Default = "none"
)

type WslChecks string

const (
	Append             WslChecks = "append"
	AssignExclusive    WslChecks = "assign-exclusive"
	AssignExpr         WslChecks = "assign-expr"
	Branch             WslChecks = "branch"
	Decl               WslChecks = "decl"
	Err                WslChecks = "err"
	Expr               WslChecks = "expr"
	For                WslChecks = "for"
	Go                 WslChecks = "go"
	If                 WslChecks = "if"
	IncDEC             WslChecks = "inc-dec"
	Label              WslChecks = "label"
	LeadingWhitespace  WslChecks = "leading-whitespace"
	Select             WslChecks = "select"
	Send               WslChecks = "send"
	Switch             WslChecks = "switch"
	TrailingWhitespace WslChecks = "trailing-whitespace"
	TypeSwitch         WslChecks = "type-switch"
	WslChecksAssign    WslChecks = "assign"
	WslChecksDefer     WslChecks = "defer"
	WslChecksRange     WslChecks = "range"
	WslChecksReturn    WslChecks = "return"
)

type SortOrder string

const (
	File     SortOrder = "file"
	Linter   SortOrder = "linter"
	Severity SortOrder = "severity"
)

// Option to pass to "go list -mod={option}".
// See "go help modules" for more information.
type ModulesDownloadMode string

const (
	Mod      ModulesDownloadMode = "mod"
	Readonly ModulesDownloadMode = "readonly"
	Vendor   ModulesDownloadMode = "vendor"
)

// Usable linter names.
type RelativePathModeEnum string

const (
	CFG     RelativePathModeEnum = "cfg"
	Gitroot RelativePathModeEnum = "gitroot"
	Gomod   RelativePathModeEnum = "gomod"
	Wd      RelativePathModeEnum = "wd"
)
