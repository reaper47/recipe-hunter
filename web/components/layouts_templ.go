// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.543
package components

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

import (
	"github.com/reaper47/recipya/internal/templates"
)

func layoutAuth(title string) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
		if !templ_7745c5c3_IsBuffer {
			templ_7745c5c3_Buffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<!doctype html><html lang=\"en\" class=\"h-full bg-indigo-100 dark:bg-gray-800\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = head(title).Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<body class=\"h-full grid place-content-center\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templ_7745c5c3_Var1.Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = toast().Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</body></html>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}

func layoutMain(title string, data templates.Data) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
		if !templ_7745c5c3_IsBuffer {
			templ_7745c5c3_Buffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var2 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var2 == nil {
			templ_7745c5c3_Var2 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<!doctype html><html lang=\"en\" class=\"h-full\" _=\"on htmx:afterSwap\n                if location.pathname is &#39;/recipes&#39; or location.pathname is &#39;/&#39; then\n                    add .active to first &lt;button/&gt; in mobile_nav then\n                    remove .active from last &lt;button/&gt; in mobile_nav then\n                    remove .md:hidden from desktop_nav then\n                    remove .hidden from mobile_nav then\n                    remove .active from first &lt;a/&gt; in recipes_sidebar_cookbooks then\n                    add .active to first &lt;a/&gt; in recipes_sidebar_recipes\n                else if location.pathname is &#39;/cookbooks&#39; then\n                    add .active to last &lt;button/&gt; in mobile_nav then\n                    remove .active from first &lt;button/&gt; in mobile_nav then\n                    remove .md:hidden from desktop_nav then\n                    remove .hidden from mobile_nav then\n                    remove .active from first &lt;a/&gt; in recipes_sidebar_recipes then\n                    add .active to first &lt;a/&gt; in recipes_sidebar_cookbooks\n                else if location.pathname is &#39;/settings&#39; or location.pathname.startsWith(&#39;/recipes/add&#39;) then\n                    add .md:hidden to desktop_nav then\n                    add .hidden to mobile_nav\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = head(title).Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<body class=\"min-h-full\"><header class=\"navbar bg-base-200 shadow-sm print:hidden\"><div class=\"navbar-start\"><a class=\"btn btn-ghost text-lg\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if data.IsAuthenticated {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(" hx-get=\"/\" hx-push-url=\"true\" hx-target=\"#content\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		} else {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(" href=\"/\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(">Recipya</a></div><div class=\"navbar-center\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if data.IsAuthenticated {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div id=\"content-title\" class=\"font-semibold hidden md:block md:text-xl\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var3 string
			templ_7745c5c3_Var3, templ_7745c5c3_Err = templ.JoinStringErrs(title)
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `layouts.templ`, Line: 60, Col: 86}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var3))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</div><button id=\"add_recipe\" class=\"btn btn-primary btn-sm hover:btn-accent\" hx-get=\"/recipes/add\" hx-target=\"#content\" hx-push-url=\"true\">Add recipe</button> <button id=\"add_cookbook\" class=\"btn btn-primary btn-sm hover:btn-accent\" hx-post=\"/cookbooks\" hx-prompt=\"Enter the name of your cookbook\" hx-target=\".cookbooks-display\" hx-swap=\"beforeend\">Add cookbook</button>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</div><div class=\"navbar-end\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if data.IsAuthenticated {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"dropdown dropdown-end\" hx-get=\"/user-initials\" hx-trigger=\"load\" hx-target=\"#user-initials\"><div tabindex=\"0\" role=\"button\" class=\"btn btn-ghost btn-circle avatar placeholder\"><div class=\"bg-neutral text-neutral-content w-10 rounded-full\"><span id=\"user-initials\">A</span></div></div><ul tabindex=\"0\" id=\"avatar_dropdown\" class=\"menu dropdown-content mt-3 z-10 p-2 shadow bg-base-100 rounded-box before:content-[&#39;&#39;] before:absolute before:right-2 before:top-[-9px] before:border-x-[15px] before:border-x-transparent before:border-b-[8px] before:border-b-[#333] dark:before:border-b-[gray]\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			if data.IsAdmin {
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<li onclick=\"document.activeElement?.blur()\"><a href=\"/admin\" hx-get=\"/admin\" hx-target=\"#content\" hx-push-url=\"true\">")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				templ_7745c5c3_Err = iconBuildingLibrary().Render(ctx, templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("Admin</a></li><div class=\"divider m-0\"></div>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<li onclick=\"document.activeElement?.blur()\"><a href=\"/settings\" hx-get=\"/settings\" hx-target=\"#content\" hx-push-url=\"true\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = iconGear().Render(ctx, templ_7745c5c3_Buffer)
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("Settings</a></li><div class=\"divider m-0\"></div><li onclick=\"about_dialog.showModal()\"><a>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = iconHelp().Render(ctx, templ_7745c5c3_Buffer)
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("About</a></li><li><a hx-post=\"/auth/logout\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = iconLogout().Render(ctx, templ_7745c5c3_Buffer)
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("Log out</a></li></ul></div>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = aboutDialog(data.About.Version).Render(ctx, templ_7745c5c3_Buffer)
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		} else {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<a href=\"/auth/login\" class=\"btn btn-ghost\">Log In</a> <a href=\"/auth/register\" class=\"btn btn-ghost\">Sign Up</a>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</div></header><div id=\"fullscreen-loader\" class=\"htmx-indicator\"></div><main class=\"inline-flex w-full\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if data.IsAuthenticated {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<aside id=\"desktop_nav\" class=\"hidden md:block\"><ul class=\"menu menu-sm bg-base-300 rounded-box h-full\" style=\"border-radius: 0\"><li id=\"recipes_sidebar_recipes\" hx-get=\"/recipes\" hx-target=\"#content\" hx-push-url=\"true\" hx-swap-oob=\"true\"><a class=\"tooltip tooltip-right active\" data-tip=\"Recipes\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = iconCubeTransparent().Render(ctx, templ_7745c5c3_Buffer)
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</a></li><li id=\"recipes_sidebar_cookbooks\" hx-get=\"/cookbooks\" hx-target=\"#content\" hx-push-url=\"true\" hx-swap-oob=\"true\"><a class=\"tooltip tooltip-right\" data-tip=\"Cookbooks\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = iconBook().Render(ctx, templ_7745c5c3_Buffer)
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</a></li></ul></aside><aside id=\"mobile_nav\" class=\"btm-nav btm-nav-xs md:hidden z-20\"><button hx-get=\"/recipes\" hx-target=\"#content\" hx-push-url=\"true\" hx-swap-oob=\"true\">Recipes</button> <button hx-get=\"/cookbooks\" hx-target=\"#content\" hx-push-url=\"true\" hx-swap-oob=\"true\">Cookbooks</button></aside>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div id=\"content\" class=\"min-h-[92.5vh] w-full\" hx-ext=\"ws\" ws-connect=\"/ws\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templ_7745c5c3_Var2.Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</div></main>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = toast().Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = toastWS("", "", false).Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<script defer>\n        document.addEventListener(\"htmx:wsBeforeMessage\", (event) => {\n            try {\n                const {type, data, fileName} = JSON.parse(event.detail.message);\n                switch (type) {\n                    case \"toast\":\n                        const {message, background} = JSON.parse(data);\n                        showToast(message, background);\n                        break;\n                    case \"file\":\n                        const decoded = atob(data);\n                        const bytes = new Uint8Array(decoded.length);\n                        for (let i = 0; i < decoded.length; i++) {\n                            bytes[i] = decoded.charCodeAt(i);\n                        }\n                        const blob = new Blob([bytes], {type: \"application/zip\"});\n                        downloadFile(blob, fileName, \"application/zip\");\n                        event.preventDefault();\n                        break;\n                }\n            } catch (_) {}\n        });\n    </script></body></html>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}

func head(title string) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
		if !templ_7745c5c3_IsBuffer {
			templ_7745c5c3_Buffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var4 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var4 == nil {
			templ_7745c5c3_Var4 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<head>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if title == "" {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<title hx-swap-oob=\"true\">Recipya</title>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		} else {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<title hx-swap-oob=\"true\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var5 string
			templ_7745c5c3_Var5, templ_7745c5c3_Err = templ.JoinStringErrs(title)
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `layouts.templ`, Line: 198, Col: 36}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var5))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(" | Recipya</title>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<meta charset=\"UTF-8\"><meta http-equiv=\"X-UA-Compatible\" content=\"IE=edge\"><meta name=\"viewport\" content=\"width=device-width, initial-scale=1.0\"><meta name=\"description\" content=\"The ultimate recipes manager for you and your family.\"><meta name=\"keywords\" content=\"Cooking, Lifestyle, Recipes, Groceries, Fast\"><link rel=\"canonical\" href=\"https://recipes.musicavis.com/\"><link rel=\"stylesheet\" href=\"/static/css/tailwind.css\"><link rel=\"stylesheet\" href=\"/static/css/app.css\"><link rel=\"icon\" href=\"/static/favicon.ico\"><script src=\"https://unpkg.com/htmx.org@1.9.10\"></script><script src=\"https://unpkg.com/hyperscript.org@0.9.11\"></script><script src=\"https://unpkg.com/htmx.org/dist/ext/ws.js\"></script><script defer>\n            const cookbooksPattern = new RegExp(\"^/cookbooks/\\\\d+$\");\n            const cookbooksSharePattern = new RegExp(\"^/c/[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$\");\n\n            const recipesPattern = new RegExp(\"^/recipes/\\\\d+(/edit)?$\");\n            const recipesSharePattern = new RegExp(\"^/r/[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$\");\n\n            const pathsShowRecipesSidebar = [\n                \"/\",\n                \"/cookbooks\",\n                \"/recipes\",\n            ];\n\n            const pathsHideAddRecipeButton = [\n                \"/admin\",\n                \"/cookbooks\",\n                \"/settings\",\n                \"/recipes/add\",\n                \"/recipes/add/manual\",\n            ];\n\n            function showAll() {\n                showAddRecipeButton();\n                showAddCookbookButton();\n                showCookbookTitle();\n                showRecipesSidebar();\n            }\n\n            function showAddRecipeButton() {\n                const isRecipe = recipesPattern.test(location.pathname) || recipesSharePattern.test(location.pathname);\n                if (isRecipe ||\n                    pathsHideAddRecipeButton.some(path => path === location.pathname) ||\n                    cookbooksPattern.test(location.pathname) ||\n                    cookbooksSharePattern.test(location.pathname)) {\n                    add_recipe?.classList.add(\"hidden\");\n                } else {\n                    add_recipe?.classList.remove(\"hidden\");\n                }\n            }\n\n            function showAddCookbookButton() {\n                if (add_cookbook) {\n                    add_cookbook.setAttribute(\"hx-target\", \"#content\");\n                    add_cookbook.setAttribute(\"hx-swap\", \"innerHTML\")\n                    htmx.process(add_cookbook);\n                }\n\n                if (location.pathname === \"/cookbooks\") {\n                    add_cookbook?.classList.remove(\"hidden\");\n                } else {\n                    add_cookbook?.classList.add(\"hidden\");\n                }\n            }\n\n            function showCookbookTitle() {\n                const cookbookTitleDiv = document.querySelector(\"#content-title\");\n                if (cookbooksPattern.test(location.pathname) ||\n                    cookbooksSharePattern.test(location.pathname)) {\n                    cookbookTitleDiv?.classList.add(\"md:block\");\n                } else {\n                    cookbookTitleDiv?.classList.remove(\"md:block\");\n                }\n            }\n\n            function showRecipesSidebar() {\n                if (pathsShowRecipesSidebar.includes(location.pathname) || cookbooksPattern.test(location.pathname)) {\n                    desktop_nav.firstElementChild.classList.remove(\"hidden\");\n                    mobile_nav.classList.remove(\"hidden\");\n                } else {\n                    desktop_nav.firstElementChild.classList.add(\"hidden\");\n                    mobile_nav.classList.add(\"hidden\");\n                }\n\n                if (recipesPattern.test(location.pathname) || recipesSharePattern.test(location.pathname) || location.pathname === \"/admin\") {\n                    desktop_nav.firstElementChild.classList.add(\"hidden\");\n                    mobile_nav.classList.add(\"hidden\");\n                } else {\n                    desktop_nav.firstElementChild.classList.remove(\"hidden\");\n                    mobile_nav.classList.remove(\"hidden\");\n                }\n            }\n\n            window.addEventListener(\"DOMContentLoaded\", () => {\n                showAll();\n                document.addEventListener(\"htmx:pushedIntoHistory\", showAll);\n            });\n\n            document.addEventListener(\"htmx:beforeProcessNode\", () => {\n                adjustAddCookbookTarget();\n            });\n\n            htmx.on('htmx:pushedIntoHistory', () => {\n                showAll();\n                document.addEventListener(\"htmx:pushedIntoHistory\", showAll);\n            });\n\n            function adjustAddCookbookTarget() {\n                if (add_cookbook) {\n                    if (document.querySelector(\".cookbooks-display\") === null) {\n                        add_cookbook.setAttribute(\"hx-target\", \"#content\");\n                        add_cookbook.setAttribute(\"hx-swap\", \"innerHTML\");\n                    } else {\n                        add_cookbook.setAttribute(\"hx-target\", \".cookbooks-display\");\n                        add_cookbook.setAttribute(\"hx-swap\", \"beforeend\");\n\n                        if (pagination && !pagination.querySelector(\"button:nth-last-child(2)\").classList.contains('btn-active')) {\n                            add_cookbook.setAttribute(\"hx-swap\", \"none\");\n                        }\n                    }\n                    htmx.process(add_cookbook);\n                }\n            }\n\n            function loadSortableJS() {\n                return loadScript(\"https://cdn.jsdelivr.net/npm/sortablejs@latest/Sortable.min.js\")\n            }\n\n            function loadScript(url) {\n                const script = document.createElement(\"script\");\n                script.src = url;\n                document.body.appendChild(script);\n\n                return new Promise((res, rej) => {\n                    script.onload = () => res();\n                    script.onerror = () => rej();\n                });\n            }\n\n            function downloadFile(data, filename, mime) {\n                const blobURL = window.URL.createObjectURL(data);\n                const a = document.createElement('a');\n                a.style.display = 'none';\n                a.href = blobURL;\n                a.setAttribute('download', filename);\n                if (typeof a.download === 'undefined') {\n                    a.setAttribute('target', '_blank');\n                }\n                document.body.appendChild(a);\n                a.click();\n                document.body.removeChild(a);\n                setTimeout(() => {\n                    window.URL.revokeObjectURL(blobURL);\n                }, 100);\n            }\n        </script></head>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}

func toast() templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
		if !templ_7745c5c3_IsBuffer {
			templ_7745c5c3_Buffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var6 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var6 == nil {
			templ_7745c5c3_Var6 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div id=\"toast_container\" class=\"toast toast-top toast-end hidden z-20 cursor-default\" _=\"on click add .hidden then call clearTimeout(timeoutToast) then set timeoutToast to null then log timeoutToast\"><div class=\"hidden alert alert-error alert-info alert-success alert-warning\"></div></div><script defer>\n        var timeoutToast = timeoutToast || null;\n        htmx.on('showToast', function (event) {\n            const {message, backgroundColor} = JSON.parse(event.detail.value);\n            showToast(message, backgroundColor);\n        });\n\n        function showToast(message, backgroundColor) {\n            const toast = document.createElement('div');\n            toast.classList.add('alert', backgroundColor);\n            toast.textContent = message;\n\n            if (toast_container.firstChild) {\n                toast_container.replaceChild(toast, toast_container.firstChild);\n            } else {\n                toast_container.appendChild(toast);\n            }\n\n            toast_container.classList.remove('hidden');\n\n            timeoutToast = setTimeout(function () {\n                toast_container.classList.add('hidden');\n                toast_container?.removeChild(toast);\n            }, 5000);\n        }\n    </script>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}

func toastWS(title, content string, isToastWSVisible bool) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
		if !templ_7745c5c3_IsBuffer {
			templ_7745c5c3_Buffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var7 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var7 == nil {
			templ_7745c5c3_Var7 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		var templ_7745c5c3_Var8 = []any{"z-20 fixed bottom-0 right-0 p-6 cursor-default", templ.KV("hidden", !isToastWSVisible)}
		templ_7745c5c3_Err = templ.RenderCSSItems(ctx, templ_7745c5c3_Buffer, templ_7745c5c3_Var8...)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div id=\"ws-notification-container\" class=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ.CSSClasses(templ_7745c5c3_Var8).String()))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\"><div class=\"bg-blue-500 text-white px-4 py-2 rounded shadow-md\"><p class=\"font-medium text-center pb-1\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var9 string
		templ_7745c5c3_Var9, templ_7745c5c3_Err = templ.JoinStringErrs(title)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `layouts.templ`, Line: 399, Col: 50}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var9))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</p>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templ.Raw(content).Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</div></div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
