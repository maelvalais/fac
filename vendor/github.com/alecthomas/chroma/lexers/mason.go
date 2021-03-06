package lexers

import (
	. "github.com/alecthomas/chroma" // nolint
)

// Mason lexer.
var Mason = Register(MustNewLexer(
	&Config{
		Name:      "Mason",
		Aliases:   []string{"mason"},
		Filenames: []string{"*.m", "*.mhtml", "*.mc", "*.mi", "autohandler", "dhandler"},
		MimeTypes: []string{"application/x-mason"},
		Priority:  0.1,
	},
	Rules{
		"root": {
			{`\s+`, Text, nil},
			{`(<%doc>)(.*?)(</%doc>)(?s)`, ByGroups(NameTag, CommentMultiline, NameTag), nil},
			{`(<%(?:def|method))(\s*)(.*?)(>)(.*?)(</%\2\s*>)(?s)`, ByGroups(NameTag, Text, NameFunction, NameTag, UsingSelf("root"), NameTag), nil},
			{`(<%\w+)(.*?)(>)(.*?)(</%\2\s*>)(?s)`, ByGroups(NameTag, NameFunction, NameTag, Using(Perl, nil), NameTag), nil},
			{`(<&[^|])(.*?)(,.*?)?(&>)(?s)`, ByGroups(NameTag, NameFunction, Using(Perl, nil), NameTag), nil},
			{`(<&\|)(.*?)(,.*?)?(&>)(?s)`, ByGroups(NameTag, NameFunction, Using(Perl, nil), NameTag), nil},
			{`</&>`, NameTag, nil},
			{`(<%!?)(.*?)(%>)(?s)`, ByGroups(NameTag, Using(Perl, nil), NameTag), nil},
			{`(?<=^)#[^\n]*(\n|\Z)`, Comment, nil},
			{`(?<=^)(%)([^\n]*)(\n|\Z)`, ByGroups(NameTag, Using(Perl, nil), Other), nil},
			{`(?sx)
                 (.+?)               # anything, followed by:
                 (?:
                  (?<=\n)(?=[%#]) |  # an eval or comment line
                  (?=</?[%&]) |      # a substitution or block or
                                     # call start or end
                                     # - don't consume
                  (\\\n) |           # an escaped newline
                  \Z                 # end of string
                 )`, ByGroups(Using(HTML, nil), Operator), nil},
		},
	},
))
