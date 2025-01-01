// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.793
package components

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

func ProdutoFormAdd() templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		if templ_7745c5c3_CtxErr := ctx.Err(); templ_7745c5c3_CtxErr != nil {
			return templ_7745c5c3_CtxErr
		}
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
		if !templ_7745c5c3_IsBuffer {
			defer func() {
				templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err == nil {
					templ_7745c5c3_Err = templ_7745c5c3_BufErr
				}
			}()
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div x-data=\"{\n        fechaForm() {\n            setTimeout(()=&gt;{\n                this.$refs.nomeInput.value = &#39;&#39;\n                this.$refs.descricaoInput.value = &#39;&#39;\n                $refs.popup.close()\n            }, 100)\n        }\n    }\"><div class=\"btn-close-form-produto\"><i x-on:click=\"fechaForm\" class=\"material-icons\">close </i></div><form hx-post=\"/actions/produto/adiciona\" hx-target=\"#lista-produtos\" hx-swap=\"innerHTML\" x-on:submit=\"fechaForm\"><label for=\"nome\">Nome: <input id=\"nome\" name=\"nome\" type=\"text\" x-ref=\"nomeInput\"></label> <label for=\"descricao\">Descrição: <input id=\"descricao\" name=\"descricao\" type=\"text\" x-ref=\"descricaoInput\"></label> <input type=\"submit\" value=\"Cadastra\"> <input type=\"reset\" value=\"Limpa\"></form></div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}

var _ = templruntime.GeneratedTemplate
