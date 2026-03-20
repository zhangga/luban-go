namespace cfg
{
    public class Tables
    {
        {{range .Tables}}
        public {{.FullName}} {{.Name}} { get; private set; }
        {{end}}

        public Tables(System.Func<string, System.IO.Stream> loader)
        {
            {{range .Tables}}
            // {{.Name}} = new {{.FullName}}(loader("{{if .Raw.InputFiles}}{{index .Raw.InputFiles 0}}{{else}}{{lower .Name}}{{end}}"));
            {{end}}
        }
    }
}