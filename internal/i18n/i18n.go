// Package i18n holds the CLI's translatable strings and a small detection
// helper. Architecture: a single Messages struct of exported string fields
// (plain text or fmt format strings, suffix *Fmt flags the latter). Each
// language declares one Messages value in its own file. Call sites read
// i18n.M.SomeField; for parameterised messages they pass it to fmt.Sprintf.
//
// Adding a field requires updating every messages_*.go file — drift is caught
// at test time by TestCatalogsComplete via reflection, so a missing translation
// fails CI instead of surfacing as a blank line at runtime.
//
// Scope (v1): CLI surface only — welcome, init wizard, chat REPL banner, usage,
// user-facing CLI errors. System prompts, internal error wrappers, and agent
// runtime telemetry stay English so model behaviour and developer logs are
// language-stable.
package i18n

import (
	"os"
	"strings"
)

type Messages struct {
	Subtitle        string
	WelcomeTitleFmt string
	NoConfigYet     string
	StartingChatFmt string
	SetKeyHint      string
	ConfigLabel     string
	ModelsLabel     string
	ConfigNotFound  string
	ConfigErrorFmt  string
	NoKey           string
	Ready           string
	GetStarted      string
	StepScaffold    string
	StepSetKey      string
	InitHint        string
	StepSetKeyHint  string
	StepChatDesc    string
	StepRunDesc     string
	HelpFooter      string
	ChatTip           string
	TurnCancelled     string
	NoSessionToResume string
	ResumeRequiresTTY string
	PickSessionLabel  string
	ResumeListHeader    string
	ResumeBusy          string
	ResumeBadIndexFmt   string
	ResumeAlreadyActive string
	ResumedTitle        string
	ChatThinking           string
	ChatThoughtForFmt      string
	ChatStatusThinkingFmt  string
	ChatStatusIdle         string
	ChatStatusPlanApproval string
	PlanApprovalPrompt     string
	ChatStatusToolApproval string
	ToolApprovalPromptFmt  string
	ToolApprovalSourceFmt  string
	ToolApprovalBuiltIn    string
	ToolApprovalImageUse   string
	DiffFoldedFmt          string
	AskTypeSomething   string
	AskTypingHint      string
	AskChatInstead     string
	ChatStatusQuestion string
	OutputStyleNone   string
	OutputStyleHeader string
	OutputStyleHint   string
	CompactionWorking string
	CompactionTitle   string
	CompactionUnit    string
	CompactionAuto    string
	CompactionManual  string
	SlashCompactDone   string
	SlashCompactFailed string
	SlashNewDone       string
	SlashNewFailed     string
	SlashTodoCleared   string
	SlashUnavailable   string
	SlashUnknown       string
	SlashHelp          string
	SlashPromptEmpty   string
	SlashMCPNone       string
	CtrlCQuitHint      string
	CompHintSlash      string
	CompHintFile       string
	CmdNew          string
	CmdCompact      string
	CmdRewind       string
	CmdTree         string
	CmdBranch       string
	CmdSwitchBranch string
	CmdResume       string
	CmdModel        string
	CmdMemory       string
	CmdForget       string
	CmdMcp          string
	CmdHooks        string
	CmdPasteImage   string
	CmdOutputStyle  string
	CmdSkill        string
	CmdVerbose      string
	CmdHelp         string
	CmdTodo         string
	CmdQuit         string
	ArgSkillList    string
	ArgSkillShow    string
	ArgSkillNew     string
	ArgSkillPaths   string
	ArgMcpAdd       string
	ArgMcpRemove    string
	ArgMcpList      string
	ArgMcpConnected string
	ArgHooksList    string
	ArgHooksTrust   string
	ArgModelCurrent string
	ListModelsHeaderFmt string
	ListModelsHint      string
	ListMemoryHeader    string
	ListMemoryNone      string
	ListSkillsHeaderFmt string
	ListSkillsNone      string
	ListHooksHeaderFmt  string
	ListHooksNone       string
	ListMcpHeader       string
	ListMcpNone         string
	SelectProvidersLabel  string
	EnterAPIKeysHeader    string
	MissingKeyIntro       string
	WroteFileFmt          string
	SetupComplete         string
	SetupCancelled        string
	TryHintFmt            string
	NextHint              string
	ConfirmReconfigureFmt string
	KeepingExisting       string
	NotOverwritingFmt     string
	UnknownCommandFmt string
	UsageRunHint      string
	ErrorPrefix       string
	WriteConfigErr    string
	WriteEnvErr       string
	SelectOneHint  string
	SelectManyHint string
	UsageBody string
}

var M = English

func DetectLanguage(override string) string {
	for _, c := range append([]string{override}, envCandidates()...) {
		if tag := normalize(c); tag != "" {
			return setLanguage(tag)
		}
	}
	return setLanguage("en")
}

func envCandidates() []string {
	keys := []string{"CODELO_LANG", "REASONIX_LANG", "LC_ALL", "LC_MESSAGES", "LANG"}
	out := make([]string, len(keys))
	for i, k := range keys {
		out[i] = os.Getenv(k)
	}
	return out
}

func setLanguage(tag string) string {
	switch tag {
	case "zh":
		M = Chinese
		return "zh"
	case "vi":
		M = Vietnamese
		return "vi"
	default:
		M = English
		return "en"
	}
}

func normalize(s string) string {
	s = strings.ToLower(strings.TrimSpace(s))
	if s == "" {
		return ""
	}
	if strings.HasPrefix(s, "zh") || strings.Contains(s, "chinese") || strings.Contains(s, "中文") {
		return "zh"
	}
	if strings.HasPrefix(s, "vi") || strings.Contains(s, "viet") || strings.Contains(s, "tiếng việt") || strings.Contains(s, "tieng viet") {
		return "vi"
	}
	if strings.HasPrefix(s, "en") || strings.Contains(s, "english") {
		return "en"
	}
	return ""
}