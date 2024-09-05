// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.771
package templates

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

func Hero() templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"text-center mt-2 mx-auto max-w-sm sm:max-w-md md:max-w-xl lg:max-w-2xl\"><span class=\"font-sans text-4xl\">♻️</span><h1 class=\"text-lg italic p-2\">Reduce, Reuse, Recycle!</h1><p class=\"text-md p-2\">Calculate your debt recycling scenario by entering your details below.</p><p class=\"text-md p-4\"><button class=\"py-2 px-4 inline-flex items-center gap-x-2 text-sm font-medium rounded-lg border border-transparent bg-blue-600 text-white hover:bg-blue-700 focus:outline-none focus:bg-blue-700 disabled:opacity-50 disabled:pointer-events-none\" _=\"on click remove .hidden from #info-modal\"><span class=\"font-sans text-lg\">ℹ️</span>More Info</button></p></div><!-- modal body --><div id=\"info-modal\" class=\"fixed inset-0 p-2 z-50 flex items-center justify-center bg-black bg-opacity-50 shadow-xl transition-all backdrop-blur-md hidden\"><div class=\"bg-stone-50 p-4 rounded-lg shadow-lg w-full max-w-4xl h-4/5 mx-4 flex flex-col\"><!-- modal header --><div class=\"flex justify-between items-center py-2 px-2 border-b\"><h2 id=\"info-modal-label\" class=\"font-bold\"><span class=\"font-sans text-lg\">ℹ️</span> Information</h2></div><!-- modal body --><div class=\"overflow-y-auto h-full text-sm sm:text-md\"><div class=\"py-2\"><h3 class=\"text-lg font-semibold\">What is debt recycling?</h3><p class=\"mt-1 text-gray-800\">Debt recycling involves refinancing or redrawing funds from your offset account to invest in income producing assets. The idea is to turn your \"bad debt\" into tax-deductible \"good debt\".</p><p class=\"mt-1 text-gray-800\">The key requirement for the interest to qualify as deductible is that the funds must be used to earn income.</p><br><p class=\"mt-1 text-gray-800\">While this can be a very effective strategy for building wealth, it's important to consider the risks, such as the possibility of losing money if your investments underperform, or worse, go to zero.</p><br><h4 class=\"text-md font-semibold text-gray-800 \">Further reading...</h4><p class=\"mt-1 text-gray-800\"><ul class=\"text-md list-disc px-4\"><li class=\"px-2 mt-2\"><a href=\"https://strongmoneyaustralia.com/debt-recycling-ultimate-guide/\" target=\"_blank\" class=\"text-blue-600 hover:underline\">Strong Money: Debt Recycling Ultimate Guide</a></li><li class=\"px-2 mt-2\"><a href=\"https://www.aussiefirebug.com/debt-recycling/\" target=\"_blank\" class=\"text-blue-600 hover:underline\">Aussie Firebug: Debt Recycling</a></li></ul></p></div><div class=\"py-2\"><h3 class=\"text-lg font-semibold text-gray-800 \">What do all of the fields mean?</h3><p class=\"mt-1 text-gray-800\">Descriptions and examples for all fields are provided below.</p><br><div class=\"mt-2 text-gray-800\"><table class=\"bg-white rounded-lg shadow-md table-auto text-xs sm:text-md mx-auto\"><tr><th class=\"border px-4 py-2\">Field</th><th class=\"border px-4 py-2\">Description</th><th class=\"border px-4 py-2\">Example</th></tr><tr><td class=\"border px-4 py-2\">Salary</td><td class=\"border px-4 py-2\">Your annual salary. This, along with any dividends income will be used to determine your bracket.</td><td class=\"border px-4 py-2\">150000</td></tr><tr><td class=\"border px-4 py-2\">Initial Investment</td><td class=\"border px-4 py-2\">The initial amount you will pay down, recycle and invest. This is the available cash you have now, e.g. in the bank or in your offset account.</td><td class=\"border px-4 py-2\">100000</td></tr><tr><td class=\"border px-4 py-2\">Annual Investment</td><td class=\"border px-4 py-2\">The annual amount you will pay down, recycle and invest. This should take into account regular payments as well as any additional payments you expect to make.</td><td class=\"border px-4 py-2\">50000</td></tr><tr><td class=\"border px-4 py-2\">Mortgage</td><td class=\"border px-4 py-2\">The total size of your mortgage loan.</td><td class=\"border px-4 py-2\">600000</td></tr><tr><td class=\"border px-4 py-2\">Mortgage Interest Rate</td><td class=\"border px-4 py-2\">The projected annual mortgage interest rate.</td><td class=\"border px-4 py-2\">5</td></tr><tr><td class=\"border px-4 py-2\">Dividend Return Rate</td><td class=\"border px-4 py-2\">The projected annual dividend return rate.</td><td class=\"border px-4 py-2\">2</td></tr><tr><td class=\"border px-4 py-2\">Capital Growth Rate</td><td class=\"border px-4 py-2\">The projected annual capital growth rate.</td><td class=\"border px-4 py-2\">8</td></tr><tr><td class=\"border px-4 py-2\">Years</td><td class=\"border px-4 py-2\">The number of years to project the scenario.</td><td class=\"border px-4 py-2\">10</td></tr><tr><td class=\"border px-4 py-2\">Country</td><td class=\"border px-4 py-2\">Select the country you live in. Used to determine your tax bracket alobng with your salary and dividend income.</td><td class=\"border px-4 py-2\">Australia</td></tr><tr><td class=\"border px-4 py-2\">Reinvest Dividends</td><td class=\"border px-4 py-2\">Whether to reinvest your dividend income via debt recycling.</td><td class=\"border px-4 py-2\">✔️</td></tr><tr><td class=\"border px-4 py-2\">Reinvest Tax Refunds</td><td class=\"border px-4 py-2\">Whether to reinvest your tax refunds via debt recycling.</td><td class=\"border px-4 py-2\">✔️</td></tr></table></div></div><div class=\"py-2\"><h3 class=\"text-lg font-semibold text-gray-800 \">DISCLAIMER</h3><p class=\"mt-1 text-gray-800\">This calculator is for illustrative purposes only and is not to be misconstrued as financial advice.</p><br><p class=\"mt-1 text-gray-800\">It does not take into account your individual needs, goals and objectives.</p><br><p class=\"mt-1 text-gray-800\">The results are simply assumptions based on the information provided and are not guaranteed to be accurate.</p><br><p class=\"mt-1 text-gray-800\">In particular, it does not take into account any changes in the market, rates of return, interest rates, tax legislation or any extraordinary occurrences that may impact your results.</p><br><p class=\"mt-1 text-gray-800\">It also does not take into account franking credits, or the (hopefully) increasing value of your property and the subsequent increase in available equity.</p><br><blockquote class=\"relative p-4\"><svg class=\"absolute -top-6 -start-2 size-12 text-gray-200 m-2\" width=\"12\" height=\"12\" viewBox=\"0 0 12 12\" fill=\"none\" aria-hidden=\"true\" version=\"1.1\" id=\"quote-svg\" xmlns=\"http://www.w3.org/2000/svg\" xmlns:svg=\"http://www.w3.org/2000/svg\"><path d=\"m 5.4018188,8.4102461 c 0,0.6966757 -0.224093,1.2612776 -0.6722789,1.6937159 -0.448186,0.408384 -1.0143195,0.612621 -1.6983915,0.612621 -0.8256143,0 -1.5096951,-0.27631 -2.0522428,-0.8288403 C 0.459953,9.3351229 0.20048119,8.5424099 0.20048119,7.5093245 c 0,-1.1051224 0.30665261,-2.1862174 0.91995791,-3.2432852 0.6369091,-1.0810951 1.5214881,-1.933962 2.653755,-2.5585917 L 4.5880064,2.9687251 C 3.9747013,3.4011632 3.4793345,3.8816469 3.1019151,4.4101853 2.7480818,4.9387237 2.5121868,5.5513442 2.3942483,6.248047 c 0.2123,-0.0961 0.459979,-0.144146 0.7430457,-0.144146 0.660486,0 1.2030336,0.2162189 1.6276247,0.648657 0.4246001,0.432438 0.6369001,0.9850037 0.6369001,1.6576881 z m 6.5105182,0 c 0,0.6966757 -0.224057,1.2612776 -0.672262,1.6937159 -0.448203,0.408384 -1.014337,0.612621 -1.6984001,0.612621 -0.8255969,0 -1.5096599,-0.27631 -2.0522434,-0.8288403 C 6.9704788,9.3351229 6.7109979,8.5424099 6.7109979,7.5093245 c 0,-1.1051224 0.3066527,-2.1862174 0.9199669,-3.2432852 0.6369091,-1.0810951 1.5215156,-1.933962 2.6537822,-2.5585917 l 0.813794,1.2612775 C 10.4852,3.4011632 9.9898789,3.8816469 9.6123959,4.4101853 9.2586079,4.9387237 9.0227486,5.5513442 8.9047291,6.248047 9.1170745,6.151947 9.3647349,6.103901 9.6478019,6.103901 c 0.6604591,0 1.2030791,0.2162189 1.6276791,0.648657 0.4246,0.432438 0.636856,0.9850037 0.636856,1.6576881 z\" fill=\"currentColor\" id=\"quote-svg-path\"></path></svg><div class=\"relative z-10\"><p class=\"text-gray-800 sm:text-md\"><em>It is better to be roughly right than precisely wrong.</em></p></div><footer class=\"mt-2\"><div class=\"text-base font-semibold text-gray-800\">John Manyard Keynes</div></footer></blockquote></div></div><!-- modal footer / close button --><hr><div class=\"mt-2 mx-auto p-2\"><button class=\"w-full py-2 px-4 bg-red-500 text-white rounded hover:bg-red-600 max-w-sm\" _=\"on click add .hidden to #info-modal\">Close</button></div></div></div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}

var _ = templruntime.GeneratedTemplate
