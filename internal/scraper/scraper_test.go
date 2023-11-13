package scraper

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/google/go-cmp/cmp"
	"github.com/reaper47/recipya/internal/models"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"testing"
)

func TestScraper(t *testing.T) {
	t.Parallel()

	testcases := []struct {
		name string
		in   string
		want models.RecipeSchema
	}{
		{
			name: "abril.com",
			in:   "https://claudia.abril.com.br/receitas/estrogonofe-de-carne/",
			want: models.RecipeSchema{
				AtContext:     atContext,
				AtType:        models.SchemaType{Value: "Recipe"},
				Category:      models.Category{Value: "Carne"},
				CookTime:      "PT30M",
				CookingMethod: models.CookingMethod{Value: "Refogado"},
				Cuisine:       models.Cuisine{Value: "Brasileira"},
				DateModified:  "2020-02-05T07:51:35-0300",
				DatePublished: "2008-10-24",
				Description: models.Description{
					Value: "Derreta a manteiga e refogue a cebola até ficar transparente. Junte a carne e tempere " +
						"com o sal. Mexa até a carne dourar de todos os lados. Acrescente a mostarda, o catchup, " +
						"a pimenta-do-reino e o tomate picado. Cozinhe até formar um molho espesso. Se necessário, " +
						"adicione água quente aos poucos. Quando o molho estiver [&hellip;]",
				},
				Keywords: models.Keywords{
					Values: "Estrogonofe de carne, Refogado, Dia a Dia, Carne, Brasileira, creme de leite, ketchup (ou catchup), pimenta-do-reino",
				},
				Image: models.Image{
					Value: "https://claudia.abril.com.br/wp-content/uploads/2020/02/receita-estrogonofe-de-carne.jpg?" +
						"quality=85&strip=info&w=620&h=372&crop=1",
				},
				Ingredients: models.Ingredients{
					Values: []string{
						"500 gramas de alcatra cortada em tirinhas",
						"1/4 xícara (chá) de manteiga",
						"1 unidade de cebola picada",
						"1 colher (sobremesa) de mostarda",
						"1 colher (sopa) de ketchup (ou catchup)",
						"1 pitada de pimenta-do-reino",
						"1 unidade de tomate sem pele picado",
						"1 xícara (chá) de cogumelo variado | variados escorridos",
						"1 lata de creme de leite",
						"sal a gosto",
					},
				},
				Instructions: models.Instructions{
					Values: []string{
						"Derreta a manteiga e refogue a cebola até ficar transparente.",
						"Junte a carne e tempere com o sal.",
						"Mexa até a carne dourar de todos os lados.",
						"Acrescente a mostarda, o catchup, a pimenta-do-reino e o tomate picado.",
						"Cozinhe até formar um molho espesso.",
						"Se necessário, adicione água quente aos poucos.",
						"Quando o molho estiver encorpado e a carne macia, adicione os cogumelos e o creme de leite.",
						"Mexa por 1 minuto e retire do fogo.",
						"Sirva imediatamente, acompanhado de arroz e batata palha.",
						"Dica: Se juntar água ao refogar a carne, frite-a até todo o líquido evaporar.",
					},
				},
				Name:     "Estrogonofe de carne",
				PrepTime: "PT30M",
				Yield:    models.Yield{Value: 4},
				URL:      "https://claudia.abril.com.br/receitas/estrogonofe-de-carne/",
			},
		},
		{
			name: "acouplecooks.com",
			in:   "https://www.acouplecooks.com/shaved-brussels-sprouts/",
			want: models.RecipeSchema{
				AtContext: atContext,
				AtType:    models.SchemaType{Value: "Recipe"},
				Name:      "Easy Shaved Brussels Sprouts",
				Description: models.Description{
					Value: "This shaved Brussels sprouts recipe make a tasty side dish that's full of texture and flavor! " +
						"Shredded Brussels are quick and crowd-pleasing.",
				},
				Keywords: models.Keywords{
					Values: "Shaved Brussels sprouts, Shaved Brussels sprouts recipe, shredded Brussel sprouts, shredded " +
						"Brussels sprouts",
				},
				Image: models.Image{
					Value: "https://www.acouplecooks.com/wp-content/uploads/2022/03/Shredded-Brussels-Sprouts-001-225x225.jpg",
				},
				URL: "https://www.acouplecooks.com/shaved-brussels-sprouts/",
				Ingredients: models.Ingredients{
					Values: []string{
						"1 pound Brussels sprouts (off the stalk)",
						"2 cloves garlic, minced",
						"1 small shallot, minced",
						"1/4 cup shredded Parmesan cheese (omit for vegan)",
						"½ teaspoon kosher salt, plus more to taste",
						"2 tablespoons olive oil",
						"1/4 cup Italian panko (optional, omit for gluten-free or use GF panko)",
					},
				},
				Instructions: models.Instructions{
					Values: []string{
						"Shave the Brussels sprouts:\n\nWith a knife: Remove any tough outer layers with your fingers. " +
							"With a large Chef’s knife, cut the Brussels sprout in half lengthwise. Place the cut " +
							"side down and thinly slice cross-wise to create shreds. Separate the shreds with your " +
							"fingers. Discard the root end.",
						"With a food processor (fastest!): Use a food processor with the shredding disc attachment blade. " +
							"(Here&#8217;s a video.)",
						"With a mandolin: Slice the whole Brussels sprouts with a mandolin, taking proper safety precautions " +
							"to keep your fingers away from the blade. (Here&#8217;s a video.)",
						"In a medium bowl, stir together the minced garlic, shallot, Parmesan cheese, and kosher salt.",
						"In a large skillet, heat the olive oil over medium high heat. Add the Brussels sprouts and cook " +
							"for 4 minutes, stirring only occasionally, until tender and browned. Stir in the Parmesan " +
							"mixture and cook additional 3 to 4 minutes until lightly browned and fragrant. Remove the " +
							"heat and if desired, stir in the panko. Taste and add additional salt as necessary.",
					},
				},
				NutritionSchema: models.NutritionSchema{
					Calories:      "149 calories",
					Carbohydrates: "14.6 g",
					Cholesterol:   "3.6 mg",
					Fat:           "9.2 g",
					Fiber:         "6.5 g",
					Protein:       "6.5 g",
					SaturatedFat:  "2.1 g",
					Sodium:        "271.1 mg",
					Sugar:         "3 g",
					TransFat:      "0 g",
				},
				PrepTime:      "PT10M",
				CookTime:      "PT7M",
				Yield:         models.Yield{Value: 4},
				Category:      models.Category{Value: "Side dish"},
				CookingMethod: models.CookingMethod{Value: "Shredded"},
				Cuisine:       models.Cuisine{Value: "Vegetables"},
				DatePublished: "2022-03-23",
			},
		},
		{
			name: "afghankitchenrecipes.com",
			in:   "http://www.afghankitchenrecipes.com/recipe/norinj-palau-rice-with-orange/",
			want: models.RecipeSchema{
				AtContext:     atContext,
				AtType:        models.SchemaType{Value: "Recipe"},
				Name:          "Norinj Palau (Rice with orange)",
				DatePublished: "2015-01-01",
				Image: models.Image{
					Value: "http://www.afghankitchenrecipes.com/wp-content/uploads/2015/01/afghan_norinj_pilaw-250x212.jpg",
				},
				Yield:    models.Yield{Value: 4},
				PrepTime: "PT10M",
				CookTime: "PT2H0M",
				Description: models.Description{
					Value: "Norinj Palau is one of traditional Afghan dishes and it has a lovely delicate flavour. " +
						"This pilau is prepared with the peel of the bitter (or Seville) oranges. It is quite a sweet dish.",
				},
				Ingredients: models.Ingredients{
					Values: []string{
						"450 g long grain white rice, preferably basmati",
						"75 ml vegetable oil",
						"2 medium onions, chopped",
						"1 medium chicken or 700–900 g lamb on the bone cut in pieces",
						"570 ml water, plus 110 ml water",
						"peel of 1 large orange",
						"50 g sugar",
						"50 g blanched and flaked almonds",
						"50 g blanched and flaked pistachios",
						"½ tsp saffron or egg yellow food colour (optional)",
						"25 ml rosewater (optional)",
						"1 tsp ground green or white cardamom seeds (optional)",
						"salt and pepper",
					},
				},
				Instructions: models.Instructions{
					Values: []string{
						"Measure out the rice and rinse several times until the water remains clear.",
						"Add fresh water and leave the rice to soak for at least half an hour.",
						"Heat the oil and add the chopped onions.",
						"Stir and fry them over a medium to high heat until golden brown and soft.",
						"Add the meat and fry until brown, turning frequently.",
						"Add 570 ml of water, salt and pepper and cook gently until the meat is tender.",
						"While the meat is cooking, wash and cut up the zest of a large orange into matchstick-sized pieces, " +
							"removing as much pith as possible.",
						"To remove any bitter taste, put the orange strips into a strainer and dip first in boiling water " +
							"and then in cold.",
						"Repeat this several times. Set aside.",
						"Make a syrup by bringing to the boil 110 ml of water and the 50 g of sugar. Add the orange peel, " +
							"the flaked almonds and pistachios to the boiling syrup.",
						"Boil for about 5 minutes, skimming off the thick froth when necessary. Strain and set aside the peel and nuts.",
						"Add the saffron and rosewater to the syrup and boil again gently for another 3 minutes.",
						"To cook the rice, strain the chicken stock (setting the meat to one side), and add the syrup.",
						"Make the syrup and stock up to 570 ml by adding extra water if necessary.",
						"The oil will be on the surface of the stock and this should also be included in the cooking of the rice.",
						"Bring the liquid to the boil in a large casserole. Drain the rice and then add it to the boiling liquid.",
						"Add salt, the nuts and the peel, reserving about a third for garnishing.",
						"Bring back to the boil, then cover with a tightly fitting lid, turn down the heat to medium and boil for " +
							"about 10 minutes until the rice is tender and all the liquid is absorbed.",
						"Add the meat, the remaining peel and nuts on top of the rice and cover with a tightly fitting lid. Put into " +
							"a preheated oven – 150°C (300°F, mark 2) – for 20–30 minutes. Or cook over a very low heat for the " +
							"same length of time.",
						"When serving, place the meat in the centre of a large dish, mound the rice over the top and then garnish " +
							"with the reserved orange peel and nuts.",
					},
				},
				Category: models.Category{Value: "Rice Dishes"},
				URL:      "http://www.afghankitchenrecipes.com/recipe/norinj-palau-rice-with-orange/",
			},
		},
		{
			name: "allrecipes.com",
			in:   "https://www.allrecipes.com/recipe/10813/best-chocolate-chip-cookies/",
			want: models.RecipeSchema{
				AtContext:     atContext,
				AtType:        models.SchemaType{Value: "Recipe"},
				Category:      models.Category{Value: "Dessert"},
				CookTime:      "PT10M",
				Cuisine:       models.Cuisine{Value: "American"},
				DateModified:  "2023-08-28T17:26:15.610-04:00",
				DatePublished: "1998-04-18T16:10:32.000-04:00",
				Description: models.Description{
					Value: "This chocolate chip cookie recipe makes delicious cookies with crisp edges and chewy middles. Try this wildly-popular cookie recipe for yourself!",
				},
				Image: models.Image{
					Value: "https://www.allrecipes.com/thmb/8xwaWAHtl_QLij6D-G0Z4B1HDVA=/1500x0/filters:no_upscale():max_bytes(150000):strip_icc()/10813-best-chocolate-chip-cookies-mfs-146-4x3-b108aceffa6043a1ac81c3c5a9b034c8.jpg",
				},
				Ingredients: models.Ingredients{
					Values: []string{
						"1 cup butter, softened",
						"1 cup white sugar",
						"1 cup packed brown sugar",
						"2 eggs",
						"2 teaspoons vanilla extract",
						"1 teaspoon baking soda",
						"2 teaspoons hot water",
						"0.5 teaspoon salt",
						"3 cups all-purpose flour",
						"2 cups semisweet chocolate chips",
						"1 cup chopped walnuts",
					},
				},
				Instructions: models.Instructions{
					Values: []string{
						"Gather your ingredients, making sure your butter is softened, and your eggs are room temperature.",
						"Preheat the oven to 350 degrees F (175 degrees C).",
						"Beat butter, white sugar, and brown sugar with an electric mixer in a large bowl until smooth.",
						"Beat in eggs, one at a time, then stir in vanilla.",
						"Dissolve baking soda in hot water. Add to batter along with salt.",
						"Stir in flour, chocolate chips, and walnuts.",
						"Drop spoonfuls of dough 2 inches apart onto ungreased baking sheets.",
						"Bake in the preheated oven until edges are nicely browned, about 10 minutes.",
						"Cool on the baking sheets briefly before removing to a wire rack to cool completely.",
						"Store in an airtight container or serve immediately and enjoy!",
					},
				},
				Name: "Best Chocolate Chip Cookies",
				NutritionSchema: models.NutritionSchema{
					Calories:       "146 kcal",
					Carbohydrates:  "19 g",
					Cholesterol:    "10 mg",
					Fat:            "8 g",
					Fiber:          "1 g",
					Protein:        "2 g",
					SaturatedFat:   "4 g",
					Sodium:         "76 mg",
					UnsaturatedFat: "0 g",
				},
				PrepTime: "PT20M",
				Tools:    models.Tools{},
				Yield:    models.Yield{Value: 48},
				URL:      "https://www.allrecipes.com/recipe/10813/best-chocolate-chip-cookies/",
			},
		},
		{
			name: "amazingribs.com",
			in:   "https://amazingribs.com/tested-recipes-chicken-recipes-crispy-grilled-buffalo-wings-recipe/",
			want: models.RecipeSchema{
				AtContext:     atContext,
				AtType:        models.SchemaType{Value: "Recipe"},
				Category:      models.Category{Value: "Appetizer"},
				CookTime:      "PT30M",
				Cuisine:       models.Cuisine{Value: "American"},
				DatePublished: "2020-01-14T18:29:00+00:00",
				Description: models.Description{
					Value: "True Buffalo wings are deep fried, but I love the flavor and convenience of cooking them on " +
						"the grill, and even smoking them first. And there is much less mess. Click here to tweet this",
				},
				Keywords: models.Keywords{
					Values: "barbecue, buffalo chicken wings, Chicken, chicken wings, grill, grilled buffalo chicken wings, " +
						"grilled chicken wings",
				},
				Image: models.Image{
					Value: "https://amazingribs.com/wp-content/uploads/2020/10/buffalo-wings.jpg",
				},
				Ingredients: models.Ingredients{
					Values: []string{
						"3 ounces cream cheese", "3 ounces quality blue cheese, crumbled",
						"1/2 cup half and half", "1/4 cup sour cream",
						"1/2 teaspoon Simon &amp; Garfunkel Seasoning or Poultry Seasoning",
						"1/2 cup melted salted butter", "2 cloves minced or pressed garlic",
						"1/2 cup Frank's Original RedHot Sauce",
						"24  whole chicken wings ((about 4 pounds (1.8 kg) for 24 whole wings))",
						"Morton Coarse Kosher Salt", "ground black pepper", "6 stalks of celery",
					},
				},
				Instructions: models.Instructions{
					Values: []string{
						"Prep. To create the blue cheese dip, take the cream cheese and the blue cheese out of the fridge and " +
							"let them come to room temp. Then smush them together with the spices. Mix in the sour cream and " +
							"half and half. Put them in a serving bowl and refrigerate. You can do this a day ahead. Cut up the " +
							"celery into 4-inch (10 cm) sections and put it back in the chiller.",
						"For the Buffalo hot sauce, melt the butter over a low heat and then adding the garlic. Let it simmer " +
							"for about a minute but don't let the garlic brown. Then add the Frank's RedHot sauce. Let them get " +
							"to know each other for at least 3 to 4 minutes. But remember, if you don't want to use the original " +
							"and you want to get creative, try one or more of the other sauces listed above. I'm partial to DC Mumbo " +
							"Sauce. Like the dip, the sauce can be made a day in advance.",
						"When it comes time to prep the wings, note that there are three distinct pieces of different thickness " +
							"and skin to meat ratio: (1) The tips (2) the flats or wingettes in the center, and (3) the drumettes " +
							"on the end that attaches to the shoulders. The thickness differences means they cook at different speeds " +
							"and finish at different times. The best thing to do is separate them into three parts with kitchen shears, " +
							"a sturdy knife, or a Chinese cleaver (my weapon of choice because the ka-chunk noise of chopping them off " +
							"is so very satisfying).",
						"The tips are almost all skin, really thin, and small enough that they often fall through the grates or " +
							"burn to a crisp. You can cook them if you wish, but I freeze them for use in making soup. Separate the " +
							"V shaped piece remaining at the joint between the flat and drumette. You will cook both these parts.",
						"Some folks like to season them with a spice rub. That works most of the time. I find that most commerci" +
							"al rubs are too salty for such thin cuts, and most have too much sugar that tends to burn during the " +
							"crisping phase. Besides, they just get lost under the sauces and dips. So I just season them with salt and " +
							"pepper. As Rachael Ray says: \"Easy peasy.\"",
						"Fire up. You can start them on a smoker if you wish, but I usually grill them. Set up the grill for 2-zone " +
							"cooking with the indirect side at about 325°F (162.8°C) to help crisp the skin and melt the fat. If you " +
							"wish, add wood to the direct side to create smoke. Use a lot of smoke.",
						"Cook. Add the wings to the indirect heat side of the grill and cook with the lid closed until the skins are " +
							"golden. That will probably take about 7 to 10 minutes per side. By then they are pretty close to done.",
						"To crisp the skin, move the wings to the direct heat side of your grill, high heat, lid open, and stand there, " +
							"turning frequently until the skin is dark golden to brown but not burnt, keeping a close eye on the " +
							"skinnier pieces, moving them to the indirect zone when they are done.",
						"Serve. Put the sauce in a big mixing bowl or pot and put it on the grill and get it warm. Stir or whisk well. " +
							"Keep warm. When the wings are done you can serve them with the sauce on the side for dipping, or just dump " +
							"them in with the sauce and toss or stir until they are coated. Then slide them onto a serving platter. Put " +
							"the celery sticks next to them, and serve with a bowl of Blue Cheese Dip. People can scoop some Blue " +
							"Cheese Sauce on their plates, and dip in the celery and wings.",
					},
				},
				Name:     "Crispy Grilled Buffalo Wings Recipe",
				PrepTime: "PT120M",
				Yield:    models.Yield{Value: 6},
				URL:      "https://amazingribs.com/tested-recipes-chicken-recipes-crispy-grilled-buffalo-wings-recipe/",
			},
		},
		{
			name: "ambitiouskitchen.com",
			in:   "https://www.ambitiouskitchen.com/lemon-garlic-salmon/",
			want: models.RecipeSchema{
				AtContext:     atContext,
				AtType:        models.SchemaType{Value: "Recipe"},
				Category:      models.Category{Value: "Dinner"},
				CookTime:      "PT18M",
				Cuisine:       models.Cuisine{Value: "American"},
				DatePublished: "2022-05-05T11:00:05+00:00",
				Description: models.Description{
					Value: "Wonderful honey lemon garlic salmon covered in an easy lemon garlic butter marinade and baked to flaky " +
						"perfection. This flavorful lemon garlic salmon recipe makes a delicious, protein packed dinner served " +
						"with your favorite salad or side dishes, and the marinade is perfect for your go-to proteins.",
				},
				Image: models.Image{
					Value: "https://www.ambitiouskitchen.com/wp-content/uploads/2021/01/Lemon-Garlic-Salmon-5.jpg",
				},
				Ingredients: models.Ingredients{
					Values: []string{
						"1 pound salmon", "2 tablespoons butter, melted",
						"2 tablespoons honey (or sub pure maple syrup)",
						"1 teaspoon dijon mustard, preferably grainy dijon", "½ lemon, juiced",
						"Zest from 1 lemon", "½ teaspoon garlic powder (or 3 cloves garlic, minced)",
						"Freshly ground salt and pepper",
					},
				},
				Instructions: models.Instructions{
					Values: []string{
						"Preheat the oven to 400 degrees F. Line a large baking sheet with parchment paper or foil and grease " +
							"lightly with olive oil or nonstick cooking spray. Place salmon skin side down.",
						"In a medium bowl, whisk together the melted butter, honey, dijon, lemon juice, lemon zest, garlic powder " +
							"and salt and pepper. Generously brush the salmon with the marinade.",
						"Place salmon in the oven and bake for 15-20 minutes or until salmon easily flakes with a fork. Mine is " +
							"always perfect at 16-18 minutes. Enjoy immediately.",
					},
				},
				Keywords: models.Keywords{
					Values: "lemon garlic butter salmon, lemon garlic salmon, lemon honey garlic salmon",
				},
				NutritionSchema: models.NutritionSchema{
					Calories:      "279 kcal",
					Carbohydrates: "9.3 g",
					Fat:           "17.2 g",
					Fiber:         "0.1 g",
					Protein:       "21.3 g",
					SaturatedFat:  "7.6 g",
					Servings:      "1",
					Sugar:         "8.2 g",
				},
				Name:     "Honey Lemon Garlic Salmon",
				PrepTime: "PT10M",
				Yield:    models.Yield{Value: 4},
				URL:      "https://www.ambitiouskitchen.com/lemon-garlic-salmon/",
			},
		},
		{
			name: "archanaskitchen.com",
			in:   "https://www.archanaskitchen.com/karnataka-style-orange-peels-curry-recipe",
			want: models.RecipeSchema{
				AtContext:     atContext,
				AtType:        models.SchemaType{Value: "Recipe"},
				Name:          "Karnataka Style Orange Peels Curry Recipe",
				Cuisine:       models.Cuisine{Value: "Karnataka"},
				DateModified:  "2023-09-02T05:30:01+0000",
				DatePublished: "2017-10-05T00:23:00+0000",
				Description: models.Description{
					Value: "Did you know that we can make a yummy curry out of Orange Peels? It is tangy, sweetish, spicy, slightly bitter and bursting with flavors. It is an unique recipe. So next time you have some guests at home, make this recipe and impress your friends and family. It is filled with flavours and tastes delicious with almost everything. Next time you eat an orange, don't throw the peels, make a curry out of it.\nServe Karnataka Style Orange Peels Curry along with Cabbage Thoran and Whole Wheat Lachha Paratha for your weekday meal. It even tastes great with Steamed Rice.\nIf you like this recipe, you can also try other Karnataka recipes such as\n\nMavina Hannina Gojju Recipe\nMavina Hannina Gojju Recipe\nKarnataka Style Bassaru Palya Recipe",
				},
				Image: models.Image{
					Value: "https://www.archanaskitchen.com/images/archanaskitchen/1-Author/Smitha-Kalluraya/" +
						"Karnataka_style_Orange_Peels_Curry_.jpg",
				},
				PrepTime: "PT15M",
				CookTime: "PT15M",
				Yield:    models.Yield{Value: 4},
				Ingredients: models.Ingredients{
					Values: []string{
						"2 Orange",
						"Tamarind, big lemon size",
						"3 tablespoon Jaggery, adjustable",
						"1 teaspoon Rasam Powder",
						"1 teaspoon Red Chilli powder",
						"1 tablespoon Rice flour",
						"Salt, to taste",
						"1 tablespoon Oil",
						"1 teaspoon Mustard seeds (Rai/ Kadugu)",
						"1/2 teaspoon Cumin seeds (Jeera)",
						"1/4 teaspoon Methi Seeds (Fenugreek Seeds)",
						"1 teaspoon White Urad Dal (Split)",
						"1 Dry Red Chilli, broken",
						"1 Green Chilli, slit",
						"Asafoetida (hing), a pinch",
						"1/4 teaspoon Turmeric powder (Haldi)",
					},
				},
				Instructions: models.Instructions{
					Values: []string{
						"To begin making the Karnataka Style Orange Peels Curry recipe, Peel out 2 oranges. Gently scrape the white " +
							"region from the orange peel using a spoon. Remove as much as possible as this will reduce the bitterness.",
						"Chop the orange peels finely.",
						"Soak tamarind in water. After it's soaked well, squeeze out tamarind water and keep aside.",
						"Heat oil in a heavy bottomed pan and temper with mustard seeds, urad dal, cumin seeds & fenugreek seeds. When " +
							"it splutters add curry leaves, red chilli powder, green chilli, hing and turmeric powder. Mix and add " +
							"chopped orange peel. Fry for 4-5 minutes.",
						"Later add tamarind water, salt, jaggery, red chilli powder and rasam powder. Add some water and " +
							"close the lid. Keep the flame on low medium and allow the peel to cook well. Mix in between.",
						"Meanwhile  in a mixing bowl, mix rice flour in water such that no lumps are formed.",
						"When orange peels are cooked well, add the rice flour mix to the curry and allow it to boil for 1-2 minutes. " +
							"Adding rice flour gives nice thickness to the gravy. If the desired consistency is attained, switch off.",
						"Serve Karnataka Style Orange Peels Curry along with Cabbage Thoran and Whole Wheat Lachha Paratha for your " +
							"weekday meal. It even tastes great with Steamed Rice.",
					},
				},
				Category: models.Category{Value: "Indian Curry Recipes"},
				Keywords: models.Keywords{
					Values: "South Indian Recipes,Indian Lunch Recipes,Orange Recipes,Karnataka Recipes",
				},
				URL: "https://www.archanaskitchen.com/karnataka-style-orange-peels-curry-recipe",
			},
		},
		{
			name: "atelierdeschefs.fr",
			in:   "https://www.atelierdeschefs.fr/fr/recette/17741-boeuf-bourguignon-traditionnel.php",
			want: models.RecipeSchema{
				AtContext:   atContext,
				AtType:      models.SchemaType{Value: "Recipe"},
				Name:        "Bœuf bourguignon traditionnel",
				CookTime:    "P0Y0M0DT0H0M10800S",
				Description: models.Description{Value: "Une vraie recette de la tradition française: des morceaux de bœuf cuits longuement dans un bouillon au vin rouge."},
				Image:       models.Image{Value: "https://adc-dev-images-recipes.s3.eu-west-1.amazonaws.com/bourguignon_3bd.jpg"},
				Ingredients: models.Ingredients{
					Values: []string{
						"1.5 kg Boeuf à braiser ( jumeau, collier, macreuse )",
						"2 pièce(s) Carotte(s)",
						"1 pièce(s) Oignon(s)",
						"30 g Farine de blé",
						"2 pièce(s) Gousse(s) d'ail",
						"1.5 l Vin de Bourgogne",
						"3 cl Huile de tournesol",
						"6 pincée(s) Sel fin",
						"6 tour(s) Moulin à poivre",
						"40 cl Fond de veau",
						"150 g Lardon(s)",
						"150 g Oignon(s) grelot",
						"150 g Champignon(s) de Paris",
						"10 g Sucre en poudre",
						"50 g Beurre doux",
						"3 cl Huile d'olive",
						"0.25 botte(s) Persil plat",
					},
				},
				Instructions: models.Instructions{Values: []string{
					"Couper et dégraisser légèrement la viande. Éplucher et tailler en gros morceaux les carottes et l'oignon. Éplucher et dégermer les gousses d'ail.\nMettre la viande et la garniture dans le vin rouge, et faire mariner toute une nuit au réfrigérateur.",
					"Égoutter la viande et la garniture en conservant le vin. Séparer la garniture et la viande. Effeuiller le persil, conserver les tiges pour la cuisson et les feuilles pour le dressage.\n\nDans une cocotte chaude, mettre l'huile de tournesol et colorer les morceaux de viande environ 1 minute de chaque côté. Ajouter la garniture aromatique, assaisonner de sel fin, puis cuire doucement pendant 3 minutes. Singer (c'est-à-dire ajouter la farine) et cuire à nouveau 1 minute tout en mélangeant pour bien incorporer la farine. Mouiller avec le vin rouge puis avec le fond de veau. Ajouter les tiges de persil et compléter avec de l'eau si nécessaire. Faire bouillir puis baisser le feu et laisser mijoter pendant 2h30.\n\nLorsque la viande est cuite, la retirer de la cocotte. Passer la sauce au chinois pour la filtrer et vérifier sa texture. Si elle est encore trop liquide, la réduire pendant quelques minutes. La goûter et l'assaisonner de sel et de poivre.",
					"Éplucher les champignons au couteau. \n\nDisposer les oignons grelots dans une poêle. Ajouter de l'eau à mi-hauteur, 20 g de beurre et 1 cuillère à soupe de sucre. Couvrir au contact avec un papier sulfurisé et cuire jusqu'à évaporation complète de l'eau. Lorsque le sucre commence à caraméliser, ajouter 1 cuillère à soupe d'eau et bien enrober les oignons de caramel.\n\nDans une casserole d'eau froide, mettre les lardons et faire bouillir pour les blanchir. Bien les égoutter, puis les colorer dans une poêle antiadhésive bien chaude. Réserver ensuite sur du papier absorbant. Dans la même poêle, mettre un filet d'huile d'olive et faire sauter les champignons pour les colorer. Réserver.",
					"Ciseler les feuilles de persil.\nDans un plat, déposer la viande, verser la sauce dessus et disposer les garnitures.",
				}},
				PrepTime: "P0Y0M0DT0H0M1200S",
				URL:      "https://www.atelierdeschefs.fr/fr/recette/17741-boeuf-bourguignon-traditionnel.php",
				Yield:    models.Yield{Value: 6},
			},
		},
		{
			name: "averiecooks.com",
			in:   "https://www.averiecooks.com/slow-cooker-beef-stroganoff/",
			want: models.RecipeSchema{
				AtContext:     atContext,
				AtType:        models.SchemaType{Value: "Recipe"},
				Category:      models.Category{Value: "Slow Cooker"},
				CookTime:      "PT8H",
				Cuisine:       models.Cuisine{Value: "American"},
				DatePublished: "2022-03-15",
				Description: models.Description{
					Value: "A comfort food classic that everyone in the family LOVES! Hearty chunks of beef, rich and flavorful beef gravy, and served over a bed of warm noodles to soak up all that goodness! The EASIEST recipe for beef stroganoff ever because your Crock-Pot truly does all the work! Set it and forget it!",
				},
				Image: models.Image{
					Value: "https://www.averiecooks.com/wp-content/uploads/2022/03/beefstroganoff-13-480x480.jpg",
				},
				Ingredients: models.Ingredients{
					Values: []string{
						"2 pounds beef stew meat or beef chuck, diced into large bite-sized pieces or chunks",
						"1/2 cup white onion, diced small", "3 to 5 cloves garlic, finely minced",
						"1 teaspoon salt", "1 teaspoon freshly ground pepper",
						"1 teaspoon beef bouillon",
						"2 to 3 sprig fresh thyme OR 1/2 teaspoon dried thyme",
						"two 10-ounce cans cream of mushroom soup",
						"2 cups low sodium beef broth, plus more if desired",
						"1 tablespoon Dijon mustard, optional", "1 tablespoon Worcestershire sauce",
						"1/2 cup heavy cream, optional for a creamier sauce; at room temperature",
						"12 ounces wide egg noodles, cooked according to package directions (or your favorite pasta or mashed potatoes)",
						"Fresh parsley, finely minced; optional for garnishing",
					},
				},
				Instructions: models.Instructions{
					Values: []string{
						"To a large 7 to 8-quart slow cooker, add the beef, onion, garlic, salt, pepper, beef bouillon, thyme, " +
							"and stir to combine; set aside.",
						"To a medium bowl, add the mushroom soup, beef broth, optional Dijon, Worcestershire sauce, and whisk " +
							"to combine.",
						"Pour the liquid over the contents in the slow cooker, stir to combine, cover the the lid, and cook on " +
							"high for 4 to 5 hours OR on low for 7 to 8 hours, or until done. Tip - At any time the beef " +
							"is slow cooking and you feel like it needs a bit more beef broth, it's fine to add a bit more, to taste.",
						"In the last 15 minutes of slow cooking, cook the egg noodles in a pot of boiling water according to " +
							"package directions; drain and set aside.* (See Notes about why I don't cook the noodles in " +
							"the slow cooker and cook them separately.)",
						"Optionally, if you want a creamier sauce, after the beef stroganoff has cooked and has cooled a bit " +
							"meaning it's not boiling nor bubbling, slowly you can add 1/2-cup heavy cream at room temperature " +
							"while whisking vigorously as you add it. Tip - Do NOT add cold cream to hot beef liquid because the " +
							"dairy proteins can separate, or break, and you will end up with a horribly ugly looking sauce after hours " +
							"and hours slow cooking. So make sure the cream is at room temp and the beef stroganoff has cooled a bit " +
							"if you are adding cream.",
						"Plate a bed of noodles, top with the beef and gravy mixture, and serve immediately. Extra beef " +
							"stroganoff will keep airtight in the fridge for up to 5 days and in the freezer for up to 4 months. " +
							"Tip - Because the noodles will continue to absorb moisture, including any of the beef gravy or sauce, " +
							"it's very important to store the beef mixture and the noodles separately in the fridge or freezer.",
					},
				},
				Name: "Slow Cooker Beef Stroganoff",
				NutritionSchema: models.NutritionSchema{
					Calories:       "575 calories",
					Carbohydrates:  "30 grams carbohydrates",
					Cholesterol:    "191 milligrams cholesterol",
					Fat:            "23 grams fat",
					Fiber:          "2 grams fiber",
					Protein:        "62 grams protein",
					SaturatedFat:   "11 grams saturated fat",
					Servings:       "1",
					Sodium:         "1288 milligrams sodium",
					Sugar:          "6 grams sugar",
					TransFat:       "1 grams trans fat",
					UnsaturatedFat: "12 grams unsaturated fat",
				},
				Keywords: models.Keywords{
					Values: "slow cooker beef stroganoff, easy beef stroganoff recipe, beef stroganoff with cream of mushroom " +
						"soup, beef stroganoff with stew meat, slow cooker beef stroganoff recipe",
				},
				PrepTime: "PT10M",
				Yield:    models.Yield{Value: 6},
				URL:      "https://www.averiecooks.com/slow-cooker-beef-stroganoff/",
			},
		},
		{
			name: "bakingmischief.com",
			in:   "https://bakingmischief.com/italian-roasted-potatoes/",
			want: models.RecipeSchema{
				AtContext:     atContext,
				AtType:        models.SchemaType{Value: "Recipe"},
				Category:      models.Category{Value: "Side"},
				CookTime:      "PT40M",
				Cuisine:       models.Cuisine{Value: "American"},
				DatePublished: "2022-02-23T01:00:11+00:00",
				Description: models.Description{
					Value: "These 3-ingredient Italian roasted potatoes are quick and simple to prep. With crispy edges " +
						"and creamy centers, they make an easy side dish that everyone will love.",
				},
				Keywords: models.Keywords{Values: "Italian roasted potatoes"},
				Image: models.Image{
					Value: "https://bakingmischief.com/wp-content/uploads/2021/11/italian-roasted-potatoes-image-square-3.jpg",
				},
				Ingredients: models.Ingredients{
					Values: []string{
						"2 pounds red potatoes (*)", "3 tablespoons olive oil",
						"2 teaspoons Italian seasoning (*)", "½ teaspoon salt",
					},
				},
				Instructions: models.Instructions{
					Values: []string{
						"Preheat oven to 425°F. Scrub potatoes, remove any blemishes, and cut them into 2-inch pieces.",
						"Pile potatoes on a baking sheet and drizzle with olive oil and sprinkle Italian seasoning and salt " +
							"over the top. Toss until well-coated. Arrange potatoes so they are evenly spaced over the " +
							"baking sheet with a cut side down.",
						"Cover the tray tightly with foil and bake on the center rack of your oven for 15 minutes.",
						"Remove and discard the foil. Raise the oven temperature to 475°F.",
						"Continue to bake the potatoes uncovered for 25 to 30 minutes, rotating the pan once halfway through, " +
							"until the potatoes are fork-tender.", "Remove from the oven, add additional salt if needed, and enjoy!",
					},
				},
				Name: "Italian Roasted Potatoes",
				NutritionSchema: models.NutritionSchema{
					Calories: "194 kcal",
					Servings: "1",
				},
				PrepTime: "PT10M",
				Yield:    models.Yield{Value: 6},
				URL:      "https://bakingmischief.com/italian-roasted-potatoes/",
			},
		},
		{
			name: "baking-sense.com",
			in:   "https://www.baking-sense.com/2022/02/23/irish-potato-farls/",
			want: models.RecipeSchema{
				AtContext:     atContext,
				AtType:        models.SchemaType{Value: "Recipe"},
				Category:      models.Category{Value: "Breakfast"},
				CookTime:      "PT10M",
				Cuisine:       models.Cuisine{Value: "Irish"},
				DatePublished: "2022-02-23T10:00:00+00:00",
				Description: models.Description{
					Value: "Have you heard of Irish Potato Farls? No? Well, if you love potato pancakes, you&#39;ll love " +
						"potato farls. They&#39;re easy to make with fresh or left over potatoes.",
				},
				Image: models.Image{
					Value: "https://www.baking-sense.com/wp-content/uploads/2022/02/potato-farls-featured.jpg",
				},
				Ingredients: models.Ingredients{
					Values: []string{
						"24 oz russet potatoes (peeled and cut into 1&quot; cubes.)",
						"1 1/2 teaspoons table salt (divided)",
						"2 oz Irish Butter (room temperature, divided)",
						"3.75 oz all purpose flour (3/4 cup)", "1/2 teaspoon baking powder",
						"1/4 teaspoon ground black pepper", "2 scallions (chopped fine)",
					},
				},
				Instructions: models.Instructions{
					Values: []string{
						"Place the potatoes in a pot of water with 1/2 teaspoon of the salt. Boil until the potatoes are tender. Drain.",
						"Pass the potatoes through a ricer back into the pot, or use a potato masher. Add 1 oz (2 tablespoons) " +
							"of the butter and the remaining salt mix until the butter is melted. Add the flour, baking powder " +
							"and pepper and stir until most of the flour is mixed in.",
						"Turn the dough out onto a lightly floured surface and knead in the chopped scallions then form the dough " +
							"into a ball. Divide the dough in half.",
						"Preheat a large skillet over medium heat. While the pan is heating, pat each half of the dough to " +
							"an 8\" round, 1/4”thick. Use flour as needed to prevent sticking. Cut the rounds into quarters. " +
							"You&#39;ll have a total of 8 farls. (See Note)",
						"Melt 1 tablespoon of the remaining butter in the pan. Fry half the farls in the butter until golden brown, " +
							"then flip and fry the other side. Cook until both sides are golden brown and the farl springs " +
							"back when pressed in the center. About 4 minutes per side.",
						"Repeat with the remaining butter and farls. Serve immediately.",
					},
				},
				Keywords: models.Keywords{Values: "pancake, potato"},
				Name:     "Irish Potato Farls Recipe",
				NutritionSchema: models.NutritionSchema{
					Calories:       "165 kcal",
					Carbohydrates:  "25 g",
					Cholesterol:    "15 mg",
					Fat:            "6 g",
					Fiber:          "2 g",
					Protein:        "3 g",
					SaturatedFat:   "4 g",
					Servings:       "1",
					Sodium:         "519 mg",
					Sugar:          "1 g",
					TransFat:       "1 g",
					UnsaturatedFat: "3 g",
				},
				PrepTime: "PT30M",
				Yield:    models.Yield{Value: 8},
				URL:      "https://www.baking-sense.com/2022/02/23/irish-potato-farls/",
			},
		},
		{
			name: "bbc.co.uk",
			in:   "https://www.bbc.co.uk/food/recipes/healthy_sausage_16132",
			want: models.RecipeSchema{
				AtContext: atContext,
				AtType:    models.SchemaType{Value: "Recipe"},
				Category:  models.Category{Value: "Main course"},
				CookTime:  "PT2H",
				Description: models.Description{
					Value: "Never thought you’d hear the words ‘healthy sausage casserole’? Well, here it is. This " +
						"all-in-one-dish dinner is packed with veggies for a perfect midweek meal.\r\n\r\nEach " +
						"serving provides 348 kcal, 34g protein, 33.5g carbohydrates (of which 18g sugars), 7g " +
						"fat (of which 2g saturates), 9.5g fibre and 1.8g salt.",
				},
				Keywords: models.Keywords{
					Values: "absolute bangers, 400-calorie dinners, cheap stews , comfort food on a budget, easy healthy dinner ideas, easy sausage suppers , healthy and filling, healthy british classics, healthy comfort food, healthy dinner, healthy family meals, healthy meals on a budget, healthy winter food, low-calorie comfort food, low-calorie, making meat go further, sausage suppers, summery sausages, the best sausage, winter stew, autumn, bonfire night, easy family dinners, winter, sausage casserole, sausage, healthy",
				},
				Image: models.Image{
					Value: "https://food-images.files.bbci.co.uk/food/recipes/healthy_sausage_16132_16x9.jpg",
				},
				Ingredients: models.Ingredients{
					Values: []string{
						"2 red peppers, seeds removed, cut into chunks",
						"2 carrots, cut into thick slices",
						"2 red onions, cut into wedges",
						"5 garlic cloves, finely sliced",
						"8 lean sausages",
						"400g tin peeled cherry tomatoes",
						"400g tin chickpeas, drained",
						"200ml/7fl oz vegetable stock",
						"1 green chilli, seeds removed, chopped",
						"1 tsp paprika",
						"2 tsp French mustard",
						"100g/3½oz frozen mixed vegetables",
					},
				},
				Instructions: models.Instructions{
					Values: []string{
						"Preheat the oven to 220C/200C Fan/Gas 7.",
						"Put the peppers, carrots, onions and garlic into a deep baking dish and roast for 20 minutes. Add the " +
							"sausages and roast for a further 10 minutes.",
						"Turn the oven down to 200C/180C Fan/Gas 6. Pour the tomatoes and chickpeas into the baking dish, then " +
							"stir in the stock, chilli and paprika. Bake for another 35 minutes.",
						"Stir in the mustard and the frozen mixed veg and return to the oven for 5 minutes. Leave to rest for " +
							"10 minutes before serving.",
					},
				},
				Name: "Healthy sausage casserole",
				NutritionSchema: models.NutritionSchema{
					Calories:      "348kcal",
					Carbohydrates: "33.5g",
					Fat:           "7g",
					Fiber:         "9.5g",
					Protein:       "34g",
					SaturatedFat:  "2g",
					Sugar:         "18g",
				},
				PrepTime: "PT30M",
				Yield:    models.Yield{Value: 4},
				URL:      "https://www.bbc.co.uk/food/recipes/healthy_sausage_16132",
			},
		},
		{
			name: "bbcgoodfood.com",
			in:   "https://www.bbcgoodfood.com/recipes/three-cheese-risotto",
			want: models.RecipeSchema{
				AtContext:     atContext,
				AtType:        models.SchemaType{Value: "Recipe"},
				Category:      models.Category{Value: "Dinner, Main course, Side dish, Supper"},
				CookTime:      "PT35M",
				Cuisine:       models.Cuisine{Value: "Italian"},
				DateModified:  "2022-11-08T14:35:36+00:00",
				DatePublished: "2014-12-04T12:17:29+00:00",
				Description: models.Description{
					Value: "Tom Kerridge's indulgently rich and cheesy risotto makes an extra-special side dish for a celebration dinner party",
				},
				Keywords: models.Keywords{
					Values: "cheesy risotto, Indulgent, rice side dish, risotto side dish, Tom Kerridge, Winter",
				},
				Image: models.Image{
					Value: "https://images.immediate.co.uk/production/volatile/sites/30/2020/08/roast-poussin-with-wild-mushroom-sauce_0-8051af1.jpg?resize=768,574",
				},
				Ingredients: models.Ingredients{
					Values: []string{
						"25g butter",
						"1 tbsp olive oil",
						"1 onion , finely chopped",
						"2 garlic cloves , finely grated",
						"200g risotto rice",
						"200ml white wine",
						"800ml warm chicken stock",
						"50g fresh parmesan (or vegetarian alternative)",
						"½ ball of mozzarella , diced",
						"pinch of cayenne pepper , to taste",
						"2 tbsp mascarpone",
					},
				},
				Instructions: models.Instructions{
					Values: []string{
						"Melt the butter and olive oil together in a large, shallow saucepan, add the onion and garlic, and cook for 5-10 mins until soft. Add the risotto rice and cook for 2-3 mins, getting a good covering in the fats and giving the rice a slightly toasted flavour.",
						"Add the white wine and cook until it has reduced away. Add the warm chicken stock, a ladleful at a time, and stir into the rice – when it has been absorbed, add more. You may not need to add all the stock, but keep adding until the rice is cooked al dente. It will take around 15 mins to get the risotto to the right consistency.",
						"Take the rice pan off the heat and stir in the cheeses, season and leave to rest for 3-4 mins. Serve with the roasted poussins, morel sauce and some wilted Baby Gem lettuce leaves.",
					},
				},
				Name: "Three-cheese risotto",
				NutritionSchema: models.NutritionSchema{
					Calories:      "451 calories",
					Carbohydrates: "42 grams carbohydrates",
					Fat:           "20 grams fat",
					Fiber:         "2 grams fiber",
					Protein:       "17 grams protein",
					SaturatedFat:  "12 grams saturated fat",
					Sodium:        "0.9 milligram of sodium",
					Sugar:         "4 grams sugar",
				},
				PrepTime: "PT20M",
				URL:      "https://www.bbcgoodfood.com/recipes/three-cheese-risotto",
			},
		},
		{
			name: "bettycrocker.com",
			in:   "https://www.bettycrocker.com/recipes/spinach-mushroom-quiche/ed3014db-7810-41d6-8e1c-cd4eed7b1db3",
			want: models.RecipeSchema{
				AtContext:    atContext,
				AtType:       models.SchemaType{Value: "Recipe"},
				Category:     models.Category{Value: "Breakfast"},
				Cuisine:      models.Cuisine{Value: "French"},
				DateCreated:  "2011-10-05",
				DateModified: "2013-04-03",
				Description: models.Description{
					Value: "Bisquick® Gluten Free mix crust topped with spinach and mushroom mixture for a tasty breakfast – " +
						"perfect if you love French cuisine.",
				},
				Keywords: models.Keywords{Values: "spinach mushroom quiche"},
				Image: models.Image{
					Value: "https://mojo.generalmills.com/api/public/content/MAHdJv1NBUeLl4-jtMq24g_gmi_hi_res_jpeg.jpeg%3Fv=2e1b9203&t=b5673970ed9e41549a020b29d456506d",
				},
				Ingredients: models.Ingredients{
					Values: []string{
						"1 cup Bisquick™ Gluten Free mix",
						"1/3 cup plus 1 tablespoon shortening",
						"3 to 4 tablespoons cold water",
						"1 tablespoon butter",
						"1 small onion, chopped (1/3 cup)",
						"1 1/2 cups sliced fresh mushrooms (about 4 oz)",
						"4 eggs",
						"1 cup milk",
						"1/8 teaspoon ground red pepper (cayenne)",
						"3/4 cup coarsely chopped fresh spinach",
						"1/4 cup chopped red bell pepper",
						"1 cup gluten-free shredded Italian cheese blend (4 oz)",
					},
				},
				Instructions: models.Instructions{
					Values: []string{
						"Heat oven to 425°F. In medium bowl, place Bisquick mix. Cut in shortening, using pastry blender " +
							"(or pulling 2 table knives through ingredients in opposite directions), until particles are size " +
							"of small peas. Sprinkle with cold water, 1 tablespoon at a time, tossing with fork until all flour " +
							"is moistened and pastry almost leaves side of bowl (1 to 2 teaspoons more water can be added if necessary).",
						"Press pastry in bottom and up side of ungreased 9-inch quiche dish or glass pie plate. Bake 12 to " +
							"14 minutes or until crust just begins to brown and is set. Reduce oven temperature to 325°F.",
						"Meanwhile, in 10-inch skillet, melt butter over medium heat. Cook onion and mushrooms in butter about " +
							"5 minutes, stirring occasionally, until tender. In medium bowl, beat eggs, milk and red pepper until " +
							"well blended. Stir in spinach, bell pepper, mushroom mixture and cheese. Pour into partially baked crust.",
						"Bake 40 to 45 minutes or until knife inserted in center comes out clean. Let stand 10 minutes before cutting.",
					},
				},
				Name: "Spinach Mushroom Quiche",
				NutritionSchema: models.NutritionSchema{
					Calories:       "260",
					Carbohydrates:  "16 g",
					Cholesterol:    "120 mg",
					Fat:            "2 ",
					Fiber:          "0 g",
					Protein:        "9 g",
					SaturatedFat:   "6 g",
					Servings:       "1",
					Sodium:         "340 mg",
					Sugar:          "4 g",
					TransFat:       "2 g",
					UnsaturatedFat: "",
				},
				PrepTime: "PT0H30M",
				Yield:    models.Yield{Value: 8},
				URL:      "https://www.bettycrocker.com/recipes/spinach-mushroom-quiche/ed3014db-7810-41d6-8e1c-cd4eed7b1db3",
			},
		},
		{
			name: "bigoven.com",
			in:   "https://www.bigoven.com/recipe/vegetable-tempura-japanese/19344",
			want: models.RecipeSchema{
				AtContext:     atContext,
				AtType:        models.SchemaType{Value: "Recipe"},
				Category:      models.Category{Value: "Main Dish"},
				CookTime:      "PT30M",
				Cuisine:       models.Cuisine{Value: "Japanese"},
				DatePublished: "2004/01/01",
				Description:   models.Description{Value: "not set"},
				Keywords: models.Keywords{
					Values: "nrm, side dish, snacks, vegetables, fry, fall, spring, summer, winter, meatless, vegetarian, " +
						"japanese,  qeethnic, contains white meat, nut free, contains gluten, red meat free, shellfish " +
						"free, contains eggs, dairy free",
				},
				Image: models.Image{
					Value: "https://bigoven-res.cloudinary.com/image/upload/h_320,w_320,c_fill/vegetable-tempura-japanese-e79b5b.jpg",
				},
				Ingredients: models.Ingredients{
					Values: []string{
						"2 cup All purpose flour sifted",
						"1 teaspoon Salt",
						"1/8 teaspoon Baking soda",
						"1 large Egg yolk",
						"2 cup Ice water",
						"Vegetable oil for frying",
						"2 medium Zucchini sliced thin",
						"1 medium Green pepper cut into strips",
						"1 large Onion sliced",
						"1/2 pound Button mushrooms",
						"1 cup Broccoli",
					},
				},
				Instructions: models.Instructions{
					Values: []string{
						"Separate the onion into rings. Steam the broccoli 5 minutes (or microwave a few minutes).  In an electric blender combine the flour, salt, baking soda, egg yolk, and water. Blend to mix. Let stand 15 minutes.  Heat 3 - 4 inches of oil in a deep heavy kettle, deep-fat fryer, or electric wok until it registers 375F / 190C on a deep-fat thermometer. Test batter consistency by dipping one piece of vegetable and letting excess drip off. There should be a light coating left on.  Dip and fry, a few at a time, in the hot oil until golden. Drain on paper towels and keep warm in the oven heated to 250F / 130C / Gas Mark  until all are cooked.",
					},
				},
				Name: "Vegetable Tempura - Japanese",
				NutritionSchema: models.NutritionSchema{
					Calories:      "300 calories",
					Carbohydrates: "60.0488662958771 g",
					Cholesterol:   "52.445 mg",
					Fat:           "2.32476678639706 g",
					Fiber:         "5.44774018914959 g",
					Protein:       "11.4823529627363 g",
					SaturatedFat:  "0.619135359073004 g",
					Servings:      "1",
					Sodium:        "50.8475852783613 mg",
					Sugar:         "54.6011261067275 g",
					TransFat:      "0.476087607733719 g",
				},
				PrepTime: "PT30M",
				Yield:    models.Yield{Value: 4},
				URL:      "https://www.bigoven.com/recipe/vegetable-tempura-japanese/19344",
			},
		},
		{
			name: "bonappetit.com",
			in:   "https://www.bonappetit.com/recipe/crispy-chicken-with-zaatar-olive-rice",
			want: models.RecipeSchema{
				AtContext:     atContext,
				AtType:        models.SchemaType{Value: "Recipe"},
				DateModified:  "2022-03-15T05:00:00.000-04:00",
				DatePublished: "2022-03-15T05:00:00.000-04:00",
				Description: models.Description{
					Value: "Give ground chicken the respect it deserves.",
				},
				Keywords: models.Keywords{
					Values: "main,dinner,quick,one-pot meals,easy,weeknight meals,healthyish,gluten-free,nut-free,ground " +
						"chicken,ground turkey,feta,olive,rice,za'atar,spinach,kale,sauté,swiss chard,castelvetrano olive,web",
				},
				Image: models.Image{
					Value: "https://assets.bonappetit.com/photos/6228bc8071b26c82f857f620/16:9/w_6208,h_3492,c_limit/Crispy-Chicken-With-Za%E2%80%99atar-Olive-Rice.jpg",
				},
				Ingredients: models.Ingredients{
					Values: []string{
						"1 lb. ground chicken or turkey",
						"½ tsp. smoked paprika",
						"1 tsp. Diamond Crystal or ½ tsp. Morton kosher salt, plus more",
						"Freshly ground black pepper",
						"3 Tbsp. extra-virgin olive oil",
						"1 cup Castelvetrano olives, smashed, pits removed",
						"3 cups cooked rice",
						"1 Tbsp. za’atar, plus more for serving",
						"2 cups coarsely chopped greens (such as spinach, kale, or chard)",
						"Zest and juice of 1 small lemon",
						"2 oz. feta, thinly sliced into planks",
						"Coarsely chopped dill (for serving)",
					},
				},
				Instructions: models.Instructions{
					Values: []string{
						"Place chicken in a medium bowl. Sprinkle paprika and 1 tsp. Diamond Crystal or ½ tsp. Morton kosher " +
							"salt over chicken; season with pepper. Gently mix with your hands to combine.",
						"Heat oil in a large nonstick skillet over medium-high. Arrange chicken in pan in a thin, even layer " +
							"and cook, undisturbed, until golden brown and crisp underneath, about 5 minutes. " +
							"Continue to cook, stirring and breaking up into bite-size pieces with a wooden " +
							"spoon, until cooked through, about 1 minute. Using a slotted spoon, transfer chicken to a " +
							"plate, leaving oil and fat behind.",
						"Add olives to same pan and cook, undisturbed, until heated through and blistered, 1–2 minutes. Add " +
							"rice and 1 Tbsp. za’atar and cook, stirring often, until slightly crisp, about 3 " +
							"minutes. Add greens and lemon juice and cook, stirring occasionally, until greens " +
							"are wilted, about 2 minutes. Remove pan from heat; stir in lemon zest, feta, and chicken. " +
							"Taste and season with more salt and pepper if needed.",
						"Transfer chicken and rice to a large shallow bowl; sprinkle with more za’atar and top with dill.",
					},
				},
				Name:  "Crispy Chicken With Za’atar-Olive Rice",
				Yield: models.Yield{Value: 4},
				URL:   "https://www.bonappetit.com/recipe/crispy-chicken-with-zaatar-olive-rice"},
		},
		{
			name: "bowlofdelicious.com",
			in:   "https://www.bowlofdelicious.com/mini-meatloaves/",
			want: models.RecipeSchema{
				AtContext:     atContext,
				AtType:        models.SchemaType{Value: "Recipe"},
				Category:      models.Category{Value: "Dinner"},
				CookTime:      "PT25M",
				Cuisine:       models.Cuisine{Value: "American"},
				DatePublished: "2021-01-19T11:41:00+00:00",
				Description: models.Description{
					Value: "Mini Meatloaves cook up in HALF the time of a whole meatloaf - an easy, fast gluten-free recipe " +
						"made with oats instead of breadcrumbs.",
				},
				Keywords: models.Keywords{Values: "Mini meatloaves"},
				Image: models.Image{
					Value: "https://www.bowlofdelicious.com/wp-content/uploads/2014/09/Mini-Meatloaves-square.jpg",
				},
				Ingredients: models.Ingredients{
					Values: []string{
						"1 onion (grated)",
						"1 cup quick oats",
						"1/2 cup milk",
						"2 eggs",
						"1 teaspoon kosher salt",
						"1/2 teaspoon black pepper",
						"2 teaspoons soy sauce ((gluten-free if necessary))",
						"2 lbs. ground beef (preferably 80/20)",
						"ketchup or barbecue sauce (for topping, optional)",
					},
				},
				Instructions: models.Instructions{
					Values: []string{
						"Preheat oven to 400 degrees F. Butter or grease a rimmed baking sheet or casserole dish.",
						"In a large bowl, add the grated onion, quick oats (1 cup), milk (1/2 cup), 2 eggs, kosher salt " +
							"(1 teaspoon), black pepper (1/2 teaspoon), and soy sauce (2 teaspoons) and mix until " +
							"well combined. Allow to sit for 3-5 minutes while the oats absorb some of the moisture.",
						"Add the ground beef (2 lbs.) to the bowl. Mix well- for best results, use hands.",
						"Divide the mixture into eight parts and form them into loaf shapes. (I like to flatten the mixture " +
							"in the bowl and use a knife to portion it out into 8 \"wedges,\" kind of like " +
							"slicing a cake, to get even amounts).",
						"At this point, you can wrap the mini meatloaves in plastic wrap or flash freeze to store in the freezer, or" +
							" refrigerate until you&#039;re ready to bake (see notes for more info on this). Otherwise, proceed to cooking.",
						"Place the mini meatloaves on the prepared baking sheet or casserole dish (as many as you want to cook - " +
							"if frozen, defrost completely before cooking). Bake at 400 degrees F for 25 minutes (or until " +
							"the internal temperature is 160 degrees).",
					},
				},
				Name: "Mini Meatloaves",
				NutritionSchema: models.NutritionSchema{
					Calories:      "356 kcal",
					Carbohydrates: "10 g",
					Cholesterol:   "121 mg",
					Fat:           "25 g",
					Fiber:         "1 g",
					Protein:       "23 g",
					SaturatedFat:  "9 g",
					Servings:      "1",
					Sodium:        "474 mg",
					Sugar:         "2 g",
					TransFat:      "1 g",
				},
				PrepTime: "PT10M",
				Yield:    models.Yield{Value: 8},
				URL:      "https://www.bowlofdelicious.com/mini-meatloaves/",
			},
		},
		{
			name: "budgetbytes.com",
			in:   "https://www.budgetbytes.com/easy-vegetable-stir-fry/",
			want: models.RecipeSchema{
				AtContext:     atContext,
				AtType:        models.SchemaType{Value: "Recipe"},
				Category:      models.Category{Value: "Dinner"},
				CookTime:      "PT10M",
				Cuisine:       models.Cuisine{Value: "Asian"},
				DatePublished: "2022-03-15T07:48:39+00:00",
				Description: models.Description{
					Value: "Vegetable stir fry is a quick and easy option for dinner, plus it&#039;s super flexible and a " +
						"great way to use up leftovers from your fridge!",
				},
				Keywords: models.Keywords{Values: "Stir Fry Recipe, vegetable stir fry"},
				Image: models.Image{
					Value: "https://www.budgetbytes.com/wp-content/uploads/2022/03/Easy-Vegetable-Stir-Fry-close.jpg",
				},
				Ingredients: models.Ingredients{
					Values: []string{
						"1/4 cup soy sauce ($0.24)",
						"1/4 cup water ($0.00)",
						"2 Tbsp brown sugar ($0.08)",
						"1 tsp toasted sesame oil ($0.10)",
						"2 cloves garlic, minced ($0.16)",
						"1 tsp grated fresh ginger ($0.10)",
						"1 Tbsp cornstarch ($0.03)",
						"3/4 lb. broccoli ($1.34)",
						"2 carrots ($0.33)",
						"8 oz. mushrooms ($1.69)",
						"8 oz. sugar snap peas ($2.99)",
						"1 small onion ($0.28)",
						"1 red bell pepper ($1.50)", "2 Tbsp cooking oil ($0.16)",
						"1 tsp sesame seeds ($0.06)",
					},
				},
				Instructions: models.Instructions{
					Values: []string{
						"Make the stir fry sauce first. Combine the soy sauce, water, brown sugar, sesame oil, garlic, ginger, " +
							"and cornstarch in a small bowl. Set the sauce aside.",
						"Chop the vegetables into similar-sized pieces. It&#39;s up to you whether you slice, dice, or cut into " +
							"any other shape you prefer.",
						"Add the cooking oil to a very large skillet or wok. Heat over medium-high. When the pan and oil are very " +
							"hot (but not smoking), add the hardest vegetables first: carrots and broccoli. Cook and " +
							"stir for about a minute, or just until the broccoli begins to turn bright green.",
						"Next, add the mushrooms and sugar snap peas. Continue to cook and stir for a minute or two more, or just " +
							"until the mushrooms begin to soften.",
						"Finally, add the softest vegetables, bell pepper and onion. Continue to cook and stir just until the onion " +
							"begins to soften.",
						"Give the stir fry sauce another brief stir, then pour it over the vegetables. Continue to cook and stir " +
							"until the sauce begins to simmer, at which point it will thicken and turn glossy. Remove " +
							"the vegetables from the heat, or continue to cook until they are to your desired doneness.",
						"Top the stir fry with sesame seeds and serve!",
					},
				},
				Name: "Easy Vegetable Stir Fry",
				NutritionSchema: models.NutritionSchema{
					Calories:       "209 kcal",
					Carbohydrates:  "27 g",
					Cholesterol:    "",
					Fat:            "9 g",
					Fiber:          "6 g",
					Protein:        "8 g",
					SaturatedFat:   "",
					Sodium:         "869 mg",
					Sugar:          "",
					TransFat:       "",
					UnsaturatedFat: ""},
				PrepTime: "PT15M",
				Tools:    models.Tools{Values: []string(nil)},
				Yield:    models.Yield{Value: 4},
				URL:      "https://www.budgetbytes.com/easy-vegetable-stir-fry/"},
		},
		{
			name: "cafedelites.com",
			in:   "https://cafedelites.com/butter-chicken",
			want: models.RecipeSchema{
				AtContext:     atContext,
				AtType:        models.SchemaType{Value: "Recipe"},
				Category:      models.Category{Value: "Dinner"},
				CookTime:      "PT30M",
				Cuisine:       models.Cuisine{Value: "Indian"},
				DatePublished: "2019-01-21T19:09:20+00:00",
				Description: models.Description{
					Value: "Butter Chicken is creamy and easy to make right at home in one pan with simple ingredients! Full of incredible flavours, it rivals any Indian restaurant! Aromatic golden chicken pieces in an incredible creamy curry sauce, this Butter Chicken recipe is one of the best you will try!",
				},
				Keywords: models.Keywords{Values: "butter chicken"},
				Image: models.Image{
					Value: "https://cafedelites.com/wp-content/uploads/2019/01/Butter-Chicken-IMAGE-64.jpg",
				},
				Ingredients: models.Ingredients{
					Values: []string{
						"28 oz (800g) boneless and skinless chicken thighs or breasts (cut into bite-sized pieces)",
						"1/2 cup plain yogurt",
						"1 1/2 tablespoons minced garlic",
						"1 tablespoon minced ginger ((or finely grated))",
						"2 teaspoons garam masala",
						"1 teaspoon turmeric",
						"1 teaspoon ground cumin",
						"1 teaspoon red chili powder",
						"1 teaspoon of salt",
						"2 tablespoons olive oil",
						"2 tablespoons ghee ((or 1 tbs butter + 1 tbs oil))",
						"1 large onion, (sliced or chopped)",
						"1 1/2 tablespoons garlic, (minced)",
						"1 tablespoon ginger, (minced or finely grated)",
						"1 1/2 teaspoons ground cumin",
						"1 1/2 teaspoons garam masala",
						"1 teaspoon ground coriander",
						"14 oz (400 g) crushed tomatoes",
						"1 teaspoon red chili powder ((adjust to your taste preference))",
						"1 1/4 teaspoons salt ((or to taste))",
						"1 cup of heavy or thickened cream ((or evaporated milk to save calories))",
						"1 tablespoon sugar",
						"1/2 teaspoon kasoori methi ((or dried fenugreek leaves))",
					},
				},
				Instructions: models.Instructions{
					Values: []string{
						"In a bowl, combine chicken with all of the ingredients for the chicken marinade; let marinate for 30 minutes to an hour (or overnight if time allows).",
						"Heat oil in a large skillet or pot over medium-high heat. When sizzling, add chicken pieces in batches of two or three, making sure not to crowd the pan. Fry until browned for only 3 minutes on each side. Set aside and keep warm. (You will finish cooking the chicken in the sauce.)",
						"Heat butter or ghee in the same pan. Fry the onions until they start to sweat (about 6 minutes) while scraping up any browned bits stuck on the bottom of the pan.",
						"Add garlic and ginger and sauté for 1 minute until fragrant, then add ground coriander, cumin and garam masala. Let cook for about 20 seconds until fragrant, while stirring occasionally.",
						"Add crushed tomatoes, chili powder and salt. Let simmer for about 10-15 minutes, stirring occasionally until sauce thickens and becomes a deep brown red colour.",
						"Remove from heat, scoop mixture into a blender and blend until smooth. You may need to add a couple tablespoons of water to help it blend (up to 1/4 cup). Work in batches depending on the size of your blender.",
						"Pour the puréed sauce back into the pan. Stir the cream, sugar and crushed kasoori methi (or fenugreek leaves) through the sauce. Add the chicken with juices back into the pan and cook for an additional 8-10 minutes until chicken is cooked through and the sauce is thick and bubbling.",
						"Garnish with chopped cilantro and serve with fresh, hot garlic butter rice and fresh homemade Naan bread!",
					},
				},
				Name: "Butter Chicken",
				NutritionSchema: models.NutritionSchema{
					Calories:      "580 kcal",
					Carbohydrates: "17 g",
					Cholesterol:   "250 mg",
					Fat:           "41 g",
					Fiber:         "3 g",
					Protein:       "36 g",
					SaturatedFat:  "19 g",
					Servings:      "1",
					Sodium:        "1601 mg",
					Sugar:         "8 g",
				},
				PrepTime: "PT15M",
				Yield:    models.Yield{Value: 5},
				URL:      "https://cafedelites.com/butter-chicken"},
		},
		{
			name: "castironketo.com",
			in:   "https://www.castironketo.net/blog/balsamic-mushrooms-with-herbed-veggie-mash/",
			want: models.RecipeSchema{
				AtContext:     atContext,
				AtType:        models.SchemaType{Value: "Recipe"},
				Category:      models.Category{Value: "Main Course"},
				CookTime:      "PT30M",
				Cuisine:       models.Cuisine{Value: "Italian"},
				DatePublished: "2022-03-06T03:40:00+00:00",
				Description: models.Description{
					Value: "This easy low-carb dinner is perfect for plant-based eaters or anyone looking to add more veggies " +
						"to their diet!",
				},
				Image: models.Image{
					Value: "https://www.castironketo.net/wp-content/uploads/2022/03/" +
						"Balsamic-Mushrooms-with-Herbed-Veggie-Mash.jpg-p1ft8631n71uvu1giuf7n1i1b1hil-1-scaled.jpg",
				},
				Ingredients: models.Ingredients{
					Values: []string{
						"10 ounces cremini mushrooms",
						"2 tablespoons olive oil",
						"3 tablespoon water",
						"3 tablespoons balsamic vinegar",
						"¼ teaspoon sea salt",
						"¼ teaspoon freshly cracked black pepper",
						"3 tablespoons unsalted dairy-free butter (or regular if not dairy-free)",
						"2 cloves garlic (minced)",
						"2 tablespoons minced fresh herbs (we used rosemary, oregano, and thyme)",
						"4 cups cauliflower florets",
						"⅔ cup unsweetened plain almond milk (or heavy cream if not dairy-free)",
						"1 ½ cups chopped kale",
						"Sea salt and freshly cracked pepper (to taste)",
					},
				},
				Instructions: models.Instructions{
					Values: []string{
						"First, make the veggie mash by heating the butter in a large 10.5” skillet over medium-high heat. " +
							"Once hot add in the garlic and fresh herbs cook for 30 seconds until fragrant. Add in the " +
							"cauliflower and heavy cream. Cover and simmer for 15 minutes until the florets are soft. Transfer " +
							"the mixture to a food processor and blend until smooth. If needed, add an additional tablespoon or " +
							"two of heavy cream to reach your desired consistency. Season with salt and pepper to taste.",
						"To the same skillet over medium heat add the kale with 2 tablespoons of water. Cover and cook 3-4 minutes " +
							"until wilted. Stir the kale into the mashed cauliflower and divide between two bowls.",
						"In the same skillet over medium-high heat, heat the olive oil. Once hot add in the garlic and mushrooms. " +
							"Add 1 tablespoon water to the skillet, cook for 5-7 minutes until the mushrooms are soft. Add " +
							"in the balsamic vinegar, salt, and black pepper. Cook another 1-2 minutes until the vinegar has " +
							"reduced and is thick. Top the bowls with the mushrooms and serve.",
					},
				},
				Name: "Balsamic Mushrooms with Herbed Veggie Mash",
				NutritionSchema: models.NutritionSchema{
					Calories:       "201 kcal",
					Carbohydrates:  "12 g",
					Cholesterol:    "",
					Fat:            "16 g",
					Fiber:          "3 g",
					Protein:        "4 g",
					SaturatedFat:   "3 g",
					Servings:       "1",
					Sodium:         "310 mg",
					Sugar:          "5 g",
					TransFat:       "2 g",
					UnsaturatedFat: "13 g",
				},
				PrepTime: "PT10M",
				Yield:    models.Yield{Value: 4},
				URL:      "https://www.castironketo.net/blog/balsamic-mushrooms-with-herbed-veggie-mash/"},
		},
		{
			name: "cdkitchen.com",
			in:   "https://www.cdkitchen.com/recipes/recs/285/MerleHaggardsRainbowStew65112.shtml",
			want: models.RecipeSchema{
				AtContext: atContext,
				AtType:    models.SchemaType{Value: "Recipe"},
				Name:      "Merle Haggard's Rainbow Stew",
				Description: models.Description{
					Value: `This colorful stew named for American country singer, Merle Haggard's song "Rainbow Stew", is ` +
						"loaded with sausage, chicken, beans, and fresh vegetables.",
				},
				Yield:    models.Yield{Value: 6},
				CookTime: "PT1H20M",
				Category: models.Category{Value: "stews"},
				Ingredients: models.Ingredients{
					Values: []string{
						"5 tablespoons canola oil, divided",
						"1 pound kielbasa, chorizo or andouille sausage, cut into 1/2-inch cubes",
						"1 pound boneless skinless chicken breasts, cut into 1-inch cubes",
						"3 cups chicken broth",
						"3 tablespoons all-purpose flour",
						"1/2 cup chopped red bell peppers",
						"1/2 cup chopped yellow bell peppers",
						"1/2 cup chopped green bell peppers",
						"1/2 cup chopped purple onion",
						"1 cup peeled and diced carrots",
						"1/2 cup chopped celery",
						"2 cloves garlic, minced",
						"1 cup peeled and cubed jicama",
						"2 tablespoons chopped parsley or cilantro",
						"1 can (16 ounce size) dark red kidney beans, rinsed and drained",
						"1 bay leaf, crumbled",
						"1 teaspoon summer savory, crumbled",
						"5 teaspoons cayenne pepper",
						"salt and pepper, to taste",
						"1/2 cup chopped green onions",
					},
				},
				Instructions: models.Instructions{
					Values: []string{
						"Heat 2 tablespoons of the oil in a Dutch oven over medium heat. Add the sausage and cook, stirring " +
							"frequently, until browned. Remove with a slotted spoon and set aside.",
						"Add the chicken to the pan and cook, stirring frequently, until browned. Remove the chicken with a " +
							"slotted spoon and add to the sausage. Drain the oil from the pan and return the chicken and " +
							"sausage to the pan.",
						"Add the broth and bring it to a simmer. Let cook until the chicken is cooked through.",
						"In a large skillet over medium heat, combine the flour and remaining oil and cook, stirring constantly, " +
							"until smooth. Stir in the bell peppers, onion, carrots, celery, garlic, jicama, and parsley. " +
							"Cook, stirring frequently, for 10 minutes.",
						"Transfer the vegetable mixture to the Dutch oven. Add the kidney beans, bay leaf, savory, and cayenne. " +
							"Bring to a boil then reduce the heat to a simmer. Let cook, uncovered, for 45 minutes, stirring frequently.",
						"Season to taste with salt and pepper. Add the green onions and mix well. Serve over rice.",
					},
				},
				NutritionSchema: models.NutritionSchema{
					Calories:      "565 calories",
					Carbohydrates: "28 grams carbohydrates",
					Fat:           "34 grams fat",
					Protein:       "36 grams protein",
				},
				URL: "https://www.cdkitchen.com/recipes/recs/285/MerleHaggardsRainbowStew65112.shtml",
			},
		},
		{
			name: "chefkoch.de",
			in:   "https://www.chefkoch.de/rezepte/1064631211795001/Knusprige-Ofenkartoffeln.html",
			want: models.RecipeSchema{
				AtContext:     atContext,
				AtType:        models.SchemaType{Value: "Recipe"},
				Category:      models.Category{Value: "Raffiniert & preiswert"},
				CookTime:      "P0DT0H40M",
				DatePublished: "2008-05-26",
				Description: models.Description{
					Value: "Knusprige Ofenkartoffeln. Über 226 Bewertungen und für raffiniert befunden. Mit ► Portionsrechner ► Kochbuch ► Video-Tipps! Jetzt entdecken und ausprobieren!",
				},
				Keywords: models.Keywords{
					Values: "Backen,Vegetarisch,Saucen,Dips,Beilage,raffiniert oder preiswert,einfach,Kartoffel,Snack",
				},
				Image: models.Image{
					Value: "https://img.chefkoch-cdn.de/rezepte/1064631211795001/bilder/1329056/crop-960x540/" +
						"knusprige-ofenkartoffeln.jpg",
				},
				Ingredients: models.Ingredients{
					Values: []string{
						"10 m.-große Kartoffeln, festkochende",
						"3 EL Olivenöl",
						"1 TL Meersalz , grobes, bei Bedarf auch mehr",
						"1 TL Paprikapulver, edelsüßes",
						"1 TL Currypulver",
						"etwas Chilipulver",
						"1 EL Rosmarin , getrocknet (frische Nadeln schmecken natürlich intensiver)",
						"etwas Pfeffer , aus der Mühle",
						"1 Becher saure Sahne",
						"1 Zitrone(n) , der Saft davon",
						"etwas Salz und Pfeffer",
						"1 EL Kräuter nach Wahl, italienische",
					},
				},
				Instructions: models.Instructions{
					Values: []string{
						"Die geschälten, geviertelten Kartoffeln gut mit Küchenkrepp abtrocknen. Die Kartoffeln in eine " +
							"große Schale geben, die Gewürze dazugeben und alles gut mit den Händen durchmengen. Nun das Öl " +
							"dazugeben und nochmals durchmengen. Die Kartoffeln auf ein Blech geben. Ich fette das Blech nicht " +
							"noch zusätzlich ein, da das Öl an den Kartoffeln bereits genügt.",
						"Im vorgeheizten Backofen bei 200 °C Ober-/Unterhitze auf mittlerer Schiene ca. 30 - 40 Min. (je " +
							"nach Größe der Kartoffelspalten) backen.",
						"Alle 10 Min. müssen die Spalten gewendet werden, damit sie von allen Seiten schön kross werden. " +
							"Wenn die Spalten anfangs etwas am Blech festbacken - nicht schlimm, einfach mit einem Pfannenwender lösen.",
						"Nun aus der sauren Sahne, dem Zitronensaft, den Gewürzen und den Kräutern einen Dip anrühren. Der " +
							"kann bei Bedarf natürlich auch mit etwas Knoblauch verfeinert werden.",
						"Dazu gibt es bei uns Fisch oder Hähnchenschenkel.",
					},
				},
				Name: "Knusprige Ofenkartoffeln",
				NutritionSchema: models.NutritionSchema{
					Calories:      "389 kcal",
					Carbohydrates: "53,93g",
					Fat:           "14,10g",
					Protein:       "8,52g",
					Servings:      "1",
				},
				PrepTime: "P0DT0H20M",
				Yield:    models.Yield{Value: 3},
				URL:      "https://www.chefkoch.de/rezepte/1064631211795001/Knusprige-Ofenkartoffeln.html",
			},
		},
		{
			name: "comidinhasdochef.com",
			in:   "https://comidinhasdochef.com/pudim-no-copinho-para-festa/",
			want: models.RecipeSchema{
				AtContext:     atContext,
				AtType:        models.SchemaType{Value: "Recipe"},
				Category:      models.Category{Value: "Sobremesa"},
				CookTime:      "PT5M",
				Cuisine:       models.Cuisine{Value: "Brasileira"},
				DatePublished: "2022-03-06T20:46:27-03:00",
				Description: models.Description{
					Value: "Você sabia que pudim pode ser um excelente doce para servir em festas? Pois, é! Praticamente todo mundo " +
						"ama um pudim, né? E,  por isso, muitas pessoas têm aprendido como fazer pudim no copinho para festa " +
						"que é um sucesso! E, se você trabalha nesta área ou adora preparar as festas da sua família, com " +
						"certeza você vai amar esta receita. O pudim fica delicioso, bem levinho, daquele jeito que todo mundo " +
						"ama e no ponto certo com esta receita, além disso, com o passo a passo que preparamos, você vai " +
						"conseguir montar todos os copinhos para servir nas festas e comemorações. Com certeza, todo mundo " +
						"vai amar e ainda querer repetir várias e várias vezes! Portanto, fique atento nesta receita de pudim " +
						"no copinho para festa, que pode ser uma solução para as comemorações e festas infantis, assim como " +
						"uma fonte de renda extra para você oferecer em seus serviços!",
				},
				Keywords: models.Keywords{Values: "Pudim no Copinho para Festa"},
				Image: models.Image{
					Value: "https://comidinhasdochef.com/wp-content/uploads/2022/03/Pudim-no-Copinho-para-Festa00.jpg",
				},
				Ingredients: models.Ingredients{Values: []string{
					"1 xícara (chá) de Água",
					"2 xícaras (chá) de açúcar",
					"500 ml de leite",
					"0 e 1/2 colheres (sopa) de essência de baunilha",
					"2 caixas de creme de leite",
					"2 caixas de leite condensado",
					"10 colheres (sopa) de Água quente",
					"2 pacotinhos de gelatina incolor sem sabor",
				},
				},
				Instructions: models.Instructions{
					Values: []string{
						"Em uma panela e no fogo médio vá adicionando o açúcar aos poucos, mexendo sem parar até derreter completamente;",
						"Adicione a água e continue mexendo até dissolver todas as pelotas formadas pelo açúcar;",
						"Quando estiver em consistência de calda bem liquida desligue o fogo e aguarde esfriar;",
						"Unte os copinhos (50 ml) com um fio de óleo e distribua toda a calda entre eles (cerca de 1 colher de sopa " +
							"em cada) e reserve.",
						"Em uma tigela dissolva toda a gelatina incolor na água quente e em seguida reserve;",
						"No liquidificador coloque o leite, a essência de baunilha, o creme de leite e o leite condensado;",
						"Bata por 2-3 minutos e acrescente a gelatina dissolvida;",
						"Bata novamente por mais 1 minuto ou até ficar bem homogêneo;",
						"Distribua a mistura entre os copinhos com a calda e para finalizar leve a geladeira por pelo menos 2 horas;",
						"Retire da geladeira e pronto, já pode servir.",
					},
				},
				Name: "Pudim no Copinho para Festa",
				NutritionSchema: models.NutritionSchema{
					Calories:       "215 kcal",
					Carbohydrates:  "36 g",
					Fat:            "6.5 g",
					Fiber:          "0 g",
					Protein:        "3.6 g",
					SaturatedFat:   "3.7 g",
					Servings:       "50",
					Sodium:         "57 mg",
					TransFat:       "0 g",
					UnsaturatedFat: "1.6 g",
				},
				PrepTime: "PT20M",
				Tools:    models.Tools{Values: []string(nil)},
				Yield:    models.Yield{Value: 50},
				URL:      "https://comidinhasdochef.com/pudim-no-copinho-para-festa/"},
		},
		{
			name: "cookeatshare.com",
			in:   "https://cookeatshare.com/recipes/balinese-bbq-pork-roast-babi-guling-81003",
			want: models.RecipeSchema{
				AtContext:     atContext,
				AtType:        models.SchemaType{Value: "Recipe"},
				Category:      models.Category{Value: "Kid Friendly"},
				CookTime:      "PT120M",
				Cuisine:       models.Cuisine{Value: "Indonesian"},
				DatePublished: "2009-07-05",
				Description: models.Description{
					Value: "My cousin Brett recently went to Bali for his honeymoon, so I decided to try the most famous Balinese " +
						"dish for July 4th this summer.  This pork came out soft and tasty like the meat you eat at a Luau, but " +
						"with a delicious complexity of flavors from the chiles, garlic, ginger, lemongrass and turmeric.  " +
						"The process of slicing the roast in the middle and stuffing it with the marinade infused the entire " +
						"roast with flavor.  Our guests raved...and were still going back for more even after dessert!  " +
						"Make sure to serve the Balinese yellow rice.  A wonderful combination!",
				},
				Image: models.Image{
					Value: "https://s3.amazonaws.com/grecipes/public/pictures/recipes/40375/balinese_pork.jpg",
				},
				Ingredients: models.Ingredients{
					Values: []string{
						"1 boneless pork shoulder roast, about 3 pounds (or bone in roast, about 3.5 pounds)",
						"4 large shallots or 2 small onions",
						"4-8 Thai chiles or 2-4 jalapenos",
						"4 cloves garlic, peeled",
						"2 Tbsp chopped fresh ginger1 Tbsp chopped fresh turmeric or 1/2 tsp ground turmeric",
						"1 Tbsp chopped fresh galangal, or 1 Tbsp more ginger",
						"3 stalks fresh lemongrass, trimmed and finely chopped or 3 large strips lemon zest",
						"1.5 tsp ground coriander",
						"1 tsp fresh ground black pepper",
						"2 Tbsp fresh lime juice (about 1 large lime)",
						"1 Tbsp firmly packed light brown sugar",
						"2 tsp salt",
						"3 Tbsp vegetable oil",
					},
				},
				Instructions: models.Instructions{
					Values: []string{
						"Combine the shallots, chiles, garlic, ginger, turmeric, galangal, lemongrass, coriander, pepper, lime " +
							"juice, sugar, salt, and 1 Tbsp vegetable oil in a mortar and pound to a smooth paste with the " +
							"pestle, or puree in a blender or food processor to a smooth paste.",
						"Heat 2 Tbsp of oil in a wok or skillet over medium heat.  Add the spice paste and saute until fragrant, " +
							"about 5 minutes.  Stir frequently to avoid splattering, and run your fan on high.  Remove from " +
							"pan and let cool. The paste can be made up to a day in advance and refrigerated in an airtight container.",
						"Trim the roast of excess fat, if any.  Using a sharp knife, make a deep slice in the center of roast, " +
							"starting and ending 3/4 inch from the ends, and cutting almost though to the other side of the " +
							"roast.  You should have a nice pocket about 6 inches long.",
						"Fill the pocket with spice paste and tie the roast back together with kitchen twine or pin it with " +
							"metal skewers.   Spread the remaining paste all over the outside of the roast.  If any spice mixture " +
							"remains, set it aside to add during the grilling process.",
						"If you have a rotisserie, this is probably the best way to cook the roast.  Preheat your grill to high " +
							"and cook for about 1.5 hours.",
						"I used the indirect cooking method.  Move the charcoals to either side of the place where you plan to " +
							"cook the roast, and place a drip pan in the middle.  If you have a gas grill where the coals are not " +
							"movable, either turn off the middle burner or just put a drip pan made of aluminum foil directly over " +
							"the center burner.  To make a foil drip pan, just tear 3 pieces of foil about 16 inches long, overlap them " +
							"so they are the width you want, then fold the edges up about 2 inches to form a make-shift drip pan.",
						"After setting up the grill, preheat it on high.  Turn the grill down and place the roast over the drip " +
							"pan.  Adjust heat or coals so the internal temperature rests at about 350 F.  This will ensure nice " +
							"browning without burning, and should result in a cook time of about 2 hours.  Turn the meat occasionally " +
							"during cooking so all sides get equally browned, and rub on additional spice mixture if any.",
						"Transfer roast to cutting board or platter and let stand for 10 minutes before removing strings and " +
							"cutting into thin slices to serve.",
					},
				},
				Name: "Balinese BBQ Pork Roast - Babi Guling",
				NutritionSchema: models.NutritionSchema{
					Calories:      "",
					Carbohydrates: "4.05 g",
					Cholesterol:   "0 g",
					Fat:           "6.92 g",
					Fiber:         "0.4 g",
					Protein:       "0.26 g",
					SaturatedFat:  "0.52 g",
					Servings:      "20 g",
					Sodium:        "777 g",
					Sugar:         "2.36 g",
					TransFat:      "0.18 g",
				},
				PrepTime: "PT30M",
				Yield:    models.Yield{Value: 0},
				URL:      "https://cookeatshare.com/recipes/balinese-bbq-pork-roast-babi-guling-81003"},
		},
		{
			name: "cookieandkate.com",
			in:   "https://cookieandkate.com/honey-butter-cornbread-recipe/",
			want: models.RecipeSchema{
				AtContext:     atContext,
				AtType:        models.SchemaType{Value: "Recipe"},
				Category:      models.Category{Value: "Baked goods"},
				CookTime:      "PT35M",
				CookingMethod: models.CookingMethod{Value: "Baked"},
				Cuisine:       models.Cuisine{Value: "American"},
				DatePublished: "2020-03-24",
				Description: models.Description{
					Value: "Make the best homemade cornbread with this easy recipe! It's fluffy on the inside, crisp around the " +
						"edges, and full of delicious honey-butter flavor. Recipe yields one large skillet of cornbread " +
						"(see recipe notes for alternate baking vessel options).",
				},
				Keywords: models.Keywords{Values: "skillet cornbread"},
				Image: models.Image{
					Value: "https://cookieandkate.com/images/2020/03/skillet-cornbread-recipe-2-2-225x225.jpg",
				},
				Ingredients: models.Ingredients{
					Values: []string{
						"1/2 cup (1 stick) unsalted butter",
						"1 1/2 cups cornmeal, medium-grind or finer*",
						"1 1/2 cups white whole wheat flour, regular whole wheat flour or all-purpose flour",
						"1 1/2 teaspoons fine sea salt",
						"2 teaspoons baking powder",
						"1/2 teaspoon baking soda",
						"3 large eggs, at room temperature**",
						"2/3 cup honey or maple syrup",
						"1 1/2 cups milk of choice (regular cow&#8217;s milk, almond or oat milk, etc.), at room temperature",
						"Optional serving suggestions: additional butter, honey or jam",
					},
				},
				Instructions: models.Instructions{
					Values: []string{
						"Preheat the oven to 375 degrees Fahrenheit. Place the butter in a large (12-inch) cast iron skillet and " +
							"place the skillet in the oven to melt the butter, about 5 to 13 minutes (keep an extra eye on " +
							"it as time goes on—we want it to get bubbly and lightly browned, but not burnt).",
						"Meanwhile, in a large bowl, combine the cornmeal, flour, salt, baking powder and baking soda. Stir to " +
							"combine, and set aside. In a medium bowl, whisk together the eggs and honey until fully blended. " +
							"Add the milk and whisk until evenly combined.",
						"Pour the liquid into the dry mixture, and stir just until moistened through (we&#8217;ll stir it more " +
							"soon). When the butter is melted and golden, use oven mitts (the skillet is crazy hot!) to remove " +
							"the skillet from the oven, and give it a gentle swirl to lightly coat about an inch up the sides.",
						"Pour the melted butter into the batter and stir just until incorporated. Pour the batter back into the " +
							"hot skillet, using a spatula to scrape all of the batter from the bowl into the skillet.",
						"Carefully return the skillet to the oven and bake until the bread is brown around the edge, springy to " +
							"the touch, and a toothpick inserted in the center comes out clean with just a few crumbs, 25 to 30 " +
							"minutes. Carefully (with oven mitts), place the skillet on a cooling rack. Let it cool for at " +
							"least 5 minutes before slicing and serving—perhaps with extra butter, honey or jam on the side.",
						"This cornbread will keep at room temperature in a sealed container for up to 3 days, or up to a week in " +
							"the refrigerator. You can also freeze it for up to 3 months. Gently reheat before serving.",
					},
				},
				Name: "Honey Butter Cornbread",
				NutritionSchema: models.NutritionSchema{
					Calories:      "200 calories",
					Carbohydrates: "30.3 g",
					Cholesterol:   "52 mg",
					Fat:           "7.8 g",
					Fiber:         "2.1 g",
					Protein:       "4.5 g",
					SaturatedFat:  "4.3 g",
					Servings:      "1",
					Sodium:        "268.9 mg",
					Sugar:         "12.9 g",
					TransFat:      "0 g",
				},
				PrepTime: "PT10M",
				Yield:    models.Yield{Value: 16},
				URL:      "https://cookieandkate.com/honey-butter-cornbread-recipe/"},
		},
		{
			name: "copykat.com",
			in:   "https://copykat.com/mcdonalds-egg-mcmuffin",
			want: models.RecipeSchema{
				AtContext:     atContext,
				AtType:        models.SchemaType{Value: "Recipe"},
				Category:      models.Category{Value: "Breakfast"},
				CookTime:      "PT5M",
				Cuisine:       models.Cuisine{Value: "American"},
				DatePublished: "2023-02-16T05:11:12+00:00",
				Description: models.Description{
					Value: "Learn how to make a McDonalds Egg McMuffin at home with this easy copycat recipe. Find out the secret " +
						"to making perfect egg rings for a breakfast sandwich.",
				},
				Keywords: models.Keywords{Values: "McDonald's Egg McMuffin"},
				Image: models.Image{
					Value: "https://copykat.com/wp-content/uploads/2020/04/McDonalds-Egg-McMuffin-Pin3.jpg",
				},
				Ingredients: models.Ingredients{
					Values: []string{
						"4 tablespoons softened butter (butter has divided uses)",
						"4 English Muffins",
						"4 slices Canadian Bacon",
						"4 eggs",
						"1/2 cup water",
						"4 slices American cheese",
					},
				},
				Instructions: models.Instructions{
					Values: []string{
						"Split open English Muffins and place them into a toaster, toast the English Muffins.",
						"In a non-stick skillet over medium heat, cook Candian bacon on both sides for about 1 to 2 minutes in " +
							"two teaspoons of butter. The bacon should begin to just brown.",
						"While the Canadian bacon is cooking, remove the English muffins from the toaster and spread softened " +
							"butter on both halves.",
						"Place the 1 slice of Canadian bacon on each English Muffin bottom.",
						"Add about 1 tablespoon of butter to the same skillet where you cooked the bacon.",
						"Place the quart-sized canning lids screw size up (or you can use an egg ring) into the skillet.",
						"Spray the canning lid with non-stick spray. Crack an egg into each of the rings.",
						"Break the yolk with a fork. Pour about 1/2 cup of water into the skillet, and place a lid on top. Cook " +
							"until the eggs are set, it should take about two minutes.",
						"Gently remove the eggs from the rings, and place one egg on each piece of Canadian bacon.",
						"Top each egg with one slice of American cheese, top cheese with the top of the English muffin.",
						"Wrap each egg McMuffin with foil or parchment paper. Wait about 30 seconds before serving.",
					},
				},
				Name: "McDonald's Egg McMuffin",
				NutritionSchema: models.NutritionSchema{
					Calories:      "420 kcal",
					Carbohydrates: "28 g",
					Cholesterol:   "229 mg",
					Fat:           "25 g",
					Fiber:         "2 g",
					Protein:       "20 g",
					SaturatedFat:  "13 g",
					Servings:      "1",
					Sodium:        "1037 mg",
					Sugar:         "1 g",
				},
				PrepTime: "PT5M",
				Yield:    models.Yield{Value: 4},
				URL:      "https://copykat.com/mcdonalds-egg-mcmuffin",
			},
		},
		{
			name: "countryliving.com",
			in:   "https://www.countryliving.com/food-drinks/a39298988/braised-turkey-wings-recipe/",
			want: models.RecipeSchema{
				AtContext:     atContext,
				AtType:        models.SchemaType{Value: "Recipe"},
				Category:      models.Category{Value: "Sunday lunch"},
				CookTime:      "PT0S",
				Cuisine:       models.Cuisine{Value: "American"},
				DatePublished: "2022-03-04T00:01:33.158861Z",
				Description: models.Description{
					Value: "This Southern dish is just bursting with flavor.",
				},
				Keywords: models.Keywords{
					Values: "American, Southern, Sunday lunch, comfort food, dinner",
				},
				Image: models.Image{
					Value: "https://hips.hearstapps.com/hmg-prod/images/braised-turkey-wings-clx040122-1646247632.jpg?crop=0.878xw:0.585xh;0,0.223xh&resize=1200:*",
				},
				Ingredients: models.Ingredients{
					Values: []string{
						"3 whole turkey wings (about 3 pounds total)",
						"Kosher salt and freshly ground black pepper",
						"2 tbsp. olive oil",
						"1 c. chopped yellow onion",
						"1/2 c. chopped carrots",
						"1/2 c. chopped celery",
						"2 cloves garlic, chopped",
						"2 tsp. chopped fresh rosemary",
						"2 tsp. chopped fresh sage",
						"1 tsp. chopped fresh thyme",
						"2 bay leaves",
						"2 tbsp. all-purpose flour",
						"4 c. turkey or chicken stock",
						"Cooked white rice, for serving",
					},
				},
				Instructions: models.Instructions{
					Values: []string{
						"Preheat oven to 350°F. Season wings with salt and pepper. Heat oil in a large Dutch oven over " +
							"medium-high heat. Add wings and cook, turning once, until golden brown, 4 to 5 minutes. " +
							"Transfer to a plate; reserve pot.",
						"Reduce heat to medium. Add onion, carrots, and celery to reserved pot. Cook, stirring occasionally, " +
							"until onion is translucent, 6 to 8 minutes. Add garlic, rosemary, sage, thyme, and bay leaves. " +
							"Cook, stirring, until garlic is fragrant, about 1 minute. Sprinkle in flour and cook, stirring, " +
							"until flour becomes a medium brown shade (like the color of caramel), 4 to 5 minutes. While " +
							"stirring, slowly pour in half of stock. Return wings to pot and pour in remaining stock until " +
							"wings are 2/3 covered by liquid. Cover and bake until wings are tender, 2 to 2 1/2 hours. Serve over rice.",
					},
				},
				Name: "Braised Turkey Wings",

				PrepTime: "PT40M",
				Yield:    models.Yield{Value: 4},
				URL:      "https://www.countryliving.com/food-drinks/a39298988/braised-turkey-wings-recipe/",
			},
		},
		{
			name: "cuisineaz.com",
			in:   "https://www.cuisineaz.com/recettes/champignons-farcis-au-fromage-brie-87449.aspx",
			want: models.RecipeSchema{
				AtContext:     atContext,
				AtType:        models.SchemaType{Value: "Recipe"},
				Category:      models.Category{Value: "Legumes Farcis"},
				CookTime:      "PT15M",
				Cuisine:       models.Cuisine{Value: "French"},
				DatePublished: "2016-06-06T14:39:27+02:00",
				Image: models.Image{
					Value: "https://img.cuisineaz.com/660x660/2016/06/06/i75661-champignons-farcis-au-fromage-brie.jpg",
				},
				Ingredients: models.Ingredients{
					Values: []string{
						"16 Champignon(s) de paris",
						"0.5 Brie",
						"1 Échalote(s)",
						"1 c. à soupe Crème fraîche",
						"1 Tranche(s) de jambon blanc",
						"2 c. à soupe Chapelure",
						"Sel poivre",
					},
				},
				Instructions: models.Instructions{
					Values: []string{
						"Etape 1",
						"Préchauffez le four à 180°C.",
						"Etape 2",
						"Enlevez le pied des champignons de Paris et nettoyez les chapeaux.",
						"Etape 3",
						"Epluchez l'échalote et coupez-la en quatre.",
						"Etape 4", "Coupez le fromage Brie en gros morceaux.",
						"Etape 5",
						"Dans le bol d'un robot mixeur, placez les morceaux d'échalote et de fromage Brie, la crème fraîche, la tranche de jambon blanc, la chapelure, du sel et du poivre. Mixez jusqu’à obtenir une crème bien lisse et homogène.",
						"Etape 6",
						"Répartissez la crème dans les champignons et disposez-les sur une plaque du four recouverte de papier sulfurisé.",
						"Etape 7",
						"Enfournez pendant 10 à 15 minutes.",
						"Etape 8",
						"Servez immédiatement accompagné de volaille.",
					},
				},
				Name:     "Champignons farcis au fromage Brie",
				PrepTime: "PT15M",
				Yield:    models.Yield{Value: 4},
				URL:      "https://www.cuisineaz.com/recettes/champignons-farcis-au-fromage-brie-87449.aspx",
			},
		},
		{
			name: "cybercook.com.br",
			in: "https://cybercook.com.br/receitas/peixes-e-frutos-do-mar/receita-de-file-de-tilapia-com-batatas-82273?" +
				"receita-do-dia",
			want: models.RecipeSchema{
				AtContext: atContext,
				AtType:    models.SchemaType{Value: "Recipe"},
				Category:  models.Category{Value: "Peixes e Frutos do Mar"},
				CookTime:  "PT1H",
				Description: models.Description{
					Value: "Já experimentou essa deliciosa receita de Filé de Tilápia com Batatas? No CyberCook você encontra " +
						"essa e outras receitas. Saiba mais!",
				},
				Image: models.Image{
					Value: "https://img.cybercook.com.br/receitas/273/file-de-tilapia-com-batatas.jpeg",
				},
				Ingredients: models.Ingredients{
					Values: []string{
						"Tilápia 500 gramas",
						"Batata 3 unidades",
						"Requeijão 1 copo",
						"Sal a gosto",
						"Azeite 1 colher (sopa)",
						"Cebola 1 unidade",
						"Salsinha 1 colher (sopa)",
					},
				},
				Instructions: models.Instructions{
					Values: []string{
						"Cozinhe as batatas al dente.",
						"Unte um refratário com azeite e requeijão, coloque as batatas e as cebolas, em cima coloque o peixe " +
							"temperado com sal e pimenta a gosto, e para finalizar acrescente o requeijão, leve ao " +
							"forno por Aproximadamente 40 minutos ou até dourar.",
						"Decore com salsinhas picadinhas",
						"Sirva a seguir.",
					},
				},
				Name: "Filé de Tilápia com Batatas",
				NutritionSchema: models.NutritionSchema{
					Calories: "274.20",
				},
				PrepTime: "PT1H",
				Tools:    models.Tools{Values: []string(nil)},
				Yield:    models.Yield{Value: 5},
				URL:      "https://cybercook.com.br/receitas/peixes-e-frutos-do-mar/receita-de-file-de-tilapia-com-batatas-82273",
			},
		},
		{
			name: "delish.com",
			in:   "https://www.delish.com/cooking/recipe-ideas/a24489879/beef-and-broccoli-recipe/",
			want: models.RecipeSchema{
				AtContext:     atContext,
				AtType:        models.SchemaType{Value: "Recipe"},
				Category:      models.Category{Value: "weeknight meals"},
				CookTime:      "PT0S",
				Cuisine:       models.Cuisine{Value: "American"},
				DateModified:  "2023-06-26T17:34:00Z",
				DatePublished: "2018-11-06T17:45:06.218596Z",
				Description: models.Description{
					Value: "A classic Chinese-American dish with thinly sliced, velveted flank steak in a rich brown sauce with tender-crisp broccoli.",
				},
				Keywords: models.Keywords{
					Values: "American, Asian, dinner, weeknight meals, beef and broccoli recipe, sirloin recipes, Chinese take " +
						"out recipes, easy weeknight dinner recipes, stir-fry recipe, beef stir-fry",
				},
				Image: models.Image{
					Value: "https://hips.hearstapps.com/hmg-prod/images/delish-230510-beef-broccoli-613-rv-index-646bca228a2b3.jpg?crop=0.502xw:1.00xh;0.250xw,0&resize=1200:*",
				},
				Ingredients: models.Ingredients{
					Values: []string{
						"2 tbsp. dry sherry or shaoxing wine",
						"2 tbsp. unseasoned rice vinegar",
						"1/2 tsp. kosher salt",
						"1/2 tsp. freshly ground black pepper",
						"1/3 c. plus 1/4 c. low-sodium soy sauce, divided",
						"2 tbsp. plus 1 1/2 tsp. cornstarch, divided",
						"2 tbsp. light brown sugar, divided",
						"1 1/2 lb. flank or skirt steak, sliced very thin against the grain",
						"4 cloves garlic",
						`1 (1/2") piece ginger, peeled`,
						"2 scallions",
						"2 small heads broccoli",
						"2/3 c. low-sodium beef broth",
						"2 tbsp. oyster sauce",
						"2 tsp. sriracha (optional)",
						"3 tbsp. vegetable oil, divided",
						"Toasted sesame seeds andwhite rice, for serving",
					},
				},
				Instructions: models.Instructions{
					Values: []string{
						"In a medium bowl, combine sherry, vinegar, salt, pepper, 1/3 cup soy sauce, 1 tablespoon plus 1 1/2 teaspoons cornstarch, and 1 tablespoon brown sugar. Add steak and toss to coat. Let sit 20 minutes.",
						"Meanwhile, finely chop garlic and ginger. Slice scallions and separate green and white parts. In a small bowl, combine garlic, ginger, and white scallion parts; reserve green parts for serving. Chop broccoli into florets and transfer to another small bowl.",
						"In a large measuring cup, whisk broth, oyster sauce, sriracha (if using), and remaining 1/4 cup soy sauce, 1 tablespoon cornstarch, and 1 tablespoon brown sugar. When ready to cook, arrange bowls of beef, garlic mixture, broccoli, and stir-fry sauce next to stove.",
						"In a large skillet over medium-high heat, heat 1 tablespoon oil. Add half of beef and cook, undisturbed, 1 minute, then stir and cook until cooked through and starting to char in some spots, about 1 minute more. Transfer to a plate. Repeat with 1 tablespoon oil and remaining beef. Discard excess marinade.",
						"Return skillet to medium heat and heat remaining 1 tablespoon oil. Add garlic mixture and cook, stirring occasionally, until fragrant, about 2 minutes. Add broccoli and cook, stirring frequently, until slightly softened, about 1 minute, then add stir-fry sauce. Cover and cook 3 minutes. Uncover, return beef to skillet, and toss to coat. Cook, tossing frequently, until warmed through and broccoli is crisp-tender, 2 to 3 minutes more.",
						"Divide rice among plates. Top with stir-fry, sesame seeds, and reserved green scallion parts.",
					},
				},
				Name: "Beef & Broccoli",
				NutritionSchema: models.NutritionSchema{
					Calories:      "609 Calories",
					Carbohydrates: "26 g",
					Cholesterol:   "111 mg",
					Fat:           "35 g",
					Fiber:         "7 g",
					Protein:       "46 g",
					SaturatedFat:  "10 g",
					Sodium:        "1918 mg",
					Sugar:         "9 g",
					TransFat:      "1 g",
				},
				PrepTime: "PT10M",
				Yield:    models.Yield{Value: 4},
				URL:      "https://www.delish.com/cooking/recipe-ideas/a24489879/beef-and-broccoli-recipe/",
			},
		},
		{
			name: "ditchthecarbs.com",
			in:   "https://www.ditchthecarbs.com/how-to-make-keto-samosa-air-fryer-oven/",
			want: models.RecipeSchema{
				AtContext:     atContext,
				AtType:        models.SchemaType{Value: "Recipe"},
				Category:      models.Category{Value: "Appetiser"},
				CookTime:      "PT10M",
				Cuisine:       models.Cuisine{Value: "Egg free"},
				DatePublished: "2022-03-17T09:35:35+00:00",
				Description: models.Description{
					Value: "Keto samosas is an Indian vegetarian dish perfect for appitizers, snacks, or even a meal.",
				},
				Keywords: models.Keywords{Values: "Keto Samosas"},
				Image: models.Image{
					Value: "https://thinlicious.com/wp-content/uploads/2022/02/Keto-Samosa-Featured-Image-Template-1200x1200-1.jpg",
				},
				Ingredients: models.Ingredients{
					Values: []string{
						"1 batch keto roti dough",
						"2 cups cauliflower (steamed and chopped)",
						"2 tbsp extra virgin olive oil ((ghee or coconut oil))",
						"½ tsp ground cumin",
						"½ tsp ground ginger",
						"1 tbsp garam masala",
						"2 cloves garlic (minced)",
						"2 tbsp jalapeno (chopped)",
						"2 stems green onions (chopped)",
						"1 tbsp lemon juice",
						"2 tbsp cilantro (chopped)",
						"+/- salt and pepper",
					},
				},
				Instructions: models.Instructions{
					Values: []string{
						"Start by maxing up a batch of keto roti dough. While the dough is resting you will want to make the filling.",
						"To make the filling heat the oil in a skillet. When the oil is hot add the seasoning(cumin, ginger, " +
							"Garam masala), garlic, and jalapeno to the skillet. Saute in the skillet for about 1 minute then " +
							"add the rest of the filling ingredients to the skillet. Let the filling cook for 3-4 minutes, stirring " +
							"occassionally. Once the filling is done remove it from the heat and let it cool completely.",
						"While the filling is cooling prepare and shape the samosas wrappers. To do this cut your roti dough " +
							"into 6 equal sections. Form each section into a ball. Roll each ball into a thin circle about the " +
							"size of a side plate between two sheets of parchment paper using a rolling pin.Then using a knife cut " +
							"each circle in half. This will give you 12 samosa wrappers. Cover and set aside until filling is " +
							"completely cool.",
						"Start by laying out one samosa wrapper with the flat cut side toward the bottom. Use your finger to " +
							"brush water in a line along half of the flat cut side. This will help the wrapper stick together " +
							"as you fold it.",
						"Next, fold the two corner edges up to make a cone. Overlap the corners with the wet edge on top of " +
							"the other. Firmly press the seam down to ensure it sticks and pinch the bottom of the cone closed.",
						"Hold the cone upright and open in your hand with the seam facing your palm. Your palm will help support " +
							"the seam while you add the filling. Spoon the filling into the cone until it is 2/3 of the way full.",
						"Using your finger brush water around the inside edge of the cone. Then press the edges of the cone " +
							"together and fold the bottom under to seal the cone and form a triangular samosa. Place the samosa " +
							"down so that it is standing up with the sealed bottom facing down. Repeast until all 12 samosas are filled.",
						"Finally, brush the samosas with olive oil and cook one of 3 ways.1) Bake in the oven on a baking tray " +
							"at 190°C/375°F for 15-18 minutes or until the samosas are golden brown.2) Arrange the samosas " +
							"in your air fryer basket and air at 190°C/375°F for 7-8 minutes. Depending on the size of your air fryer " +
							"you may need to cook the samosas in batches.3) Add an inch of olive oil to a skillet and fry the " +
							"samosas over medium-high heat for 3-4 minutes on each side. Only cook 3-4 samosas at a time. Let " +
							"the samosas cool for a few minutes after cooking and enjoy.",
					},
				},
				Name: "Easy Keto Samosas Recipe (Air Fryer Recipe)",
				NutritionSchema: models.NutritionSchema{
					Calories:      "69.6 kcal",
					Carbohydrates: "7.5 g",
					Fat:           "4.5 g",
					Fiber:         "3.8 g",
					Protein:       "1.2 g",
					SaturatedFat:  "",
					Servings:      "1",
					Sodium:        "30.7 mg",
					Sugar:         "0.9 g",
				},
				PrepTime: "PT30M",
				Yield:    models.Yield{Value: 12},
				URL:      "https://www.ditchthecarbs.com/how-to-make-keto-samosa-air-fryer-oven/",
			},
		},
		{
			name: "domesticate-me.com",
			in:   "https://domesticate-me.com/10-summer-cocktail-recipes/",
			want: models.RecipeSchema{
				AtContext:     atContext,
				AtType:        models.SchemaType{Value: "Recipe"},
				Category:      models.Category{Value: "Drinks"},
				DatePublished: "2021-05-28T16:11:33+00:00",
				Description: models.Description{
					Value: "Made with muddled strawberries, thyme, lemon, vodka, and St. Germain, this refreshing Strawberry " +
						"Thyme Cooler is perfect for all your summer celebrations!",
				},
				Image: models.Image{
					Value: "https://domesticate-me.com/wp-content/uploads/2019/08/Strawberry-Thyme-Cooler-Cocktail-1.jpg",
				},
				Ingredients: models.Ingredients{
					Values: []string{
						"1/3 cup hulled and diced strawberries",
						"1 sprig of fresh thyme",
						"½ ounce fresh lemon juice",
						"1 ounce St. Germain",
						"2 ounces vodka",
						"2-3 ounces club soda (depending on personal taste)",
						"Sliced strawberries",
						"Thyme sprig",
					},
				},
				Instructions: models.Instructions{
					Values: []string{
						"Muddle the strawberries, thyme sprig, and lemon juice in a cocktail shaker. Add the St. Germain, " +
							"vodka, and some ice, and shake vigorously to combine.",
						"Strain the cocktail into a glass with ice and top with club soda.",
						"Garnish with sliced strawberries and a sprig of thyme. Bottoms up!",
					},
				},
				Name:     "Strawberry Thyme Cooler and 9 Other Summer Cocktail Recipes",
				PrepTime: "PT5M",
				Yield:    models.Yield{Value: 1},
				URL:      "https://domesticate-me.com/10-summer-cocktail-recipes/",
			},
		},
		/*{
			name: "downshiftology.com",
			in:   "https://downshiftology.com/recipes/baked-chicken-breasts/",
			want: models.RecipeSchema{
				AtContext:     atContext,
				AtType:        models.SchemaType{Value: "Recipe"},
				Category:      models.Category{Value: "Main Course"},
				CookTime:      "PT20M",
				Cuisine:       models.Cuisine{Value: "American"},
				DatePublished: "2021-01-10T08:00:48+00:00",
				Description: models.Description{
					Value: "Perfectly baked chicken breaststhat are juicy, tender, easy, and oh so flavorful! All you need " +
						"is a drizzle of olive oil, a special seasoning mix, and a few insider tips for these super " +
						"tasty, no-fail chicken breasts.",
				},
				Keywords: models.Keywords{
					Values: "baked chicken, baked chicken breasts, baked chicken recipe",
				},
				Image: models.Image{
					Value: "https://i2.wp.com/www.downshiftology.com/wp-content/uploads/2021/01/Baked-Chicken-Breasts-10.jpg",
				},
				Ingredients: models.Ingredients{
					Values: []string{
						"4 boneless skinless chicken breasts",
						"1 tablespoon olive oil (or avocado oil)",
						"1 teaspoon kosher salt",
						"1 teaspoon paprika",
						"1/2 teaspoon garlic powder",
						"1/2 teaspoon dried thyme (or oregano, parsley or other herbs)",
						"1/4 teaspoon freshly ground black pepper",
					},
				},
				Instructions: models.Instructions{
					Values: []string{
						"Preheat your oven to 425F/220C. In a small bowl, mix together the paprika, garlic powder, thyme, " +
							"salt and pepper until combined.",
						"Lightly coat the chicken breasts in olive oil, and then generously rub the spice mix on both sides " +
							"of the chicken.",
						"Place the chicken breasts in a baking dish and cook for 20-25 minutes, depending on size (see chart " +
							"above). Let the chicken rest for a few minutes to allow the juices to redistribute within " +
							"the meat, then serve.",
					},
				},
				Name: "Best Baked Chicken Breast",
				NutritionSchema: models.NutritionSchema{
					Calories:      "163 kcal",
					Carbohydrates: "1 g",
					Cholesterol:   "72 mg",
					Fat:           "7 g",
					Fiber:         "1 g",
					Protein:       "24 g",
					SaturatedFat:  "1 g",
					Servings:      "1",
					Sodium:        "713 mg",
					Sugar:         "1 g",
				},
				PrepTime: "PT5M",
				Yield:    models.Yield{Value: 4},
				URL:      "https://downshiftology.com/recipes/baked-chicken-breasts/",
			},
		},*/
		{
			name: "dr.dk",
			in:   "https://www.dr.dk/mad/opskrift/nytarskage-med-champagne-kransekagebund-solbaer-og-chokoladepynt",
			want: models.RecipeSchema{
				AtContext:     atContext,
				AtType:        models.SchemaType{Value: "Recipe"},
				Name:          "Nytårskage med champagne, kransekagebund, solbær og chokoladepynt",
				DatePublished: "2022-01-01T21:00:00+00:00",
				Image: models.Image{
					Value: "https://asset.dr.dk/ImageScaler03/?url=http%3A%2F%2Fmad-recipe-pictures-dr-dk.s3.amazonaws.com%2F" +
						"prod%2Frecipe%2Fnytarskage-anelise-169-1640951740.jpg",
				},
				Description: models.Description{
					Value: "Smuk nytårskage med urvisere og masser af smag, der passer perfekt til nytårsaften.",
				},
				URL: "https://www.dr.dk/mad/opskrift/nytarskage-med-champagne-kransekagebund-solbaer-og-chokoladepynt",
				Ingredients: models.Ingredients{
					Values: []string{
						"30 g æggehvider",
						"50 g sukker",
						"Fintrevet citronskal",
						"200 g bagemarcipan",
						"15 g sukker",
						"15 g ristede, hakkede smuttede mandler",
						"1 nip salt",
						"50 g mørk chokolade",
						"1 spsk. frysetørret solbærpulver",
						"1 tsk. pufsukker",
						"140 g hvid chokolade",
						"1/2 blad husblas",
						"2 æggeblommer",
						"15 g sukker",
						"80 ml piskefløde",
						"80 g solbærpuré",
						"2 tsk. citronsaft",
						"1/2 tsk. citronsyre",
						"75 g sukker",
						"40 g glukosesirup",
						"30 ml tør champagne",
						"1/2 tsk. citronsyre",
						"50 g æggehvider",
						"1 spsk. sukker",
						"2 tsk. solbærpulver",
						"2 blade husblas",
						"400 g gold chokolade",
						"160 ml tør champagne",
						"40 ml citronsaft",
						"4 dl piskefløde",
						"1 tsk. solbærpulver",
						"5 blade husblas",
						"150 g hvid chokolade",
						"100 g kondenseret mælk",
						"51 g solbærpuré",
						"54 g vand",
						"150 g sukker",
						"100 g glukosesirup",
						"Minimum 800 g mørk chokolade",
					},
				},
				Instructions: models.Instructions{
					Values: []string{
						"Kransekagebund",
						"Varm marcipan kort i mikroovn, rør alle ingredienser sammen i røremaskinen indtil den samler sig.",
						"Fordel dejen i en kagering på ca. 16ø og bag den I ca. 15 minutter ved 200C, til den er let gylden. Lad den køle ned.",
						"\n",
						"Mandel-solbær-knas",
						"Smelt sukkeret ved lav blus i en gryde, indtil den er let gylden. Tag gryden af varme og rør mandler og salt i.",
						"Hæld karamelliseret mandler i et stk bagepapir og lad den køle helt ned. Hak mandlerne fint. Smelt chokoladen op til 45C, tilsæt solbærpulver, pufsukker og de hakket mandler i.",
						"Fordel knaset oven på kransekagebunden til et jævnt fint lag. Stil kagen i fryseren mens du arbejder videre.",
						"\n",
						"Solbærcremeux",
						"Kom chokoladen i en skål og sæt den til side. Sæt husblas i koldt vand i ca. 5 minutter.",
						"Pisk æggeblommer og sukker let sammen i en anden skål.",
						"Varm fløde, citronsyre og solbærpuré op til kogepunktet og hæld den i æggeblandingen under piskning. Hæld blandingen tilbage i gryden og kog cremen op til 85C under konstant omrøring med en silikoneske.",
						"Tag gryden af varmen og sigt cremen.",
						"Vrid husblas fri fra vand og smelt den i cremen. Hæld cremen over chokoladen og rør midt i skålen indtil chokoladen og cremen er blevet homogen.",
						"Smag til med citronsaft og stablen cremeuxen. Fordel den ovenpå kransekagebund (chokolade knas skal være ned i bunden) i bageringen og sæt den i fryseren.",
						"\n",
						"Champagne-flødebolleskum",
						"Bring 75g sukker, glukosesirup, champagne og citronsyre i kog op til 118C.",
						"Pisk æggehvider og 1 spsk sukker næsten stive i en skål og hæld den varme sukkerlage ned i æggehviderne i en tynd stråle under piskning. Pisk videre, til skummet er sejt og fast. Ved slutning tilsæt solbærpulver og pisk færdig.",
						"Fordel skummet over solbærcremeux, gør overfladen glat med en paletkniv og stil kagen i fryseren igen.",
						"\n",
						"Champagnemousse",
						"1/2 tsk citronsyre Udblød husblassen i koldt vand. Smelt chokoladen op til 45C.",
						"Bring champagne, citronsyre og citronsaft til kogepunktet og tag gryden straks af varmen.",
						"Vrid husblas fri for vand og rør den i den varme champagne. Hæld champagnen over chokoladen, mens du rør " +
							"i midten af skålen. Fortsæt indtil massen samler sig.",
						"Pisk fløde til let skum, ved skummet over chokoladen ad 3 omgange.",
						"Fordel champagnemousse i en silikoneforme og tryk solbærindlæg ned i moussen - bunden skal være oppe. Lad " +
							"kagen fryse indtil den skal glazes.",
						"\n",
						"Solbærglaze",
						"Læg husblad i koldt vandbad i ca. 5 minutter. Kom chokolade i en høj kande.",
						"Bring solbærpuré, vand, sukker og glukosesirup i en gryde og koge den op til kogepunktet. Tag gryden af " +
							"varmen og rør kondenseret mælk i.",
						"Vrid husblassen fri for vand og rør den ud i den varme væske. Hæld væsken over chokoladen og lad den " +
							"træk i et par minutter.",
						"Stavblend glaze indtil det er samlet og ensartet. Tilsæt guldstøv og stavblend igen.",
						"Dæk overflade med film og lad gazen køle ned til 32 grader.",
						"Placér den frosne kage på en rist og glaze kagen.",
						"\n",
						"Forberedelse af chokoladepynt",
						"Klip et langt stykke plast på 63 cm til som et bylandskab og et andet stykke plast på 30 cm",
						"Klip to viser af stift kageplast",
						"Forberedelse af en halvkugle Ø10 cm pudses med vat",
						"Temperere guld kakaofarve sammen med guldstøv og dup med en svamp et fyrværkeri mønster i halvkuglen, så " +
							"det ligner en guldregn",
						"Temperere 800 g mørk chokolade",
						"Fyld hele halvkuglen op med chokolade og lade det sidde i 1-2 min for at sørge for, at skallen bliver tyk",
						"Bank alt det overskydende chokolade ud på et stykke bagepapir og sæt den til at størkne",
						"Fordel et jævnt lag chokolade ud på plasten der forestiller et bylandskab og på 30 cm plast",
						"Beklæd en 20 cm kagering med plast og sæt bylandskabet rundt herom, som et stort bånd",
						"Udstik med en Ø3 cm ring små knapper af den næsten størknet chokolade på 30 cm plast og befri dem, når " +
							"chokoladen er kølet helt af",
						"De runde cirkler pensles med guldstøv og efterfølgende skrives der med romertal fra 1-12",
						"Befri halvkuglen fra formen og sæt til side",
						"Lav visere af chokolade",
						"\n",
						"Samling",
						"Den frosne kage befries fra formen og sættes på en drejeskive",
						"Kagen betrækkes med et lag smørcreme, som glattes tyndt ud",
						"Kagen sættes tilbage på frost i 10 min",
						"Betræk med et nyt lag, som glattes helt jævnt ud",
						"Sæt kagen tilbage på frost 10-15 min",
						"Spray kagen med cremet velvet spray farve",
						"Betræk kagen med chokoladebånd",
						"Pynt toppen med guldtal oven på små chokoladekugler og placere halvkuglen i midten af kagen",
					},
				},
			},
		},
		{
			name: "eatingbirdfood.com",
			in:   "https://www.eatingbirdfood.com/cinnamon-rolls/",
			want: models.RecipeSchema{
				AtContext:     atContext,
				AtType:        models.SchemaType{Value: "Recipe"},
				Category:      models.Category{Value: "Breakast"},
				CookTime:      "PT25M",
				Cuisine:       models.Cuisine{Value: "American"},
				DatePublished: "2022-12-15T08:05:22+00:00",
				Description: models.Description{
					Value: "Make cinnamon rolls from scratch with this easy recipe that&#039;s perfect for beginners! " +
						"They&#039;re soft, gooey, and made with bread flour, which gives them the perfect fluffy " +
						"texture. Overnight instructions included.",
				},
				Keywords: models.Keywords{Values: "cinnamon rolls"},
				Image: models.Image{
					Value: "https://www.eatingbirdfood.com/wp-content/uploads/2022/12/fluffy-cinnamon-rolls-hero.jpg",
				},
				Ingredients: models.Ingredients{
					Values: []string{
						"1 0.25 oz. pkg. active dry yeast",
						"1 cup warm milk (around 105°–115°F)",
						"3 cups bread flour (plus more for rolling dough)",
						"1 teaspoon sea salt",
						"2 Tablespoons granulated sugar",
						"3 Tablespoons unsalted butter, melted (plus more for greasing)",
						"1 large egg (at room temperature)",
						"½ cup brown sugar (packed )",
						"1 Tablespoon ground cinnamon",
						"¼ cup unsalted butter (softened)",
						"4 oz cream cheese (full fat, softened to room temperature)",
						"¼ cup Greek yogurt (I used plain full fat)",
						"2 Tablespoons maple syrup",
						"1 teaspoon vanilla extract",
					},
				},
				Instructions: models.Instructions{
					Values: []string{
						"Heat milk in a saucepan or in the microwave (about 40-50 seconds) until warm, but not hot. Around " +
							"(105°–115°F). Stir yeast into warm almond milk until dissolved. Let stand 10 minutes.",
						"In a stand mixer with the paddle attachment, add flour, salt and sugar. Mix to combine. With mixer " +
							"running at low speed, add melted butter, egg and yeast mixture. Increase speed to medium-low, " +
							"and mix 2 minutes until dough starts to form.",
						"Switch to the dough hook attachment and knead the dough in the stand mixer at medium-low speed for " +
							"5 minutes, or until dough is smooth. Increase speed to medium and mix 2 minutes. Kneading is " +
							"done when dough makes a slapping sound as it hits the side of the bowl. Dough temperature " +
							"should be close to 90°F. If dough is too sticky add a little more flour.",
						"Once combined, place dough in oiled mixing bowl and cover with plastic wrap and a warm towel. Let " +
							"rise 1-2 hours, or until doubled in volume. Time will depend on how warm your house is.",
						"While dough is rising, make cinnamon sugar filling by stirring together brown sugar and cinnamon " +
							"in a small bowl. Grease 13 x 9-inch baking pan or round 9.5-inch pie pan with butter.",
						"Once dough has doubled in size, sprinkle extra flour on your surface and rolling pin and roll " +
							"dough into 14 x 12-inch rectangle.",
						"Spread softened butter onto dough with your fingers or a knife.",
						"Sprinkle cinnamon sugar mixture over butter and press down slightly with your hands.",
						"Starting at the top, roll the dough toward you into a large log, moving back and forth down " +
							"the line of dough (in a “typewriter” motion) and always rolling toward you. Roll it " +
							"tightly as you go so the rolls will be nice and neat. When it’s all rolled, pinch the " +
							"seam closed and turn the roll over so that the seam is facing down. Cut roll crosswise " +
							"into 12 1-inch-thick pieces and place on prepared baking pan.",
						"Cover, and let rise in warm place 45 minutes, or until doubled in size.",
						"Preheat oven to 350°F. Bake cinnamon rolls for 20-25 minutes, or until golden brown, cooked " +
							"through and no longer doughy. I baked mine for 22 minutes.",
						"While cinnamon rolls are cooling, make cream cheese frosting by adding cream cheese, greek yogurt, " +
							"maple syrup and vanilla to a large mixing bowl. Using a hand mixer on medium speed, whip " +
							"all the ingredients together until smooth and fluffy, scraping down the sides as needed. " +
							"Alternatively, you can use a stand mixer.",
						"Spread frosting over warm cinnamon rolls and serve.",
					},
				},
				Name: "Fluffy Cinnamon Rolls",
				NutritionSchema: models.NutritionSchema{
					Calories:       "269 kcal",
					Carbohydrates:  "37 g",
					Cholesterol:    "43 mg",
					Fat:            "11 g",
					Fiber:          "1 g",
					Protein:        "6 g",
					SaturatedFat:   "6 g",
					Servings:       "1",
					Sodium:         "252 mg",
					Sugar:          "14 g",
					UnsaturatedFat: "3 g",
				},
				PrepTime: "PT30M",
				Yield:    models.Yield{Value: 12},
				URL:      "https://www.eatingbirdfood.com/cinnamon-rolls/",
			},
		},
		{
			name: "eatingwell.com",
			in:   "https://www.eatingwell.com/recipe/7887715/lemon-chicken-piccata/",
			want: models.RecipeSchema{
				AtContext:     atContext,
				AtType:        models.SchemaType{Value: "Recipe"},
				Cuisine:       models.Cuisine{Value: ""},
				DateModified:  "2023-09-19T12:07:24.154-04:00",
				DatePublished: "2021-02-04T15:20:10.000-05:00",
				Description: models.Description{
					Value: "This chicken piccata recipe has a bright, briny flavor, is made from ingredients you likely have on hand, and goes with everything from chicken to tofu to scallops. Bonus: It&#39;s lower in calories than a lot of other pan sauces.",
				},
				Image: models.Image{
					Value: "https://www.eatingwell.com/thmb/vQS6R5pXm1TqsYvOvn2WW0UVrIw=/1500x0/filters:no_upscale():max_bytes(150000):strip_icc()/chik-picatta-2000-872203d28060486397039a9ad5d2b118.jpg",
				},
				Ingredients: models.Ingredients{
					Values: []string{
						"1.25 pounds boneless, skinless chicken breasts, trimmed",
						".5 teaspoon salt",
						".25 teaspoon ground pepper",
						"2 tablespoons extra-virgin olive oil",
						"1 medium shallot, minced",
						"3 cloves garlic, minced",
						"2 teaspoons all-purpose flour",
						".5 cup low-sodium chicken broth",
						".5 cup dry white wine",
						"2 tablespoons lemon juice",
						"1 tablespoon butter",
						"1 tablespoon capers, rinsed",
						"2 tablespoons chopped fresh parsley",
					},
				},
				Instructions: models.Instructions{
					Values: []string{
						"Remove tenders from chicken and reserve for another use. Place the chicken breasts between 2 pieces of plastic wrap and gently pound with a meat mallet, rolling pin or small skillet to an even thickness of about ½ inch. Pat the chicken dry and sprinkle with salt and pepper.",
						"Heat oil in a large skillet over medium-high heat. Add the chicken and cook, flipping once, until well browned on both sides, 6 to 8 minutes. Continue to cook, flipping often, until an instant-read thermometer inserted in the thickest part registers 165°F, about 3 minutes more. Transfer to a clean cutting board and tent with foil to keep warm.",
						"Reduce heat to medium and add shallot to the pan. Cook, stirring often, until softened, 1 to 2 minutes. Add garlic and cook, stirring, until fragrant, about 1 minute. Sprinkle with flour and cook, stirring, for 1 minute. Stir in broth and wine, scraping up any browned bits. Simmer until reduced by half, 3 to 5 minutes. Remove from heat and stir in lemon juice, butter, capers and parsley. Serve the chicken with the sauce.",
					},
				},
				Name: "Lemon Chicken Piccata",
				NutritionSchema: models.NutritionSchema{
					Calories:       "264 kcal",
					Carbohydrates:  "7 g",
					Cholesterol:    "70 mg",
					Fat:            "13 g",
					Protein:        "24 g",
					SaturatedFat:   "4 g",
					Sodium:         "382 mg",
					Sugar:          "1 g",
					UnsaturatedFat: "0 g",
				},
				Yield: models.Yield{Value: 4},
				URL:   "https://www.eatingwell.com/recipe/7887715/lemon-chicken-piccata/",
			},
		},
		{
			name: "eatsmarter.com",
			in:   "https://eatsmarter.com/recipes/vietnamese-chicken-cabbage-salad",
			want: models.RecipeSchema{
				AtContext:     atContext,
				AtType:        models.SchemaType{Value: "Recipe"},
				Category:      models.Category{Value: "Lunch"},
				Cuisine:       models.Cuisine{Value: "Asian, Vietnamese"},
				DatePublished: "2016-10-07",
				Description: models.Description{
					Value: "Light and refreshing Vietnamese Chicken Cabbage Salad with crunchy peanuts",
				},
				Image: models.Image{
					Value: "https://images.eatsmarter.com/sites/default/files/styles/max_size/public/" +
						"vietnamese-chicken-cabbage-salad-580858.jpg",
				},
				Ingredients: models.Ingredients{
					Values: []string{
						"24 ozs Chicken broth",
						"22 ozs Chicken breasts",
						"3  carrots (10-12 ounces)",
						"1 bunch Radish",
						"1 Red chili pepper",
						"1  Chinese cabbage",
						"1 bunch mixed Fresh herbs (such as mint, basil)",
						"1  Lime",
						"3 Tbsps chopped Peanuts",
						"2 Tbsps vegetable oil",
						"peppers",
					},
				},
				Instructions: models.Instructions{
					Values: []string{
						"Bring broth to a boil, add chicken and simmer for about 20 minutes on low heat.",
						"Peel carrots and cut into thin, long strips. Rinse and dry radishes, cut into thin slices. " +
							"Rinse and chop chile pepper.",
						"Rinse and dry cabbage, cut into fine strips. Rinse and shake dry herbs, pluck off leaves.",
						"Remove chicken from broth and cool. Squeeze lime juice.",
						"Cut meat into thin strips, mix with prepared vegetables and herbs, drizzle with " +
							"lime juice and oil, season with salt and pepper and sprinkle with nuts. Serve.",
					},
				},
				Name:  "Vietnamese Chicken Cabbage Salad",
				Tools: models.Tools{Values: []string(nil)},
				Yield: models.Yield{Value: 4},
				URL:   "https://eatsmarter.com/recipes/vietnamese-chicken-cabbage-salad",
			},
		},
		/*{
			name: "eatwhattonight.com",
			in:   "https://eatwhattonight.com/2021/11/diced-chicken-with-spicy-chilies-%e8%be%a3%e5%ad%90%e9%b8%a1%e4%b8%81/#wpzoom-recipe-card",
			want: models.RecipeSchema{
				AtContext:     atContext,
				AtType:        models.SchemaType{Value: "Recipe"},
				Category:      models.Category{Value: "Sides"},
				CookTime:      "PT20M",
				Cuisine:       models.Cuisine{Value: "Chinese"},
				DatePublished: "2021-11-30T13:14:26+08:00",
				Description:   models.Description{Value: ""},
				Keywords:      models.Keywords{Values: "Chicken"},
				Image: models.Image{
					Value: "http://eatwhattonight.com/wp-content/uploads/2021/11/Spicy-Chicken_1-scaled.jpg",
				},
				Ingredients: models.Ingredients{
					Values: []string{
						"150g diced boneless chicken leg, de-skinned",
						"12 pcs dried chilies, soaked in water to soften, cut into sections",
						"2 tbsp cooking oil",
						"1 tsp Szechuan peppercorns",
						"2 small bulbs of garlic, sliced",
						"5-6 thin slices of ginger",
						"2-3 tbsp water",
						"1 tbsp cooking wine",
						"Some sesame seeds",
						"1 3/4 tsp light soya sauce",
						"1 tbsp ginger juice",
						"1/2 tsp sugar",
						"Pinch of pepper",
						"4 tsp corn flour",
					},
				},
				Instructions: models.Instructions{
					Values: []string{
						"Marinate chicken, cover and set aside in the fridge for an hour.",
						"Heat up 1 tbsp cooking oil and panfry marinated chicken till cooked and remove from flame (helps " +
							"to release excess oils from chicken).",
						"Add to air-fryer to grill for further 5 mins at 160 degrees C or till browned. Remove and set aside.",
						"Heat up balance cooking oil and sauté ginger and garlic slices.",
						"Add Szechuan pepper and dried chilies and stir-fry to bring out the aroma.",
						"Add water if it starts to get too dry. Add chicken pieces and stir-fry to mix well.",
						"Add cooking wine and stir-fry till chicken are dry. Off heat and sprinkle sesame seeds all over. " +
							"Mix well and serve hot immediately to enjoy.",
					},
				},
				Name:     "Diced Chicken with Spicy Chilies 辣子鸡丁",
				PrepTime: "PT10M",
				Yield:    models.Yield{Value: 1},
				URL: "https://eatwhattonight.com/2021/11/diced-chicken-with-spicy-chilies-" +
					"%e8%be%a3%e5%ad%90%e9%b8%a1%e4%b8%81/#wpzoom-recipe-card",
			},
		},*/
		{
			name: "epicurious.com",
			in:   "https://www.epicurious.com/recipes/food/views/olive-oil-cake",
			want: models.RecipeSchema{
				AtContext:     atContext,
				AtType:        models.SchemaType{Value: "Recipe"},
				DateModified:  "2017-11-01T10:56:00.000-04:00",
				DatePublished: "2017-11-01T10:56:00.000-04:00",
				Description: models.Description{
					Value: "Even die-hard butter devotees admit that olive oil makes exceptionally good cakes. " +
						"EVOO is liquid at room temperature, so it lends superior moisture over time. In fact, " +
						"olive oil cake only improves the longer it sits—this dairy-free version will keep on " +
						"your counter for days (not that it’ll last that long).",
				},
				Keywords: models.Keywords{
					Values: "cake,amaretto,vermouth,grand marnier,italian,cake flour,almond flour,lemon,vanilla,snack,breakfast,nut free,baking,stand mixer,bon appétit,web",
				},
				Image: models.Image{
					Value: "https://assets.epicurious.com/photos/5a05db121a9e232c87581a7f/16:9/w_2000,h_1125,c_limit/olive-oil-cake-recipe-BA-111017.jpg",
				},
				Ingredients: models.Ingredients{
					Values: []string{
						"1¼ cups plus 2 tablespoons extra-virgin olive oil; plus more for pan",
						"1 cup plus 2 tablespoons sugar; plus more",
						"2 cups cake flour",
						"⅓ cup almond flour or meal or fine-grind cornmeal",
						"2 teaspoons baking powder",
						"½ teaspoon baking soda",
						"½ teaspoon kosher salt",
						"3 tablespoons amaretto, Grand Marnier, sweet vermouth, or other liqueur",
						"1 tablespoon finely grated lemon zest",
						"3 tablespoon fresh lemon juice",
						"2 teaspoons vanilla extract",
						"3 large eggs",
						"A 9\"-diameter springform pan",
					},
				},
				Instructions: models.Instructions{
					Values: []string{
						"Preheat oven to 400°F. Drizzle bottom and sides of pan with oil and use your fingers to coat. " +
							"Line bottom with a round of parchment paper and smooth to eliminate air bubbles; coat parchment " +
							"with more oil. Generously sprinkle pan with sugar and tilt to coat in an even layer; tap out excess. " +
							"Whisk cake flour, almond flour, baking powder, baking soda, and salt in a medium bowl to combine " +
							"and eliminate any lumps. Stir together amaretto, lemon juice, and vanilla in a small bowl.",
						"Using an electric mixer on high speed (use whisk attachment if working with a stand mixer), beat " +
							"eggs, lemon zest, and 1 cup plus 2 Tbsp. sugar in a large bowl until mixture is very light, thick, " +
							"pale, and falls off the whisk or beaters in a slowly dissolving ribbon, about 3 minutes if using " +
							"a stand mixer and about 5 minutes if using a hand mixer. With mixer still on high speed, gradually " +
							"stream in 1¼ cups oil and beat until incorporated and mixture is even thicker. Reduce mixer speed " +
							"to low and add dry ingredients in 3 additions, alternating with amaretto mixture in 2 additions, " +
							"beginning and ending with dry ingredients. Fold batter several times with a large rubber spatula, " +
							"making sure to scrape the bottom and sides of bowl. Scrape batter into prepared pan, smooth top, " +
							"and sprinkle with more sugar.",
						"Place cake in oven and immediately reduce oven temperature to 350°F. Bake until top is golden brown, " +
							"center is firm to the touch, and a tester inserted into the center comes out clean, 40–50 minutes. " +
							"Transfer pan to a wire rack and let cake cool in pan 15 minutes.",
						"Poke holes all over top of cake with a toothpick or skewer and drizzle with remaining 2 Tbsp. oil; " +
							"let it absorb. Run a thin knife around edges of cake and remove ring from pan. Slide cake onto " +
							"rack and let cool completely. For the best flavor and texture, wrap cake in plastic and let sit at " +
							"room temperature at least a day before serving.",
						"Cake can be baked 4 days ahead. Store tightly wrapped at room temperature.",
					},
				},
				Name:  "Olive Oil Cake",
				Yield: models.Yield{Value: 8},
				URL:   "https://www.bonappetit.com/recipe/olive-oil-cake",
			},
		},
		{
			name: "expressen.se",
			in:   "https://alltommat.expressen.se/recept/saftiga-choklad--och-apelsinbullar/",
			want: models.RecipeSchema{
				AtContext: atContext,
				AtType:    models.SchemaType{Value: "Recipe"},
				Name:      "Saftiga choklad- och apelsinbullar",
				Description: models.Description{
					Value: `Goda små "fjärilsbullar" med choklad och krämig apelsinfyllning. För att få bullarnas fina fjärilsliknande form skärs degrullen i skivor som trycks ihop i mitten. Spritsa apelsinfyllningen i mitten av varje bulle. Supergott och lyxigt!`,
				},
				Image: models.Image{
					Value: "https://static.cdn-expressen.se/images/45/cd/45cd6c5649004e1fa957e891d581fa49/1x1/1920@80.jpg",
				},
				Ingredients: models.Ingredients{
					Values: []string{
						"50 g jäst",
						"5 dl mjölk",
						"200 g smör",
						"2 dl strösocker",
						"1 msk hela kardemummakärnor",
						"1.5 tsk salt",
						"16 dl vetemjöl",
						"200 g smör",
						"2 dl strösocker",
						"3 msk kakao",
						"2 tsk vaniljsocker",
						"0.5 dl strösocker",
						"1.25 dl mjölk",
						"0.5 dl apelsin",
						"1 apelsin",
						"17 g majsstärkelse (Maizena)",
						"40 g äggulor",
						"1 krm salt",
						"5 g smör",
						"1 ägg",
						"1 krm salt",
						"droppar vatten",
						"4 msk pärlsocker",
					},
				},
				Instructions: models.Instructions{
					Values: []string{
						"Smula ner jästen i en bunke, gärna tillhörande en köksassistent. Tillsätt mjölken och blanda sedan " +
							"i resten av ingredienserna. Arbeta degen i ca 10 min. Lägg en bakduk över bunken och låt degen " +
							"jäsa i 30min. Gör under tiden fyllningen och apelsinkrämen.",
						"Chokladfyllning: Rör ihop smör, strösocker, kakao och vaniljsocker. Om den känns för tjock kan du " +
							"mikra den ett par sekunder.",
						"Apelsinkräm: Blanda ihop alla ingredienserna utom smöret i en kastrull. Låt det sjuda och vispa " +
							"under tiden. Dra av kastrullen från värmen när krämen börjar tjockna och vispa i smöret. " +
							"Passera krämen genom en sil. Fyll en spritspåse med apelsinkrämen. Låt den svalna.",
						"Stjälp upp degen på en mjölad arbetsbänk. Kavla ut degen till en rektangel, 25x65 cm. Bred ut " +
							"chokladfyllningen på hela ytan. Rulla ihop degen från långsidan till en rulle. Skär rullen " +
							"i bitar, ca 3 cm breda. Tryck till bitarna på mitten med ett grillspett eller en rund pinne " +
							"så att snittytorna viks in mot mitten som en fjäril.",
						"Låt bullarna jäsa under en bakduk i 1–1 ½ timme. Sätt ugnen på 220grader.",
						"Gör ett hål med fingret i mitten på varje bulle och spritsa i apelsinfyllningen. Vispa ihop ägg, " +
							"salt och vatten med en gaffel. Pensla bullarna med äggblandningen och strö över pärlsocker. " +
							"Grädda bullarna i 7–8min, låt dem svalna på ett galler.",
					},
				},
				Keywords: models.Keywords{Values: "sections/recept"},
				Yield:    models.Yield{Value: 22},
				URL:      "https://alltommat.expressen.se/recept/saftiga-choklad--och-apelsinbullar/",
			},
		},
		{
			name: "fifteenspatulas.com",
			in:   "https://www.fifteenspatulas.com/guacamole/",
			want: models.RecipeSchema{
				AtContext:     atContext,
				AtType:        models.SchemaType{Value: "Recipe"},
				Category:      models.Category{Value: "Appetizer"},
				Cuisine:       models.Cuisine{Value: "Mexican"},
				DatePublished: "2023-06-05T09:44:00+00:00",
				Description: models.Description{
					Value: "This Homemade Guacamole has the perfect texture and combination of flavors, with chunky " +
						"mashed avocados mixed with fresh lime juice, jalapeno, white onion, tomatoes, and cilantro.",
				},
				Keywords: models.Keywords{Values: "guacamole"},
				Image: models.Image{
					Value: "https://www.fifteenspatulas.com/wp-content/uploads/2013/06/Guacamole-Fifteen-Spatulas-8.jpg",
				},
				Ingredients: models.Ingredients{
					Values: []string{
						"3 ripe avocados",
						"1/4 cup seeded and diced tomato",
						"2 tbsp finely chopped white onion",
						"2 tbsp minced jalapeno",
						"1 tbsp freshly squeezed lime juice",
						"1/2 tsp salt (or to taste)",
						"2 tbsp chopped cilantro",
					},
				},
				Instructions: models.Instructions{
					Values: []string{
						"Cut the avocados in half and remove the pits.",
						"Scoop the avocado flesh out into a bowl, and mash the avocado with a fork, leaving plenty " +
							"of chunky, unmashed bits of avocado.",
						"Add the tomato, onion, jalapeno, lime juice, and salt, then gently stir to combine.",
						"Gently fold in the cilantro.",
						"Taste the guacamole and adjust to your tastes (you may desire more salt, or more acidity), " +
							"then serve. Enjoy!",
					},
				},
				Name: "Guacamole",
				NutritionSchema: models.NutritionSchema{
					Calories:      "168 kcal",
					Carbohydrates: "10 g",
					Fat:           "15 g",
					Fiber:         "7 g",
					Protein:       "2 g",
					SaturatedFat:  "2 g",
					Servings:      "1",
					Sodium:        "307 mg",
					Sugar:         "2 g",
				},
				PrepTime: "PT10M",
				Yield:    models.Yield{Value: 4},
				URL:      "https://www.fifteenspatulas.com/guacamole/",
			},
		},
		{
			name: "finedininglovers.com",
			in:   "https://www.finedininglovers.com/recipes/main-course/szechuan-chicken",
			want: models.RecipeSchema{
				AtContext: atContext,
				AtType:    models.SchemaType{Value: "Recipe"},
				Category: models.Category{
					Value: "Main Course",
				},
				Description: models.Description{
					Value: "<p>Szechuan Chicken is a spicy, crispy chicken recipe from Sichuan Region in China: discover " +
						"the original recipe and try to make it at home.</p>",
				},
				Image: models.Image{
					Value: "https://www.finedininglovers.com/sites/g/files/xknfdk626/files/2022-03/TP_SZECHUAN_CHICKEN_COM.jpg",
				},
				Ingredients: models.Ingredients{
					Values: []string{
						"Chicken",
						"White onion",
						"Red pepper",
						"Yellow bell peppers",
						"Chilli",
						"Ginger",
						"Cane sugar",
						"Cornstarch",
						"Garlic powder",
						"Dark soy sauce",
						"Sesame Oil",
						"Chicken stock",
						"Szechuan pepper",
						"Sunflower oil",
						"Kosher salt",
					},
				},
				Instructions: models.Instructions{
					Values: []string{
						"Start with the chicken. Cut it into pieces of equal size that are not too small so that the meat " +
							"can have the same cooking time. Put the cornstarch in a bowl (set aside a spoon) and toss " +
							"the chicken in the flour.\nIn a non-stick pan",
						"pour a drizzle of seed oil and when hot",
						"add the floured chicken. Cook the chicken nuggets for 3/4 minutes so that they become crunchy " +
							"and amber on the outside. In the meantime",
						"wash the peppers",
						"remove the stalk",
						"seeds and internal filaments",
						"then cut them into small pieces. Peel the onion and cut this into small pieces of the same size " +
							"as the peppers.\nWhen the chicken is well browned",
						"remove it from the pan and keep it warm in a covered bowl. In the same pan",
						"add the peppers",
						"onion",
						"whole chillies",
						"garlic powder and freshly grated ginger. Season with salt and bring to the fire to dry the " +
							"vegetables. It will take about 5 minutes on high heat. If you have decided to add it",
						"now is the time to add the Szechuan pepper.\nMeanwhile",
						"prepare the sauce by combining the soy sauce",
						"sesame oil",
						"chicken broth",
						"sugar and a tablespoon of cornstarch in a bowl. Mix very well.\nAdd the chicken to the vegetables " +
							"and cook for another 5 minutes",
						"then add the sauce and mix well. After 30 seconds your super creamy and spicy Szechuan chicken " +
							"is ready. Serve it alone or accompanied with white rice. If you like",
						"you can add chopped fresh parsley or coriander. Tip & Tricks The use of corn starch instead of normal flour",
						"both in the flour and in the preparation of the sauce",
						"gives the dish a creamy texture. Origins This recipe has its origins in the Szechuan region",
						"a place where every dish is prepared with spicy hints. Also typical of this area is Szechuan pepper",
						"a fragrant berry that is added to dishes dried or in powder. Variants For a slightly less " +
							"spicy Szechuan chicken",
						"you can remove the seeds from the peppers or add less during cooking. To give a note of freshness " +
							"you can add the grated lime or lemon zest just before serving.",
					},
				},
				Keywords: models.Keywords{Values: "Tried and Tasted,Channel: Food & Drinks"},
				Name:     "Szechuan Chicken",
				NutritionSchema: models.NutritionSchema{
					Servings: "4",
				},
				Yield: models.Yield{Value: 4},
				URL:   "https://www.finedininglovers.com/recipes/main-course/szechuan-chicken",
			},
		},
		{
			name: "fitmencook.com",
			in:   "https://fitmencook.com/rosemary-blue-cheese-turkey-sliders/",
			want: models.RecipeSchema{
				AtContext:     atContext,
				AtType:        models.SchemaType{Value: "Recipe"},
				CookTime:      "PT20M",
				DateModified:  "2023-05-16T19:59:35+00:00",
				DatePublished: "2021-09-05T18:05:31+00:00",
				Keywords:      models.Keywords{Values: "meal prep,meat,turkey"},
				Image: models.Image{
					Value: "https://fitmencook.com/wp-content/uploads/2021/09/Rosemary-Blue-Cheese-Turkey-Sliders-2.jpg",
				},
				Ingredients: models.Ingredients{
					Values: []string{
						"Meat",
						"1 1/2 lb 93% lean ground turkey",
						"1 1/2 tablespoons The Fit Cook Everyday Blend",
						"2 tablespoons fresh rosemary, finely chopped",
						"1/3 cup blue cheese crumble",
						"pinch of sea salt & pepper",
						"\n",
						"Sliders",
						"8 (wheat) slider buns",
						"2 roma tomatoes, sliced",
						"1 medium cucumber, diagonally sliced",
						"8 tablespoons Dijon mustard",
						"8 Romaine lettuce leaves",
						"\n",
						"Quick Caramelized Onions (OPTIONAL)",
						"3 tablespoons olive oil",
						"1 large white onion, sliced",
						"pinch of sea salt & peppe",
						"\n",
					},
				},
				Instructions: models.Instructions{
					Values: []string{
						"In a bowl, mix together the ingredients for the turkey sliders.  Use an ice cream scoop or 1/4 " +
							"cup measuring cup to scoop to measure out each slider to keep them uniform. Alternatively " +
							"you can measure out full sized burgers.",
						"Set a nonstick skillet on medium heat and once hot, spray with avocado or olive oil, then add " +
							"the slider patties. Cook for 4 – 6 minutes on each side, or until the top/bottom are " +
							"browned and the slider is cooked through.",
						"In a carbon steel skillet on medium heat and once hot, add olive oil and onion. As the onions " +
							"saute and caramelize, add a pinch of sea salt to draw out sweetness. If desired, reduce " +
							"the onions in 1/2 cup of (chicken/veggie/beef) broth, white wine, or water. Continue " +
							"cooking until the onions are “wilted,” soft and golden brown, about 15 minutes.",
						"Build the sliders! Toast the buns then add Dijon, lettuce, tomato, cucumber, caramelized onions " +
							"and the burger!",
					},
				},
				Name: "Rosemary Blue Cheese Turkey Sliders",
				NutritionSchema: models.NutritionSchema{
					Calories:      "330cal",
					Carbohydrates: "24g",
					Fat:           "15g",
					Fiber:         "2g",
					Protein:       "21g",
					Sodium:        "670mg",
					Sugar:         "5g",
				},
				PrepTime: "PT5M",
				URL:      "https://fitmencook.com/rosemary-blue-cheese-turkey-sliders/",
			},
		},
		{
			name: "food.com",
			in:   "https://www.food.com/recipe/jim-lahey-s-no-knead-pizza-margherita-382696",
			want: models.RecipeSchema{
				AtContext:     atContext,
				AtType:        models.SchemaType{Value: "Recipe"},
				Category:      models.Category{Value: "Lunch/Snacks"},
				CookTime:      "PT6M",
				DatePublished: "2009-07-24T13:27Z",
				Description: models.Description{
					Value: "This is a great recipe for a simple, thin crust pizza.  It's from Jim Lahey (of no-knead " +
						"bread fame) who now runs a popular NYC pizzeria called Co.  The recipe was printed in New " +
						"York Magazine (Jul 12, 2009).  If you don't have a pizza stone, this works well in a cast " +
						"iron skillet.  The recipe requires very little time and effort but the dough must be started the day before.",
				},
				Keywords: models.Keywords{
					Values: "Cheese,Vegetable,European,Low Cholesterol,Healthy,Kid Friendly,Kosher,Broil/Grill,< 60 Mins," +
						"Oven,Beginner Cook,Easy,Inexpensive",
				},
				Image: models.Image{
					Value: "https://img.sndimg.com/food/image/upload/q_92,fl_progressive,w_1200,c_scale/v1/img/recipes/38/" +
						"26/96/WW42zdf3SqiQ13mzb97U_0S9A9717.jpg",
				},
				Ingredients: models.Ingredients{
					Values: []string{
						"3  cups  all-purpose flour or 3  cups  bread flour, more for dusting",
						"1/4 teaspoon  instant yeast",
						"1 1/2 teaspoons  salt",
						"1 1/4 cups  water",
						"1   vine-ripened tomatoes (about 5 oz.) or 1    heirloom tomato (about 5 oz.)",
						"1  pinch  salt",
						"1/4 teaspoon  extra virgin olive oil",
						"5  tablespoons  tomato sauce",
						"2  ounces  buffalo mozzarella (about 1/4 ball)",
						"basil leaves",
						"1  tablespoon  extra virgin olive oil",
						"salt",
					},
				},
				Instructions: models.Instructions{
					Values: []string{
						"To make dough: In a large bowl, mix the flour, yeast, and salt. Add water and stir until blended " +
							"(the dough will be very sticky). Cover with plastic wrap and let rest for 12 to 24 hours in a warm " +
							"spot, about 70 degrees.",
						"Place the dough on a lightly floured work surface and sprinkle the top with flour. Fold the dough " +
							"over on itself once or twice, cover loosely with plastic wrap, and let rest for 15 minutes.",
						"Shape the dough into 3 or 4 balls, depending on how thick you want the crust. Generously sprinkle " +
							"a clean cotton towel with flour and cover the dough with it. Let the dough rise for 2 hours.",
						"To make sauce: Blanch tomato for 5 seconds in boiling water and quickly remove. Allow to cool to " +
							"the touch. Peel the skin with your hands and quarter the tomato. Remove the jelly and seeds, and reserve " +
							"in a strainer or fine sieve. Strain the jelly to remove seeds, and combine resulting liquid in a " +
							"bowl with the flesh of the tomatoes. Proceed to crush the tomatoes with your hands. Add salt " +
							"and olive oil and stir.",
						"To make pizza: Place pizza stone on the middle rack of the oven and preheat on high broil. Stretch " +
							"or toss the dough into a disk approximately 10 inches in diameter. Pull rack out of oven and place the " +
							"dough on top of the preheated pizza stone. Drizzle 5 generous tablespoons of sauce over the dough, and " +
							"spread evenly. Try to keep the sauce about &frac12; inch away from the perimeter of the dough. Break " +
							"apart or slice the buffalo mozzarella and arrange over the dough. Return rack and pizza stone to the " +
							"middle of the oven and broil for approximately 6 minutes. Remove and top with basil, olive oil, and salt.",
					},
				},
				Name: "Jim Lahey’s No-Knead Pizza Margherita",
				NutritionSchema: models.NutritionSchema{
					Calories:      "569.4",
					Carbohydrates: "98.9",
					Cholesterol:   "15",
					Fat:           "10.5",
					Fiber:         "4.3",
					Protein:       "17.9",
					SaturatedFat:  "3.4",
					Sodium:        "1472",
					Sugar:         "2.7",
				},
				PrepTime: "PT30M",
				URL:      "https://www.food.com/recipe/jim-lahey-s-no-knead-pizza-margherita-382696",
			},
		},
		{
			name: "food52.com",
			in:   "https://food52.com/recipes/7930-orecchiette-with-roasted-butternut-squash-kale-and-caramelized-red-onion",
			want: models.RecipeSchema{
				AtContext:     atContext,
				AtType:        models.SchemaType{Value: "Recipe"},
				Category:      models.Category{Value: "Appetizer"},
				CookTime:      "PT1H0M",
				Cuisine:       models.Cuisine{Value: "Italian"},
				DateModified:  "2021-04-17 19:58:03 -0400",
				DatePublished: "2010-11-21 19:54:38 -0500",
				Description: models.Description{
					Value: "This recipe is for the butternut squash lover. This orecchiette recipe is yet another reason " +
						"to love squash (and kale, and adorably shaped pasta).",
				},
				Keywords: models.Keywords{
					Values: "Onion, Kale, Goat Cheese, Butternut Squash, Sage, Milk/Cream, Nutmeg, Pasta, Fall",
				},
				Image: models.Image{
					Value: "https://images.food52.com/DTUlTEnTqJKQiU_rOoW8NCh9Dhc=/1200x1200/" +
						"7072c89a-4a7f-412b-b60e-f24d4fcdd1eb--2014-1014_orecchiette-with-roasted-" +
						"butternut-squash-kale-carmelized-onion-012.jpg",
				},
				Ingredients: models.Ingredients{
					Values: []string{
						"1 large butternut squash, cut into small cubes, divided",
						"4 tablespoons extra-virgin olive oil, divided",
						"Kosher salt and pepper",
						"1 pinch cayenne pepper",
						"1/4 teaspoon ground nutmeg",
						"1 red onion, sliced thinly",
						"1/2 pound orecchiette",
						"1 or 2 garlic cloves, minced",
						"2 cups chicken broth, divided",
						"1 bunch kale",
						"1/2 cup white wine",
						"1/2 cup heavy cream",
						"1 ounce goat cheese, optional",
						"1 tablespoon chopped fresh sage",
						"Parmesan cheese, to serve",
					},
				},
				Instructions: models.Instructions{
					Values: []string{
						"Preheat oven to 425° F. Toss all but 1 cup of the butternut squash with 1 tablespoon olive " +
							"oil, salt, pepper, a pinch of cayenne pepper, and the nutmeg. Roast until butternut " +
							"squash pieces are tender and caramelized, about 30 minutes. Set aside.",
						"Heat 1 tablespoon olive oil in a medium saucepan over low heat. Cook sliced red onions " +
							"until caramelized, about 30 minutes. Set aside.",
						"Heat a pot of water over high heat until boiling. Salt water generously. Cook orecchiette " +
							"according to package instructions until al dente.",
						"Meanwhile, heat another tablespoon of olive oil in a heavy pan over medium-high heat. Cook " +
							"the remaining cup of butternut squash for approximately 3 minutes. Add garlic and " +
							"cook for another minute. Add 1/2 cup of the chicken broth and cook until broth is almost completely absorbed.",
						"Remove the middle stems from the kale and roughly chop the leaves. Add kale to butternut " +
							"squash and stir until kale has softened. Add caramelized red onions.",
						"Add white wine and cook for 2 minutes. Add remaining chicken broth and reduce, about 10 minutes. ",
						"Turn heat to low and add the heavy cream. When the pasta is al dente, add it to the pan with the " +
							"sauce. Add the roasted butternut squash.",
						"Loosen sauce with pasta water if needed. Sprinkle with goat cheese (optional), sage, and Parmesan cheese.",
					},
				},
				Name:     "Orecchiette With Roasted Butternut Squash, Kale, &amp;amp; Caramelized Red Onion",
				PrepTime: "PT0H15M",
				Yield:    models.Yield{Value: 4},
				URL: "https://food52.com/recipes/7930-orecchiette-with-roasted-butternut-squash-kale-" +
					"and-caramelized-red-onion",
			},
		},
		{
			name: "foodandwine.com",
			in:   "https://www.foodandwine.com/recipes/garlic-salmon-with-sheet-pan-potatoes",
			want: models.RecipeSchema{
				AtContext:     atContext,
				AtType:        models.SchemaType{Value: "Recipe"},
				Category:      models.Category{Value: "Dinner"},
				DateModified:  "2023-08-02T10:26:55.248-04:00",
				DatePublished: "2022-03-25T10:20:00.000-04:00",
				Description: models.Description{
					Value: "For this sheet pan dinner, baby potatoes, red onion, and spring onions get a head start in a hot oven, before they are joined by a side of salmon, slathered with mustard and drizzled with toasted garlic oil, which cooks alongside the vegetables for a seamless final presentation. Sommelier Erin Miller, of Charlie Palmer&#39;s Dry Creek Kitchen in Healdsburg, California, who provided the inspiration for this dish, notes that it tastes even better when served with a great wine. She recommends a glass of Hirsch Vineyards Raschen Ridge Sonoma Coast Pinot Noir, noting, &#34;The bright acidity of the Hirsch Pinot Noir is a perfect foil for the fresh, fatty fish and flavors of garlic and lemon.&#34;",
				},
				Image: models.Image{
					Value: "https://www.foodandwine.com/thmb/n_1_sUoINm5kjE771S6sKwU7Xw0=/1500x0/filters:no_upscale():max_bytes(150000):strip_icc()/sizzling-garlic-salmon-with-sheet-pan-potatoes-FT-RECIPE0422-eb5c9402ddd44d81aa471177c60bcfaa.jpg",
				},
				Ingredients: models.Ingredients{
					Values: []string{
						"1.5 pounds baby yellow potatoes, halved lengthwise",
						"1 large red onion, cut into 1/2-inch wedges",
						"5 spring onions (about 6 ounces), trimmed and halved lengthwise",
						"0.333 cup plus 2 tablespoons extra-virgin olive oil, divided",
						"2 teaspoons kosher salt, divided",
						"0.5 teaspoon black pepper, divided",
						"0.25 cup Dijon mustard, divided",
						"1 (2-pound) skin-on or skinless salmon side",
						"3 garlic cloves, thinly sliced (about 1 1/2 tablespoons)",
						"2 tablespoons chopped fresh tarragon",
						"2 tablespoons chopped fresh chives",
						"Lemon wedges, for serving",
					},
				},
				Instructions: models.Instructions{
					Values: []string{
						"Preheat oven to 425°F. Toss together potatoes, red onion, spring onions, and 2 tablespoons oil on a " +
							"large rimmed baking sheet; spread in an even layer. Sprinkle evenly with 1 teaspoon salt " +
							"and 1/4 teaspoon pepper. Roast in preheated oven until vegetables begin to brown, about 20 " +
							"minutes. Remove from oven, and reduce oven temperature to 325°F.",
						"Add 1 tablespoon mustard to vegetable mixture on baking sheet; toss to coat. Push vegetables to long " +
							"edges of baking sheet. Place salmon, skin side down, lengthwise in middle of baking sheet. " +
							"Spread salmon with remaining 3 tablespoons mustard; sprinkle with remaining 1 teaspoon salt " +
							"and remaining 1/4 teaspoon pepper.",
						"Heat remaining 1/3 cup oil in a large skillet over medium-high. Add garlic; cook, stirring often, " +
							"until garlic is fragrant and light golden brown, about 2 minutes. Pour hot oil mixture over " +
							"salmon on baking sheet. Roast at 325°F until salmon flakes easily with a fork and vegetables " +
							"are tender, 12 to 15 minutes. Remove from oven. Transfer salmon to a platter; sprinkle with " +
							"tarragon and chives. Transfer vegetables to a bowl. Serve salmon and vegetables with lemon wedges.",
					},
				},
				Name:  "Sizzling Garlic Salmon with Sheet Pan Potatoes",
				Yield: models.Yield{Value: 4},
				URL:   "https://www.foodandwine.com/recipes/garlic-salmon-with-sheet-pan-potatoes",
			},
		},
		{
			name: "foodrepublic.com",
			in:   "https://www.foodrepublic.com/recipes/hand-cut-burger/",
			want: models.RecipeSchema{
				AtContext: atContext,
				AtType:    models.SchemaType{Value: "Recipe"},
				Name:      "How To Cook A Hand-Cut Burger",
				Description: models.Description{
					Value: "When you don't have a meat grinder, but still want a nice, juicy burger, this recipe for a hand-cut burger has a trick you'll use over and over again.",
				},
				Image: models.Image{
					Value: "https://www.foodrepublic.com/img/gallery/hand-cut-burger/intro-import.jpg",
				},
				DateModified:  "2018-06-07T23:13:57+00:00",
				DatePublished: "2018-06-08T15:00:40+00:00",
				Ingredients: models.Ingredients{
					Values: []string{
						"1 (1 1/2-pound) boneless rib-eye steak (preferably dry-aged)",
						"2 cups unsalted butter or rendered beef tallow, plus 2 tablespoons unsalted butter",
						"1 white onion",
						"1 (5-ounce) piece horseradish (2 to 3 inches)",
						"2 tablespoons buttermilk",
						"kosher salt",
						"4 Pain de Mie Buns or 8 slices soft slab bread",
						"16 to 24 dill pickle slices",
					},
				},
				Instructions: models.Instructions{
					Values: []string{
						"Chill the steak in the freezer until firm to the touch but not frozen, 15 to 20 minutes. Cut the steak into 1⁄4-inch-thick slices, then slice into 1⁄4-inch-thick strips, and then into 1⁄4-inch cubes. Remove the sinew and connective tissue but keep the fat.",
						"Divide the beef into four equal balls. Put a sheet of plastic wrap over a 4-inch ring mold on a cutting board or other hard surface. Put a ball in the middle of the mold and gently press down with the palm of your hand, forming a patty that is 4 inches wide. Pop it out with the plastic wrap. Put the patties on a large dish or small baking sheet and refrigerate until ready to cook.",
						"Melt 2 cups of the butter in a pot over medium heat. (Why yes, that is a lot of butter, but it’s used to fully submerge the onion while it cooks; you will not eat 2 cups of butter in this burger.) Add the onion, turn the heat to low, and gently cook at a bare simmer until the onion is tender, about 20 minutes. The onion should be cooked but still al dente, so there’s some texture and a slight hit of sharpness yet not enough that you’ll taste onion the rest of the day. Remove the onion from the butter and drain on a paper towel.",
						"While the onion cooks, make a horseradish sauce. In a bowl, mix the grated horseradish with the buttermilk and a pinch of salt. Stir to combine and refrigerate until ready to use.",
						"Before you begin cooking the burgers, get the buns toasting. Heat a cast-iron skillet or similar surface over medium-low heat. Slice the buns in half horizontally. Smear the remaining 2 tablespoons of butter on the buns and place, butter side down, on the hot surface, working in batches if necessary. Toast until golden brown, 6 to 8 minutes, adjusting the heat if necessary. You want to do your best to time their completion to the burger cooking.",
						"While the buns toast, cook the patties. Heat a cast-iron skillet or grill over high heat. Use a spatula to handle the patty—it will be loose, so be careful. Salt both sides of each patty and put them on the hot skillet. Cook on one side, about 1 minute, then flip the patties and cook until rare, another minute.",
						"Place a patty on a bottom bun and top with some pickles and onions. Slather 1 1⁄2 teaspoons horseradish sauce on the top bun and cap it off. Repeat.",
					},
				},
				URL: "https://www.foodrepublic.com/recipes/hand-cut-burger/",
			},
		},
		{
			name: "forksoverknives.com",
			in:   "https://www.forksoverknives.com/recipes/vegan-snacks-appetizers/crispy-buffalo-cauliflower-bites/",
			want: models.RecipeSchema{
				AtContext: atContext,
				AtType:    models.SchemaType{Value: "Recipe"},
				Category: models.Category{
					Value: "Appetizers",
				},
				CookTime:      "PT0D0H35M",
				DatePublished: "2017-01-27 15:29:56",
				Description: models.Description{
					Value: "It took a lot of trial and error to find the right coating that would not draw out the moisture " +
						"and would make the florets crisp, so I am pleased that it has turned out to be a very " +
						"simple recipe. You will not need to add salt as the sauces have enough salt to season " +
						"them. Either a smoky barbecue sauce or Frank’s hot sauce would work well, but if you " +
						"are like me and prefer sweet and spicy, then try a little bit of both. Serve with " +
						"ranch or Caesar dressing on the side if you wish, or whip up a batch of Spinach Ranch " +
						"Dip. Note: The buffalo cauliflower bites will get softer once they are coated with " +
						"the sauce, so hold off tossing until the very last minute",
				},
				Image: models.Image{
					Value: "https://www.forksoverknives.com/uploads/FOK_Coliflower8384-WP.jpg?auto=webp",
				},
				Ingredients: models.Ingredients{
					Values: []string{
						"⅔ cup brown rice flour",
						"2 tablespoons almond flour",
						"1 tablespoon tomato paste",
						"2 teaspoons garlic powder",
						"2 teaspoons onion powder",
						"2 teaspoons smoked paprika",
						"1 teaspoon dried parsley",
						"1 head cauliflower, cut into 2-inch florets",
						"⅓ cup Frank’s hot sauce or barbecue sauce",
						"<a href=\"https://www.forksoverknives.com/recipes/spinach-ranch-dip/#gs.MasyuTc\" target=\"_blank\" " +
							"rel=\"noopener\">Spinach Ranch Dip</a>",
					},
				},
				Instructions: models.Instructions{
					Values: []string{
						"Preheat oven to 450°F. Line 2 baking sheets with parchment paper.",
						"Combine the brown rice flour, almond flour, tomato paste, garlic powder, onion powder, paprika, " +
							"parsley, and ⅔ cup of water in a blender. Puree until the batter is smooth and thick. Transfer " +
							"to a bowl and add the cauliflower florets; toss until the florets are well coated with the batter.",
						"Arrange the cauliflower in a single layer on the prepared baking sheets, making sure that the " +
							"florets do not touch one another. Bake for 20 to 25 minutes, until crisp on the edges. They " +
							"will not get crispy all over while still in the oven.",
						"Remove from the heat and let stand for 3 minutes to crisp up a bit more. Transfer to a bowl " +
							"and drizzle with the sauce. Serve immediately.",
					},
				},
				Name:     "Crispy Buffalo Cauliflower Bites",
				PrepTime: "PT0D0H0M",
				Yield:    models.Yield{Value: 6},
				URL:      "https://www.forksoverknives.com/recipes/vegan-snacks-appetizers/crispy-buffalo-cauliflower-bites/",
			},
		},
		{
			name: "franzoesischkochen.de",
			in:   "https://www.franzoesischkochen.de/navettes-aus-marseille/",
			want: models.RecipeSchema{
				AtContext:     atContext,
				AtType:        models.SchemaType{Value: "Recipe"},
				Category:      models.Category{Value: "Alte französische Rezepte"},
				CookTime:      "PT30M",
				Cuisine:       models.Cuisine{Value: "French"},
				DateModified:  "2022-03-09T15:14:09+02:00",
				DatePublished: "2022-03-09T15:14:09+02:00",
				Description:   models.Description{Value: "Ein einfaches Rezept mit Schritt-für-Schritt-Fotos und vielen Tipps über das Thema: Navettes aus Marseille"},
				Keywords:      models.Keywords{Values: "Alte französische Rezepte,Einfachste Rezepte,In der Boulangerie,Kekse &amp; Plätzchen,Provence,Traditionelle Rezepte,Typisch französische Kuchen"},
				Image: models.Image{
					Value: "https://www.franzoesischkochen.de/wp-content/uploads/2022/01/Navette-orangenbluettenwasser.jpg",
				},
				Ingredients: models.Ingredients{
					Values: []string{
						"1 Ei",
						"160g Mehl T55",
						"20 g Olivenöl",
						"60 g Honig",
						"1 TL Orangenblütenwasser",
						"1 kleine Prise Salz. Milch zum Pinseln.",
					},
				},
				Instructions: models.Instructions{
					Values: []string{
						"1- Das Ei mit dem Honig, Orangenblütenwasser und Olivenöl rühren....",
						"2- Kleine Teigkugels (25 bis 30g ) nehmen, wie Knete in die Länge ausrollen. Sie sollten eine " +
							"ovale Form oder besser eine Schiff-Form bekommen. Die geformten Navette auf " +
							"einem mit Backpapier belegten Backblech legen. Die Navettes in der Mitte entlang anschneiden.",
						"3- Backen: 180°C Umluft für 12 bis 15 Minuten. (es kommt darauf an, ob Ihr eure Navettes brauner " +
							"mögt wie ich oder lieber hell!)",
					},
				},
				Name:     "Navettes aus Marseille",
				PrepTime: "PT60M",
				URL:      "https://www.franzoesischkochen.de/navettes-aus-marseille/",
			},
		},
		{
			name: "giallozafferano.com",
			in:   "https://www.giallozafferano.com/recipes/Christmas-spice-cookies.html",
			want: models.RecipeSchema{
				AtContext:     atContext,
				AtType:        models.SchemaType{Value: "Recipe"},
				Category:      models.Category{Value: "Sweets and desserts"},
				CookTime:      "PT15M",
				DateModified:  "2022-12-10 00:00:00",
				DatePublished: "2022-12-10 00:00:00",
				Description: models.Description{
					Value: "Christmas spice cookies are shortcrust pastry sweets flavored with vanilla, ginger and cinnamon garnished with a white chocolate ganache!",
				},
				Keywords: models.Keywords{
					Values: "recipes, recipe, italian cuisine, how to cook, Christmas spice cookies",
				},
				Image: models.Image{
					Value: "https://www.giallozafferano.com/images/260-26068/Christmas-spice-cookies_1200x800.jpg",
				},
				Ingredients: models.Ingredients{
					Values: []string{
						"Flour 00 2 &frac14; cups",
						"Butter cold ½ cup",
						"Brown sugar ½ cup",
						"Eggs 1",
						"Honey ¾ tbsp",
						"Fine salt &frac14; tbsp",
						"Vanilla extract 1 &frac14; tsp",
						"Powdered ginger 1 &frac14; tsp",
						"Cinnamon powder 1 &frac14; tsp",
						"Baking powder 1 tsp",
						"White chocolate 1 &frac14; cup",
						"Heavy cream cold ½ cup",
						"Vanilla extract 1 tbsp",
						"Powdered ginger to taste",
						"Chopped hazelnuts to taste",
						"Colored sprinkles",
					},
				},
				Instructions: models.Instructions{
					Values: []string{
						"To prepare the Christmas spice cookies , start with the shortcrust pastry: in a large bowl, pour the flour, the diced butter 1 , the brown sugar 2 and the cinnamon 3 .",
						"Also add the ginger powder 4 , the salt 5 and the baking powder 6 .",
						"Start kneading with your hands 7 until you get a \"crumble\" consistency 8 . At this point add the honey 9 .",
						"Also add the egg 10 and the vanilla extract 11 . Continue kneading 12 .",
						"You will have to obtain a smooth and homogeneous mixture 13 . Transfer it on a surface and compact it to form a loaf 14 ,and place it on a sheet of parchment paper 15 .",
						"Cover with another sheet of parchment paper and immediately roll out the pastry to a thickness of about .20\" (5 mm) 16 . Transfer to the fridge and let it rest for at least 30 minutes. After that, make the cookies using a pastry ring with a diameter of 2.33\" (6 cm) 17 . Place them gradually on a baking sheet lined with parchment paper, spacing them apart 18 : with these doses you will get about 30 biscuits. Bake in a preheated static oven at 340 \u00b0 F (170\u00b0C) for about 15-20 mins.",
						"Meanwhile prepare the ganache. Coarsely chop the white chocolate 19 and melt it in the microwave or in a bain-marie. Pour the vanilla extract 20 , the ginger powder and the cold cream 21 .",
						"Mix well and transfer into a pastry bag without a nozzle 22 . Place in the refrigerator to firm up. As soon as they are ready, take the cookies out of the oven and let them cool completely 23 . Once cold you can decorate them with the ganache 24 . If the ganache is too solid, keep it at room temperature for a few minutes.",
						"Garnish with chopped hazelnuts 25 or colored sprinkles 26 . Christmas spice cookies are ready to enjoy 27 !",
					},
				},
				Name:     "Christmas spice cookies",
				PrepTime: "PT30M",
				Yield:    models.Yield{Value: 30},
				URL:      "https://www.giallozafferano.com/recipes/Christmas-spice-cookies.html",
			},
		},
		{
			name: "gimmesomeoven.com",
			in:   "https://www.gimmesomeoven.com/miso-chocolate-peanut-butter-cornflake-bars-gimme-some-oven/",
			want: models.RecipeSchema{
				AtContext:     atContext,
				AtType:        models.SchemaType{Value: "Recipe"},
				DatePublished: "2022-02-18",
				Description: models.Description{
					Value: "These no-bake miso chocolate peanut butter cornflake bars are quick and easy to whip up and " +
						"ridiculously delicious. See notes above for modifications to make this recipe gluten-free and/or vegan.",
				},
				Keywords: models.Keywords{Values: ""},
				Image: models.Image{
					Value: "https://www.gimmesomeoven.com/wp-content/uploads/2022/02/Peanut-Butter-Cornflake-Bars-8-1-225x225.jpg",
				},
				Ingredients: models.Ingredients{
					Values: []string{
						"1 tablespoon coconut oil",
						"1 cup natural creamy peanut butter",
						"1/4 cup honey",
						"2 tablespoons white (shiro) miso paste",
						"1 teaspoon vanilla extract",
						"5 cups (5 ounces) cornflakes",
						"1 1/2 cups (9 ounces) semisweet chocolate chips",
						"flaky sea salt",
					},
				},
				Instructions: models.Instructions{
					Values: []string{
						"Line a 9 x 9-inch baking dish with parchment paper.",
						"Heat the coconut oil in a large saucepan over medium-low heat until melted. Add peanut butter, honey, " +
							"miso paste, vanilla extract, and stir steadily for 2 minutes or until lightly warmed. Remove from " +
							"heat and stir in the cornflakes until they are evenly coated with the peanut butter sauce. Transfer the " +
							"cornflake mixture to the prepared pan and use a silicone spatula (or the flat bottom of a glass or a " +
							"measuring cup) to press the mixture down firmly and evenly until it is very compact.",
						"Heat the chocolate chips in a double boiler (or in the microwave in 10-second intervals) until completely " +
							"melted, being careful not to overcook and burn the chocolate. Immediately spread the melted chocolate " +
							"in an even layer over the cornflake mixture, using a spoon to create pretty swirls if you&#8217;d like.",
						"Refrigerate the bars for 2 to 3 hours or until firm.",
						"Serve. When you’re ready to serve the bars, carefully lift up the parchment and transfer the entire batch " +
							"to a cutting board. Use a chef’s knife to carefully cut the bars into squares. Serve immediately and enjoy!",
					},
				},
				Name:     "Miso Chocolate Peanut Butter Cornflake Bars",
				PrepTime: "PT20M",
				Yield:    models.Yield{Value: 16},
				URL:      "https://www.gimmesomeoven.com/miso-chocolate-peanut-butter-cornflake-bars-gimme-some-oven/",
			},
		},
		{
			name: "globo.com",
			in:   "https://receitas.globo.com/cheesecake-com-geleia-de-frutas-vermelhas-do-bbb-22.ghtml",
			want: models.RecipeSchema{
				AtContext:     atContext,
				AtType:        models.SchemaType{Value: "Recipe"},
				Category:      models.Category{Value: "Tortas e bolos"},
				CookingMethod: models.CookingMethod{Value: "Americana"},
				Cuisine:       models.Cuisine{Value: "Americana"},
				DateModified:  "2023-03-07T19:38:51.176Z",
				DatePublished: "2022-03-30T19:43:06.164Z",
				Description: models.Description{
					Value: "Veja como fazer cheesecake com geleia de frutas vermelhas. Receita é feita em camadas, sendo a massa " +
						"feita com biscoito maisena e manteiga; recheio com leite condensado, cream cheese, ovos, sal e creme " +
						"de leite e cobertura com geleia caseira de frutas vermelhas, feita com morango, blueberry, amora, " +
						"framboesa, água e açúcar.",
				},
				Keywords: models.Keywords{
					Values: "cheesecake, lanche da tarde, recepção, aniversário",
				},
				Image: models.Image{
					Value: "https://s2-receitas.glbimg.com/XjnpBAqPQSKGTlZ6fYujfyRZ0lA=/1200x/smart/filters:cover():strip_icc()/i.s3.glbimg.com/v1/AUTH_1f540e0b94d8437dbbc39d567a1dee68/internal_photos/bs/2022/s/3/rsLexpSU6nXAgmuhfKNw/cheesecake-com-geleia-de-frutas-vermelhas-bbb22-1.jpg",
				},
				Ingredients: models.Ingredients{
					Values: []string{
						"1 pacote de biscoito maisena ",
						"3 colheres de sopa de margarina",
						"1 lata de leite condensado ",
						"300 gramas de cream cheese ",
						"2 ovos",
						"1 pitada de sal",
						"1 caixa de creme de leite ",
						"Meia xícara de morango",
						"Meia xícara de blueberry, também conhecido como mirtilo",
						"Meia xícara de amora",
						"Meia xícara de framboesa ",
						"1 xícara de água ",
						"1 xícara de açúcar ",
					},
				},
				Instructions: models.Instructions{
					Values: []string{
						"Com auxílio de um mixer, processador ou liquidificador triture os biscoitos até formar uma farofinha " +
							"bem fininha.",
						"Coloque a farofinha numa tigela e misture com a manteiga, até formar uma massa bem úmida.",
						"Forre uma forma de fundo removível com a massa.",
						"Leve ao forno com a temperatura de 180 graus Celsius por aproximadamente 10 minutos ou até ficar douradinha.",
						"Bata o leite condensado, o creme de leite, os ovos e o sal na batedeira. Reserve.",
						"Coloque todos os ingredientes na panela para cozinhar em fogo baixo, e deixe até formar uma geleia. " +
							"Deixe esfriar.",
						"Ao retirar a massa da cheesecake do forno deixe esfriar.",
						"Quando a massa estiver fria, acrescente o recheio feito previamente.",
						"Volte a torta para o forno por 25 a 30 minutos a 180 graus Celsius ou até ficar douradinha.",
						"Deixe a torta esfriar e coloque a cobertura de geleia de frutas vermelhas.",
						"Leve a cheesecake para gelar por mais ou menos duas horas.",
					},
				},
				Name:  "Cheesecake com geleia de frutas vermelhas do 'BBB 22'",
				Yield: models.Yield{Value: 4},
				URL:   "https://receitas.globo.com/cheesecake-com-geleia-de-frutas-vermelhas-do-bbb-22.ghtml",
			},
		},
		{
			name: "gonnawantseconds.com",
			in:   "https://www.gonnawantseconds.com/beef-tomato-macaroni-soup/#wprm-recipe-container-15941",
			want: models.RecipeSchema{
				AtContext:     atContext,
				AtType:        models.SchemaType{Value: "Recipe"},
				Category:      models.Category{Value: "Soup"},
				CookTime:      "PT30M",
				Cuisine:       models.Cuisine{Value: "American"},
				DatePublished: "2022-03-25T05:00:29+00:00",
				Description: models.Description{
					Value: "This simple but satisfying, hearty Beef and Tomato MacaroniSoup will be a repeat visitor to your dining table when the temperatures drop and appetites grow.",
				},
				Keywords: models.Keywords{
					Values: "beef soup recipes, ground beef recipes, ground beef soup recipes, macaroni soup recipes, tomato soup recipes",
				},
				Image: models.Image{
					Value: "https://www.gonnawantseconds.com/wp-content/uploads/2022/03/Beef-and-Tomato-Macaroni-Soup-01.jpg",
				},
				Ingredients: models.Ingredients{
					Values: []string{
						"2 tablespoons vegetable oil",
						"1 medium yellow onion, (finely chopped)",
						"1 green bell pepper, (finely chopped)",
						"2 cloves garlic, (minced)",
						"1 pound ground beef",
						"2 teaspoons chili powder",
						"2 teaspoons dried oregano",
						"1 teaspoon salt",
						"1/2 teaspoon black pepper",
						"2 (10.75-ounce) cans condensed cream of tomato soup",
						"1 (15-ounces) can diced tomatoes (undrained)",
						"32 ounces beef broth",
						"4 cups water",
						"2 cups uncooked pasta",
					},
				},
				Instructions: models.Instructions{
					Values: []string{
						"Heat the vegetable oil (2 tablespoons) in a large pot over medium-high heat. Add the onions (1), green " +
							"bell pepper (1), and garlic (2 cloves) and saute until the onion mixture begins to soften " +
							"about 5-6 minutes.",
						"Add the ground beef (1 pound), crumbling with a wooden spoon, cook until there is no longer any pink. " +
							"Drain off excess fat.",
						"Add the chili powder (2 teaspoons), oregano (2 teaspoons), salt (1 teaspoon), and pepper (1/2 teaspoon) " +
							"and cook over medium heat for 1-2 minutes.",
						"Add condensed cream of tomato soup (2 (10.75-ounce) cans), diced tomatoes with their juice (1 (15-ounces) " +
							"can), beef broth (32 ounces), and water (4 cups).",
						"Bring to a boil, add pasta (2 cups). Reduce heat and cover and simmer until the pasta is just al dente. " +
							"Adjust seasoning and serve.",
					},
				},
				Name: "Beef and Tomato Macaroni\u00a0Soup",
				NutritionSchema: models.NutritionSchema{
					Calories:       "829 kcal",
					Carbohydrates:  "79 g",
					Cholesterol:    "80.5 mg",
					Fat:            "40 g",
					Fiber:          "8.5 g",
					Protein:        "36 g",
					SaturatedFat:   "16 g",
					Servings:       "1",
					Sodium:         "2643 mg",
					Sugar:          "9.5 g",
					UnsaturatedFat: "7 g",
				},
				PrepTime: "PT10M",
				Yield:    models.Yield{Value: 4},
				URL:      "https://www.gonnawantseconds.com/beef-tomato-macaroni-soup/#wprm-recipe-container-15941",
			},
		},
		{
			name: "greatbritishchefs.com",
			in:   "https://www.greatbritishchefs.com/recipes/babecued-miso-poussin-recipe",
			want: models.RecipeSchema{
				AtContext:     atContext,
				AtType:        models.SchemaType{Value: "Recipe"},
				DatePublished: "2017-03-30T00:00:00Z",
				DateModified:  "2021-07-16T15:18:02.087Z",
				Description: models.Description{
					Value: "Scott Hallsworth's tasty barbecued poussin recipe is packed with bold Japanese flavours.",
				},
				Name: "Barbecued miso poussin with lemon, garlic and chilli dip",
				Image: models.Image{
					Value: "https://media-cdn2.greatbritishchefs.com/media/hpsovny5/img68297.whqc_1426x713q80.jpg",
				},
				Category: models.Category{Value: "Main"},
				CookTime: "PT60M",
				Ingredients: models.Ingredients{
					Values: []string{
						"2 poussin",
						"90g of brown miso paste",
						"50g of caster sugar",
						"40ml of mirin",
						"40ml of sake",
						"1 green chilli, finely chopped",
						"100ml of sake",
						"2 tbsp of dark soy sauce",
						"1 red chilli, finely chopped",
						"2 tsp Tabasco green",
						"2 tsp garlic purée, fresh",
						"2 tsp yuzu juice",
						"2 tbsp of lemon juice",
						"2 tbsp of olive oil",
						"100g of daikon radish",
						"50g of carrots",
						"50g of cucumber",
						"6 mint leaves",
						"10 coriander leaves, with a bit of stem left on",
						"10ml of yuzu juice",
						"soy sauce, to taste",
						"extra virgin olive oil, to taste",
					},
				},
				Instructions: models.Instructions{
					Values: []string{
						"To begin, make the den miso for the marinade. Prepare a bain marie by filling a pan of water and sitting a " +
							"heatproof bowl snugly over the top (ensuring the bottom of the bowl doesn’t touch the water). " +
							"Whisk together the miso, sugar, mirin and sake and pour into the bowl. Cook over a high heat for about " +
							"20 minutes, stirring continuously. Remove and chill",
						"Make the marinade by mixing 100ml of the den miso and the chillies together. Use a sharp knife to cut each " +
							"poussin clean in half and make a couple of score marks, one in the fat part of the drumsticks " +
							"and one in the thighs. Marinate in the miso-chilli marinade for at least 6 hours and up to 12 hours",
						"To make the dip, whisk together all the ingredients, except the oil. Slowly whisk in the oil until " +
							"emulsified. This will keep in the fridge for a month",
						"To make the salad, thinly slice the daikon on a Japanese mandoline and layer the slices in piles of 5 or " +
							"6. Using a knife, shred very thinly. Do the same with the carrot and cucumber and mix together. " +
							"Add the mint and coriander leaves and drizzle with the yuzu, soy sauce and extra virgin olive oil",
						"Set up your barbecue and get the charcoal very hot. Once the flames start to die down a little and the " +
							"embers begin to glow, put your poussins on the grill. If you’re concerned about the poussins not " +
							"being cooked through enough and burning, take off the barbecue and finish cooking in a hot oven, about " +
							"180°C/gas mark 4 for 8–10 minutes. To check the birds are done, insert a thin metal skewer or the sharp " +
							"end of small knife into the thickest part of the thigh, pause for a couple of seconds, then hold the " +
							"skewer to your lip; if its scorching hot they're done",
						"Once cooked, serve with the daikon salad and the dip on the side",
					},
				},
				Keywords: models.Keywords{Values: "easy"},
				URL:      "https://www.greatbritishchefs.com/recipes/babecued-miso-poussin-recipe",
				Yield:    models.Yield{Value: 4},
			},
		},
		{
			name: "halfbakedharvest.com",
			in:   "https://www.halfbakedharvest.com/louisiana-style-chicken-and-rice/",
			want: models.RecipeSchema{
				AtContext:     atContext,
				AtType:        models.SchemaType{Value: "Recipe"},
				Category:      models.Category{Value: "Main Course"},
				CookTime:      "PT50M",
				Cuisine:       models.Cuisine{Value: "American"},
				DatePublished: "2022-03-23T02:00:00+00:00",
				Description: models.Description{
					Value: "One Skillet Louisiana Style Chicken and Rice: has a variety of flavors and textures, yet it&#39;s all " +
						"made in ONE skillet with pantry staple ingredients!",
				},
				Keywords: models.Keywords{Values: "one skillet"},
				Image: models.Image{
					Value: "https://www.halfbakedharvest.com/wp-content/uploads/2022/03/One-Skillet-Louisiana-Style-Chicken-and" +
						"-Rice-1.jpg",
				},
				Ingredients: models.Ingredients{
					Values: []string{
						"2 tablespoons extra virgin olive oil",
						"1 pound boneless chicken breasts or thighs",
						"2 tablespoons cajun seasoning",
						"kosher salt and black pepper",
						"6 tablespoons salted butter",
						"1 lemon, sliced",
						"1/2 cup dry broken spaghetti or angel hair pasta",
						"1 cup long grain rice",
						"1 medium yellow onion, sliced",
						"2 bell peppers, sliced",
						"3-4 cups low sodium chicken broth",
						"3 cloves garlic, chopped",
						"1/2 cup fresh tenders herbs, cilantro + parsley",
						"chili flakes",
					},
				},
				Instructions: models.Instructions{
					Values: []string{
						"Preheat the oven to 425° F.",
						"In a large oven-safe skillet, combine the olive oil, chicken, and cajun seasoning, toss to coat. Set the skillet over high heat. Sear on both sides until golden, 3-5 minutes. During the last 2 minutes of cooking, add 1 tablespoon of butter and lemon slices. Remove everything from the skillet.",
						"Add the rice and pasta. Cook until the rice is toasted, about 1 minute. Add the onion and peppers and continue to cook another 3-4 minutes, then pour in 3 cups broth. Season with salt and pepper. Bring to a boil.",
						"Slide the chicken, lemon slices, and any juices left on the plate back into the skillet. Bring to a boil. Cover the skillet and turn the heat down to the lowest setting possible. Allow the rice to cook 10 minutes, until most of the liquid has cooked into the rice, but not all of it. If needed add more broth. Bake, uncovered for 10-15 minutes or until the chicken is cooked through.",
						"Meanwhile, melt together 5 tablespoons butter, the garlic, and a pinch of chili flakes. Cook until the butter is browning. Stir in the mixed herbs.",
						"Serve the chicken and rice drizzled with garlic butter and topped with fresh herbs.",
					},
				},
				Name: "Skillet Louisiana Style Chicken and Rice",
				NutritionSchema: models.NutritionSchema{
					Calories: "547 kcal",
					Servings: "1",
				},
				PrepTime: "PT20M",
				Yield:    models.Yield{Value: 6},
				URL:      "https://www.halfbakedharvest.com/louisiana-style-chicken-and-rice/",
			},
		},
		{
			name: "hassanchef.com",
			in:   "https://www.hassanchef.com/2022/10/dragon-chicken.html",
			want: models.RecipeSchema{
				AtContext:     atContext,
				AtType:        models.SchemaType{Value: "Recipe"},
				Category:      models.Category{Value: "appetizer"},
				CookTime:      "PT10M",
				Cuisine:       models.Cuisine{Value: "Indo Chinese"},
				DatePublished: "2022-10-14",
				Description: models.Description{
					Value: "Dragon Chicken an appetizer or snacks of Indian Chinese cuisines where deep fried chicken strips are stir fried with a spicy combination of sauces and herbs",
				},
				Keywords: models.Keywords{Values: "Dragon Chicken"},
				Image: models.Image{
					Value: "https://blogger.googleusercontent.com/img/b/R29vZ2xl/AVvXsEjVPyaqbaCDbK5VdlCoe93-7wQjDmM4jVCrnuGlne0QDqUKlwfzat-Z2RS7GSFujClIpZUZIn7Q0-J75jr4LFCkJu_OwOc-YTIw30WnvpC0lH9vhMGjSDE-FmIvvg0m6dv2KlFRo1YcfA804XBHPp1AeOpf0tA0qoMFzWKHo4tSjUtrL_TJ5a7HP24w/s4623/IMG_20220906_222143.webp",
				},
				Ingredients: models.Ingredients{
					Values: []string{
						"1 large Chicken breast",
						"1 teaspoon red chilly paste",
						"1 teaspoon ginger garlic paste",
						"1/4 teaspoon green chilly paste",
						"2 tablespoon cornflour",
						"1 tablespoon beaten egg or egg white",
						"1/2 of each red, yellow and green bell pepper cut into julienne",
						"1/2 of a onion cut into thin slices",
						"Oil for deep frying",
						"1/2 teaspoon chopped garlic",
						"1/2 teaspoon chopped ginger",
						"1 whole red chilly cut into pieces",
						"10 - 12 roasted or golden fried cashew nuts",
						"1/2 teaspoon red chilly sauce",
						"1/2 teaspoon red chilly paste",
						"1/3 teaspoon pepper powder",
						"1/3 teaspoon madras curry powder",
						"Salt as taste",
						"Seasoning powder(optional)",
						"Some chopped green spring onions",
						"1 tablespoon slurry (cornflour)",
						"2 tablespoon cooking oil",
					},
				},
				Instructions: models.Instructions{
					Values: []string{
						"Dragon Chicken is made using boneless chicken breast pieces. Here I have taken only one single chicken breast and sliced it into two parts. Then cut them into thin strips around half inch width. Now add half teaspoon red chilly paste, ginger garlic paste, around 1/4 teaspoon green chilli paste, 2 tablespoon corn flour, 2 tablespoon beaten egg and salt as per your taste. Mix them nicely so that all the chicken strips are get evenly coated. Cover and let them marinated for 5 to 6 minutes to absorb the flavours.",
						"Cut the bell peppers and capsicum into julienne shape and onion into thin slices. Finely chopped the ginger and garlic. You can avoid to use the bell peppers if not available at your pantry, capsicum and onion also yield a good result. You can either roast the cashew nuts or fry them till golden brown. Roughly chopped the green spring onions and cut the whole red chilly into two or three parts.",
						"Heat a kadai or small deep vessel pan with enough cooking oil under medium heat flame. When the oil become hot lower the gas flame to low and carefully add the chicken strips one by one into the hot oil. Fry them until become crisp golden from all sides. Remove them on a absorbent paper to absorb any excess oil.",
						"Heat a non stick pan or Chinese wok with one tablespoon oil under medium flame heat gas fire. Add the bell peppers, capsicum, onion and pieces of whole red chilly. Stir fry them for few seconds in high flame heat. Then add chopped onion, chopped ginger, garlic, green chilly paste and cashew nuts. Stir fry them for one minute.",
						"Now add red chilly paste, red chilly sauce, tomato sauce and give them a good mix.",
						"Add some water and cook the sauces for few seconds. Then add soya sauce, white pepper powder, salt and madras curry powder and cook further for few seconds. Now add the fried chicken strips and mix well with the sauces so that the chicken pieces are get well coated with them.",
						"Pour 2 tablespoon corn flour slurry over the chicken mixture and mix well. Finally sprinkle some chopped spring onions and give a final mix. Remove on a serving plate and serve hot.",
					},
				},
				Name: "Dragon Chicken",
				NutritionSchema: models.NutritionSchema{
					Calories: "640 cal",
					Fat:      "34 g",
					Servings: "1",
				},
				PrepTime: "PT20M",
				Yield:    models.Yield{Value: 1},
				URL:      "https://www.hassanchef.com/2022/10/dragon-chicken.html",
			},
		},
		{
			name: "headbangerskitchen.com",
			in:   "https://headbangerskitchen.com/recipe/keto-chicken-adobo/",
			want: models.RecipeSchema{
				AtContext:     atContext,
				AtType:        models.SchemaType{Value: "Recipe"},
				Category:      models.Category{Value: "Main Dish"},
				CookTime:      "PT25M",
				Cuisine:       models.Cuisine{Value: "Filipino"},
				DatePublished: "2021-10-06T17:23:00+00:00",
				Description: models.Description{
					Value: "A Keto version of a classic Filipino chicken adobo",
				},
				Image: models.Image{
					Value: "https://headbangerskitchen.com/wp-content/uploads/2021/10/CHICKENADOBO-Vertical2.jpg",
				},
				Ingredients: models.Ingredients{
					Values: []string{
						"6 Chicken Thighs (Bone in and Skin On)",
						"1/2 tbsp Avocado Oil",
						"8 cloves garlic (I used small cloves)",
						"1/2 tbsp Whole Black Peppercorns",
						"4 small bay leaves",
						"1/2 tsp Black Pepper Powder",
						"60 ml dark soya sauce or coconut aminos",
						"80 ml cane vinegar or apple cider vinegar",
						"2 tsp Keto sweetener (1:1 sugar substitute)",
						"spring onion greens for garnish",
						"salt as needed",
					},
				},
				Instructions: models.Instructions{
					Values: []string{
						"Get a large, deep skillet on the stove and add a good glug of oil to it. I’m using avocado oil.",
						"Once the oil is hot, place your chicken thighs in the pan skin side down. Let the chicken cook till the " +
							"skin is crispy, about 3-4 minutes.",
						"Using the flat of a knife, smash the eight cloves of garlic.",
						"Flip the chicken over and add in the aromats – the smashed garlic, the whole bay leaves and peppercorns – " +
							"and let everything fry till fragrant, about a minute or two.",
						"Deglaze the pan with a splash of water, then add in the black pepper, soy sauce and apple cider vinegar " +
							"and give it all a good mix. Flip the chicken again and let everything come to a boil.",
						"Cook the chicken for about 10 minutes, flipping the pieces halfway through.",
						"Check the seasoning in the sauce, then add about two teaspoons worth of Keto sweetener. You may need to " +
							"adjust this depending on the sweetener you’re using, so taste, taste, taste before adding any more.",
						"After your 10 minutes are up, remove the chicken from the sauce, and let the sauce continue reducing. You " +
							"want it to become a sticky, syrupy consistency.",
						"Once the sauce is reduced, add the chicken back in and flip it a few times till it’s completely basted and " +
							"bathed in that sauce. Finish with some spring onion greens and serve over plain cauli rice.",
					},
				},
				Name:     "Keto Chicken Adobo",
				PrepTime: "PT5M",
				Yield:    models.Yield{Value: 6},
				URL:      "https://headbangerskitchen.com/recipe/keto-chicken-adobo/",
			},
		},
		{
			name: "hellofresh.com",
			in:   "https://www.hellofresh.com/recipes/creamy-shrimp-tagliatelle-5a8f0fcbae08b52f161b5832",
			want: models.RecipeSchema{
				AtContext:     atContext,
				AtType:        models.SchemaType{Value: "Recipe"},
				Category:      models.Category{Value: "main course"},
				Cuisine:       models.Cuisine{Value: "American"},
				DatePublished: "2018-02-22T18:45:31+00:00",
				Description: models.Description{
					Value: "Pronto! Pronto! You can make this dinner recipe with the lightning speed of an Italian race car. Thanks to fresh tagliatelle, which cooks faster than the dried kind, you arrive at al dente perfection in a matter of minutes. The shrimp and heirloom tomatoes only need a quick toss in the pan, too, becoming tender on the count of uno, due, tre.",
				},
				Keywords: models.Keywords{Values: "Spicy,Dinner Ideas"},
				Image: models.Image{
					Value: "https://img.hellofresh.com/f_auto,fl_lossy,h_640,q_auto,w_1200/hellofresh_s3/image/5a8f0fcbae08b52f" +
						"161b5832-033c9a4a.jpg",
				},
				Ingredients: models.Ingredients{
					Values: []string{
						"2 clove Garlic",
						"2 unit Scallions",
						"1 unit Chili Pepper",
						"10 ounce Heirloom Grape Tomatoes",
						"1 unit Lemon",
						"10 ounce Shrimp",
						"9 ounce Tagliatelle Pasta",
						"4 tablespoon Sour Cream",
						"1 teaspoon Olive Oil",
						"2 tablespoon Butter",
						"Salt",
						"Pepper",
					},
				},
				Instructions: models.Instructions{
					Values: []string{
						"Wash and dry all produce. Bring a large pot of salted water to a boil. Mince garlic. Trim, then thinly " +
							"slice scallions, keeping greens and whites separate. Finely mince chili, removing seeds and ribs " +
							"for less heat. Halve tomatoes. Cut lemon into wedges. Rinse shrimp and pat dry with a paper towel.",
						"Heat a drizzle of olive oil in large pan over medium-high heat. Add garlic, scallion whites, and chili " +
							"(to taste). Cook until fragrant, about 30 seconds. Add shrimp and cook, tossing, until starting to " +
							"turn pink but not quite cooked through, 1-2 minutes. Season with salt and pepper.",
						"Once water is boiling, add tagliatelle to pot. (TIP: If any noodles are stuck together, separate them " +
							"first.) Cook, stirring occasionally, until al dente, 4-5 minutes. Carefully scoop out and reserve ¼ " +
							"cup pasta cooking water, then drain.",
						"Meanwhile, add tomatoes to pan with shrimp. Cook, tossing, until wilted and juicy, 2-3 minutes. Season " +
							"with salt and pepper. Remove from heat and set aside until pasta is ready. TIP: If you like it extra " +
							"hot, add any remaining chili (to taste) at this point.",
						"Once tagliatelle is done cooking, return pan with shrimp and tomatoes to medium heat and add tagliatelle " +
							"and 2 TBSP butter. Toss to combine and melt butter. Season with salt and pepper.",
						"Remove pan from heat and stir in sour cream, a squeeze of lemon, and as much pasta cooking water as needed " +
							"to reach a saucy consistency. Season with salt and pepper. Divide between plates or bowls and garnish " +
							"with scallion greens. Serve with lemon wedges on the side for squeezing over.",
					},
				},
				Name: "Creamy Shrimp Tagliatelle with Heirloom Tomatoes, Garlic, and Chili",
				NutritionSchema: models.NutritionSchema{
					Calories:      "750 kcal",
					Carbohydrates: "86 g",
					Cholesterol:   "350 mg",
					Fat:           "27 g",
					Fiber:         "5 g",
					Protein:       "50 g",
					SaturatedFat:  "12 g",
					Sodium:        "880 mg",
					Sugar:         "9 g",
				},
				Yield: models.Yield{Value: 2},
				URL:   "https://www.hellofresh.com/recipes/creamy-shrimp-tagliatelle-5a8f0fcbae08b52f161b5832",
			},
		},
		/*{
			name: "homechef.com",
			in:   "https://www.homechef.com/meals/farmhouse-fried-chicken",
			want: models.RecipeSchema{
				AtContext: atContext,
				AtType:    models.SchemaType{Value: "Recipe"},
				Image: models.Image{
					Value: "https://homechef.imgix.net/https%3A%2F%2Fasset.homechef.com%2Fuploads%2Fmeal%2Fplated%2F2504%2F2504" +
						"FarmhouseFriedChicken_Ecomm__1_of_1_.jpg?ixlib=rails-1.1.0&w=600&auto=format&s=136cb76781125f3880aa8edd214bfae7",
				},
				Name: "Farmhouse Fried Chicken",
				URL:  "https://www.homechef.com/meals/farmhouse-fried-chicken",
				Description: models.Description{
					Value: "This stick-to-your-ribs satisfying country classic is an indulgence you've earned. The crispy comfort " +
						"that only fried chicken can supply is accompanied by mashed potatoes and sweet corn. While that “other” " +
						"chicken has you eating out of a bucket, this homey treat transports you to an idyllic country farmhouse " +
						"on the prairie. Yee-Haw!",
				},
				Yield: models.Yield{Value: 2},
				NutritionSchema: models.NutritionSchema{
					Calories:      "970",
					Carbohydrates: "71g",
					Fat:           "56g",
					Protein:       "45g",
					Sodium:        "2070mg",
				},
				Instructions: models.Instructions{
					Values: []string{
						"Cut potato into 1/2\" pieces. Bring a small pot with potato pieces and enough water to cover to a boil. " +
							"Reduce to a simmer and cook until fork-tender, 12-15 minutes. Drain potatoes in a colander and " +
							"return to pot. Add half the butter , 1/4 the cream (reserve remaining of each for gravy), 1/2 tsp. " +
							"olive oil , and a pinch of salt . Mash until desired consistency is reached. Cover and set aside. " +
							"While potato cooks, prepare ingredients.",
						"Trim and thinly slice green onions on an angle. Heat canola oil in a medium pan over medium heat, 5 " +
							"minutes. While oil heats, pat chicken breasts dry, and season both sides with a pinch of pepper . " +
							"Combine mayonnaise and 2 tsp. water in a mixing bowl. Place chicken breading in another mixing bowl. " +
							"Dip one chicken breast in mayonnaise-water mixture, then coat completely in chicken breading, shaking " +
							"off any excess. Repeat with second chicken breast.",
						"Line a plate with a paper towel. Test oil temperature by adding a pinch of chicken breading to it. It " +
							"should sizzle gently. If it browns immediately, turn heat down and let oil cool. If it doesn't brown, " +
							"increase heat. Lay chicken breasts away from you in hot oil and flip every 3-5 minutes until golden " +
							"brown and chicken reaches a minimum internal temperature of 165 degrees, 10-14 minutes. Transfer chicken " +
							"to towel-lined plate. Rest at least 5 minutes. While chicken rests, cook corn.",
						"Place another small pot over medium heat. Add 1 tsp. olive oil and corn to hot pot. Stir occasionally " +
							"until warmed through, 4-5 minutes. Transfer corn to a plate and season with a pinch of salt and pepper . " +
							"Wipe pot clean and reserve.",
						"Return pot used to cook corn to medium heat. Add green onions (reserve a pinch for garnish) and remaining " +
							"cream and bring to a simmer. Once simmering, stir often until slightly thickened, 3-5 minutes. Remove " +
							"from burner and swirl in remaining butter . Season with a pinch of pepper . If desired, slice chicken " +
							"into 1/2\" pieces. Plate dish as pictured on front of card, pouring gravy over chicken and garnishing " +
							"potatoes with reserved green onions. Bon appétit!",
					},
				},
				Ingredients: models.Ingredients{
					Values: []string{
						"2 Russet Potatoes",
						"13 oz. Boneless Skinless Chicken Breasts",
						"6 fl. oz. Canola Oil",
						"4 oz. Light Cream",
						"½ cup Chicken Breading",
						"3 oz. Corn Kernels",
						"2 Green Onions",
						"0.84 oz. Mayonnaise",
						"⅗ oz. Butter",
					},
				},
			},
		},*/
		{
			name: "hostthetoast.com",
			in:   "https://hostthetoast.com/guinness-beef-stew-with-cheddar-herb-dumplings/",
			want: models.RecipeSchema{
				AtContext:     atContext,
				AtType:        models.SchemaType{Value: "Recipe"},
				CookTime:      "PT3H",
				DatePublished: "2014-03-18",
				Image: models.Image{
					Value: "https://hostthetoast.com/wp-content/uploads/2014/03/Guinness-Beef-Stew-16-225x225.jpg",
				},
				Ingredients: models.Ingredients{
					Values: []string{
						"¼ pound bacon",
						"2 pounds boneless beef chuck, chopped into bite-sized pieces",
						"Kosher salt and black pepper",
						"4 sticks celery, chopped",
						"3 large carrots, chopped",
						"1 large onion, chopped",
						"4 cloves garlic, minced",
						"2 large potatoes or parsnips, diced",
						"1 turnip, diced",
						"3 ounces tomato paste",
						"1 (12 ounce) bottle Guinness",
						"4 cups low sodium chicken broth",
						"2 tablespoons Worcestershire sauce",
						"1 bay leaf",
						"3 sprigs thyme",
						"1 tablespoon cornstarch, or as needed",
						"½ pound cremini mushrooms, sliced (optional)",
						"Chopped parsley",
						"1 ½ cups self-rising flour",
						"1/2 teaspoon garlic powder",
						"1/3 cup shortening",
						"3/4 cup shredded Irish sharp cheddar",
						"2/3 cup milk",
						"2 tablespoons mixed fresh herbs such as parsley, chives, and thyme, chopped",
					},
				},
				Instructions: models.Instructions{
					Values: []string{
						"Cook the bacon in a large, oven-safe, heavy-based pot or high-walled saute pan over medium heat.",
						"Remove the bacon, crumble, and set aside, but leave the bacon fat in the pot. Season the beef with salt and " +
							"pepper and fry in the bacon fat until browned on all sides. Remove the beef from the pan and set aside.",
						"In the same pot, fry the onion, celery, and carrots until soft and fragrant, adding a little oil if necessary.",
						"Add garlic and fry for another 30 seconds. Stir in the tomato paste.",
						"Pour in the Guinness and Worcestershire sauce. Allow to come to a simmer and stir with a wooden spoon, " +
							"scraping up the browned bits from the bottom of the pot.",
						"Add the beef back to the pot and pour in the chicken broth. Add the bay leaf and thyme.",
						"Reduce to a simmer and cover. Simmer for 1 1/2 hours. Add the potatoes or parsnips and the turnip. Simmer " +
							"for another ½ hour, or until the vegetables are tender.",
						"Remove the bay leaf and thyme branches. If the stew is still thin, mix a tablespoon of cornstarch with a " +
							"tablespoon of cold water to form a slurry. Mix the slurry into the stew and bring the mixture to a boil. " +
							"Reduce to a simmer again, stirring occasionally, and add in the mushrooms if desired. Cook for 10 minutes, " +
							"uncovered, until the stew thickens and the mushrooms are cooked through. Stir the bacon back in. Preheat " +
							"the oven to 350°F.",
						"Stir together the self-rising flour and garlic powder in a medium bowl. Cut in the shortening until " +
							"mixture resembles coarse crumbs. Stir in the cheddar cheese, then add the milk and stir until the dry " +
							"ingredients are moistened.",
						"Make small balls with the dough and place them on top of the stew, leaving them room to expand-- " +
							"they grow a lot as they cook. Place the stew in the oven uncovered and bake until the dumplings are " +
							"browned and cooked through, about 30 to 40 minutes.",
						"Garnish the stew with parsley and serve.",
					},
				},
				Name:     "Guinness Beef Stew with Cheddar Herb Dumplings",
				PrepTime: "PT30M",
				Yield:    models.Yield{Value: 6},
				URL:      "https://hostthetoast.com/guinness-beef-stew-with-cheddar-herb-dumplings/",
			},
		},
		{
			name: "indianhealthyrecipes.com",
			in:   "https://www.indianhealthyrecipes.com/mango-rice-mamidikaya-pulihora/",
			want: models.RecipeSchema{
				AtContext:     atContext,
				AtType:        models.SchemaType{Value: "Recipe"},
				Category:      models.Category{Value: "Main"},
				CookTime:      "PT25M",
				Cuisine:       models.Cuisine{Value: "South Indian"},
				DatePublished: "2016-04-01T07:34:51+00:00",
				Description: models.Description{
					Value: "This Mango rice is a traditional South Indian dish made with precooked rice, raw green unripe " +
						"mangoes tempering spices and curry leaves. It tastes slightly tangy, hot and flavorful.",
				},
				Keywords: models.Keywords{Values: "mango rice, mango rice recipe"},
				Image: models.Image{
					Value: "https://www.indianhealthyrecipes.com/wp-content/uploads/2022/04/mango-rice-recipe.jpg",
				},
				Ingredients: models.Ingredients{
					Values: []string{
						"2 cups rice",
						"1 raw unripe green mango ((Sour, medium sized))",
						"1 to 2 sprigs curry leaves",
						"3 to 4 green chilies ( slit or chopped)",
						"1 to 2 dried red chilies (broken)",
						"⅖ to 1 teaspoon salt ((adjust to taste))",
						"¼ teaspoon turmeric ((prefer organic))",
						"¼ cup peanuts (or cashewnuts)",
						"1 tablespoon chana dal ((bengal gram))",
						"1 tablespoon urad dal ((skinned split black gram))",
						"1 teaspoon mustard seeds",
						"1 inch ginger (chopped, sliced, grated )",
						"3 tablespoons oil",
						"1 pinch hing ((⅛ teaspoon asafetida))",
					},
				},
				Instructions: models.Instructions{
					Values: []string{
						"Cook rice to grainy texture: Add rice to a bowl &amp; rinse it a few times. Pour 4 cups water &amp; " +
							"place the bowl in a pressure cooker. Cover the bowl &amp; pressure cook for 3 whistles.",
						"When the pressure drops, remove the rice &amp; cool completely.",
						"Wash, peel and grate or chop the mango. You can also grate in a processor. Update: Lately I have started " +
							"to make this with cooked mango. I just peel, cut and add them to a bowl. Cover and place it in " +
							"the pressure cooker over the rice bowl (PIP). Once cooked I mash it and use as mentioned below.",
						"Heat 1 tablespoon oil in a pan and fry the peanuts on a medium heat until aromatic and golden. Remove them " +
							"to a plate for later.",
						"Pour 2 tablespoons more oil and heat it. Add chana dal, urad dal, mustard seeds and dried red chilli.",
						"When the lentils turn light golden, add ginger, green chilies &amp; curry leaves. Fry till the curry leaves " +
							"become crisp, then add hing.",
						"Add mango, salt and turmeric.Saute for 2 to 3 minutes. Cook covered until the mango turns mushy, completely " +
							"soft &amp; pulpy.  (skip this with cooked mango.)",
						"Add this to the cooked rice little by little and begin to mix. Taste test and add more mango mixture as " +
							"required. Adjust salt and oil at this stage.",
						"Transfer mango rice to serving plates and garnish with roasted peanuts.",
					},
				},
				Name: "Mango Rice Recipe",
				NutritionSchema: models.NutritionSchema{
					Calories:      "636 kcal",
					Carbohydrates: "83 g",
					Fat:           "28 g",
					Fiber:         "7 g",
					Protein:       "11 g",
					SaturatedFat:  "13 g",
					Servings:      "1",
					Sodium:        "28 mg",
					Sugar:         "1 g",
				},
				PrepTime: "PT10M",
				Yield:    models.Yield{Value: 4},
				URL:      "https://www.indianhealthyrecipes.com/mango-rice-mamidikaya-pulihora/",
			},
		},
		{
			name: "innit.com",
			in:   "https://www.innit.com/meal/504/8008/Salad%3A%20Coconut-Pineapple-Salad",
			want: models.RecipeSchema{
				AtContext:     atContext,
				AtType:        models.SchemaType{Value: "Recipe"},
				Category:      models.Category{Value: "Salads and Sides"},
				CookTime:      "PT7M",
				DatePublished: "2022-02-12",
				Image: models.Image{
					Value: "https://www.innit.com/meal-service/en-US/images/Meal-Salad%3A%20Coconut_Pineapple_Salad_" +
						"1529953193419_480x480.png",
				},
				Ingredients: models.Ingredients{
					Values: []string{
						"2 Fresh Mexican Limes",
						"2 cups Pineapple",
						"1/2 cup Mint",
						"2 cups Jasmine Rice",
						"2 cups Canned Coconut Milk",
						"2 tsp Kosher Salt",
						"1/4 tsp Korean Chili Flakes",
						"2 Tbsp Extra Virgin Olive Oil",
						"1 cup Macadamia Nuts",
					},
				},
				Instructions: models.Instructions{
					Values: []string{
						"Combine ingredients in pot & bring to a boil.",
						"Cover & simmer for 15 minutes.",
						"Remove from heat & steam with lid for 5 minutes.",
						"Dice pineapple.",
						"Combine ingredients in a large bowl. Mix well.",
						"Plate on platter or bowl & garnish with macadamia nuts, mint & chili flakes.",
					},
				},
				Name: "Coconut Pineapple Rice",
				NutritionSchema: models.NutritionSchema{
					Calories:       "880 kcal",
					Carbohydrates:  "88 g",
					Cholesterol:    "0 mg",
					Fat:            "56 g",
					Fiber:          "7 g",
					Protein:        "12 g",
					SaturatedFat:   "28 g",
					Sodium:         "1190 mg",
					Sugar:          "10 g",
					UnsaturatedFat: "28 g",
				},
				PrepTime: "PT28M",
				Yield:    models.Yield{Value: 4},
				URL:      "https://www.innit.com/meal/504/8008/Salad%3A%20Coconut-Pineapple-Salad",
			},
		},
		{
			name: "inspiralized.com",
			in:   "https://inspiralized.com/vegetarian-zucchini-noodle-pad-thai/",
			want: models.RecipeSchema{
				AtContext:     atContext,
				AtType:        models.SchemaType{Value: "Recipe"},
				CookTime:      "PT15M",
				DatePublished: "2014-05-05T12:00:03+00:00",
				Description: models.Description{
					Value: "Make quick and healthy zucchini noodle pad thai with eggs, hoisin sauce, peanuts and spiralized " +
						"zucchini for dinner tonight.",
				},
				Image: models.Image{
					Value: "https://inspiralized.com/wp-content/uploads/2014/05/IMG_9863-copy-1.jpg",
				},
				Ingredients: models.Ingredients{
					Values: []string{
						"2 whole eggs",
						"1/4 cup roasted salted peanuts",
						"1/2 tbsp peanut oil (or oil of choice)",
						"1 garlic clove (minced)",
						"1 shallot (minced)",
						"1 tbsp coconut flour",
						"1 tbsp roughly chopped cilantro + whole cilantro leaves to garnish",
						"2 medium zucchinis (Blade C)",
						"For the sauce:",
						"2 tbsp freshly squeezed lime juice",
						"1 tbsp fish sauce (or hoisin sauce, if you're strict vegetarian)",
						"1/2 tbsp soy sauce",
						"1 tbsp chili sauce (I used Thai chili garlic sauce)",
						"1 tsp honey",
					},
				},
				Instructions: models.Instructions{
					Values: []string{
						"Scramble the eggs and set aside.",
						"Place all of the ingredients for the sauce into a bowl, whisk together and set aside.",
						"Place the peanuts into a food processor and pulse until lightly ground (no big peanuts should " +
							"remain, but it shouldn't be powdery). Set aside.",
						"Place a large skillet over medium heat. Add in oil, garlic and shallots. Cook for about 1-2 minutes, " +
							"stirring frequently, until the shallots begin to soften. Add in the sauce and whisk quickly " +
							"so that the flour dissolves and the sauce thickens. Cook for 2-3 minutes or until sauce is " +
							"reduced and thick.",
						"Once the sauce is thick, add in the zucchini noodles and cilantro and stir to combine thoroughly.",
						"Cook for about 2 minutes or until noodles soften and then add in the scrambled eggs and ground peanuts. " +
							"Cook for about 30 seconds, tossing to fully combine.",
						"Plate onto dishes and garnish with cilantro leaves. Serve with lime wedges.",
					},
				},
				Name:     "Vegetarian Zucchini Noodle Pad Thai",
				PrepTime: "PT10M",
				Yield:    models.Yield{Value: 2},
				URL:      "https://inspiralized.com/vegetarian-zucchini-noodle-pad-thai/",
			},
		},
		{
			name: "jamieoliver.html.com",
			in:   "https://www.jamieoliver.com/recipes/chicken-recipes/thai-green-chicken-curry/",
			want: models.RecipeSchema{
				AtContext:     atContext,
				AtType:        models.SchemaType{Value: "Recipe"},
				Category:      models.Category{Value: "Mains"},
				Cuisine:       models.Cuisine{Value: "https://schema.org/LowLactoseDiet"},
				DatePublished: "2015-09-16",
				Description: models.Description{
					Value: "This deliciously fragrant Thai green curry really packs a flavour punch.",
				},
				Keywords: models.Keywords{
					Values: "chicken, mushroom, dairy-free, poultry, vegetable, thai green, curry, chicken thighs, paste, chicken " +
						"curry, thai, thai green curry, vegetables, One-pan recipes, Curry, Chicken, Stewing, Dinner Party",
				},
				Image: models.Image{
					Value: "https://img.jamieoliver.com/jamieoliver/recipe-database/oldImages/large/1575_2_1437576282.jpg?tr=w-800,h-800",
				},
				Ingredients: models.Ingredients{
					Values: []string{
						"750 g skinless free-range chicken thighs",
						"groundnut oil",
						"400 g mixed oriental mushrooms",
						"1 x 400g tin of light coconut milk",
						"1 organic chicken stock cube",
						"6 lime leaves",
						"200 g mangetout",
						"½ a bunch fresh Thai basil",
						"2 limes",
						"4 cloves of garlic",
						"2 shallots",
						"5cm piece of ginger",
						"2 lemongrass stalks",
						"4 green Bird's eye chillies",
						"1 teaspoon ground cumin",
						"½ a bunch of fresh coriander",
						"2 tablespoons fish sauce",
					},
				},
				Instructions: models.Instructions{
					Values: []string{
						"<ol class=\"recipeSteps\"><li>To make the curry paste, peel, roughly chop and place the garlic, " +
							"shallots and ginger into a food processor. </li><li>Trim the lemongrass, remove the " +
							"tough outer leaves, then finely chop and add to the processor. Trim and add the chillies " +
							"along with the cumin and half the coriander (stalks and all). Blitz until finely chopped, " +
							"add the fish sauce and blitz again. </li><li>Slice the chicken into 2.5cm strips. Heat 1 " +
							"tablespoon of oil in a large pan on a medium heat, add the chicken and fry for 5 to 7 " +
							"minutes, or until just turning golden, then transfer to a plate. </li><li>Tear the " +
							"mushrooms into even pieces. Return the pan to a medium heat, add the mushrooms and " +
							"fry for 4 to 5 minutes, or until golden. Transfer to a plate using a slotted spoon. </li>" +
							"<li>Reduce the heat to medium-low and add the Thai green paste for 4 to 5 minutes, stirring " +
							"occasionally. </li><li>Pour in the coconut milk and 200ml of boiling water, crumble in the " +
							"stock cube and add the lime leaves. Turn the heat up and bring gently to the boil, then simmer " +
							"for 10 minutes, or until reduced slightly.</li><li>Stir in the chicken and mushrooms, reduce " +
							"the heat to low and cook for a further 5 minutes, or until the chicken is cooked through, adding " +
							"the mangetout for the final 2 minutes. </li><li>Season carefully to taste with sea salt and " +
							"freshly ground black pepper. Pick, roughly chop and stir through the basil leaves and remaining " +
							"coriander leaves. Serve with lime wedges and steamed rice.</li></ol>",
					},
				},
				Name: "Thai green chicken curry",
				NutritionSchema: models.NutritionSchema{
					Calories:      "285 calories",
					Carbohydrates: "6.1 g carbohydrate",
					Fat:           "16.2 g fat",
					Fiber:         "2.2 g fibre",
					Protein:       "28.9 g protein",
					SaturatedFat:  "6.5 g saturated fat",
					Sodium:        "1.0 g salt",
					Sugar:         "4.2 g sugar",
				},
				Yield: models.Yield{Value: 6},
				URL:   "https://www.jamieoliver.com/recipes/chicken-recipes/thai-green-chicken-curry/",
			},
		},
		{
			name: "jimcooksfoodgood.com",
			in:   "https://jimcooksfoodgood.com/recipe-weeknight-pad-thai/",
			want: models.RecipeSchema{
				AtContext:     atContext,
				AtType:        models.SchemaType{Value: "Recipe"},
				Category:      models.Category{Value: "Main Dish"},
				CookTime:      "PT15M",
				Cuisine:       models.Cuisine{Value: "American"},
				DatePublished: "2023-05-09T12:58:13+00:00",
				Description:   models.Description{Value: "Quick easy and delicious"},
				Keywords:      models.Keywords{Values: "#healthyrecipe"},
				Image: models.Image{
					Value: "https://jimcooksfoodgood.com/wp-content/uploads/2021/05/8DBE2045-ED9A-4B03-90F8-9B2114FC742C-scaled.jpeg",
				},
				Ingredients: models.Ingredients{
					Values: []string{
						"8 ounces rice noodles",
						"3 tablespoons tamarind ((or 2 tablespoons more of both lime juice and brown sugar))",
						"1/2 cups soy sauce",
						"4 tablespoons brown sugar",
						"2 tablespoons Sriracha",
						"2 limes ((one for juice, one for wedges))",
						"2 green onions",
						"2 shallots",
						"3 eggs",
						"4 garlic cloves",
						"1 cup bean sprouts",
						"2 cups Chopped Broccoli",
						"1/2 c roasted peanuts ((coarsely chopped))",
						"3 tablespoons cooking oil",
						"Optional: 1 pound of cooked protein ((shrimp, tofu, etc))",
						"Optional: Sriracha Mayo",
					},
				},
				Instructions: models.Instructions{
					Values: []string{
						"Bring a pot of water to boil and cook rice noodles according to package directions, shy just one minute. " +
							"Drain and set aside.",
						"In a separate bowl, combine tamarind, soy sauce, brown sugar, sriracha, and the juice of one lime together.",
						"In a very large pan over medium heat, add one tablespoon of the oil. Add the eggs and scramble until just " +
							"set, and set aside.",
						"Slice the shallots and green onions thinly, and mince the garlic. In the same pan, add the remainder of " +
							"the oil, still over medium heat. Add green onions, shallots, garlic and broccoli. Sauté until " +
							"broccoli is cooked through, 4-5 minutes.",
						"Add the noodles to the pan and pour on the sauce. Toss to coat all noodles. Add the eggs, bean sprout, and " +
							"your cooked protein. Sprinkle peanuts on top, and serve along with a wedge of lime and Sriracha Mayo.",
					},
				},
				Name: "Pad Thai",
				NutritionSchema: models.NutritionSchema{
					Calories: "389 kcal",
					Servings: "4",
				},
				PrepTime: "PT15M",
				Yield:    models.Yield{Value: 4},
				URL:      "https://jimcooksfoodgood.com/recipe-weeknight-pad-thai/",
			},
		},
		{
			name: "joyfoodsunshine.com",
			in:   "https://joyfoodsunshine.com/peanut-butter-frosting/",
			want: models.RecipeSchema{
				AtContext:     atContext,
				AtType:        models.SchemaType{Value: "Recipe"},
				Category:      models.Category{Value: "condiment"},
				Cuisine:       models.Cuisine{Value: "American"},
				DatePublished: "2022-02-16T02:36:00+00:00",
				Description: models.Description{
					Value: "This peanut butter frosting recipe is easy to make in 5 minutes. It&#039;s silky smooth, made with " +
						"more peanut butter than butter and is flavored with vanilla &amp; sea salt. It pipes well and tastes " +
						"delicious on top of chocolate cupcakes and brownies and chocolate cake.",
				},
				Keywords: models.Keywords{
					Values: "how to make peanut butter frosting, peanut butter buttercream, peanut butter frosting, peanut butter " +
						"frosting recipe",
				},
				Image: models.Image{
					Value: "https://joyfoodsunshine.com/wp-content/uploads/2022/02/peanut-butter-frosting-recipe-3.jpg",
				},
				Ingredients: models.Ingredients{
					Values: []string{
						"¾ cup creamy peanut butter",
						"½ cup salted butter (softened)",
						"½ teaspoon pure vanilla extract",
						"¼ teaspoon fine sea salt",
						"2 cups powdered sugar",
						"1-2 tablespoons whole milk (room temperature)",
					},
				},
				Instructions: models.Instructions{
					Values: []string{
						"In the bowl of a standing mixer fitted with the paddle attachment, or in a large bowl with a handheld " +
							"mixer, beat together peanut butter and butter until smooth.",
						"Add vanilla and sea salt and beat until combined.",
						"Add powdered sugar, 1 cup at a time, and beat until fully incorporated after each addition.",
						"Add 1 tablespoon whole milk and beat. If necessary, add an additional 1 tablespoon milk to achieve your " +
							"desired consistency.",
						"Use to frost a chocolate cake, chocolate cupcakes, brownies, etc.",
					},
				},
				Name: "Peanut Butter Frosting Recipe",
				NutritionSchema: models.NutritionSchema{
					Calories:       "181 kcal",
					Carbohydrates:  "17 g",
					Cholesterol:    "15 mg",
					Fat:            "12 g",
					Fiber:          "1 g",
					Protein:        "3 g",
					SaturatedFat:   "5 g",
					Servings:       "2 TBS",
					Sodium:         "143 mg",
					Sugar:          "16 g",
					TransFat:       "1 g",
					UnsaturatedFat: "6 g",
				},
				PrepTime: "PT5M",
				Yield:    models.Yield{Value: 16},
				URL:      "https://joyfoodsunshine.com/peanut-butter-frosting/",
			},
		},
		{
			name: "justataste.com",
			in:   "https://www.justataste.com/mini-sour-cream-doughnut-muffins-recipe/",
			want: models.RecipeSchema{
				AtContext:     atContext,
				AtType:        models.SchemaType{Value: "Recipe"},
				Category:      models.Category{Value: "Breakfast"},
				CookTime:      "PT16M",
				Cuisine:       models.Cuisine{Value: "American"},
				DatePublished: "2022-03-10T09:59:00+00:00",
				Description: models.Description{
					Value: "Two breakfast favorites join forces in a family-friendly recipe for Mini Sour Cream Doughnut " +
						"Muffins rolled in cinnamon-sugar.",
				},
				Keywords: models.Keywords{
					Values: "cinnamon, doughnut, sour cream, sugar, vanilla extract",
				},
				Image: models.Image{
					Value: "https://www.justataste.com/wp-content/uploads/2021/12/mini-sour-cream-doughnut-muffins.jpg",
				},
				Ingredients: models.Ingredients{
					Values: []string{
						"Cooking spray",
						"1 cup all-purpose flour",
						"1 teaspoon baking powder",
						"1/4 teaspoon baking soda",
						"1/4 teaspoon salt",
						"3 Tablespoons unsalted butter, at room temp",
						"3 Tablespoons vegetable oil",
						"1 cup sugar, divided",
						"1 large egg",
						"1/2 cup sour cream",
						"1 teaspoon vanilla extract",
						"1 1/2 teaspoons ground cinnamon",
					},
				},
				Instructions: models.Instructions{
					Values: []string{
						"Preheat the oven to 350°F. Grease a nonstick mini muffin pan with cooking spray.",
						"In a medium bowl, whisk together the flour, baking powder, baking soda and salt. Set the mixture aside.",
						"In the bowl of a stand mixer fitted with the paddle attachment, beat together the butter, vegetable oil " +
							"and 1/2 cup sugar until well combined, about 2 minutes. Add the egg and beat until combined.",
						"Add the flour mixture, sour cream and vanilla extract and beat just until combined.",
						"Using a small ice cream scoop (or two spoons), scoop out heaping 1-tablespoon portions of the batter into " +
							"the prepared muffin pan.",
						"Bake the muffins for 16 to 22 minutes until pale golden. While the muffins bake, in a medium bowl, whisk " +
							"together the remaining 1/2 cup sugar and cinnamon.",
						"Remove the muffins from the oven and let them cool for 2 minutes in the pan before transferring them in " +
							"batches into the cinnamon-sugar mixture, tossing to coat. Repeat the coating process with the " +
							"remaining muffins then serve.",
						"It’s important to toss the muffins in the cinnamon-sugar mixture while they are hot to ensure the " +
							"cinnamon-sugar will stick.",
						"★Did you make this recipe? Don&#39;t forget to give it a star rating below!",
					},
				},
				Name: "Mini Sour Cream Doughnut Muffins",
				NutritionSchema: models.NutritionSchema{
					Calories:       "122 kcal",
					Carbohydrates:  "17 g",
					Cholesterol:    "17 mg",
					Fat:            "6 g",
					Fiber:          "1 g",
					Protein:        "1 g",
					SaturatedFat:   "2 g",
					Servings:       "1",
					Sodium:         "57 mg",
					Sugar:          "11 g",
					TransFat:       "1 g",
					UnsaturatedFat: "3 g",
				},
				PrepTime: "PT10M",
				Yield:    models.Yield{Value: 18},
				URL:      "https://www.justataste.com/mini-sour-cream-doughnut-muffins-recipe/",
			},
		},
		{
			name: "justonecookbook.com",
			in:   "https://www.justonecookbook.com/teriyaki-tofu-bowl/",
			want: models.RecipeSchema{
				AtContext:     atContext,
				AtType:        models.SchemaType{Value: "Recipe"},
				Category:      models.Category{Value: "Main Course"},
				CookTime:      "PT25M",
				Cuisine:       models.Cuisine{Value: "Japanese"},
				DatePublished: "2022-03-21T05:00:00+00:00",
				Description: models.Description{
					Value: "Smothered with sweet-savory homemade teriyaki sauce, this crispy Pan-Fried Teriyaki Tofu Bowl is amazingly easy and delicious!  It‘s also a great way to incorporate tofu into your weekly menu rotation.",
				},
				Keywords: models.Keywords{Values: "teriyaki sauce, tofu"},
				Image: models.Image{
					Value: "https://www.justonecookbook.com/wp-content/uploads/2022/03/Teriyaki-Tofu-Bowl-6768-I.jpg",
				},
				Ingredients: models.Ingredients{
					Values: []string{
						"14 oz medium-firm tofu (momen dofu) ((1 block))",
						"⅓ cup potato starch or cornstarch",
						"3 Tbsp neutral oil ((divided))",
						"¼ cup sake",
						"¼ cup mirin",
						"¼ cup soy sauce",
						"4 tsp sugar",
						"2 servings cooked Japanese short-grain rice ((typically 1⅔ cups (250 g) per donburi serving))",
						"1 green onion/scallion",
						"½ tsp toasted white sesame seeds",
					},
				},
				Instructions: models.Instructions{
					Values: []string{
						"Before You Start: For the steamed rice, please note that 1½ cups (300 g, 2 rice cooker cups) of uncooked Japanese short-grain rice yield 4⅓ cups (660 g) of cooked rice, enough for 2 donburi servings (3⅓ cups, 500 g). See how to cook short-grain rice with a rice cooker, pot over the stove, Instant Pot, or donabe.",
						"Open the package of 14 oz medium-firm tofu (momen dofu) and drain out the water.Next, wrap the tofu block in a paper towel (or tea towel) and place it on a plate or tray. Now, press the tofu: First, put another tray or plate or even a cutting board on top of the tofu block to evenly distribute the weight. Then, place a heavy item* (I used a marble mortar but a can of food works) on top to apply pressure.Let it sit for at least 30 minutes before using. *The weighted item should not be so heavy that it will crumble or crush the tofu block but heavy enough that it will press out the tofu&#39;s liquid.",
						"While draining the tofu, you can cook the rice or a side dish. For this recipe, I also prepare this blanched broccoli recipe.",
						"Gather all the ingredients.",
						"To make the homemade teriyaki sauce, whisk the ¼ cup sake, ¼ cup mirin, ¼ cup soy sauce, and 4 tsp sugar in a (microwave-safe) medium bowl. If the sugar doesn‘t dissolve easily, microwave it for 30 seconds and whisk well. Set aside.",
						"Cut 1 green onion/scallion diagonally into thin slices.",
						"After 30 minutes of draining the tofu, remove the paper towel and transfer the tofu to the cutting board. First, cut the tofu block in half widthwise.",
						"Next, cut the tofu into roughly ¾-inch (2-cm) cubes.",
						"Put ⅓ cup potato starch or cornstarch in a shallow tray or bowl and gently coat the tofu cubes with the potato starch. Set aside.",
						"Heat a large frying pan on medium-high heat. When it‘s hot, add 1½ Tbsp of the 3 Tbsp neutral oil (keep the rest for the next batch) and distribute it evenly. Add the first batch of tofu cubes to the pan, placing them about 1 inch (2.5 cm) apart from each other so it‘s easy to rotate the tofu cubes without sticking to each other.",
						"Fry the cubes on one side until golden brown, then turn them to fry the next side. Repeat until all sides are brown and crispy. Transfer the fried tofu cubes to a plate or tray lined with a paper towel.",
						"Add the next batch of uncooked tofu to the pan and fry until crispy and golden brown on all sides. Add more of the remaining oil as needed to help brown the tofu faster.",
						"Remove all the fried tofu to the plate/tray.",
						"Wipe off any remaining oil in the pan with a paper towel. Then, transfer the tofu back into the pan.",
						"Add the teriyaki sauce to the pan; the sauce will start to thicken immediately. Quickly toss the tofu cubes in the sauce to coat, then turn off the heat and remove the pan from the stove. Tip: The sauce will continue to thicken with the residual heat, so if you want to keep some sauce in the pan, be sure to turn off the heat as soon as the tofu is coated.",
						"Divide 2 servings cooked Japanese short-grain rice into individual large (donburi) bowls. Serve the tofu and blanched broccoli over the steamed rice. Garnish the tofu with green onions and ½ tsp toasted white sesame seeds.",
						"You can keep the leftovers in an airtight container and store in the refrigerator for 3 days. Since the texture of the tofu changes when frozen, I don‘t recommend storing the tofu in the freezer.",
					},
				},
				Name: "Pan-Fried Teriyaki Tofu Bowl",
				NutritionSchema: models.NutritionSchema{
					Calories:       "443 kcal",
					Carbohydrates:  "27 g",
					Fat:            "23 g",
					Fiber:          "3 g",
					Protein:        "21 g",
					SaturatedFat:   "3 g",
					Servings:       "1",
					Sodium:         "979 mg",
					Sugar:          "10 g",
					TransFat:       "1 g",
					UnsaturatedFat: "20 g",
				},
				PrepTime: "PT5M",
				Yield:    models.Yield{Value: 2},
				URL:      "https://www.justonecookbook.com/teriyaki-tofu-bowl/",
			},
		},
		{
			name: "kennymcgovern.com",
			in:   "https://kennymcgovern.com/chicken-noodle-soup",
			want: models.RecipeSchema{
				AtContext:     atContext,
				AtType:        models.SchemaType{Value: "Recipe"},
				Category:      models.Category{Value: "Soup"},
				CookTime:      "PT5M",
				Cuisine:       models.Cuisine{Value: "Chinese"},
				DatePublished: "2022-03-27T18:12:02+00:00",
				Keywords: models.Keywords{
					Values: "noodles, Soup, noodle soup, chicken, chicken noodle soup, chicken soup",
				},
				Image: models.Image{
					Value: "https://i0.wp.com/kennymcgovern.com/wp-content/uploads/2022/03/chicken-noodle-soup.jpg?fit=685%2C643&ssl=1",
				},
				Ingredients: models.Ingredients{
					Values: []string{
						"40g thin lucky boat noodles (soaked and drained, drained weight)",
						"275 ml light chicken stock",
						"Dash light soy sauce",
						"1/4 teaspoon sea salt",
						"1/4 teaspoon MSG",
						"Pinch white pepper",
						"50 grams raw chicken breast, thinly sliced (or 1 small handful cooked shredded chicken breast or thigh meat)",
						"1 spring onion (finely sliced)",
						"Dash sesame oil (optional, see notes)",
					},
				},
				Instructions: models.Instructions{
					Values: []string{
						"Prepare the noodles and place 40g of prepared noodles in a soup bowl. Set aside.",
						"Put the light chicken stock, light soy sauce, sea salt, MSG and white pepper in a saucepan, bring to the " +
							"boil then reduce to a simmer.",
						"Add the sliced chicken to the soup and simmer for about 3 minutes until the chicken is cooked through. " +
							"Pour the chicken soup over the prepared noodles in the bowl. Garnish with the sliced spring onion, " +
							"drizzle with sesame oil (if using) and serve.",
					},
				},
				Name:     "Chicken Noodle Soup",
				PrepTime: "PT5M",
				Yield:    models.Yield{Value: 1},
				URL:      "https://kennymcgovern.com/chicken-noodle-soup",
			},
		},
		{
			name: "kingarthurbaking.com",
			in:   "https://www.kingarthurbaking.com/recipes/sourdough-zucchini-bread-recipe",
			want: models.RecipeSchema{
				AtContext:     atContext,
				AtType:        models.SchemaType{Value: "Recipe"},
				Category:      models.Category{Value: ""},
				CookTime:      "PT1H5M",
				DatePublished: "June 3, 2021 at 2:13pm",
				Description: models.Description{
					Value: "This delicious whole grain zucchini bread makes wonderful use of excess sourdough starter you might otherwise discard. Paired with summer’s avalanche of zucchini, it’s one loaf that solves two kitchen conundrums!",
				},
				Keywords: models.Keywords{
					Values: "Quick bread, Lemon, Raisin, Sourdough, Spice, Breakfast & brunch",
				},
				Image: models.Image{
					Value: "https://www.kingarthurbaking.com/sites/default/files/2021-06/sourdough-zucchini-bread_0521.jpg",
				},
				Ingredients: models.Ingredients{
					Values: []string{
						"3/4 cup (170g) sourdough starter fed (ripe) or unfed (discard)",
						"1/2 cup (99g) granulated sugar",
						"1/4 cup (85g) honey",
						"6 tablespoons (75g) vegetable oil",
						"2 large eggs",
						"1/4 teaspoon nutmeg",
						"1 1/2 teaspoons lemon zest (grated rind)",
						"1 1/2 teaspoons King Arthur Pure Vanilla Extract",
						"1 cup (113g) King Arthur White Whole Wheat Flour",
						"3/4 cup (90g) King Arthur Unbleached All-Purpose Flour",
						"1/2 teaspoon baking soda",
						"1 teaspoon baking powder",
						"1 teaspoon table salt",
						"2 cups (242g to 300g) grated zucchini somewhere between firmly and lightly packed",
						"3/4 cup (85g) chopped walnuts lightly toasted",
						"3/4 cup (128g) raisins currants or dried cranberries",
					},
				},
				Instructions: models.Instructions{
					Values: []string{
						"Preheat the oven to 350°F. Lightly grease a 9” x 5” quick bread pan or 12” x 4” tea loaf pan.",
						", In a large bowl, stir together the starter, sugar, honey, oil, eggs, nutmeg, lemon zest, and vanilla until thoroughly combined.",
						", In a separate medium bowl, whisk together the flours, baking soda, baking powder, and salt; stir into the wet ingredients.",
						", Stir in the grated zucchini, then the nuts and fruit. Transfer the batter to the prepared pan, smoothing the top.",
						", Bake the bread in the 9” x 5” pan for 45 minutes. Tent with foil and bake for an additional 20 minutes, until a thin paring knife inserted in the center comes out clean. For bread in a tea loaf pan, bake for 40 minutes before tenting, then bake for another 20 minutes, or until the loaf tests done.",
						", Remove the bread from the oven and cool in the pan on a rack.",
						", Store bread, well wrapped, at room temperature for up to three days; freeze for longer storage.",
					},
				},
				Name: "Sourdough Zucchini Bread",
				NutritionSchema: models.NutritionSchema{
					Calories:       "279 calories",
					Carbohydrates:  "33g",
					Cholesterol:    "23g",
					Fat:            "13g",
					Fiber:          "2g",
					Protein:        "5g",
					SaturatedFat:   "2g",
					Servings:       "",
					Sodium:         "202mg",
					Sugar:          "21g",
					TransFat:       "0g",
					UnsaturatedFat: "",
				},
				PrepTime: "PT30M",
				Tools:    models.Tools{Values: []string(nil)},
				Yield:    models.Yield{Value: 16},
				URL:      "https://www.kingarthurbaking.com/recipes/sourdough-zucchini-bread-recipe",
			},
		},
		{
			name: "kochbar.de",
			in:   "https://www.kochbar.de/rezept/465773/Spargelsalat-Fruchtig.html",
			want: models.RecipeSchema{
				AtContext:     atContext,
				AtType:        models.SchemaType{Value: "Recipe"},
				Category:      models.Category{Value: "Hauptspeise"},
				Cuisine:       models.Cuisine{Value: "Internationale Küche"},
				DatePublished: "2013-04-20T18:30:20+02:00",
				Description:   models.Description{Value: "lauwarmer Spargel-Salat"},
				Keywords: models.Keywords{
					Values: "Spargelsalat Fruchtig, Spargel grün frisch, Spargel weiss frisch, Mango frisch",
				},
				Image: models.Image{
					Value: "https://ais.kochbar.de/kbrezept/465773_670587/1200x1200/spargelsalat-fruchtig-rezept.jpg",
				},
				Ingredients: models.Ingredients{
					Values: []string{
						"1 kg Spargel grün frisch",
						"1 Kg Spargel weiss frisch",
						"1 Stück Mango frisch",
						"1 Stück Orange frisch",
						"4 El Olivenöl",
						"Gourmet-Pfeffer aus meinem KB",
						"Salz",
						"Zucker",
						"Räucherlachs",
					},
				},
				Instructions: models.Instructions{
					Values: []string{
						"für 2 Personen als Hauptspeise \r\nfür 4 Personen als Vorspeise",
						"1. Spargel schälen und in 4-5 Stücke schneiden",
						"3. Spargel inm Salzwasser und wenig Zucker bissfest Kochen. Wasser wegschütten",
						"4. leicht abkühlen lassen und in der Zwischenzeit die Mango schälen und in kleine Würfel schneiden",
						"Dressing",
						"Saft von 1 Orange in eine Schüssel geben und das Olivenöl hinzufügen gut verrühren und mit Salz und " +
							"Pfeffer abschmecken. Die Spargeln darin wenden und ein wenig ziehen lassen.",
						"Schön Anrichten und mit Lachs garnieren ANSTELLE Lachs passen auch wunderbar Crevetten dazu.",
					},
				},
				Name: "Spargelsalat Fruchtig",
				NutritionSchema: models.NutritionSchema{
					Calories:      "97 kcal",
					Carbohydrates: "1,87273 g",
					Fat:           "9,23273 g",
					Protein:       "1,78182 g",
					Servings:      "100 g",
				},
				Yield: models.Yield{Value: 2},
				URL:   "https://www.kochbar.de/rezept/465773/Spargelsalat-Fruchtig.html",
			},
		},
		{
			name: "koket.se",
			in:   "https://www.koket.se/mitt-kok/tommy-myllymaki/myllymakis-toast-skagen",
			want: models.RecipeSchema{
				AtContext: atContext,
				AtType:    models.SchemaType{Value: "Recipe"},
				Description: models.Description{
					Value: "Toast skagen är en klassisk förrätt på årets festdag - nyårsafton. Tommys variant görs med hemslagen " +
						"majonnäs, pepparrot och löjrom.",
				},
				Image: models.Image{
					Value: "https://img.koket.se/standard-mega/myllymakis-toast-skagen-2.jpg",
				},
				Ingredients: models.Ingredients{
					Values: []string{
						"1 kg räkor med skal (gärna färska av fin kvalitet)",
						"2 äggulor",
						"2 tsk senap",
						"1 msk vitvinsvinäger",
						"6 dl matolja",
						"1 kruka dill",
						"10 cm färsk pepparrot, skalad",
						"4 skivor vitt bröd (ej levain)",
						"smör, till stekning",
						"50 g löjrom",
						"1 citron",
					},
				},
				Instructions: models.Instructions{
					Values: []string{
						"Skala alla räkor och ställ åt sidan.",
						"Gör en majonnäs genom att lägga ner äggulor, senapen och vinägern i en bunke. Tillsätt matoljan i en " +
							"tunn stråle medan du vispar hela tiden. Använd elvisp eller handvisp. När majonnäsen är tjock " +
							"och du ser dragen/spåren av vispen i majonnäsen är den klar.",
						"Lägg alla räkor i en bunke, tillsätt fint plockad dill och blanda ner lite majonnäs i taget.",
						"Tillsätt lite riven pepparrot och smaka av. Slå på mer majonnäs för en rinnigare röra eller mer pepparrot " +
							"för mer sting.",
						"Ta fram brödet och skär ut önskad form utan att ta med kanterna, använd en skål eller ett glas som mall " +
							"om ni vill ha runda bröd. Stek sedan gyllene i smör.",
						"Lägg upp bröden på tallrik, toppa med skagenröra och en rejäl klick löjrom. Avsluta med en dillkvist och " +
							"en citronskiva.",
					},
				},
				Name:  "Myllymäkis toast skagen",
				Yield: models.Yield{Value: 4},
				URL:   "https://www.koket.se/mitt-kok/tommy-myllymaki/myllymakis-toast-skagen",
			},
		},
		{
			name: "kuchnia-domowa.pl",
			in:   "https://www.kuchnia-domowa.pl/przepisy/dodatki-do-dan/548-mizeria",
			want: models.RecipeSchema{
				AtContext: atContext,
				AtType:    models.SchemaType{Value: "Recipe"},
				Category:  models.Category{Value: "Dodatki do dań"},
				Cuisine:   models.Cuisine{Value: "Polska"},
				Description: models.Description{
					Value: "Lekka surówka do obiadu ze świeżego ogórka, śmietany lub jogurtu oraz koperku. Bardzo prosta, idealnie nadająca się do wielu dań obiadowych. Mizeria najsmaczniejsza jest z ziemniakami najlepiej młodymi i jakimś mięsem np. kotletem mielonym lub schabowym.\nMy najbardziej lubimy kremową mizerię z miękkimi, cienkimi plasterkami ogórka doprawioną nie tylko solą i pieprzem, ale również (aby była słodko- winna) sokiem z cytryny i cukrem. A jak u Ciebie przygotowuje się mizerię?",
				},
				Keywords: models.Keywords{Values: "przepis, mizeria, surówka z ogórków, mizeria z octem i śmietaną, tradycyjna mizeria, klasyczna mizeria, domowa mizeria"},
				Image: models.Image{
					Value: "https://kuchnia-domowa.pl/images/content/548/mizeria.jpg",
				},
				Ingredients: models.Ingredients{
					Values: []string{
						"600 g świeżych ogórków gruntowych (lub długich, szklarniowych)*",
						"300 g gęstej, kwaśnej śmietany 18% lub jogurtu typu greckiego",
						"1 łyżeczka soli",
						"1 łyżka soku z cytryny (lub niepełna łyżka octu jabłkowego)",
						"1 łyżeczka cukru",
						"czarny pieprz mielony",
						"1 łyżka drobno posiekanego koperku",
					},
				},
				Instructions: models.Instructions{
					Values: []string{
						"Ogórki umyć, osuszyć, obrać i pokroić w jak najcieńsze plasterki.",
						"Plasterki umieścić w misce i posypać 1 łyżeczką soli. Wymieszać i pozostawić na ok. 15 minut.",
						"W międzyczasie śmietanę przełożyć do miseczki. Przyprawić sokiem z cytryny, cukrem, pieprzem i posiekanym " +
							"koperkiem. Wymieszać.",
						"Po 15 minutach odlać wodę, którą puściły ogórki. (Lekko je odcisnąć, ale nie za mocno, aby mizeria nie " +
							"wyszła za sucha).",
						"Dodać przygotowaną śmietanę i wymieszać.",
					},
				},
				Name:  "Mizeria",
				Yield: models.Yield{Value: 4},
				URL:   "https://www.kuchnia-domowa.pl/przepisy/dodatki-do-dan/548-mizeria",
			},
		},
		{
			name: "kwestiasmaku.com",
			in:   "https://www.kwestiasmaku.com/przepis/muffiny-czekoladowe-z-maslem-orzechowym",
			want: models.RecipeSchema{
				AtContext:     atContext,
				AtType:        models.SchemaType{Value: "Recipe"},
				DateModified:  "2022-11-05T09:43:22+01:00",
				DatePublished: "2022-03-24T19:55:17+01:00",
				Description: models.Description{
					Value: "Mocno kakaowe muffiny wzmocnione dodatkową dawką czekolady w postaci dropsów czekoladowych (lub " +
						"posiekanej czekolady). Dla miłośników masła orzechowego dodajemy do nich po łyżeczce masła " +
						"orzechowego i rozprowadzamy je w czekoladowej masie za pomocą wykałaczki.\nZ przepisu otrzymamy " +
						"od 14 do 16 muffinków. Nakładamy do foremek tyle ciasta aby nie wypływało na zewnątrz podczas " +
						"pieczenia i nie robił się \"grzybek\". W związku z tym, że możemy mieć różne wielkości foremek, " +
						"najlepiej wypełniać foremki surowym ciastem do 2/3 ich objętości. Pozostawiamy w ten sposób miejsce " +
						"na wyrośnięcie ciasta i otrzymamy kształtne babeczki.\n",
				},
				Image: models.Image{
					Value: "https://www.kwestiasmaku.com/sites/v123.kwestiasmaku.com/files/muffiny-czekoladowe-z-maslem-" +
						"orzechowym-00.jpg",
				},
				Ingredients: models.Ingredients{
					Values: []string{
						"150 g masła",
						"150 g dropsów czekoladowych (np. z ciemnej czekolady) lub 150 g czekolady deserowej lub gorzkiej",
						"300 g mąki",
						"2 łyżeczki proszku do pieczenia",
						"1/2 łyżeczki sody oczyszczonej",
						"3 łyżki kakao",
						"1 szklanka (200 g) cukru",
						"1 łyżka cukru wanilinowego",
						"2 duże jajka (L)",
						"200 ml mleka",
						"ok. 5 - 6 łyżek masła orzechowego",
						"15 - 18 papilotek",
						"metalowa forma na muffiny z wgłębieniami",
					},
				},
				Instructions: models.Instructions{
					Values: []string{
						"Piekarnik nagrzać do 180 stopni C.\u00a0Masło roztopić i przestudzić. Czekoladę pokroić na kawałeczki.",
						"Mąkę przesiać do miski razem z proszkiem do pieczenia, sodą i kakao, dokładnie wymieszać. Dodać cukier " +
							"oraz cukier wanilinowy i ponownie wymieszać.",
						"W drugiej misce rozmiksować jajka z mlekiem (rózgą lub mikserem).",
						"Do sypkich składników dodać masę jajeczną i krótko zamieszać łyżką. Dodać roztopione masło i wymieszać " +
							"do połączenia składników, pod koniec dodając 2/3 ilości dropsów czekoladowych.",
						"Masę wyłożyć do papilotek umieszczonych w formie na muffiny, na wierzch wyłożć po łyżeczce masła " +
							"orzechowego na każdą muffinkę.",
						"Wykałaczką zrobić \"ósemkę\" w cieście mieszając delikatnie masę czekoladową z masłem orzechowym. Wierzch " +
							"posypać pozostałą 1/3 dropsów czekoladowych.",
						"Wstawić do piekarnika (można piec na raty, w 2 partiach) i piec\u00a0przez około 20 -\u00a023 minuty, " +
							"do suchego patyczka.",
					},
				},
				Name:  "Muffiny czekoladowe z masłem orzechowym",
				Yield: models.Yield{Value: 15},
				URL:   "https://www.kwestiasmaku.com/przepis/muffiny-czekoladowe-z-maslem-orzechowym",
			},
		},
		{
			name: "lecremedelacrumb.com",
			in:   "https://www.lecremedelacrumb.com/instant-pot-pot-roast-potatoes/",
			want: models.RecipeSchema{
				AtContext:     atContext,
				AtType:        models.SchemaType{Value: "Recipe"},
				Category:      models.Category{Value: "Main Course"},
				CookTime:      "PT80M",
				Cuisine:       models.Cuisine{Value: "American"},
				DatePublished: "2018-01-19T11:11:53+00:00",
				Description: models.Description{
					Value: "Juicy and tender instant pot pot roast and potatoes with gravy makes the perfect family-friendly " +
						"dinner. This easy one pot dinner recipe will please even the picky eaters!",
				},
				Keywords: models.Keywords{
					Values: "instant pot pot roast, pot roast and potatoes",
				},
				Image: models.Image{
					Value: "https://www.lecremedelacrumb.com/wp-content/uploads/2018/01/instant-pot-beef-roast-103.jpg",
				},
				Ingredients: models.Ingredients{
					Values: []string{
						"3-5 pound beef chuck roast (see notes for instructions from frozen)",
						"1 tablespoon oil",
						"1 teaspoon salt",
						"1 teaspoon onion powder",
						"1 teaspoon garlic powder",
						"½ teaspoon black pepper",
						"½ teaspoon smoked paprika (optional)",
						"1 pound baby red potatoes",
						"4 large carrots (chopped into large chunks, see note for using baby carrots)",
						"1 large yellow onion (chopped)",
						"4 cups beef broth",
						"2 tablespoons worcestershire sauce",
						"¼ cup water",
						"2 tablespoons corn starch",
					},
				},
				Instructions: models.Instructions{
					Values: []string{
						"Turn on your instant pot and set it to \"saute\". In a small bowl stir together salt, pepper, garlic " +
							"powder, onion powder, and smoked paprika. Rub mixture all over the roast to coat all sides.",
						"Drizzle oil in instant pot, wait about 30 seconds, then use tongs to place roast in the pot. Do not " +
							"move it for 3-4 minutes until well-seared and browned. Use tongs to turn the roast onto another " +
							"side for 3-4 minutes, repeating until all sides are browned.",
						"Switch instant pot to \"pressure cook\" on high and set to 60-80 minutes (60 for a 3 pound roast, 80 " +
							"for a 5 pound roast. see notes if using baby carrots). Add potatoes, onions, and carrots to pot " +
							"(just arrange them around the roast) and pour beef broth and worcestershire sauce over everything. " +
							"Place lid on the pot and turn to locked position. Make sure the vent is set to the sealed position.",
						"When the cooking time is up, do a natural release for 10 minutes (don't touch anything on the pot, just " +
							"let it de-pressurize on it's own for 10 minutes). After 10 minutes, turn vent to the venting " +
							"release position and allow all of the steam to vent and the float valve to drop down before removing " +
							"the lid.",
						"Transfer the roast, potatoes, onions, and carrots to a platter and shred the roast with 2 forks into " +
							"chunks. Use a handheld strainer to scoop out bits from the broth in the pot. Set instant pot to \"soup\" " +
							"setting. Whisk together the water and corn starch. Once broth is boiling, stir in corn starch mixture " +
							"until the gravy thickens. Add salt, pepper, and garlic powder to taste.",
						"Serve gravy poured over roast and veggies and garnish with fresh thyme or parsley if desired.",
					},
				},
				Name: "Instant Pot Pot Roast Recipe",
				NutritionSchema: models.NutritionSchema{
					Calories:      "133 kcal",
					Carbohydrates: "23 g",
					Fat:           "3 g",
					Fiber:         "3 g",
					Protein:       "4 g",
					SaturatedFat:  "1 g",
					Servings:      "1",
					Sodium:        "1087 mg",
					Sugar:         "5 g",
				},
				PrepTime: "PT20M",
				Yield:    models.Yield{Value: 6},
				URL:      "https://www.lecremedelacrumb.com/instant-pot-pot-roast-potatoes/",
			},
		},
		/*{
			name: "lekkerensimpel.com",
			in:   "https://www.lekkerensimpel.com/gougeres/",
			want: models.RecipeSchema{
				AtContext: atContext,
				AtType:    models.SchemaType{Value: "Recipe"},
				Name:      "Gougères",
				Category:  models.Category{Value: "Snacks"},
				Yield:     models.Yield{Value: 4},
				URL:       "https://www.lekkerensimpel.com/gougeres/",
				PrepTime:  "PT20M",
				CookTime:  "PT25M",
				Description: models.Description{
					Value: "Vandaag een receptje uit de Franse keuken, namelijk deze gougères. Gougères zijn een soort " +
						"hartige kaassoesjes, erg lekker! We hadden een tijdje geleden een stuk gruyère kaas " +
						"gekocht bij de kaasboer, meer uit nood want parmezaanse had hij op dat moment even niet. " +
						"Inmiddels lag de kaas al een tijdje in de koelkast en moesten we er toch echt wat mee gaan " +
						"doen. Iemand tipte ons dat we echt eens gougères moesten maken en eerlijk gezegd hadden we " +
						"er nooit eerder van gehoord. Een kleine speurtocht bracht ons uiteindelijk bij een recept " +
						"van ‘The Guardian – How to make the perfect gougères‘. We zijn ermee aan de slag gegaan en " +
						"zie hier het resultaat! \n\nNog meer van dit soort lekkere snacks en borrelhapjes vind je " +
						"in onze categorie tapas recepten en tussen de high-tea recepten.",
				},
				Ingredients: models.Ingredients{
					Values: []string{
						"90 gr gruyere kaas",
						"125 ml water",
						"40 gr boter",
						"75 gr bloem",
						"3 eieren",
						"nootmuskaat",
						"zout",
					},
				},
				Instructions: models.Instructions{
					Values: []string{
						"Verwarm de oven voor op 200 graden. Klop vervolgens 2 eieren los in een beker. Doe het water, de boter " +
							"en een snuf zout in een pan en laat de boter al roerend smelten. Zet het ‘vuur’ laag en doe " +
							"de bloem erbij. Zelf doen we de bloem eerst door een zeef zodat er geen kleine klontjes meer " +
							"inzitten. Roer de bloem door het botermengsel totdat er een soort deeg ontstaat. Haal de pan " +
							"van het vuur en mix het deeg, bij voorkeur met een mixer, een minuut of 3-4. Voeg dan de " +
							"helft van het losgeklopte ei toe, even goed mengen en dan kan de andere helft erbij. Mix " +
							"daarna nog de nootmuskaat en geraspte gruyère door het deeg. Bekleed een bakplaat met " +
							"bakpapier. Schep met twee lepels kleine bolletjes deeg op de bakplaat of gebruik hiervoor " +
							"een spuitzak. Smeer de bovenkant in met een beetje losgeklopt ei, bestrooi met nog geraspte " +
							"gruyère kaas en dan kan de bakplaat de oven in voor 20-25 minuten. Eet smakelijk!",
						"Bewaar dit recept op Pinterest !",
					},
				},
				DatePublished: "2021-09-28T04:00:00+00:00",
				DateModified:  "2021-09-21T08:22:19+00:00",
			},
		},*/
		{
			name: "littlespicejar.com",
			in:   "https://littlespicejar.com/starbucks-pumpkin-loaf/",
			want: models.RecipeSchema{
				AtContext:     atContext,
				AtType:        models.SchemaType{Value: "Recipe"},
				Category:      models.Category{Value: "Bread & Baking"},
				CookTime:      "PT55M",
				Cuisine:       models.Cuisine{Value: "American"},
				DatePublished: "2021-11-09",
				Description: models.Description{
					Value: "Learn how to make an easy delicious copycat Starbucks Pumpkin Loaf right at home! This pumpkin " +
						"bread is studded with roasted pepitas and loaded with spices and so much pumpkin goodness!",
				},
				Keywords: models.Keywords{Values: ""},
				Image: models.Image{
					Value: "https://littlespicejar.com/wp-content/uploads/2021/11/Copycat-Starbucks-Pumpkin-Loaf-8-720x720.jpg",
				},
				Ingredients: models.Ingredients{
					Values: []string{
						"3 ½ cups all-purpose flour",
						"1 tablespoon cinnamon",
						"2 teaspoons EACH: baking soda AND ground ginger",
						"1 teaspoon EACH: baking powder AND ground allspice",
						"½ teaspoon EACH: ground cloves, ground cardamom, AND ground nutmeg",
						"¾ teaspoon kosher salt",
						"1 cup EACH: granulated sugar AND coconut oil (or other; see post)",
						"1 ½ cups light brown sugar",
						"1 (15-ounce) can pumpkin puree",
						"4 large eggs, room temperature*",
						"1 tablespoon vanilla extract",
						"Zest of 1 orange (optional)",
						"4-5 tablespoons roasted pepitas, for topping",
					},
				},
				Instructions: models.Instructions{
					Values: []string{
						"PREP: Position a rack in the center of the oven and preheat the oven to 350ºF. Spray two 8 ½ x 4 ½ " +
							"(or 9x5) bread pans with cooking spray, you can also line with parchment if you’d like; set aside for now.",
						"DRY INGREDIENTS: Add the dry ingredients: flour, baking soda baking powder, all the spices, and salt " +
							"to a medium bowl. Whisk to combine; set aside for now.",
						"WET INGREDIENTS: Add the granulated sugar, brown sugar, and oil to a large bowl. Whisk to combine, then " +
							"add the pumpkin puree, eggs, vanilla, and orange zest and combine to whisk until all the eggs have " +
							"been incorporated into the wet batter. Don't be alarmed if the batter splits or curdles! It's totally fine!",
						"BREAD BATTER: Add the dry ingredients into the wet ingredients in two batches, stirring just long enough " +
							"so each batch of flour is incorporated. Do not over-mix or you’ll end up with dry bread!",
						"BAKE: Divide the batter into the to pans, taking care to only fill each pan about ¾ of the way full. The " +
							"bread will rise significantly! Smooth out the batter then sprinkle with the pepitas. Bake the bread for " +
							"52-62 minutes or until a toothpick inserted in the center of the loaf comes out clean. Cool the pans on a " +
							"wire baking rack for at least 10 minutes before removing from the pan and allowing the bread to cool further.",
					},
				},
				Name:     "The Best Starbucks Pumpkin Loaf Recipe (Copycat)",
				PrepTime: "PT15M",
				Yield:    models.Yield{Value: 2},
				URL:      "https://littlespicejar.com/starbucks-pumpkin-loaf/",
			},
		},
		{
			name: "livelytable.com",
			in:   "https://livelytable.com/bbq-ribs-on-the-charcoal-grill/",
			want: models.RecipeSchema{
				AtContext:     atContext,
				AtType:        models.SchemaType{Value: "Recipe"},
				Category:      models.Category{Value: "main dish"},
				CookTime:      "PT2H30M",
				CookingMethod: models.CookingMethod{Value: "grilled"},
				Cuisine:       models.Cuisine{Value: "BBQ"},
				DatePublished: "2019-07-25",
				Description: models.Description{
					Value: "Nothing says summer like grilled BBQ ribs! These baby back ribs on the charcoal grill are " +
						"simple, delicious, and sure to please a crowd! (gluten-free, dairy-free, nut-free)",
				},
				Keywords: models.Keywords{Values: "BBQ ribs, ribs on the charcoal grill"},
				Image: models.Image{
					Value: "https://livelytable.com/wp-content/uploads/2019/07/ribs-on-charcoal-grill-2-225x225.jpg",
				},
				Ingredients: models.Ingredients{
					Values: []string{
						"1 rack baby back pork ribs",
						"1/3 cup BBQ spice rub",
						"water",
						"BBQ sauce of choice (optional)",
					},
				},
				Instructions: models.Instructions{
					Values: []string{
						"Prepare fire in the charcoal grill. Remove the grates, place a pile of charcoal on one side of the grill " +
							"only. On the other side, place a small foil pan filled with water. Start the fire and return " +
							"the grates to the grill. Let the grill get to a low temperature (about 275°F.) You may also add " +
							"pieces of wood to the charcoal for a more smoky flavor.",
						"While the fire is heating, prepare ribs. Turn ribs over so that the bone side is facing up. Remove the " +
							"membrane along the back by sliding a dull knife (such as a butter knife) under the membrane " +
							"along the last bone until you get under the membrane. Hold on tight, and pull it until the whole " +
							"thing is removed from the rack of ribs.",
						"Rub ribs all over with spice rub. Once fire is ready, place the ribs on indirect heat - the side of the " +
							"grill that has the foil pan. Cover and cook about 2 hours, watching to make sure the fire is " +
							"maintained at a steady low temperature, adding charcoal as needed, and rotating the rack of ribs " +
							"roughly every 30 minutes so that different edges of the rack are turned toward the hot side.",
						"After 1 1/2 to 2 hours, remove ribs and wrap in foil. Return to the grill for another 30 minutes or so.",
						"When ribs are done, you can either remove them from the foil and place back on the grill, meat side down, " +
							"for a little char, or place them meat side up and brush with barbecue sauce in layers, waiting " +
							"about 5 minutes between layers. Or simply remove them from the grill to a cutting board, slice, and serve!",
					},
				},
				Name: "BBQ Ribs on the Charcoal Grill",
				NutritionSchema: models.NutritionSchema{
					Calories:      "416 calories",
					Carbohydrates: "8.9 g",
					Cholesterol:   "122.9 mg",
					Fat:           "26.3 g",
					Fiber:         "0.8 g",
					Protein:       "36.1 g",
					SaturatedFat:  "9.1 g",
					Servings:      "2",
					Sodium:        "512.8 mg",
					Sugar:         "5.1 g",
					TransFat:      "0.2 g",
				},
				PrepTime: "PT10M",
				Yield:    models.Yield{Value: 1},
				URL:      "https://livelytable.com/bbq-ribs-on-the-charcoal-grill/",
			},
		},
		{
			name: "lovingitvegan.com",
			in:   "https://lovingitvegan.com/vegan-buffalo-chicken-dip/",
			want: models.RecipeSchema{
				AtContext:     atContext,
				AtType:        models.SchemaType{Value: "Recipe"},
				Category:      models.Category{Value: "Appetizer"},
				CookTime:      "PT20M",
				Cuisine:       models.Cuisine{Value: "American"},
				DatePublished: "2022-01-21T14:31:28+00:00",
				Description: models.Description{
					Value: "This baked vegan buffalo chicken dip is rich, creamy and so cheesy. It&#39;s packed with spicy " +
						"flavor and makes the perfect crowd pleasing party dip.",
				},
				Keywords: models.Keywords{
					Values: "vegan buffalo chicken dip, vegan buffalo dip",
				},
				Image: models.Image{
					Value: "https://lovingitvegan.com/wp-content/uploads/2022/01/Vegan-Buffalo-Chicken-Dip-Square.jpg",
				},
				Ingredients: models.Ingredients{
					Values: []string{
						"1 1/2 cups Raw Cashews ((225g) Soaked in hot water for 1 hour)",
						"2 Tablespoons Lemon Juice (Freshly Squeezed)",
						"1/2 cup Canned Coconut Cream ((120ml) Unsweetened)",
						"1 teaspoon Distilled White Vinegar",
						"1 teaspoon Salt",
						"1 teaspoon Onion Powder",
						"1 teaspoon Vegan Chicken Spice (or Vegan Poultry Seasoning)",
						"1/2 cup Vegan Buffalo Sauce ((120ml))",
						"3/4 cup Nutritional Yeast ((45g))",
						"14 ounce Can Artichoke Hearts (in Brine or Water, (1 can) Drained and sliced into quarters)",
						"1/3 cup Spring Onions (Chopped)",
					},
				},
				Instructions: models.Instructions{
					Values: []string{
						"Soak the cashews. Place the cashews into a bowl. Pour boiling hot water from the kettle over the top of " +
							"the cashews to submerge them. Leave the cashews to soak for 1 hour and then drain and rinse.",
						"Preheat the oven to 375°F (190°C).",
						"Add the soaked cashews, lemon juice, coconut cream, distilled white vinegar, salt, onion powder, vegan " +
							"chicken spice, vegan buffalo sauce and nutritional yeast to the blender and blend until smooth.",
						"Transfer the blended mix to a mixing bowl.",
						"Add chopped artichoke hearts and chopped spring onions and gently fold them in.",
						"Transfer to an oven safe 9-inch round dish and smooth down.",
						"Bake for 20 minutes until lightly browned on top.",
						"Serve topped with chopped spring onions with tortilla chips, crackers, breads or veggies for dipping.",
					},
				},
				Name: "Vegan Buffalo Chicken Dip",
				NutritionSchema: models.NutritionSchema{
					Calories:       "214 kcal",
					Carbohydrates:  "13 g",
					Fat:            "16 g",
					Fiber:          "3 g",
					Protein:        "8 g",
					SaturatedFat:   "7 g",
					Servings:       "1",
					Sodium:         "938 mg",
					Sugar:          "2 g",
					UnsaturatedFat: "8 g",
				},
				PrepTime: "PT10M",
				Yield:    models.Yield{Value: 8},
				URL:      "https://lovingitvegan.com/vegan-buffalo-chicken-dip/",
			},
		},
		{
			name: "madensverden.dk",
			in:   "https://madensverden.dk/durumboller-nemme-italienske-boller-med-durum-mel/",
			want: models.RecipeSchema{
				AtContext:     atContext,
				AtType:        models.SchemaType{Value: "Recipe"},
				Category:      models.Category{Value: "Brød"},
				CookTime:      "PT12M",
				Cuisine:       models.Cuisine{Value: "Italiensk"},
				DatePublished: "2023-04-13T06:00:58+00:00",
				Description: models.Description{
					Value: "Her er min bedste opskrift på durumboller. De klassiske italienske boller, der bages med durummel, " +
						"og som er gode til alt fra morgenbordet til en sandwich.",
				},
				Keywords: models.Keywords{Values: "boller, hjemmebagt"},
				Image: models.Image{
					Value: "https://madensverden.dk/wp-content/uploads/2019/01/durumboller-opskrift.jpg",
				},
				Ingredients: models.Ingredients{
					Values: []string{
						"50 gram gær",
						"500 gram vand",
						"500 gram durum mel",
						"150 gram manitoba hvedemel",
						"10 gram bageenzymer",
						"12 gram salt",
					},
				},
				Instructions: models.Instructions{
					Values: []string{
						"Lunkent vand vejes af og kommes i skålen til røremaskine. Smuldr gæren i, og rør det ud med en ske.",
						"Sæt skålen fast på røremaskinen, og monter dejkrogen på maskinen.",
						"Durummel, manitoba mel, bageenzymer og salt vejes af i en skål. Bageenzymer kan som sagt udelades, og får " +
							"stadig gode og luftige durumboller.",
						"Tilsæt gradvist melblandingen i dejen, og ælt til sidst i mindst 10 minutter ved middel hastighed.",
						"Tag dejen ud på bordet, hvor den strækkes aflang og foldes sammen.",
						"Lægges tilbage i skålen, luk til med husholdningsfilm og lad dejen hæve i 15 minutter.",
						"Beklæd en bageplade med bagepapir.",
						"Tag dejen op, og den skal ikke æltes nu. Del i 12 lige store durumboller, som sættes på bagepladen. Drys " +
							"bollerne med lidt durum mel, og læg en dyb bradepande eller bageplade ovenpå.",
						"Lad durumbollerne efterhæve i 40 minutter.",
						"Forvarm ovnen til 250 grader over- og undervarme.",
						"Bag dine durumboller i cirka 12 minutter - eller indtil de er færdige.",
					},
				},
				Name: "Durumboller - nemme italienske boller med durum mel",
				NutritionSchema: models.NutritionSchema{
					Calories: "170 kcal",
					Servings: "1",
				},
				PrepTime: "PT15M",
				Yield:    models.Yield{Value: 12},
				URL:      "https://madensverden.dk/durumboller-nemme-italienske-boller-med-durum-mel/",
			},
		},
		{
			name: "marthastewart.com",
			in:   "https://www.marthastewart.com/1539828/lemon-glazed-sheet-cake",
			want: models.RecipeSchema{
				AtContext:     atContext,
				AtType:        models.SchemaType{Value: "Recipe"},
				DateModified:  "2020-11-17T14:30:47.000-05:00",
				DatePublished: "2019-05-13T17:09:07.000-04:00",
				Description: models.Description{
					Value: "This sweet, tangy Lemon-Glazed Sheet Cake has a delicate crumb and serves a crowd.",
				},
				Keywords: models.Keywords{Values: "lemon-glazed sheet cake, cake, dessert, lemon"},
				Image: models.Image{
					Value: "https://www.marthastewart.com/thmb/qGEGYT4Q4D1RG7fxw3rl448Vax8=/1500x0/filters:no_upscale():max_bytes(150000):strip_icc()/lemon-glazed-sheet-cake-0519-fe9760b1-horiz-365892a83bdb4ce3bae434e35f6656fe.jpgitok28jrW8YB",
				},
				Ingredients: models.Ingredients{
					Values: []string{
						"10 tablespoons unsalted butter, room temperature, plus more for pan",
						"1.333 cups sugar",
						"1.5 teaspoons baking powder",
						"0.25 teaspoon baking soda",
						"1.25 teaspoons kosher salt",
						"2 teaspoons grated lemon zest, plus 1 tablespoon fresh juice",
						"2 large eggs, room temperature",
						"1.667 cups cake flour (not self-rising), such as Swans Down",
						"0.5 cup whole milk, room temperature",
						"0.75 cup sugar",
						"0.25 cup cornstarch",
						"0.5 teaspoon kosher salt",
						"1.5 teaspoons grated lemon zest, plus ⅔ cup fresh juice (from about 3 lemons)",
						"4 large egg yolks",
						"4 tablespoons unsalted butter, room temperature",
						"Candied lemon zest , for serving (optional)",
					},
				},
				Instructions: models.Instructions{
					Values: []string{
						"Cake: Preheat oven to 350 degrees. Butter a 9-by-13-inch baking pan. Beat butter with sugar, baking " +
							"powder, baking soda, salt, and lemon zest on medium-high speed until light and fluffy, " +
							"about 3 minutes. Add eggs, one at a time, beating to combine after each addition and scraping " +
							"down sides as needed. Beat in lemon juice. Beat in flour in three additions, alternating with " +
							"milk and beginning and ending with flour. Scrape batter into prepared pan, smoothing top " +
							"with an offset spatula.",
						"Bake until a tester inserted in center comes out with a few moist crumbs, 30 to 35 minutes. Transfer " +
							"pan to a wire rack; let cool completely.",
						"Glaze: Combine sugar, cornstarch, salt, and lemon zest in a saucepan. Whisk in yolks, then 1 2/3 cups " +
							"water, lemon juice, and butter. Bring to a boil over mediumhigh heat, whisking constantly, and cook, " +
							"still whisking, 1 minute. Strain through a fine-mesh sieve into a heatproof bowl. Let stand 30 minutes, " +
							"whisking occasionally.",
						"Poke 20 holes in cake with a skewer; pour glaze over top. Refrigerate at least 2 hours and up to " +
							"overnight. Decorate with candied zest. (Leftovers can be refrigerated, wrapped in plastic, up to 3 days.)",
					},
				},
				Name:     "Lemon-Glazed Sheet Cake",
				PrepTime: "PT35M",
				URL:      "https://www.marthastewart.com/1539828/lemon-glazed-sheet-cake",
			},
		},
		{
			name: "matprat.no",
			in:   "https://www.matprat.no/oppskrifter/tradisjon/vafler/",
			want: models.RecipeSchema{
				AtContext:     atContext,
				AtType:        models.SchemaType{Value: "Recipe"},
				Category:      models.Category{Value: "Dessert"},
				Cuisine:       models.Cuisine{Value: "Europa"},
				DatePublished: "01.04.2014",
				Description: models.Description{
					Value: "Vafler er alltid en suksess! Sett frem syltet&#248;y, r&#248;mme, sm&#248;r, sukker og brunost. Da  f&#229;r alle sine &#248;nsker oppfylt. Verdens beste vafler!",
				},
				Keywords: models.Keywords{
					Values: "vafler, hvetemel, melk, egg, smør, sukker, malt kardemomme, vaffel, vaffelrøre, vafler, vafler, vafler med bær, vafler, vafler, vafler, vaffeloppskrifter, oppskrift på vafler, vaffel, vaffelkake, vaffler, vaffelrøre, den store vaffeldagen",
				},
				Image: models.Image{Value: "https://images.matprat.no/uveqekyypv"},
				Ingredients: models.Ingredients{
					Values: []string{
						"4 dl hvetemel",
						"1 dl sukker",
						"1 ts bakepulver",
						"1 ts malt kardemomme",
						"4 dl melk",
						"3 stk. egg",
						"100 g smeltet smør",
					},
				},
				Instructions: models.Instructions{
					Values: []string{
						"Ha alt det t&oslash;rre i en bolle og spe med litt av melken om gangen. R&oslash;r godt mellom hver gang " +
							"for &aring; f&aring; en glatt r&oslash;re uten melklumper.",
						"R&oslash;r inn eggene og tilsett smeltet sm&oslash;r. La r&oslash;ren svelle i 1/2 time. Juster r&oslash;" +
							"ren med litt vann eller melk om den er for tykk.",
						"Stek vaflene og server dem gjerne varme.",
					},
				},
				Name:  "Vafler",
				Yield: models.Yield{Value: 1},
				URL:   "https://www.matprat.no/oppskrifter/tradisjon/vafler/",
			},
		},
		{
			name: "melskitchencafe.com",
			in:   "https://www.melskitchencafe.com/grilled-rosemary-ranch-chicken/",
			want: models.RecipeSchema{
				AtContext:     atContext,
				AtType:        models.SchemaType{Value: "Recipe"},
				CookTime:      "PT12M",
				DatePublished: "2021-08-23T04:00:00+00:00",
				Image: models.Image{
					Value: "https://www.melskitchencafe.com/wp-content/uploads/2021/08/rosemary-ranch-chicken6.jpg",
				},
				Ingredients: models.Ingredients{
					Values: []string{
						"1/4 cup olive oil",
						"1/2 cup ranch salad dressing (see note)",
						"1/4 cup Worcestershire sauce",
						"2 tablespoons fresh lemon juice or red wine vinegar",
						"1 tablespoon finely chopped fresh rosemary",
						"1 teaspoon salt",
						"1/4 teaspoon black pepper",
						"1 1/2 - 2 pounds chicken breasts (see note)",
					},
				},
				Instructions: models.Instructions{
					Values: []string{
						"For the marinade, whisk together all the ingredients except the chicken until well-combined.",
						"Place chicken in shallow dish and pour the marinade over the chicken, turning the chicken to coat " +
							"evenly. Cover the dish and refrigerate for 8-12 hours.",
						"Grill the chicken on medium-high, about 3-4 minutes per side (exact time will depend on the " +
							"thickness of the chicken) until cooked through and an instant-read thermometer registers " +
							"165 degrees F at the thickest part of the chicken. Tent the chicken with foil and let rest " +
							"for 10 minutes before slicing and serving.",
					},
				},
				Name: "Rosemary Ranch Chicken",
				NutritionSchema: models.NutritionSchema{
					Calories:      "259 kcal",
					Carbohydrates: "2 g",
					Cholesterol:   "100 mg",
					Fat:           "13 g",
					Fiber:         "1 g",
					Protein:       "32 g",
					SaturatedFat:  "2 g",
					Servings:      "1",
					Sodium:        "718 mg",
					Sugar:         "1 g",
				},
				PrepTime: "PT500M",
				Yield:    models.Yield{Value: 6},
				URL:      "https://www.melskitchencafe.com/grilled-rosemary-ranch-chicken/",
			},
		},
		{
			name: "mindmegette.hu",
			in:   "https://www.mindmegette.hu/karamellas-lavasuti.recept/",
			want: models.RecipeSchema{
				AtContext: atContext,
				AtType:    models.SchemaType{Value: "Recipe"},
				Category:  models.Category{Value: "Aprósütemény"},
				CookTime:  "PT1H",
				Description: models.Description{
					Value: "A karamellás lávasüti egy igen csábító édesség a hét szinte bármely napján. Éppen ezért nem csak " +
						"különleges alkalmakkor kell elővenni a receptet, hanem amikor csak kedved tartja.",
				},
				Keywords: models.Keywords{
					Values: "lávasüti,lávasütemény,recept,karamella,desszert,karamellás lávasüti",
				},
				Image: models.Image{
					Value: "https://www.mindmegette.hu/images/219/O/karamelles_lava_sutik.jpg",
				},
				Ingredients: models.Ingredients{
					Values: []string{
						"0.66 bögre 1,5%-os tej",
						"0.33 bögre habtejszín",
						"0.5 bögre sózatlan vaj",
						"1 bögre kristálycukor",
						"0.5 bögre barnacukor",
						"1 tk só",
						"1 tk vaníliakivonat",
						"10 ek vaj",
						"0.75 bögre barnacukor",
						"1 db tojás",
						"1 tk vaníliakivonat",
						"2.5 bögre liszt",
						"1.5 tk sütőpor",
						"5 tk őrölt fahéj",
						"0.5 tk só",
						"fahéjas cukor",
					},
				},
				Name:     "Karamellás lávasüti receptje  |  Mindmegette.hu",
				PrepTime: "PT1H",
				Yield:    models.Yield{Value: 16},
				URL:      "https://www.mindmegette.hu/karamellas-lavasuti.recept/",
			},
		},
		{
			name: "minimalistbaker.com",
			in:   "https://minimalistbaker.com/adaptogenic-hot-chocolate-mix/",
			want: models.RecipeSchema{
				AtContext:     atContext,
				AtType:        models.SchemaType{Value: "Recipe"},
				Category:      models.Category{Value: "Beverage"},
				Cuisine:       models.Cuisine{Value: "Dairy-Free"},
				DatePublished: "2020-03-01T03:00:00+00:00",
				Description: models.Description{
					Value: "Low-sugar hot chocolate mix with raw cacao powder and adaptogens like reishi mushroom, maca, " +
						"ashwagandha, and he shou wu. The perfect low-caffeine cozy beverage to replace coffee or matcha.",
				},
				Keywords: models.Keywords{Values: "hot chocolate mix"},
				Image: models.Image{
					Value: "https://minimalistbaker.com/wp-content/uploads/2020/01/Adaptogenic-Hot-Chocolate-Mix-SQUARE.jpg",
				},
				Ingredients: models.Ingredients{
					Values: []string{
						"2/3 cup unsweetened cacao powder",
						"5 tsp Ashwagandha",
						"10 tsp Reishi mushroom powder",
						"10 tsp Maca",
						"2/3 cup Tocos* ((a.k.a. rice bran solubles))",
						"1 ¼ tsp ground cinnamon ((optional))",
						"2 ½ tsp He Shou Wu ((optional))",
						"Sweetener of choice ((we prefer stevia or coconut sugar to taste*))",
						"10 ounces very hot water ((or favorite dairy-free milk))",
						"3 Tbsp Adaptogenic Hot Chocolate Mix",
						"1 tsp coconut butter ((optional // if using dairy-free milk, you can omit))",
						"1 scoop collagen* ((optional // see notes for vegan option))",
					},
				},
				Instructions: models.Instructions{
					Values: []string{
						"To a large jar or container add cacao powder, ashwagandha, reishi mushroom powder, maca, tocos, ground " +
							"cinnamon, and he shou wu (optional). At this point you can keep it unsweetened and sweeten to taste " +
							"per batch, or you can add sweetener. We opted for a bit of stevia (we like Trader Joe’s stevia " +
							"packets, and we added about 5, so 1/2 a packet / serving). However, not all stevia is made equal, " +
							"and commonly they’re much sweeter, so add little by little! Alternatively, you could sweeten with " +
							"coconut sugar. We’d recommend 1-2 tsp / serving, so as the recipe is written, roughly 3-7 Tbsp.",
						"Will keep stored at room temperature (preferably in a cool, dark place) up to 3 months.",
						"To a high-speed blender (or small blender), add hot water (or dairy-free milk), hot chocolate mix (3 Tbsp " +
							"per 1 serving), coconut butter (optional), and collagen (optional). Add sweetener of choice if not " +
							"added when making mix. Note: If using dairy-free milk in place of water, you can opt to skip the " +
							"tocos and coconut butter.",
						"Blend on high until creamy and frothy — about 1 minute. Taste and adjust flavor as needed, adding more " +
							"sweetener to taste, cacao for rich chocolate flavor, maca for malty flavor, or coconut butter " +
							"for coconut flavor / butteriness. The adaptogens / mushrooms can be a bit on the bitter side, " +
							"so adding more coconut butter, cacao, and sweetener will offset this.",
						"Serve and enjoy immediately. You can also make this in a big batch for the week and reheat throughout " +
							"the week as needed either on the stovetop in a saucepan, or in our go-to milk frother. Leftovers " +
							"will keep in the refrigerator up to 3-4 days (though best when fresh). Not freezer friendly.",
					},
				},
				Name: "Adaptogenic Hot Chocolate Mix",
				NutritionSchema: models.NutritionSchema{
					Calories:      "183 kcal",
					Carbohydrates: "22.7 g",
					Fat:           "8.8 g",
					Fiber:         "4 g",
					Protein:       "4 g",
					SaturatedFat:  "3.3 g",
					Servings:      "1",
					Sugar:         "3.3 g",
				},
				PrepTime: "PT5M",
				Yield:    models.Yield{Value: 10},
				URL:      "https://minimalistbaker.com/adaptogenic-hot-chocolate-mix/",
			},
		},
		{
			name: "misya.info",
			in:   "https://www.misya.info/ricetta/grigliata-di-carne.htm",
			want: models.RecipeSchema{
				AtContext:     atContext,
				AtType:        models.SchemaType{Value: "Recipe"},
				Category:      models.Category{Value: "Secondi di carne"},
				CookTime:      "PT25M",
				Cuisine:       models.Cuisine{Value: "Italiana"},
				DatePublished: "2022-04-03",
				Description: models.Description{
					Value: "La grigliata di carne è un grande classico del pranzo di Pasquetta. Se quest'anno avremo la " +
						"fortuna di una Pasquetta senza pioggia potreste sfruttare solo la parte iniziale del procedimento " +
						"per poi procedere con la cottura sul barbecue. In alternativa, direi che potete sfruttare la " +
						"cottura su bistecchiera come ho fatto io: non sarà proprio la stessa cosa, ma vi assicuro che " +
						"la grigliata di carne sarà comunque buonissima e ricordatevi sempre che l'importante è riuscire " +
						"a godervi la compagnia dei vostri cari ;)",
				},
				Keywords: models.Keywords{
					Values: "grigliata di carne,ricetta grigliata di carne",
				},
				Image: models.Image{
					Value: "https://www.misya.info/wp-content/uploads/2022/03/grigliata-di-carne.jpg",
				},
				Ingredients: models.Ingredients{
					Values: []string{
						"4 salsicce e/o spiedini",
						"4 fette di pancetta non stagionata",
						"4 braciole di maiale",
						"pomodorini",
						"patate",
						"sale",
						"pepe in grani",
						"2 spicchi di aglio",
						"1 rametto di rosmarino",
						"olio di oliva extravergine",
						"pepe rosa",
						"aglio",
						"sale",
						"olio di oliva extravergine",
						"paprica",
						"birra",
						"sale",
					},
				},
				Instructions: models.Instructions{
					Values: []string{
						"Preparate le 3 marinate, semplicemente unendo in 3 contenitori diversi i vari ingredienti.",
						"Mettete la pancetta nella marinata al rosmarino, le braciole in quella al pepe rosa e le salsicce/spiedini " +
							"in quelle alla birra.Coprite con pellicola trasparente e lasciate riposare per 2 ore in frigorifero.",
						"Lavate le patate e bollitele per circa 25-30 minuti, finché non saranno quasi cotte, quindi scolatele, " +
							"pelatele e tagliatele a spicchi.Fate arroventare una bistecchiera sul fuoco, scolate la carne " +
							"dalla marinata e procedete con la cottura: iniziate con salsicce e spiedini che hanno una " +
							"cottura più lunga, poi aggiungete le braciole e solo alla fine la pancetta, insieme con patate " +
							"e pomodorini (ben lavati e asciugati), che insaporiranno a dovere.",
						"La grigliata di carne è pronta, servitela subito.",
					},
				},
				Name:     "Grigliata di carne",
				PrepTime: "PT15M",
				Yield:    models.Yield{Value: 4},
				URL:      "https://www.misya.info/ricetta/grigliata-di-carne.htm",
			},
		},
		{
			name: "momsdish.com",
			in:   "https://momsdish.com/khinkali",
			want: models.RecipeSchema{
				AtContext:     atContext,
				AtType:        models.SchemaType{Value: "Recipe"},
				Category:      models.Category{Value: "Main Course"},
				CookTime:      "PT15M",
				Cuisine:       models.Cuisine{Value: "georgian"},
				DatePublished: "2021-03-10T03:23:33+00:00",
				Description: models.Description{
					Value: "Khinkali are super flavorful, meat-filled dumplings that are similar to soup dumplings. They reheat well, making them great for meal prep and can even be frozen! Both the dough and filling are easy to make and they’re fun to assemble.",
				},
				Keywords: models.Keywords{Values: "dumpling recipe, georgian dumplings, khinkali recipe"},
				Image: models.Image{
					Value: "https://cdn.momsdish.com/wp-content/uploads/2021/02/Khinkali-Recipe-Georgian-Dumplings-018-scaled.jpg",
				},
				Ingredients: models.Ingredients{
					Values: []string{
						"4 cups all-purpose flour",
						"2 tsp salt",
						"2 eggs",
						"1 cup water",
						"1 lb ground beef",
						"1 lb ground chicken",
						"1 medium onion (minced)",
						"1 tsp ground black pepper",
						"1 tbsp salt (adjust to taste)",
						"1 tbsp herbs ((optional))",
					},
				},
				Instructions: models.Instructions{
					Values: []string{
						"In a large bowl combine the flour with the salt. Make a well in the center and add the eggs. Whisk together using a fork.",
						"Add the water to the center and fold the flour into the liquid. Knead the dough by hand until it feels elastic.",
						"Cover the kneaded dough and let it rest for at least 30 minutes.",
						"Combine the beef with the chicken, minced onion, salt and pepper.",
						"Roll out the dough as thin as you possibly can.",
						"Cut the dough into 3 inch circles. Place a dollop of the meat filling in the center.",
						"Pull the edge over the filling and pinch all around, forming little pockets with meat.",
						"Bring a large pot of water to a boil. Add a few of the Khinkali at a time. Once they float to the top, give them 2-4 minutes to simmer.",
						"Remove them from the water. Serve with butter, and fresh herbs.",
					},
				},
				Name: "Khinkali Recipe (Georgian Dumplings)",
				NutritionSchema: models.NutritionSchema{
					Calories:      "127 kcal",
					Carbohydrates: "13 g",
					Cholesterol:   "35 mg",
					Fat:           "5 g",
					Fiber:         "1 g",
					Protein:       "7 g",
					SaturatedFat:  "2 g",
					Servings:      "1",
					Sodium:        "412 mg",
					Sugar:         "1 g",
					TransFat:      "1 g",
				},
				PrepTime: "PT30M",
				Yield:    models.Yield{Value: 30},
				URL:      "https://momsdish.com/khinkali",
			},
		},
		{
			name: "momswithcrockpots.com",
			in:   "https://momswithcrockpots.com/crockpot-cornbread/",
			want: models.RecipeSchema{
				AtContext:     atContext,
				AtType:        models.SchemaType{Value: "Recipe"},
				Category:      models.Category{Value: "Bread"},
				CookTime:      "PT120M",
				DatePublished: "2018-01-01T14:50:58+00:00",
				Description: models.Description{
					Value: "Save your oven space and make this delicious Crockpot Cornbread Recipe right in your slow cooker. " +
						"Comes out perfect every time!",
				},
				Keywords: models.Keywords{Values: "holiday"},
				Image: models.Image{
					Value: "https://momswithcrockpots.com/wp-content/uploads/2018/01/Crockpot-Cornbread-1-2.jpg",
				},
				Ingredients: models.Ingredients{
					Values: []string{
						"1 Cup Cornmeal",
						"1 Cup Flour",
						"1 tablespoons Baking Powder",
						"1 tablespoons Sugar",
						"1/2 tsp salt",
						"1/4 cup melted butter ( or vegetable oil)",
						"1 cup Buttermilk (or milk substitute)",
						"1 egg (beaten)",
						"2 tablespoons Honey",
					},
				},
				Instructions: models.Instructions{
					Values: []string{
						"In a medium bowl mix together all of your dry ingredients.",
						"Stir in the milk, butter/oil, egg, and honey until a batter forms.",
						"Pour your cornbread batter into a parchment paper lined 5 qt or larger crockpot.",
						"Cover and cook on high for 1 1/2 to 2 hours.",
						"Remove from crock and serve warm.",
					},
				},
				Name:     "Crockpot Cornbread",
				PrepTime: "PT5M",
				Yield:    models.Yield{Value: 6},
				URL:      "https://momswithcrockpots.com/crockpot-cornbread/",
			},
		},
		/*{
			name: "monsieur-cuisine.com",
			in:   "https://fr.monsieur-cuisine.com/nl/recipe/concentraat-voor-runderbouillon",
			want: models.RecipeSchema{
				AtContext:     atContext,
				AtType:        models.SchemaType{Value: "Recipe"},
				Name:          "Little chocolate puddings with a molten centre",
				Yield:         models.Yield{Value: 8},
				PrepTime:      "PT15M",
				CookTime:      "PT42M",
				DatePublished: "2018-11-06",
				DateCreated:   "2018-11-06",
				Image: models.Image{
					Value: "https://www.monsieur-cuisine.com/fileadmin/_processed_/d/7/csm_24979_Rezeptfoto_925b5_dec6d60bed.jpg",
				},
				Ingredients: models.Ingredients{
					Values: []string{
						"200 g dark chocolate",
						"200 g butter",
						"6 eggs (medium)",
						"250 g sugar",
						"1 pinch of salt",
						"120 g plain flour (type 405)",
						"40 g cocoa powder",
					},
				},
				Instructions: models.Instructions{
					Values: []string{
						"Brush 8 ramekins with butter and keep them in the freezer for 30 minutes.",
						"Pre-heat the oven to 210 °C. Cut the chocolate into pieces and put it into the blender jug, then, with " +
							"the measuring beaker in place, chop for 5 seconds/speed setting 8. Add the butter and, with the " +
							"measuring beaker in place, melt for 4 minutes/speed setting 2/60 °C. Transfer the mixture to another " +
							"vessel and clean out the blender jug.",
						"Put the eggs, sugar and salt into the blender jug and, with the measuring beaker in place, mix for " +
							"2 minutes/speed setting 5. Then, the measuring beaker not inserted, mix for 60 seconds/speed " +
							"setting 3, pouring the chocolate and butter mixture slowly in through the filler opening. Add the " +
							"flour and cocoa powder and, with the measuring beaker in place, mix in for 30 seconds/speed setting " +
							"3. Transfer the mixture to the ramekins and bake on the second-bottom shelf of the oven for 12 minutes. " +
							"Serve the puddings immediately.",
					},
				},
				NutritionSchema: models.NutritionSchema{
					Calories:      "2366 kj / 565 kcal",
					Protein:       "9 g",
					Carbohydrates: "55 g",
					Fat:           "35 g",
				},

				URL: "https://www.monsieur-cuisine.com/en/recipes/detail/little-chocolate-puddings-with-a-molten-centre/",
			},
		},*/
		{
			name: "motherthyme.com",
			in:   "https://www.motherthyme.com/2018/06/blt-pasta-salad.html",
			want: models.RecipeSchema{
				AtContext:     atContext,
				AtType:        models.SchemaType{Value: "Recipe"},
				DatePublished: "2018-06-29",
				Description: models.Description{
					Value: "Everything you love about a BLT tossed in this easy and delicious BLT Pasta Salad! If you like BLT's, " +
						"you're going to love this!",
				},
				Image: models.Image{
					Value: "https://www.motherthyme.com/wp-content/uploads/2018/06/BLT-PASTA-SALAD-4-225x225.jpg",
				},
				Ingredients: models.Ingredients{
					Values: []string{
						"1 pound bacon (cooked and crumbled)",
						"1 pound penne pasta or similar (cooked, drained and cooled)",
						"2-3 cups chopped iceberg lettuce",
						"1-2 cup halved cherry tomatoes",
						"1 cup shredded cheddar cheese",
						"2 cups croutons (optional)",
						"1 cup mayonnaise",
						"1/4 cup sour cream",
						"1/4 cup milk",
						"1 tablespoons sugar",
						"1 tablespoon grated Parmesan cheese",
						"1/2 teaspoon dried chives",
						"1/2 teaspoon garlic powder",
						"1/4 teaspoon onion powder",
						"1/2 teaspoon liquid smoke (optional but gives it a hint of smoky flavor)",
						"1/2 teaspoon salt (plus more to taste)",
						"1/4 teaspoon pepper",
					},
				},
				Instructions: models.Instructions{
					Values: []string{
						"In a large bowl toss together bacon, pasta, lettuce, tomatoes and cheese.",
						"Season with a pinch of salt and pepper.",
						"In a small bowl or mason jar mix together mayonnaise, sour cream, milk, sugar, Parmesan cheese, chives, " +
							"garlic powder, onion powder, liquid smoke, salt and pepper until combined.",
						"Toss about 1/2 cup of the dressing with salad then chill salad.",
						"Before serving toss with remaining dressing, top with croutons and season with additional salt and pepper.",
					},
				},
				Name: "BLT Pasta Salad",
				URL:  "https://www.motherthyme.com/2018/06/blt-pasta-salad.html",
			},
		},
		{
			name: "mybakingaddiction.com",
			in:   "https://www.mybakingaddiction.com/pistachio-pudding-cake/",
			want: models.RecipeSchema{
				AtContext:     atContext,
				AtType:        models.SchemaType{Value: "Recipe"},
				Category:      models.Category{Value: "Cake"},
				CookTime:      "PT45M",
				DatePublished: "2022-03-10T12:36:34+00:00",
				Description: models.Description{
					Value: "Pistachio Pudding Cake is a simple bundt cake to make any time of year. Made with a cake mix and " +
						"pistachio pudding mix, this cake can be topped with a simple glaze or any number of frostings for " +
						"a delicious crowd-pleasing dessert.",
				},
				Keywords: models.Keywords{
					Values: "bundt cake, Cake, cake mix, dessert, recipe, st patrick's day",
				},
				Image: models.Image{
					Value: "https://www.mybakingaddiction.com/wp-content/uploads/2022/03/overhead-view-sliced-pistachio-cake.jpg",
				},
				Ingredients: models.Ingredients{
					Values: []string{
						"1 package yellow cake mix (15.25 ounces)",
						"1 package instant pistachio pudding mix (3.4 ounces)",
						"3/4 cup sour cream",
						"3/4 cup vegetable oil",
						"3 large eggs (lightly beaten)",
						"2 teaspoons pure vanilla extract",
						"1/2 cup water",
						"½ cup roughly chopped pistachios (plus extra for garnish)",
						"1 cup powdered sugar",
						"½ teaspoon pure vanilla extract",
						"1-2 tablespoons milk",
					},
				},
				Instructions: models.Instructions{
					Values: []string{
						"Preheat oven to 350°F. Spray a 10- or 12-cup bundt cake pan with baking spray. Set aside.",
						"In a large bowl, add the cake mix, pudding mix, sour cream, vegetable oil, eggs, vanilla extract, and " +
							"water. Beat on medium speed with an electric mixer for 2 minutes. Fold in the chopped pistachios.",
						"Pour the batter into the prepared cake pan. Bake for 45-50 minutes or until a toothpick inserted into " +
							"the cake comes out clean. Allow to cool in the pan for 20 minutes before turning out onto a " +
							"wire rack to cool completely.",
						"Whisk together the glaze ingredients. Drizzle over the cooled cake and top with additional chopped pistachios.",
					},
				},
				Name: "Pistachio Pudding Cake",
				NutritionSchema: models.NutritionSchema{
					Calories:       "425 kcal",
					Carbohydrates:  "55 g",
					Cholesterol:    "50 mg",
					Fat:            "21 g",
					Fiber:          "1 g",
					Protein:        "4 g",
					SaturatedFat:   "5 g",
					Servings:       "1",
					Sodium:         "454 mg",
					Sugar:          "37 g",
					TransFat:       "0.2 g",
					UnsaturatedFat: "15 g",
				},
				PrepTime: "PT5M",
				Yield:    models.Yield{Value: 12},
				URL:      "https://www.mybakingaddiction.com/pistachio-pudding-cake/",
			},
		},
		{
			name: "mykitchen101.com",
			in:   "https://mykitchen101.com/%e5%8e%9f%e5%91%b3%e7%89%9b%e6%b2%b9%e8%9b%8b%e7%b3%95/",
			want: models.RecipeSchema{
				AtContext:     atContext,
				AtType:        models.SchemaType{Value: "Recipe"},
				Category:      models.Category{Value: "Cake"},
				Cuisine:       models.Cuisine{Value: "baking"},
				CookTime:      "PT55M",
				Keywords:      models.Keywords{Values: "传统牛油蛋糕, 原味牛油蛋糕, 牛油蛋糕"},
				Name:          "原味牛油蛋糕 (传统牛油蛋糕)",
				DatePublished: "2017-10-19T02:26:58+00:00",
				Description: models.Description{
					Value: "牛油蛋糕是许多烘培初学者必学的蛋糕。这个原味牛油蛋糕食谱没有添加任何人造香精，所以蛋糕有着浓郁的牛油香味。食谱采用的是分蛋法来制作，所以蛋糕的组织比较细腻。",
				},
				Image: models.Image{
					Value: "https://mykitchen101.com/wp-content/uploads/2017/10/plain-butter-cake-mykitchen101-feature.jpg",
				},
				Ingredients: models.Ingredients{
					Values: []string{
						"250 g \u514b \u725b\u6cb9 (\u6709\u76d0\uff0c\u5ba4\u6e29)",
						"100 g \u514b \u7ec6\u7802\u7cd6",
						"\u00bd tsp \u8336\u5319 \u7ec6\u76d0",
						"4 \u4e2a \u86cb\u9ec4 (A\u7ea7)",
						"280 g \u514b \u4f4e\u7b4b\u9762\u7c89",
						"1 tsp \u8336\u5319 \u6ce1\u6253\u7c89 (baking powder)",
						"130 ml \u6beb\u5347 \u725b\u5976",
						"4 \u4e2a \u86cb\u767d (A\u7ea7)",
						"100 g \u514b \u7ec6\u7802\u7cd6",
					},
				},
				Instructions: models.Instructions{
					Values: []string{
						"把泡打粉加入面粉拌匀后过筛备用。",
						"把烤炉预热至160°C/320°F (可在烤炉里放一盘水，约2杯，以增加湿度，可避免蛋糕表面裂开)。",
						"把蛋白放入干净的搅拌碗，以低速打至细泡。 慢慢的把砂糖加入，继续以低速打至干性发泡 (搅拌器拉起蛋白霜会形成直立状)。(温馨提示：蛋白和用来打蛋白的器具里绝对不能有任何油份和水份，建议用洗碗剂把蛋的外壳和器具清洗干净后用干净的布抹干。用一个干净的小碗来把蛋黄和蛋白分开，如果蛋黄破裂了，那个蛋白就不要用了。)",
						"把牛油、糖和盐混合，然后用中速打至乳白状。把蛋黄一个一个加入拌匀 (完全混合后才加入另一个)。",
						"分4次把面粉和牛奶轮流加入，然后继续搅拌至均匀。",
						"把⅓的蛋白霜加入面糊中，用搅拌器轻轻混合。再把其余的蛋白霜分2次加入，用刮刀轻轻拌入至完全混合。把搅拌碗敲击桌面数次以震破大气泡。 (温馨提示： 将蛋白霜和面糊混合时，动作要轻而快。橡皮刮刀从碗底往上翻转，翻拌至均匀，不可画圈搅拌，避免蛋白霜消泡。)",
						"把面糊倒入已铺上不沾烤盘纸的8吋(20-cm)方形烤盘，用刮刀把表面稍微抹平，让烤盘跌下桌面数次以震破大气泡。",
						"以160°C/320°F烘烤45分钟后, 调至180°C/355°F继续烘烤10分钟，或至完全熟透 (用木签插入蛋糕中心不粘到面糊即可)。(温馨提醒：由于每个烤炉的温度不一样，建议的时间只供参考，请依个自的烤炉调整烘烤的时间。)",
						"出炉后，让蛋糕在烤盘里冷却10分钟后才脱模。把牛油蛋糕放在铁架上至完全冷却。",
					},
				},
				NutritionSchema: models.NutritionSchema{
					Calories:       "165 kcal",
					Carbohydrates:  "17 g",
					Cholesterol:    "55 mg",
					Fat:            "10 g",
					Fiber:          "0.3 g",
					Protein:        "3 g",
					SaturatedFat:   "6 g",
					Servings:       "1",
					Sodium:         "145 mg",
					Sugar:          "9 g",
					TransFat:       "0.3 g",
					UnsaturatedFat: "4 g",
				},
				PrepTime: "PT35M",
				Yield:    models.Yield{Value: 24},
				URL:      "https://mykitchen101.com/%e5%8e%9f%e5%91%b3%e7%89%9b%e6%b2%b9%e8%9b%8b%e7%b3%95/",
			},
		},
		{
			name: "mykitchen101en.com",
			in:   "https://mykitchen101en.com/plain-butter-cake/",
			want: models.RecipeSchema{
				AtContext:     atContext,
				AtType:        models.SchemaType{Value: "Recipe"},
				Category:      models.Category{Value: "Cake"},
				CookTime:      "PT55M",
				Cuisine:       models.Cuisine{Value: "Baking"},
				DatePublished: "2017-10-19T07:07:37+00:00",
				Name:          "Plain Traditional Butter Cake",
				Description: models.Description{
					Value: "Butter cake is a popular recipe among baking beginners. This butter cake recipe does not have any artificial flavour added, thus the cake has a rich buttery flavour.",
				},
				Keywords: models.Keywords{Values: "butter cake, marble butter cake recipe, plain butter cake"},
				Image: models.Image{
					Value: "https://mykitchen101en.com/wp-content/uploads/2017/10/plain-butter-cake-mykitchen101en-feature.jpg",
				},
				URL: "https://mykitchen101en.com/plain-butter-cake/",
				Ingredients: models.Ingredients{
					Values: []string{
						"250 g butter ((salted, room temperature))",
						"100 g fine granulated sugar ((\u00bd cup) )",
						"\u00bd tsp salt",
						"4 egg yolks ((grade A/size: L))",
						"280 g low protein flour (cake flour) ((2 cups) )",
						"1 tsp baking powder",
						"130 ml milk ((\u00bd cup + 2 tsps) )",
						"4 egg whites ((grade A/size: L))",
						"100 g fine granulated sugar ((\u00bd cup) )",
					},
				},
				Instructions: models.Instructions{
					Values: []string{
						"Mix and sieve together baking powder and low protein flour.",
						"Preheat oven to 160°C/320°F (you may add a tray of water in the oven to increase the humidity, which can help preventing the cake surface from cracking).",
						"In a clean mixing bowl, whisk egg whites on low speed until tiny bubbles. Add in sugar gradually, continue whisking on low speed until stiff peak (the peaks will point straight up and hold). (Reminder: Egg whites and equipment used to beat egg whites must be grease-free and water-free. It is advisable to clean the egg shells and equipment with dishwasher detergent, then dry with a clean cloth. Use a clean small bowl to separate the egg yolk and egg white, if the egg yolk cracks, don’t use that egg white.)",
						"Combine butter, sugar and salt, then beat over medium speed until light and fluffy. Add in egg yolks gradually, beat until fully combined before adding another.",
						"Add in flour and milk alternatively in 4 batches, mix until well combined.",
						"Add ⅓ part of meringue to batter, mix gently using balloon whisk. Add the remaining meringue in 2 batches, fold in gently using rubber spatula until well mixed. Tap mixing bowl on countertop for a few times to burst large bubbles. (Reminder: When folding meringue into batter, do it quick but gentle to avoid deflating the meringue. Scrape the bottom of mixing bowl with rubber spatula and fold over, repeat folding motion until just well mixed.)",
						"Pour the batter into greased and lined 8″ (20-cm) square baking pan, smooth the surface with spatula, then drop the cake pan on countertop for a few times to burst large bubbles.",
						"Bake at 160°C/320°F for 45 minutes, then increase to 180°C/355°F and continue baking for 10 minutes, or until fully cooked (wooden stick inserted in the centre of the cake comes out clean). (Reminder: The heat for different oven is different, the suggested time is only for reference, adjust the baking time base on your oven if necessary.)",
						"Allow the butter cake to cool in the baking pan for 10 minutes before unmoulding. Let the butter cake cools completely on a wire rack.",
					},
				},
				NutritionSchema: models.NutritionSchema{
					Calories:       "165 kcal",
					Carbohydrates:  "17 g",
					Cholesterol:    "55 mg",
					Fat:            "10 g",
					Fiber:          "0.3 g",
					Protein:        "3 g",
					SaturatedFat:   "6 g",
					Servings:       "1",
					Sodium:         "145 mg",
					Sugar:          "9 g",
					TransFat:       "0.3 g",
					UnsaturatedFat: "4 g",
				},
				PrepTime: "PT35M",
				Yield:    models.Yield{Value: 24},
			},
		},
		{
			name: "myplate.gov",
			in:   "https://www.myplate.gov/recipes/supplemental-nutrition-assistance-program-snap/20-minute-chicken-creole",
			want: models.RecipeSchema{
				AtContext: atContext,
				AtType:    models.SchemaType{Value: "Recipe"},
				Name:      "20-Minute Chicken Creole",
				Image: models.Image{
					Value: "https://myplate-prod.azureedge.us/sites/default/files/styles/recipe_525_x_350_/public/2020-10/" +
						"Chicken%20Creole.jpg?itok=IX9jnbBD",
				},
				Yield:    models.Yield{Value: 8},
				CookTime: "PT20M",
				Description: models.Description{
					Value: "This Creole-inspired dish uses chili sauce and cayenne pepper to spice it up. Tomatoes, green pepper, " +
						"celery, onions and garlic spices also surround the chicken with delicious color. This main dish can " +
						"be cooked on the stovetop or with an electric skillet.",
				},
				Ingredients: models.Ingredients{
					Values: []string{
						"1 tablespoon vegetable oil",
						"2 chicken breasts (skinless, boneless)",
						"1 can diced tomatoes (14 1/2 ounces)",
						"1 cup chili sauce",
						"1 green pepper (chopped, large)",
						"2 celery stalks (chopped)",
						"1 onion (chopped)",
						"2 garlic cloves (minced)",
						"1 teaspoon dried basil",
						"1 teaspoon parsley (dried)",
						"1/4 teaspoon cayenne pepper",
						"1/4 teaspoon salt",
					},
				},
				Instructions: models.Instructions{
					Values: []string{
						"Wash hands with soap and water.",
						"Heat pan over medium-high heat (350 °F in an electric skillet). Add vegetable oil and chicken and " +
							"cook until the chicken reaches an internal temperature of 165 °F\u00a0(3-5 minutes).",
						"Reduce heat to medium (300 °F in electric skillet).",
						"Add tomatoes with juice, chili sauce, green pepper, celery, onion, garlic, basil, parsley, cayenne " +
							"pepper, and salt.",
						"Bring to a boil; reduce heat to low and simmer, covered for 10-15 minutes.",
						"Serve over hot, cooked rice or whole wheat pasta.",
						"Refrigerate leftovers within 2\u00a0hours.",
					},
				},
				NutritionSchema: models.NutritionSchema{
					Calories:      "77",
					Fat:           "3 g",
					SaturatedFat:  "0 g",
					Cholesterol:   "21 mg",
					Sodium:        "255 mg",
					Carbohydrates: "6 g",
					Fiber:         "2 g",
					Sugar:         "3 g",
					Protein:       "8 g",
				},
				URL: "https://www.myplate.gov/recipes/supplemental-nutrition-assistance-program-snap/20-minute-chicken-creole",
			},
		},
		{
			name: "myrecipes.com",
			in:   "https://www.myrecipes.com/recipe/quick-easy-nachos",
			want: models.RecipeSchema{
				AtContext:     atContext,
				AtType:        models.SchemaType{Value: "Recipe"},
				Category:      models.Category{Value: "Appetizer"},
				CookTime:      "PT15M",
				Cuisine:       models.Cuisine{Value: "TexMex"},
				DateModified:  "2023-09-26T14:52:05.511-04:00",
				DatePublished: "2006-02-03T07:37:15.000-05:00",
				Description: models.Description{
					Value: "These classic Tex-Mex nachos are loaded to the max! Avoid soggy nachos by briefly baking them before topping with cheese, seasoned beef, refried beans, guacamole, and salsa. They&#39;re a great snack, party appetizer, or even casual weeknight dinner.",
				},
				Image: models.Image{
					Value: "https://www.simplyrecipes.com/thmb/_38VUZIotH7LHCImZlAMMtlBl50=/1500x0/filters:no_upscale():max_bytes(150000):strip_icc()/__opt__aboutcom__coeus__resources__content_migration__simply_recipes__uploads__2019__04__Nachos-LEAD-5-ab0842bd5c3a492b989240cca869cefb.jpg",
				},
				Ingredients: models.Ingredients{
					Values: []string{
						"For the spice mix:",
						"2 tablespoons chili powder",
						"1 1/2 teaspoons kosher salt",
						"1 teaspoon granulated garlic",
						"1 teaspoon granulated onion",
						"1 teaspoon ground cumin",
						"1/2 teaspoon dried oregano",
						"1/4 teaspoon black pepper",
						"Pinch of cayenne pepper (optional)",
						"For the nachos:",
						"1 teaspoon vegetable oil",
						"1 pound ground beef (80:20 lean-to-fat ratio)",
						"16 ounces (2 cups) refried beans, canned or homemade",
						"1/4 cup water",
						"1 large bag of tortilla chips",
						"4 ounces cheddar cheese, grated (about 2 cups), plus more for topping",
						"4 ounces Colby Jack cheese, grated (about 2 cups), plus more for topping",
						"1 cup pico de gallo, store-bought or homemade , plus more for topping",
						"1/4 cup chopped cilantro",
						"1 sliced jalapeño (pickled or fresh)",
						"Optional toppings:",
						"Guacamole",
						"Salsa",
						"Sour cream",
						"Canned black olives",
						"Sliced green onions",
						"Shredded lettuce",
						"Corn",
						"Hot sauce",
					},
				},
				Instructions: models.Instructions{
					Values: []string{
						"Preheat the oven to 350°F.",
						"Make the taco spice blend: Combine all of the spices (chili powder through cayenne) together in a small bowl.",
						"Make the beef and bean topping: Heat the vegetable oil on medium high heat until it begins to shimmer. Add the ground beef to the pan and season it with all of the taco spice blend. As the meat cooks, use a spoon to break the meat up into crumbles. Cook for about 8 minutes until the meat has browned and drain the fat using a colander. Return the meat to the pan and add the refried beans and the water. Heat the mixture until the beans are smooth and warmed through. Reduce the heat to low and keep the beef-bean mixture warm while you prepare the chips.",
						"Toast the chips: On a 13x18-inch oven-safe platter or sheet pan, arrange the tortilla chips in a single layer, overlapping them slightly. Toast the chips in the preheated oven for 5 minutes, or just until you begin to smell their aroma.",
						"Assemble and bake the nachos: Carefully remove the pan from the oven and top with one half of the shredded cheeses. Allow the heat from the chips to melt the cheese slightly before topping the chips with the beef and bean mixture. Sprinkle the remaining cheese over the beef and return the pan to the oven for 5 minutes, or until the cheese has fully melted.",
						"Top and serve: Top the nachos with the pico de gallo, chopped cilantro, jalapeño slices, or any of your preferred toppings. Serve hot. Did you love the recipe? Give us some stars and leave a comment below!",
					},
				},
				Keywords: models.Keywords{
					Values: "Quick and Easy, Nachos, Refried Beans, Super Bowl, TexMex, Tortilla, Gluten-Free, Appetizer, Snack, Game Day",
				},
				Name: "The Best Nachos",
				NutritionSchema: models.NutritionSchema{
					Calories:       "1237 kcal",
					Carbohydrates:  "40 g",
					Cholesterol:    "305 mg",
					Fat:            "75 g",
					Fiber:          "7 g",
					Protein:        "98 g",
					SaturatedFat:   "29 g",
					Servings:       "Serves 6",
					Sodium:         "1432 mg",
					Sugar:          "2 g",
					UnsaturatedFat: "0 g",
				},
				PrepTime: "PT15M",
				Yield:    models.Yield{Value: 6},
				URL:      "https://www.myrecipes.com/recipe/quick-easy-nachos",
			},
		},
		{
			name: "nourishedbynutrition.com",
			in:   "https://nourishedbynutrition.com/fudgy-gluten-free-tahini-brownies/",
			want: models.RecipeSchema{
				AtContext:     atContext,
				AtType:        models.SchemaType{Value: "Recipe"},
				CookTime:      "PT-472163H35M40S",
				DatePublished: "2022-02-09",
				Description: models.Description{
					Value: "Rich and fudgy gluten-free tahini brownies that just happen to be also be grain-free and nut-free! " +
						"These tahini brownies make for the perfect healthier chocolate dessert!",
				},
				Image: models.Image{
					Value: "https://nourishedbynutrition.com/wp-content/uploads/2022/02/Fudgy-Tahini-Brownies-5-of-7.jpg",
				},
				Ingredients: models.Ingredients{
					Values: []string{
						"1 cup tahini",
						"½ cup maple syrup",
						"2 eggs",
						"1 teaspoon vanilla",
						"⅓ cup cocoa powder",
						"½ teaspoon baking soda",
						"¼ teaspoon salt",
						"⅓ cup chocolate chips",
					},
				},
				Instructions: models.Instructions{
					Values: []string{
						"Preheat oven to 350ºF. Line an 8×8-inch baking pan with parchment paper.",
						"In a large bowl, combine tahini, maple syrup, eggs, and vanilla; whisk well to combine. The mixture will " +
							"thicken quite a bit. Add cocoa powder, baking soda, and salt. Continue to mix until the mixture is smooth.",
						"Melt the chocolate chips in the microwave for 90 seconds, stopping every 30 seconds to mix (this can also " +
							"be done on the stovetop). Add the melted chocolate to the batter and mix to combine.",
						"Transfer mixture to prepared baking pan. Bake for 23 to 25 minutes, or until a toothpick inserted in the" +
							" center comes out mostly clean. Sprinkle with flaky salt.",
						"Let brownies cool completely in the pan. Lift parchment to remove the brownies from the pan. Cut into 12-16 " +
							"squares.",
					},
				},
				Name:     "Fudgy Tahini Brownies",
				PrepTime: "PT10M",
				Yield:    models.Yield{Value: 12},
				URL:      "https://nourishedbynutrition.com/fudgy-gluten-free-tahini-brownies/",
			},
		},
		{
			name: "nutritionbynathalie.com",
			in:   "https://www.nutritionbynathalie.com/single-post/2020/07/30/Mexican-Cauliflower-Rice",
			want: models.RecipeSchema{
				AtContext:     atContext,
				AtType:        models.SchemaType{Value: "Recipe"},
				DateModified:  "2020-07-30T18:25:45.140Z",
				DatePublished: "2020-07-30T18:25:45.140Z",
				Description: models.Description{
					Value: "Ingredients: • 1 bag fresh or frozen cauliflower rice (if using fresh cauliflower rice, add olive oil, avocado oil or coconut oil to pan) • 1-2 Tbsp olive oil • 1/4 teaspoon turmeric • 1/4 teaspoon cayenne pepper (optional) • 1/2 teaspoon garlic powder • 3/4 cup salsa • vegan chive or scallion cream cheese • fresh cilantro, chopped • sea salt and pepper to taste Directions: Heat a pan on medium heat with oil. Add the cauliflower and allow it to cook for about 5 minutes (should be nearly fully co",
				},
				Name: "Mexican Cauliflower Rice",
				Image: models.Image{
					Value: "https://static.wixstatic.com/media/d3b5ba_7ae468273837425aa869486557b06bac~mv2.jpg/v1/fill/w_837,h_1000,al_c,q_85,usm_0.66_1.00_0.01/d3b5ba_7ae468273837425aa869486557b06bac~mv2.jpg",
				},
				Ingredients: models.Ingredients{
					Values: []string{
						"1 bag fresh or frozen cauliflower rice (if using fresh cauliflower rice, add olive oil, avocado oil " +
							"or coconut oil to pan)",
						"1-2 Tbsp olive oil",
						"1/4 teaspoon turmeric",
						"1/4 teaspoon cayenne pepper (optional)",
						"1/2 teaspoon garlic powder",
						"3/4 cup salsa",
						"vegan chive or scallion cream cheese",
						"fresh cilantro, chopped",
						"sea salt and pepper to taste",
					},
				},
				Instructions: models.Instructions{
					Values: []string{
						"Heat a pan on medium heat with oil.",
						"Add the cauliflower and allow it to cook for about 5 minutes (should be nearly fully cooked).",
						"Turn heat down to low and add turmeric, cayenne, garlic powder, salsa, salt and pepper and continue to" +
							" cook until done (about 2-3 more minutes).",
						"Stir in vegan cream cheese and cilantro. Serve immediately and enjoy!",
					},
				},
				URL: "https://www.nutritionbynathalie.com/single-post/2020/07/30/Mexican-Cauliflower-Rice",
			},
		},
		{
			name: "nytimes.com",
			in: "https://cooking.nytimes.com/recipes/8357-spaghetti-with-fried-eggs?action=click&module=" +
				"Collection%20Band%20Recipe%20Card&region=Easy%20Easter%20Dinner%20Recipes&pgType=" +
				"supercollection&rank=2",
			want: models.RecipeSchema{
				AtContext: atContext,
				AtType:    models.SchemaType{Value: "Recipe"},
				Category: models.Category{
					Value: "dinner, easy, for two, quick, weeknight, pastas, main course",
				},
				CookingMethod: models.CookingMethod{Value: ""},
				Cuisine:       models.Cuisine{Value: "italian"},
				Description: models.Description{
					Value: "Here's a quick and delicious pasta dish to make when you have little time, and even less " +
						"food in the house. All you need is a box of spaghetti, four eggs, olive oil and garlic " +
						"(Parmesan is a delicious, but optional, addition).",
				},
				Keywords: models.Keywords{Values: "egg, spaghetti, fall, vegetarian"},
				Image: models.Image{
					Value: "https://static01.nyt.com/images/2021/03/22/dining/spaghetti-with-fried-eggs/spaghetti-with-fried-eggs-mediumSquareAt3X.jpg",
				},
				Ingredients: models.Ingredients{
					Values: []string{
						"Salt",
						"1/2 pound thin spaghetti",
						"6 tablespoons extra virgin olive oil or lard",
						"2 large cloves garlic, lightly smashed and peeled",
						"4 eggs",
						"Freshly ground black pepper",
						"Freshly grated Parmesan or pecorino cheese, optional",
					},
				},
				Instructions: models.Instructions{
					Values: []string{
						"Bring a pot of salted water to the boil. Start the sauce in the next step, and start cooking the " +
							"pasta when the water boils.",
						"Combine garlic and 4 tablespoons of the oil in a small skillet over medium-low heat. Cook the garlic, " +
							"pressing it into the oil occasionally to release its flavor; it should barely color on both " +
							"sides. Remove the garlic, and add the remaining oil.",
						"Fry the eggs gently in the oil, until the whites are just about set and the yolks still quite runny. " +
							"Drain the pasta, and toss with the eggs and oil, breaking up the whites as you do. (The eggs " +
							"will finish cooking in the heat of the pasta.) Season to taste, and serve immediately, with " +
							"cheese if you like.",
					},
				},
				Name: "Spaghetti With Fried Eggs",
				NutritionSchema: models.NutritionSchema{
					Calories:       "",
					Carbohydrates:  "58 grams",
					Cholesterol:    "",
					Fat:            "34 grams",
					Fiber:          "3 grams",
					Protein:        "17 grams",
					SaturatedFat:   "6 grams",
					Sodium:         "381 milligrams",
					Sugar:          "2 grams",
					TransFat:       "0 grams",
					UnsaturatedFat: "26 grams",
				},
				Yield: models.Yield{Value: 2},
				URL: "https://cooking.nytimes.com/recipes/8357-spaghetti-with-fried-eggs?action=click&module=" +
					"Collection%20Band%20Recipe%20Card&region=Easy%20Easter%20Dinner%20Recipes&pgType=supercollection&rank=2",
			},
		},
		{
			name: "ohsheglows.com",
			in:   "https://ohsheglows.com/2017/11/23/bread-free-stuffing-balls/",
			want: models.RecipeSchema{
				AtContext:     atContext,
				AtType:        models.SchemaType{Value: "Recipe"},
				Category:      models.Category{Value: "Vegan"},
				CookTime:      "PT23M",
				Cuisine:       models.Cuisine{Value: "Canadian"},
				DatePublished: "2018-12-21 23:37:54",
				Description: models.Description{
					Value: `My recipe tester Nicole likes to call these “bread-free stuffing balls," and I think I would have to agree! These festive bites have all the flavours of traditional stuffing, but they’re protein-packed, bite-sized, and gluten-free as well. This is a new and improved version of my popular Lentil Mushroom Walnut Balls recipe. I've streamlined the procedure and provided a make-ahead version in the Tips below. This recipe moves quickly using quite a few components, so my advice is to gather all of the ingredients and do as much prep as you can before you begin. If you aren't a cranberry sauce fan, my Vegan Mushroom Gravy is a nice option too!`,
				},
				Keywords: models.Keywords{
					Values: "Vegan, Gluten-Free, Soy-Free, Budget Friendly, Freezer Friendly, Kid Friendly, Make-Ahead, " +
						"Party Favourite",
				},
				Image: models.Image{
					Value: "https://ohsheglows.com/gs_images/2018/10/Bread-Free-Stuffing-Balls-00724.jpg",
				},
				Ingredients: models.Ingredients{
					Values: []string{
						"1 tablespoon (15 mL) extra-virgin olive oil",
						"1 (8-ounce/225 g) package cremini mushrooms*",
						"3 large garlic cloves, minced",
						"2 cups (50 g) stemmed kale leaves",
						"1/2 cup (50 g) gluten-free rolled oats",
						"1 (14-ounce/398 mL) can lentils, drained and rinsed",
						"1 cup (100 g) walnut halves**",
						"1 teaspoon (5 mL) dried thyme (or 2 teaspoons fresh)",
						"1/2 teaspoon dried oregano",
						"1/4 teaspoon dried rosemary (or 1/2 teaspoon fresh, minced)",
						"1/3 cup (40 g) dried cranberries, finely chopped",
						"1 tablespoon (15 mL) ground flax",
						"2 tablespoons (30 mL) water",
						"2 1/2 teaspoons (12.5 mL) sherry vinegar",
						"3/4 to 1 teaspoon fine sea salt, to taste",
						"Freshly ground black pepper, to taste",
						"2 cups (210 g) fresh or frozen cranberries",
						"1 large (230 g) ripe pear, peeled and finely chopped",
						"1/2 cup (125 mL) pure maple syrup",
						"Small pinch fine sea salt",
					},
				},
				Instructions: models.Instructions{
					Values: []string{
						"Preheat the oven to 350°F (180°C) and line a baking sheet with parchment paper.",
						"Add the oil to a large pot and turn heat to medium. Finely chop the mushrooms until they’re roughly " +
							"the size of peas. Add chopped mushrooms to the pot along with minced garlic and a pinch of salt. " +
							"Stir until combined. Sauté for about 6 to 8 minutes, until the water from the mushrooms cooks off, " +
							"reducing heat to low if necessary to prevent burning.",
						"Meanwhile, tear the kale into large pieces and place into a food processor. Pulse (do not process) the " +
							"kale until finely chopped (pieces roughly the size of almonds), being careful not to overprocess it. " +
							"Remove and place into a bowl for later.",
						"To the processor (no need to clean it out!), add the rolled oats. Process the oats until they’re finely " +
							"chopped and resemble coarse flour, about 30 seconds.",
						"Add the drained lentils and walnuts to the processor bowl with the oat flour. Pulse the mixture, stopping" +
							" to check on it every few pulses, until it’s coarsely chopped. Be sure not to overprocess it into a " +
							"paste as you still want a lot of texture and crunchy walnut pieces. Set aside.",
						"To the pot with the mushrooms and garlic, add the herbs and sauté for 30 seconds until fragrant. Stir in " +
							"the kale and chopped dried cranberries, then turn off the heat.",
						"Stir the flax and water together in a small cup (no need to let it sit).",
						"Now add all of the food processor contents, vinegar, and flax mixture to the pot. Stir until thoroughly " +
							"combined. The dough should be heavy and dense. Add salt and pepper to taste.",
						"With lightly wet hands, shape and roll about 14 to 15 balls, roughly 3 to 4 tablespoons of dough each. " +
							"Place them on the prepared baking sheet about two inches apart.",
						"Bake for 22 to 24 minutes, until golden on the bottom and firm to touch. Remove and let cool for 5 minutes.",
						"While the Bread-Free Stuffing Balls are baking, make the Cranberry-Pear Sauce. Add the cranberries, pear, " +
							"maple syrup, and salt to a medium pot. Bring to a low boil over high heat and then reduce to medium. " +
							"Simmer uncovered for 10 to 20 minutes until thickened. Use a potato masher to mash up the pear near " +
							"the end of cooking, if desired.",
						"Leftover balls can be refrigerated in an airtight container for a few days. To reheat, add oil to a s" +
							"killet and fry over medium heat, tossing occasionally, until heated through.",
					},
				},
				Name: "Bread-Free Stuffing Balls",
				NutritionSchema: models.NutritionSchema{
					Calories:      "140 calorie",
					Carbohydrates: "18 grams",
					Fat:           "6 grams",
					Fiber:         "2 grams",
					Protein:       "4 grams",
					SaturatedFat:  "0.5 grams",
					Servings:      "1",
					Sodium:        "160 milligrams",
					Sugar:         "9 grams",
				},
				PrepTime: "PT30M",
				Yield:    models.Yield{Value: 14},
				URL:      "https://ohsheglows.com/2017/11/23/bread-free-stuffing-balls/",
			},
		},
		{
			name: "onceuponachef.com",
			in:   "https://www.onceuponachef.com/recipes/perfect-basmati-rice.html",
			want: models.RecipeSchema{
				AtContext:     atContext,
				AtType:        models.SchemaType{Value: "Recipe"},
				Category:      models.Category{Value: "Vegetables & Sides"},
				CookTime:      "PT0M",
				Cuisine:       models.Cuisine{Value: "Indian"},
				DatePublished: "2013-12-05T16:29:22-05:00",
				Description: models.Description{
					Value: "This recipe makes tender and fluffy basmati rice every time.",
				},
				Keywords: models.Keywords{Values: "All Seasons, Rice"},
				Image: models.Image{
					Value: "https://www.onceuponachef.com/images/2013/12/perfect-basmati-rice-1200x1496.jpg",
				},
				Ingredients: models.Ingredients{
					Values: []string{
						"1 cup basmati rice (preferably imported from India or Pakistan)",
						"1¾ cups water",
						"1½ tablespoons unsalted butter",
						"½ teaspoon salt",
					},
				},
				Instructions: models.Instructions{
					Values: []string{
						"Place the rice in a fine mesh strainer. Place under cold running water, swishing the rice with your hand, for 1 to 2 " +
							"minutes to release excess starch. (Alternatively, place the rice in a medium bowl and add enough water to cover by 2 inches. Using your hands, gently swish " +
							"the grains to release any excess starch. Carefully pour off the water, leaving the rice in the " +
							"bowl. Repeat four times, or until the water runs almost clear. Use a fine mesh strainer to drain the rice.)",
						"In a medium pot, bring the rice, water, butter, and salt to a boil. Cover the pot with a tight fitting " +
							"lid, then turn the heat down to a simmer and cook for 15 to 20 minutes, until all of the water i" +
							"s absorbed and the rice is tender. If the rice is still too firm, add a few more tablespoons of water " +
							"and continue cooking for a few minutes more. Remove the pan from the heat and allow it to sit covered " +
							"for 5 minutes. Fluff the rice with a fork and serve.",
						"<strong>Freezer-Friendly Instructions:</strong> This rice can be frozen in an airtight container for up " +
							"to 3 months. (Putting it in a flat layer in sealable plastic bags works well as it will take up " +
							"less space in the freezer.) No need to thaw before reheating; remove it from the freezer and reheat " +
							"in the microwave with 1 to 2 tablespoons of water.",
					},
				},
				Name: "Perfect Basmati Rice",
				NutritionSchema: models.NutritionSchema{
					Calories:      "207",
					Carbohydrates: "37 g",
					Cholesterol:   "11 mg",
					Fat:           "5 g",
					Fiber:         "1 g",
					Protein:       "3 g",
					SaturatedFat:  "3 g",
					Sodium:        "120 mg",
					Sugar:         "0 g",
				},
				PrepTime: "PT0M",
				Yield:    models.Yield{Value: 4},
				URL:      "https://www.onceuponachef.com/recipes/perfect-basmati-rice.html",
			},
		},
		{
			name: "paleorunningmomma.com",
			in:   "https://www.paleorunningmomma.com/grain-free-peanut-butter-granola-bars-vegan-paleo-option/",
			want: models.RecipeSchema{
				AtContext:     atContext,
				AtType:        models.SchemaType{Value: "Recipe"},
				Category:      models.Category{Value: "Breakfast"},
				Cuisine:       models.Cuisine{Value: "Gluten-free"},
				DatePublished: "2021-03-26T18:11:21+00:00",
				Description: models.Description{
					Value: "These no bake, chewy peanut butter granola bars are a breeze to make and so addicting!  They’re " +
						"gluten free and grain free, with both vegan and paleo options.  Perfect for quick snacks, these " +
						"homemade granola bars are perfect right out of the fridge or freezer.",
				},
				Keywords: models.Keywords{
					Values: "bars, egg-free, no-bake, nut butter, paleo, peanut butter, vegan",
				},
				Image: models.Image{
					Value: "https://www.paleorunningmomma.com/wp-content/uploads/2021/03/peanut-butter-granola-bars-5.jpg",
				},
				Ingredients: models.Ingredients{
					Values: []string{
						"1 1/2 cups pecan halves ( or walnuts)",
						"1 1/2 cups almonds",
						"1 cup unsweetened coconut flakes",
						"1/2 tsp salt",
						"2 tablespoons organic coconut oil (melted (use refined for neutral flavor))",
						"2/3 cup peanut butter or other nut butter like almond (cashew, walnut, or sunflower butter (for paleo))",
						"1/4 cup + 2 Tbsp pure maple syrup (or raw honey (for paleo))",
						"1 tsp pure vanilla extract",
						"2/3 cup mini chocolate chips (dairy free )",
					},
				},
				Instructions: models.Instructions{
					Values: []string{
						"Place the nuts in a food processor and pulse several times to “chop” them into a crumbly texture - a " +
							"few larger pieces are a good thing - don’t overmix!",
						"Transfer the nuts to a large mixing bowl and stir in coconut flakes and salt to evenly combine.",
						"Place the melted coconut oil in a medium bowl and whisk in the peanut butter and honey or maple syrup. " +
							"Once mixture is smooth and well combined, stir in the vanilla.",
						"Pour the wet mixture into the large bowl with the dry ingredients and stir to fully combine - I used a " +
							"silicone spatula for this step. Thoroughly mix to make sure all the dry mixture is coated. Once " +
							"coated, gently stir in the chocolate chips.",
						"Line an 8 x 8” or 9 x 9” square pan with parchment paper along the bottom and sides, with extra up the " +
							"sides for easy removal. Transfer mixture in and press down, using your hands, or another piece of " +
							"parchment paper to get it packed tightly into the pan.",
						"Cover the top with parchment or plastic wrap, then set in the freezer for at least 1 hour to firm, or " +
							"longer if you have time.",
						"Remove pan from freezer and grab two ends of the parchment paper to remove the bars, set on a cutting board.",
						"Using a long very sharp knife, cut into 15-20 bars. You can wrap them in parchment individually storing " +
							"in the fridge (for up to two weeks) or freezer for longer.",
						"Bars will start to melt around room temp due to the coconut oil, so they’ll need to be kept chilled to " +
							"stay firm. Enjoy!",
					},
				},
				Name: "Grain Free Peanut Butter Granola Bars {Vegan, Paleo Option}",
				NutritionSchema: models.NutritionSchema{
					Calories:      "249 kcal",
					Carbohydrates: "12 g",
					Cholesterol:   "1 mg",
					Fat:           "21 g",
					Fiber:         "3 g",
					Protein:       "6 g",
					SaturatedFat:  "7 g",
					Servings:      "1",
					Sodium:        "106 mg",
					Sugar:         "7 g",
					TransFat:      "1 g",
				},
				PrepTime: "PT15M",
				Yield:    models.Yield{Value: 20},
				URL:      "https://www.paleorunningmomma.com/grain-free-peanut-butter-granola-bars-vegan-paleo-option/",
			},
		},
		{
			name: "panelinha.com.br",
			in:   "https://www.panelinha.com.br/receita/Frango-ao-curry",
			want: models.RecipeSchema{
				AtContext:     atContext,
				AtType:        models.SchemaType{Value: "Recipe"},
				Category:      models.Category{Value: "Aves"},
				Cuisine:       models.Cuisine{Value: "Prática"},
				DatePublished: "2000-05-13",
				Description: models.Description{
					Value: "A lista de vantagens desta receita é longa: fácil, rápida, tem poucos ingredientes e muito sabor. E " +
						"tem mais: você pode preparar bem antes da hora de servir. Graças ao caldinho delicioso de creme de leite, " +
						"maçãs e especiarias, o frango segue macio, macio mesmo depois de requentado. Agora, o inegociável: investir " +
						"em um curry de qualidade, já que é ele que dá todo o sabor ao preparo.",
				},
				Image: models.Image{
					Value: "https://i.panelinha.com.br/i1/228-q-5378-frango-ao-curry-com-maca.webp",
				},
				Ingredients: models.Ingredients{
					Values: []string{
						"2 filés de peito de frango (cerca de 500g)",
						"2 maçãs fuji",
						"1 cebola",
						"2 dentes de alho",
						"2 colheres (sopa) de curry",
						"1 xícara (chá) de creme de leite fresco",
						"2 colheres (sopa) de azeite",
						"sal e pimenta-do-reino moída na hora a gosto",
						"folhas de coentro a gosto para servir",
					},
				},
				Instructions: models.Instructions{
					Values: []string{
						"Corte os filés de frango em cubos de 3 cm e transfira para uma tigela. Tempere com 1 colher (chá) de sal, " +
							"o curry e pimenta a gosto e mantenha em temperatura ambiente enquanto prepara os outros ingredientes — assim " +
							"ele absorve melhor o sabor do curry e perde o gelo antes de ir para a panela.",
						"Descasque e pique fino a cebola e os dentes de alho. Descasque e corte as maçãs em fatias de 1,5 cm; " +
							"descarte o miolo com as sementes, corte as fatias em tiras e as tiras em cubos.",
						"Leve uma panela média ao fogo médio. Quando aquecer, regue com o azeite, adicione a cebola e tempere com " +
							"uma pitada de sal. Refogue por cerca de 3 minutos, até murchar. Junte o alho e mexa bem por 1 minuto " +
							"para perfumar.",
						"Acrescente os cubos de frango e refogue por cerca de 1 minuto, até perderem a aparência crua — evite" +
							" refogar em excesso para que o frango não resseque; ele vai terminar de cozinhar com o restante dos " +
							"ingredientes.",
						"Junte os cubos de maçã, regue com o creme de leite e misture bem. Assim que começar a ferver, abaixe o" +
							" fogo e deixe cozinhar por cerca de 10 minutos, até que o frango e a maçã estejam cozidos e o molho " +
							"levemente encorpado. Sirva a seguir com folhas de coentro a gosto.",
					},
				},
				Name:  "Frango ao curry com maçã",
				Yield: models.Yield{Value: 4},
				URL:   "https://www.panelinha.com.br/receita/Frango-ao-curry",
			},
		},
		{
			name: "paninihappy.com",
			in:   "https://paninihappy.com/why-you-need-this-pumpkin-muffin-recipe/",
			want: models.RecipeSchema{
				AtContext: atContext,
				AtType:    models.SchemaType{Value: "Recipe"},
				Name:      "My Favorite Pumpkin Muffins",
				PrepTime:  "PT10M",
				CookTime:  "PT20M",
				Description: models.Description{
					Value: "Because I’ve made them many times over the years and they’re the best pumpkin muffins I’ve tasted — fluffy, flavorful, unfussy, nice doming.\n\nBecause even though baking with pumpkin can be kind of a seasonal fad, it’s a delicious one, so pumpkin on!\n\nBecause now that it’s October you’re undoubtedly going to need to bring a crowd-pleasing, autumn-appropriate dish to school/work/church/soccer, etc. Or you’re simply going to want one at home on a brisk autumn afternoon.\n\nBecause you can easily double the recipe for a big group — in fact, the orignal recipe from Erin Cooks (one of my earliest favorite food blogs) makes a whole lot of muffins — or make them in mini muffin pans or mini loaf pans for lunch boxes or cute gifts.\n\nBecause pumpkin + cake mix does not equal a recipe (yeah, I said it — sorry, Pinterest!).\n\nBecause you probably have all the ingredients on hand already (especially in October, because pumpkin time).\n\nBecause they can pass for breakfast, dessert or even a side dish — versatility awaits!\n\nBecause friends love to receive the occasional pumpkin muffin surprise on their doorstep.\n\nBecause baking these muffins doubles as an awesome home fragrance for your kitchen.\n\nBecause…oh, just turn on the oven and make ’em!",
				},
				Image: models.Image{
					Value: "https://paninihappy.com/wp-content/uploads/2014/10/pumpkin-muffins-490.jpg",
				},
				Yield: models.Yield{Value: 12},
				Ingredients: models.Ingredients{
					Values: []string{
						"1 1/2 cups flour",
						"1 teaspoon baking soda",
						"1 teaspoon baking powder",
						"1 teaspoon cinnamon",
						"1/2 teaspoon salt",
						"1/4 teaspoon ground ginger",
						"1/4 teaspoon ground nutmeg",
						"2 eggs",
						"3/4 cup sugar",
						"1 cup pumpkin purée",
						"3/4 cup vegetable oil",
					},
				},
				Instructions: models.Instructions{
					Values: []string{
						"Heat the oven to 400ºF.",
						"Combine all of the dry ingredients in a medium bowl. Beat the eggs, sugar, pumpkin, and oil until smooth. " +
							"Pour the pumpkin mixture into the dry ingredients and mix just until blended.",
						"Grease a muffin tin or fill your tin with cupcake papers. Fill the wells with the batter until they are " +
							"2/3 of the way full. Bake for 16-20 minutes. Cool 5 minutes and then complete the cooling process" +
							" on a wire rack.",
					},
				},
				URL: "https://paninihappy.com/why-you-need-this-pumpkin-muffin-recipe/",
			},
		},
		{
			name: "practicalselfreliance.com",
			in:   "https://practicalselfreliance.com/zucchini-relish/",
			want: models.RecipeSchema{
				AtContext:     atContext,
				AtType:        models.SchemaType{Value: "Recipe"},
				CookTime:      "PT10M",
				DatePublished: "2022-08-08",
				Description: models.Description{
					Value: "Zucchini relish is a flavorful topping for summer grilling, and the perfect way to use up extra " +
						"zucchini from the garden.",
				},
				Image: models.Image{
					Value: "https://creativecanning.com/wp-content/uploads/2021/02/Zucchini-Relish-61-720x720.jpg",
				},
				Ingredients: models.Ingredients{
					Values: []string{
						"2 cups zucchini, diced (about 3 medium)",
						"1 cup onion, diced (about 1 medium)",
						"1 cup red bell pepper, diced (about 2 small or 1 large)",
						"2 Tablespoons Salt (pickling and canning salt, or kosher salt)",
						"1 3/4 cups sugar",
						"2 teaspoons celery seed (whole)",
						"1 teaspoon mustard seed (whole)",
						"1 cup cider vinegar (5% acidity)",
						"Pickle Crisp Granules (optional, helps veggies stay firm after canning)",
					},
				},
				Instructions: models.Instructions{
					Values: []string{
						"Wash vegetables.",
						"Remove stem and blossom ends from zucchini and dice into 1/4 to 1/2 inch pieces. Measure 2 cups.",
						"Peel and dice onion. Measure 1 cup.",
						"Stem and seed peppers, then dice. Measure 1 cup.",
						"Important, don't skip this step! Combine diced vegetables in a large bowl and sprinkle salt over the top. " +
							"Stir gently to distribute the salt, then add water until vegetables are completely submerged. " +
							"Allow the vegetables to soak in the saltwater for 2 hours, then drain completely.",
						"Prepare a water bath canner (optional, only if canning).",
						"In a separate saucepan or stockpot, bring vinegar, sugar, and spices to a gentle simmer (180 degrees F). " +
							"Do not add salt, the salt is only used to soak veggies before draining.",
						"Add drained vegetables to the simmering vinegar/spices and gently simmer for 10 minutes.",
						"Pack hot relish into prepared half-pint or pint jars, leaving 1/2 inch headspace.",
						"If not canning, just seal jars and allow them to cool on the counter before storing in the refrigerator.",
						"If canning, de-bubble jars, wipe rims, and adjust headspace to ensure 1/2 inch. Seal with 2 part canning lids.",
						"Process in a water bath canner for 10 minutes, then turn off the heat. Allow the jars to sit in the canner " +
							"for another 5 minutes to cool slightly, then remove the jars to cool on a towel on the counter.",
						"Leave the jars undisturbed for 24 hours, then check seals. Store any unsealed jars in the refrigerator for " +
							"immediate use. Properly canned and sealed jars should maintain peak quality on the pantry shelf " +
							"for 12-18 months.",
					},
				},
				Name:     "Zucchini Relish Recipe for Canning",
				PrepTime: "PT2H10M",
				Yield:    models.Yield{Value: 4},
				URL:      "https://creativecanning.com/zucchini-relish/",
			},
		},
		{
			name: "primaledgehealth.com",
			in:   "https://www.primaledgehealth.com/slow-cooker-crack-chicken/",
			want: models.RecipeSchema{
				AtContext:     atContext,
				AtType:        models.SchemaType{Value: "Recipe"},
				Category:      models.Category{Value: "Appetizers"},
				CookTime:      "PT360M",
				Cuisine:       models.Cuisine{Value: "American"},
				DatePublished: "2022-03-09T09:12:00+00:00",
				Description: models.Description{
					Value: "Cheesy Ranch chicken, also known as Crack Chicken, is a low-carb slow cooker recipe with 5 " +
						"ingredients! Easily made into a dip or dinner, depending on how you serve it, this " +
						"ultra-creamy combo of chicken, cheese, and bacon is sure to please!",
				},
				Keywords: models.Keywords{
					Values: "Cheesy Ranch Chicken, Keto Crack Chicken, Slow Cooker Crack Chicken",
				},
				Image: models.Image{
					Value: "https://www.primaledgehealth.com/wp-content/uploads/2022/03/crack-chicken.jpg",
				},
				Ingredients: models.Ingredients{
					Values: []string{
						"3 lbs chicken thigh (or breast)",
						"12 oz cream cheese",
						"1 cup cheddar cheese (shredded)",
						"8 oz bacon (cooked and crumbled)",
						"2 1-oz Ranch seasoning packets (or follow the DIY option below)",
						"2 tsp parsley (dried)",
						"2 tsp dill (dried)",
						"2 tsp chives (dried)",
						"2 tsp onion powder",
						"2 tsp garlic powder",
						"1 tsp salt",
						"½ tsp ground black pepper",
					},
				},
				Instructions: models.Instructions{
					Values: []string{
						"Add chicken to slow cooker and cover with seasoning. Top with cream cheese. Use the back of a spoon " +
							"to smear the cream cheese over the meat.",
						"Cover the slow cooker with a lid. Cook on LOW for 6-8 hours or HIGH 3-4 hours.",
						"Once cooking time finishes, shred the chicken by using two forks and pulling the meat apart. Remove " +
							"bones if needed. Stir well and thoroughly coat the chicken with sauce.",
						"Fry bacon in a pan over medium heat on the stove. Remove and chop or crumble into small pieces. Cover " +
							"the chicken with bacon and shredded cheese. Put the lid back on and continue cooking on HIGH for " +
							"15 minutes until cheese melts.",
						"Remove from heat. Mix bacon and cheese into the chicken or serve as is with them resting on top. Garnish " +
							"with a tablespoon or two of fresh parsley if desired!",
					},
				},
				Name: "Slow Cooker Crack Chicken (Cheesy Ranch Chicken)",
				NutritionSchema: models.NutritionSchema{
					Calories:       "603 kcal",
					Carbohydrates:  "3 g",
					Cholesterol:    "247 mg",
					Fat:            "39 g",
					Fiber:          "1 g",
					Protein:        "37 g",
					SaturatedFat:   "23 g",
					Servings:       "1",
					Sodium:         "689 mg",
					Sugar:          "1 g",
					TransFat:       "1 g",
					UnsaturatedFat: "30 g",
				},
				PrepTime: "PT5M",
				Yield:    models.Yield{Value: 8},
				URL:      "https://www.primaledgehealth.com/slow-cooker-crack-chicken/",
			},
		},
		{
			name: "przepisy.pl",
			in:   "https://www.przepisy.pl/przepis/placki-ziemniaczane",
			want: models.RecipeSchema{
				AtContext:     atContext,
				AtType:        models.SchemaType{Value: "Recipe"},
				Category:      models.Category{Value: "Placki ziemniaczane"},
				DatePublished: "2010-11-12T14:29:28.000Z",
				Description: models.Description{
					Value: "Jedni zajadają się plackami ziemniaczanymi posypanymi cukrem, inni z ochotą dodają do nich gęstą, kwaśną śmietanę, a jeszcze inni najbardziej cenią przepisy na placki ziemniaczane polane mięsnym gulaszem. Niewątpliwie najlepsze placki ziemniaczane to te usmażone na złoto, których chrupiąca, lekko przypieczona skórka skrywa miękki i delikatny środek. Niezależnie od tego, jakie dodatki należą do Twoich kulinarnych faworytów, poznaj naszą wersję przyrządzania tych przysmaków. Rumiane placki ziemniaczane to coś zdecydowanie więcej niż tylko potrawa charakterystyczna dla barów mlecznych.\r\nDla kogo? \r\nPrzepis na placki ziemniaczane przypadnie do gustu smakoszom ziemniaczanych potraw. Do stołu chętnie zasiądą również zwolennicy tradycyjnych dań. Kuchnia jak u mamy? Z plackami ziemniaczanymi według naszego przepisu ten efekt uzyskasz bez problemów. Masz ochotę na odrobinę nowości? Do chrupiących placuszków dodaj łososia albo krewetki – nie pożałujesz!\r\nNa jaką okazję? \r\nPlacki ziemniaczane to idealna propozycja na poskromienie większego głodu. Świetnie sprawdzają się w jesiennych i zimowych miesiącach jako rozgrzewający obiad. Jak zrobić placki ziemniaczane w takiej odsłonie? Sos grzybowy, żurawina i oscypek to sezonowe, obowiązkowe dodatki.\r\nCzy wiesz, że? \r\nZiemniaków nie musisz trzeć na drobnej tarce. Wystarczy, że zetrzesz je na grubych oczkach. Jak zrobić placki ziemniaczane tak, aby nie rozpadły się podczas smażenia, do masy ziemniaczanej dodaj jajko i odrobinę mąki. Ich postrzępione brzegi nie tylko pięknie się prezentują, ale też genialnie smakują i chrupią.\r\nDla urozmaicenia: \r\nPamiętaj, że przepis na placki ziemniaczane możesz dowolnie urozmaicić, serwując je z ulubionymi składnikami. Jak zrobić placki ziemniaczane, które zaskoczą wszystkich? Jeśli lubisz eksperymentować, do masy (oprócz czosnku) dorzuć ulubione zioła lub drobno pokrojone warzywa. Dzięki szpinakowi, papryce chili lub cebuli będą aromatyczne i kolorowe.",
				},
				Keywords: models.Keywords{
					Values: "Na co dzień, Ziemniaki, Warzywa, Jajka, Łagodne, Bez mięsa",
				},
				Image: models.Image{
					Value: "https://s3.przepisy.pl/przepisy3ii/img/variants/800x0/placki-ziemniaczane.jpg",
				},
				Ingredients: models.Ingredients{
					Values: []string{
						"1 kilogram ziemniaki",
						"1 sztuka cebula",
						"2 sztuki jajka",
						"1 sztuka Przyprawa w Mini kostkach Czosnek Knorr",
						"1 szczypta gałka muszkatołowa",
						"1 szczypta sól",
						"3 łyżki mąka",
					},
				},
				Instructions: models.Instructions{
					Values: []string{
						"Obierz ziemniaki, zetrzyj na tarce. Odsącz masę przez sito. Zetrzyj cebulę na tarce.",
						"Dodaj do ziemniaków cebulę, jajka, gałkę muszkatołową oraz mini kostkę Knorr.",
						"Wymieszaj wszystko dobrze, dodaj mąkę, aby nadać masie odpowiednią konsystencję.",
						"Rozgrzej na patelni olej, nakładaj masę łyżką. Smaż placki z obu stron na złoty brąz i od razu podawaj.",
					},
				},
				Name:     "Placki ziemniaczane",
				PrepTime: "PT40M",
				Yield:    models.Yield{Value: 8},
				URL:      "https://www.przepisy.pl/przepis/placki-ziemniaczane",
			},
		},
		{
			name: "purelypope.com",
			in:   "https://purelypope.com/sweet-chili-brussel-sprouts/",
			want: models.RecipeSchema{
				AtContext:     atContext,
				AtType:        models.SchemaType{Value: "Recipe"},
				DateModified:  "2020-09-17T22:05:26+00:00",
				DatePublished: "2020-05-21T00:35:12+00:00",
				Name:          "Sweet Chili Brussel Sprouts",
				Yield:         models.Yield{Value: 4},
				PrepTime:      "PT10M",
				CookTime:      "PT32M",
				Image: models.Image{
					Value: "https://i0.wp.com/purelypope.com/wp-content/uploads/2020/05/IMG_5412-1-scaled.jpg",
				},
				Ingredients: models.Ingredients{
					Values: []string{
						"2 cups brussel sprouts, stems removed & cut in half",
						"2 tbsp coconut aminos",
						"1 tbsp sriracha",
						"1/2 tbsp maple syrup",
						"1 tsp sesame oil",
						"Everything bagel seasoning or sesame seeds, to top",
					},
				},
				Instructions: models.Instructions{
					Values: []string{
						"Preheat oven to 350 degrees.",
						"Whisk the sauce (coconut aminos, sriracha, maple syrup & sesame oil) together in a large bowl.",
						"Toss in brussel sprouts and coat mixture evenly over the brussels.",
						"Roast for 30 minutes.",
						"Turn oven to broil for 2-3 minutes to crisp (watch carefully to not burn.)",
						"Top with everything or sesame seeds.",
					},
				},
				URL: "https://purelypope.com/sweet-chili-brussel-sprouts/",
			},
		},
		{
			name: "purplecarrot.com",
			in: "https://www.purplecarrot.com/recipe/gnocchi-al-pesto-with-charred-green-beans-lemon-zucchini-bc225f0b-" +
				"1985-4d94-b05b-a78de295b2da?plan=chefs_choice",
			want: models.RecipeSchema{
				AtContext:     atContext,
				AtType:        models.SchemaType{Value: "Recipe"},
				DatePublished: "2020-05-13T06:59:57.162-04:00",
				Image: models.Image{
					Value: "https://images.purplecarrot.com/uploads/product/image/2017/_1400_700_GnocchiAlPestowithCharredGreenBeans_LemonZucchini_WEBHERO-5d97b356980badd112987864879c4f71.jpg",
				},
				Ingredients: models.Ingredients{
					Values: []string{
						"1 zucchini, trimmed and peeled lengthwise into ribbons",
						"1 lemon, half juiced, half cut into wedges (divided)",
						"1 tsp Aleppo pepper flakes",
						"6 oz green beans, cut in half",
						"10 oz fresh gnocchi",
						"¼ cup vegan basil pesto",
						"1 tbsp + 2 tsp olive oil*",
						"Salt and pepper*",
						"*Not included",
						"Ingredients are listed for 2 servings. If you're making 4 servings, please double your ingredients by using both meal kit bags provided.",
						"For full ingredient list, see Nutrition.",
					},
				},
				Instructions: models.Instructions{
					Values: []string{
						"1 - Prepare the zucchini.",
						"2 - Char the green beans.",
						"3 - Cook the gnocchi.",
						"4 - Finish the gnocchi.",
						"5 - Serve.",
					},
				},
				Name: "Gnocchi Al Pesto with Charred Green Beans & Lemon Zucchini",
				NutritionSchema: models.NutritionSchema{
					Calories: "540 cal",
					Fat:      "22.0 g",
				},
				PrepTime: "PT30M",
				Yield:    models.Yield{Value: 2},
				URL:      "https://www.purplecarrot.com/recipe/gnocchi-al-pesto-with-charred-green-beans-lemon-zucchini-bc225f0b-1985-4d94-b05b-a78de295b2da?plan=chefs_choice",
			},
		},
		{
			name: "rachlmansfield.com",
			in:   "https://rachlmansfield.com/delicious-crispy-rice-salad-gluten-free/",
			want: models.RecipeSchema{
				AtContext:     atContext,
				AtType:        models.SchemaType{Value: "Recipe"},
				CookTime:      "PT15M",
				DatePublished: "2022-04-03",
				Description: models.Description{
					Value: "This Crispy Rice Salad is such an easy and flavorful recipe to make for lunch or dinner. This salad " +
						"" +
						"is vegan, gluten-free and craveable.",
				},
				Image: models.Image{
					Value: "https://rachlmansfield.com/wp-content/uploads/2022/03/IMG_8796-scaled.jpg",
				},
				Ingredients: models.Ingredients{
					Values: []string{
						"1/2 cup basmati rice, uncooked",
						"1/2 cup broth (code RACHL)",
						"1/4 cup filtered water",
						"1 large head of tuscan kale",
						"Olive oil to massage kale",
						"2 garlic cloves, chopped",
						"1/4 cup chopped sweet onion",
						"1 edamame beans (not in the shells)",
						"1/4 cup scallions, chopped",
						"1/3 cup cherry peppers, sliced",
						"1/2 cup kimchi, chopped",
						"1/3 cup roasted unsalted peanuts",
						"1.5 tablespoons coconut aminos",
						"1 tablespoon sesame oil",
						"Salt and pepper to taste",
					},
				},
				Instructions: models.Instructions{
					Values: []string{
						"Prepare the rice per package but do a mix of the broth and water (optimal flavor!)",
						"While the rice cooks, remove stems from the kale and chop the kale",
						"Add kale to a large mixing bowl and massage with your hands with olive oil to cut the bitterness",
						"Warm a large skillet with oil, garlic and onion and cook for 3-5 minutes or until fragrant",
						"Add in the rice and press down to form a large &#8220;rice pancake&#8221; of sorts",
						"Cook on medium heat for about 8 minutes then start to stir it to crisp the other side of the rice (do " +
							"not cover the rice or it won&#8217;t crisp!)",
						"Remove rice from pan once crisped and add to mixing bowl with the kale and add in the edamame, scallions, " +
							"peppers, kimchi, peanuts and mix",
						"Dress with coconut aminos and sesame oil and salt and pepper and enjoy!",
						"You can also add some cooked salmon, chicken or any additional protein if you&#8217;d like",
					},
				},
				Name:     "Delicious Crispy Rice Salad (gluten-free)",
				PrepTime: "PT5M",
				Yield:    models.Yield{Value: 2},
				URL:      "https://rachlmansfield.com/delicious-crispy-rice-salad-gluten-free/",
			},
		},
		{
			name: "rainbowplantlife.com",
			in:   "https://rainbowplantlife.com/livornese-stewed-beans/",
			want: models.RecipeSchema{
				AtContext:     atContext,
				AtType:        models.SchemaType{Value: "Recipe"},
				Category:      models.Category{Value: "Dinner"},
				CookTime:      "PT60M",
				Cuisine:       models.Cuisine{Value: "Italian"},
				DatePublished: "2022-01-04T04:56:15+00:00",
				Description: models.Description{
					Value: "These Tuscan Stewed Beans are the ultimate rustic Italian comfort food! Made with simple pantry-friendly ingredients like onions, garlic, tomato paste and white beans, but big on gourmet Italian flavor. It&#039;s cozy and indulgent yet wholesome, vegan, and gluten-free.",
				},
				Keywords: models.Keywords{Values: "italian beans, italian white bean stew, stewed beans, tuscan beans"},
				Image: models.Image{
					Value: "https://rainbowplantlife.com/wp-content/uploads/2022/01/Livornese-stewed-beans-5-of-5.jpg",
				},
				Ingredients: models.Ingredients{
					Values: []string{
						"1/4 cup (56 mL) extra virgin olive oil",
						"1 medium yellow onion, (chopped)",
						"2 medium or large carrots, (peeled and finely chopped)",
						"2 celery ribs, (diced)",
						"4 garlic cloves, (finely chopped)",
						"½ tsp red pepper flakes",
						"1/4 cup (4g) flat-leaf parsley leaves and tender stems, (minced)",
						"1 tablespoon minced fresh sage",
						"4 1/2 tablespoons (67g) tomato paste ((in a tube, not a can)*)",
						"¾ cup (180 mL) dry white wine**",
						"1 28-ounce (800g) can whole peeled tomatoes, ( crushed by hand)",
						"1 teaspoon kosher salt, (plus more to taste)",
						"Freshly cracked black pepper",
						"1 bay leaf",
						"1 1/2 cups (360 mL) vegetable broth, plus more as desired",
						"2 (15-ounce/425g) cans cannellini beans, drained and rinsed",
						"½ cup (8g) fresh basil, (slivered***)",
					},
				},
				Instructions: models.Instructions{
					Values: []string{
						"Heat the olive oil in a Dutch oven over medium heat. Once the oil is hot, add the onion, and season with " +
							"a pinch or two of salt and pepper. Cook for 7 to 8 minutes, until golden, stirring occasionally. " +
							"Add in the carrot, celery, and garlic, with another pinch of salt and cook for 3 to 4 minutes. " +
							"Add the red pepper flakes, parsley, and sage and cook until fragrant, about 1 minute.",
						"Add the tomato paste and cook, stirring almost continuously, for 1 to 2 minutes, until it&#39;s a bit " +
							"darker in color.",
						"Pour the white wine in and deglaze the pan, scraping up any browned bits stuck to the bottom of the pot. " +
							"Allow wine to simmer rapidly for 3 minutes, or until mostly evaporated and it no longer smells like " +
							"wine, stirring often.",
						"Add tomatoes along with their juices, bay leaf, 1 teaspoon kosher salt, and several cracks of black pepper. " +
							"Cook at a rapid simmer, stirring fairly often, until the tomatoes are fully broken down and most of " +
							"the liquid has evaporated, 12 to 13 minutes.",
						"Add the veggie broth and 2 cans of beans. Reduce the heat to low, cover the pan, and maintain a decent " +
							"simmer for 30 minutes, stirring once in a while. If you want the stew to be thicker, towards the end " +
							"of cooking, use the back of a wooden spoon or a spatula to gently smash a small portion of the beans.",
						"Taste, adding a pinch of sugar if needed (if your tomatoes are good-quality, it should not be necessary). " +
							"Remove the bay leaf. Finish with chopped basil. Season to taste, adding salt and pepper as needed.",
					},
				},
				Name: "Tuscan Stewed Beans",
				NutritionSchema: models.NutritionSchema{
					Calories:       "472 kcal",
					Carbohydrates:  "59 g",
					Fat:            "16 g",
					Fiber:          "14 g",
					Protein:        "18 g",
					SaturatedFat:   "2 g",
					Servings:       "1",
					Sodium:         "1117 mg",
					Sugar:          "7 g",
					UnsaturatedFat: "13 g",
				},
				PrepTime: "PT15M",
				Yield:    models.Yield{Value: 4},
				URL:      "https://rainbowplantlife.com/livornese-stewed-beans/",
			},
		},
		{
			name: "realsimple.com",
			in:   "https://www.realsimple.com/food-recipes/browse-all-recipes/sheet-pan-chicken-and-sweet-potatoes",
			want: models.RecipeSchema{
				AtContext:     atContext,
				AtType:        models.SchemaType{Value: "Recipe"},
				DateModified:  "2018-07-05T14:07:52.000-04:00",
				DatePublished: "2016-12-07T11:48:40.000-05:00",
				Description: models.Description{
					Value: "Get the recipe for Sheet Pan Chicken and Sweet Potatoes.",
				},
				Image: models.Image{
					Value: "https://www.realsimple.com/thmb/8gMeQAdUxCc8bTx33CFY4cdH7PU=/1500x0/filters:no_upscale():max_bytes(150000):strip_icc()/sheet-pan-chicken-sweet-potatoes_0-d610f954ea1e46179f961d536abc8f32.jpg",
				},
				Ingredients: models.Ingredients{
					Values: []string{
						"4 bone-in, skin-on chicken leg quarters (about 2 lb.)",
						"2 medium sweet potatoes, peeled and cut into 1-in. wedges",
						"1 teaspoon chopped fresh sage",
						"0.75 teaspoon kosher salt, plus more to taste",
						"0.5 teaspoon black pepper, plus more to taste",
						"3 tablespoons olive oil, divided",
						"3 slices bacon",
						"3 cups baby watercress",
						"1 tablespoon fresh lemon juice",
					},
				},
				Instructions: models.Instructions{
					Values: []string{
						"Preheat oven to 450°F. Arrange the chicken and sweet potatoes side by side in a single layer on a " +
							"large rimmed baking sheet. Season with the sage, salt, and pepper and drizzle with 2 " +
							"tablespoons of the oil, tossing to coat. Lay the bacon on top of the sweet potatoes.",
						"Roast until a meat thermometer inserted into the thickest portion of a thigh registers 165°F, 20 to " +
							"25 minutes.",
						"Meanwhile, toss together the watercress, lemon juice, and the remaining 1 tablespoon of olive oil and " +
							"season to taste with salt and pepper.",
						"Serve the chicken with the sweet potatoes and salad, with the bacon crumbled over the top.",
					},
				},
				Name: "Sheet Pan Chicken and Sweet Potatoes",
				NutritionSchema: models.NutritionSchema{
					Calories:       "533 kcal",
					Carbohydrates:  "22 g",
					Cholesterol:    "181 mg",
					Fat:            "34 g",
					Protein:        "34 g",
					SaturatedFat:   "9 g",
					Sodium:         "670 mg",
					Sugar:          "5 g",
					UnsaturatedFat: "0 g",
				},
				URL: "https://www.realsimple.com/food-recipes/browse-all-recipes/sheet-pan-chicken-and-sweet-potatoes",
			},
		},
		{
			name: "recettes.qc.ca",
			in:   "https://www.recettes.qc.ca/recettes/recette/yakisoba-nouille-sautees-a-la-japonaise",
			want: models.RecipeSchema{
				AtContext:     atContext,
				AtType:        models.SchemaType{Value: "Recipe"},
				Category:      models.Category{Value: "Pâtes alimentaires"},
				CookTime:      "PT20M",
				DatePublished: "2015-07-28T21:44:00-04:00",
				Description: models.Description{
					Value: "Recette de Yakisoba (nouilles sautées à la japonaise)",
				},
				Keywords: models.Keywords{Values: "pates alimentaires"},
				Image: models.Image{
					Value: "https://m1.quebecormedia.com/emp/rqc_prod/recettes_du_quebec-_-45fe466bb6b64f799cc2ce9ab8db72f66d46ef08-_-yakisoba.jpg",
				},
				Ingredients: models.Ingredients{
					Values: []string{
						"250 g nouilles soba ou à ramen",
						"300 g porc haché",
						"1 cuillère à table huile de sésame",
						"2 cuillères à table huile de pépins de raisin",
						"1 moyen oignon coupé en 8",
						"1 gousse d'ail",
						"500 g chou coupé en fines lanières",
						"1 poivron coupé en fines tranches",
						"2 cuillères à table gingembre rouge mariné (beni-shoga)",
						"2 cuillères à table algues ao-nori séchées en poudre",
						"1 cuillère à table sucre",
						"60 mL mirin",
						"2 cuillères à table saké",
						"60 mL sauce soja japonaise",
					},
				},
				Instructions: models.Instructions{
					Values: []string{
						"Faites cuire les nouilles dans de l’eau bouillante en les gardant fermes. Égouttez-les.",
						"Préparez la sauce : mettez les ingrédients de la sauce dans une petite casserole et chauffez jusqu’à ce que le sucre soit dissous.",
						"Faites chauffer l’huile de sésame et 1 cuillerée d’huile de pépins de raisin dans un wok. Faites-y revenir le porc jusqu’à ce qu’il soit légèrement doré. Réservez.",
						"Ajoutez le reste de l’huile dans le wok et faites sauter l’oignon et l’ail jusqu’à ce que l’oignon blondisse. Ajoutez le chou et le poivron. Cuisez jusqu’à ce qu’ils soient tendres. Ajoutez les nouilles, le porc, le gingembre mariné et la sauce. Mélangez et réchauffez. Servez et parsemez avec les ao-nori.",
					},
				},
				Name:     "Yakisoba (nouilles sautées à la japonaise)",
				PrepTime: "PT15M",
				Yield:    models.Yield{Value: 4},
				URL:      "https://www.recettes.qc.ca/recettes/recette/yakisoba-nouille-sautees-a-la-japonaise",
			},
		},
		{
			name: "recipetineats.com",
			in:   "https://www.recipetineats.com/chicken-sharwama-middle-eastern/",
			want: models.RecipeSchema{
				AtContext:     atContext,
				AtType:        models.SchemaType{Value: "Recipe"},
				Category:      models.Category{Value: "Chicken"},
				CookTime:      "PT10M",
				Cuisine:       models.Cuisine{Value: "Arabic"},
				DatePublished: "2022-02-06T06:47:00+00:00",
				Description: models.Description{
					Value: "Recipe video above. The smell when this is cooking is outrageous! The marinade is very quick to " +
						"prepare and the chicken can be frozen in the marinade, then defrosted prior to cooking. " +
						"Best cooked on the outdoor grill / BBQ, but I usually make it on the stove. Serve with " +
						"Yogurt Sauce (provided) or the Tahini sauce in this recipe. Add a simple salad and " +
						"flatbread laid out on a large platter, then let everyone make their own wraps!",
				},
				Keywords: models.Keywords{Values: "Chicken Shawarma, shawarma"},
				Image: models.Image{
					Value: "https://www.recipetineats.com/wp-content/uploads/2022/02/Chicken-Shawarma-Wrap_2-SQ.jpg",
				},
				Ingredients: models.Ingredients{
					Values: []string{
						"1 kg / 2 lb chicken thigh fillets (, skinless and boneless (Note 3))",
						"1 large garlic clove (, minced (or 2 small cloves))",
						"1 tbsp ground coriander",
						"1 tbsp ground cumin",
						"1 tbsp ground cardamon",
						"1 tsp ground cayenne pepper ((reduce to 1/2 tsp to make it not spicy))",
						"2 tsp smoked paprika",
						"2 tsp salt",
						"Black pepper",
						"2 tbsp lemon juice",
						"3 tbsp olive oil",
						"1 cup Greek yoghurt",
						"1 clove garlic (, crushed)",
						"1 tsp cumin",
						"Squeeze of lemon juice",
						"Salt and pepper",
						"4 - 5 flatbreads ((Lebanese or pita bread orhomemade soft flatbreads))",
						"Sliced lettuce ((cos or iceberg))",
						"Tomato slices",
						"Red onion (, finely sliced)",
						"Cheese (, shredded (optional))",
						"Hot sauce of choice ((optional))",
					},
				},
				Instructions: models.Instructions{
					Values: []string{
						"Marinade chicken - Combine the marinade ingredients in a large ziplock bag. Add the chicken, seal, the " +
							"massage from the outside with your hands to make sure each piece is coated. Marinate 24 " +
							"hours (minimum 3 hours).",
						"Yogurt Sauce - Combine the Yogurt Sauce ingredients in a bowl and mix. Cover and put in the fridge " +
							"until required (it will last for 3 days in the fridge).",
						"Preheat stove or BBQ - Heat a large non-stick skillet with 1 tablespoon over medium high heat, or " +
							"lightly brush a BBQ hotplate/grills with oil and heat to medium high. (See notes for baking)",
						"Cook chicken - Place chicken in the skillet or on the grill and cook the first side for 4 to 5 minutes " +
							"until nicely charred. Turn and cook the other side for 3 to 4 minutes (the 2nd side takes less time).",
						"Rest - Remove chicken from the grill and cover loosely with foil. Set aside to rest for 5 minutes.",
						"Slice chicken and pile onto platter alongside flatbreads, Salad and the Yoghurt Sauce (or dairy free " +
							"Tahini sauce from this recipe).",
						"To make a wrap, get a piece of flatbread and smear with Yoghurt Sauce. Top with a bit of lettuce and " +
							"tomato and Chicken Shawarma. Roll up and enjoy!",
					},
				},
				Name: "Chicken Shawarma (Middle Eastern)",
				NutritionSchema: models.NutritionSchema{
					Calories:       "275 kcal",
					Carbohydrates:  "1.1 g",
					Cholesterol:    "140 mg",
					Fat:            "16.2 g",
					Protein:        "32.9 g",
					SaturatedFat:   "3.2 g",
					Servings:       "183 g",
					Sodium:         "918 mg",
					UnsaturatedFat: "13 g",
				},
				PrepTime: "PT10M",
				Yield:    models.Yield{Value: 4},
				URL:      "https://www.recipetineats.com/chicken-sharwama-middle-eastern/",
			},
		},
		{
			name: "redhousespice.com",
			in:   "https://redhousespice.com/pork-fried-rice/",
			want: models.RecipeSchema{
				AtContext:     atContext,
				AtType:        models.SchemaType{Value: "Recipe"},
				Category:      models.Category{Value: "Main Course"},
				CookTime:      "PT8M",
				Cuisine:       models.Cuisine{Value: "Chinese"},
				DatePublished: "2022-03-26T17:48:51+00:00",
				Description: models.Description{
					Value: "Delicious pork fried rice made in less than 20 minutes. Enjoy the mix of fluffy rice, tender pork " +
						"and crunchy veggies coated with umami-filled seasoning.",
				},
				Keywords: models.Keywords{Values: "Pork, Rice, Stir-fry"},
				Image: models.Image{
					Value: "https://redhousespice.com/wp-content/uploads/2022/03/chinese-pork-fried-rice-1-scaled.jpg",
				},
				Ingredients: models.Ingredients{
					Values: []string{
						"2 tbsp oyster sauce (see note 1 for substitutes)",
						"1 tbsp light soy sauce",
						"½ tbsp dark soy sauce",
						"⅛ tsp ground white pepper",
						"2 tbsp neutral cooking oil (divided)",
						"2 eggs, lightly beaten",
						"1 cup minced pork (about 225g/8oz)",
						"1 small onion, diced",
						"1 tbsp minced garlic",
						"1 tsp minced ginger",
						"½ cup peas (about 50g/1.8oz)",
						"½ cup carrot, diced (about 50g/1.8oz)",
						"3 cups cold cooked white rice (about 400g/14oz (see note 2))",
						"Scallions, finely chopped (for garnishing)",
					},
				},
				Instructions: models.Instructions{
					Values: []string{
						"In a small bowl, mix oyster sauce, light soy sauce, dark soy sauce and white pepper. Set aside.",
						"Heat an empty, well-seasoned wok over high heat until smoking hot. Add 1 tablespoon of oil (see note 3 " +
							"if using other cookware). Swirl to coat a bigger perimeter.",
						"Pour in the beaten egg. Once it begins to set at the bottom, stir to help the running part flow. Use a " +
							"spatula to scramble so that it turns into small pieces. Transfer out and set aside.",
						"Pour the remaining 1 tablespoon of oil into the wok. Add minced pork. Spread and flatten it to ensure " +
							"maximum contact with the wok. Wait for the bottom part to get lightly browned. Then flip and stir " +
							"to fry it thoroughly.",
						"Once the pork loses its pink colour, add onion, garlic and ginger. Fry until the onion becomes a little " +
							"transparent.",
						"Stir in peas and carrots. Retain the high heat to fry for 30 seconds or so. Add rice and return the egg " +
							"to the wok. Cook for a further 30-40 seconds.",
						"Pour the sauce mixture over. Toss and stir constantly to ensure an even coating. Once all the ingredients " +
							"are piping hot, turn off the heat. Sprinkle scallions over and give everything a final toss.",
					},
				},
				Name: "Pork Fried Rice (猪肉炒饭)",
				NutritionSchema: models.NutritionSchema{
					Calories: "455 kcal",
					Servings: "1",
				},
				PrepTime: "PT10M",
				Yield:    models.Yield{Value: 3},
				URL:      "https://redhousespice.com/pork-fried-rice/",
			},
		},
		{
			name: "reishunger.de",
			in:   "https://www.reishunger.de/rezepte/rezept/440/chicken-tikka-masala",
			want: models.RecipeSchema{
				AtContext:     atContext,
				AtType:        models.SchemaType{Value: "Recipe"},
				Category:      models.Category{Value: "Grillen"},
				Cuisine:       models.Cuisine{Value: "Indisch"},
				DatePublished: "2015-07-25T21:55:34+00:00",
				Description: models.Description{
					Value: "Hier kommt der Briten liebstes Gericht: Chicken Tikka Masala zusammen mit einem indischen Biryani " +
						"Reis und Minzjoghurt! Geflügel ist gerade im Sommer eine tolle Sache und dieses Gericht lässt " +
						"sich problemlos für mehrere Personen kochen!",
				},
				Image: models.Image{
					Value: "https://cdn.reishunger.com/chicken-tikka-masala.jpg?quality=85",
				},
				Ingredients: models.Ingredients{
					Values: []string{
						"200g Basmati Reis Pusa",
						"2 TL Indian Tikka Marinade",
						"2 EL Bio Cashewkerne",
						"2 EL Rosinen",
						"4-5 Minzblätter",
						"4-5 Minzblätter",
						"2 EL Tomatenmark",
						"300g Hähnchenbrustfilet",
						"4 Hähnchenkeulen",
						"100g Joghurt",
						"40g Butter",
						"4 Knoblauchzehen",
						"3 Schalotten",
						"2 Chillischoten",
						"1 Becher Creme Fraiche",
						"1 kleines Stück Ingwer",
						"1 Limette",
						"0.25 Gurke",
					},
				},
				Instructions: models.Instructions{
					Values: []string{
						"Den Knoblauch und den Ingwer mit einer feinen Küchenreibe in eine große Schüssel reiben. \nDie zwei Schalotten in grobe Scheiben schneiden und die Chilischoten in der Mitte teilen, nun kann man die Kerne entfernen.",
						"Die gehackten Zutaten kommen zu den Übrigen in die Schüssel. Anschließend den Saft einer Limette, 250 Gramm Joghurt, einen guten Schuss Olivenöl sowie die Gewürzmischung samt Salz und Pfeffer dazu. \nAlles zusammen mit dem Hühnchen vermengen und dieses mindestens eine halbe Stunde in den Kühlschrank stellen.",
						"Als Nächstes die marinierten Stücke der Brust aus der Schüssel nehmen und diese auf Spieße stecken. Den Grill auf direkte, starke Hitze vorheizen und die Spieße direkt über die Glut legen, um Sie von allen Seiten anzugrillen. \nSind die Spieße scharf angegrillt (Ihr könnt die Stücke auch in der Pfanne braten), diese vom Grill nehmen und sie kurz bei Seite legen.",
						"In der Zwischenzeit kann man eine Pfanne auf den Grill stellen und dort die Butter, eine grob gehackte Schalotte, das Tomatenmark sowie den Rest der Marinade hineingeben. \nAlles zusammen wird unter direkter starker Hitze angebraten und anschließend mit 500 ml Wasser aufgegossen.",
						"Ist alles etwas eingekocht gebt Ihr die Hähnchenstücke und Euren Becher Creme Fraiche dazu. \nNun wird die Pfanne in den indirekten Bereich gestellt und alles wird etwa 20-25 Minuten bei geschlossenem Deckel geschmort.",
						"Der Reis wird zusammen mit der Gewürzpaste in einen Topf gegeben, anschließend kommt die 1,5 Fache Menge an Wasser dazu (im Zweifel mit Tassen abmessen). \nNun lasst Ihr den Reis zusammen mit dem Wasser und der Gewürzmischung etwa 10 Minuten quellen ehe Ihr Ihn auf den Herd oder die Seitenkochplatte stellt.",
						"Den Reis kurz aufkochen, die Flamme zurückstellen und den Reis unter gelegentlichem Rühren garen lassen. Wenn das Wasser verschwunden ist, ist er fertig.\nJetzt kann man die Rosinen und die Cashewkerne dazu geben und den Reis eventuell mit noch etwas Gewürzpaste und Salz abschmecken.",
						"Für den Minzjoghurt die übrigen 250 Gramm Joghurt nehmen und in eine Schüssel geben. \nDie Gurke in der Mitte durchschneiden und das Kerngehäuse entfernen. Anschließend in feine Würfel schneiden. Die Minze in feine Streifen schneiden und zusammen mit den Gurkenwürfeln zum Joghurt geben. Mit Salz und Pfeffer abschmecken.",
					},
				},
				Name:  "Chicken Tikka Masala",
				Yield: models.Yield{Value: 3},
				URL:   "https://www.reishunger.de/rezepte/rezept/440/chicken-tikka-masala",
			},
		},
		{
			name: "rezeptwelt.de",
			in:   "https://www.rezeptwelt.de/vorspeisensalate-rezepte/haehnchen-nuggets/y3duba6e-e2d56-608317-cfcd2-vjez4wd6",
			want: models.RecipeSchema{
				AtContext:     atContext,
				AtType:        models.SchemaType{Value: "Recipe"},
				Cuisine:       models.Cuisine{Value: "Europäisch, Spanisch"},
				Name:          "Hähnchen-Nuggets",
				DateModified:  "2014-05-28",
				DatePublished: "2014-05-26",
				Description: models.Description{
					Value: "Hähnchen-Nuggets, ein Rezept der Kategorie Vorspeisen/Salate. Mehr Thermomix ® Rezepte auf www.rezeptwelt.de",
				},
				Image: models.Image{
					Value: "https://de.rc-cdn.community.thermomix.com/recipeimage/y3duba6e-e2d56-608317-cfcd2-vjez4wd6/57e6b699-53cb-4229-9398-1eb1ec70245e/main/haehnchen-nuggets.jpg",
				},
				Ingredients: models.Ingredients{
					Values: []string{
						"200 g Hähnchenbrust, ohne haut, in Stücken",
						"1/2 TL Salz",
						"1/2 TL Knoblauch, granuliert, oder eine Knoblauchzehe",
						"2 Scheiben Toastbrot (ohne Kruste)",
						"60 g Frischkäse",
						"60 g Milch",
						"1 Ei, mit 50 g Wasser verquirlt",
						"100 g Paniermehl",
						"Öl zum Frittieren",
					},
				},
				Instructions: models.Instructions{
					Values: []string{
						"Hähnchenfleisch, Salz, Knoblauch in den Mixtopf geben und 5 Sek./Stufe 7 zerkleinern.",
						"Brot in Stücken mit Frischkäse und Milch zugeben und 10 Sek./Stufe 7 vermischen.",
						"Fleischmischung aus dem Mixtopf nehmen, walnussgroße Bällchen formen und leicht mit dem Boden des Messbechers flach drücken. Jedes Nugget zuerst in Ei und dann in Paniermehl wenden. Öl in einer tiefen Pfanne erhitzen. Nuggets darin goldbraun frittieren und auf Küchenkrepp abtropfen lassen.",
					},
				},
				Keywords: models.Keywords{Values: "einfach,europaisch,spanisch,vorspeise,braten,snack,"},
				PrepTime: "PT20M",
				Yield:    models.Yield{Value: 20},
				URL:      "https://www.rezeptwelt.de/vorspeisensalate-rezepte/haehnchen-nuggets/y3duba6e-e2d56-608317-cfcd2-vjez4wd6",
			},
		},
		{
			name: "sallysbakingaddiction.com",
			in:   "https://sallysbakingaddiction.com/breakfast-pastries/",
			want: models.RecipeSchema{
				AtContext:     atContext,
				AtType:        models.SchemaType{Value: "Recipe"},
				Category:      models.Category{Value: "Breakfast"},
				CookTime:      "PT20M",
				CookingMethod: models.CookingMethod{Value: "Baking"},
				Cuisine:       models.Cuisine{Value: "Danish"},
				DatePublished: "2020-08-01",
				Description: models.Description{
					Value: "These homemade breakfast pastries use a variation of classic Danish pastry dough. We're working the " +
						"butter directly into the dough, which is a different method from laminating it with separate layers " +
						"of butter. Make sure the butter is very cold before beginning. This recipe yields 2 pounds of dough.",
				},
				Keywords: models.Keywords{Values: "breakfast pastries, danishes, pastry"},
				Image: models.Image{
					Value: "https://sallysbakingaddiction.com/wp-content/uploads/2020/06/breakfast-pastries-2-225x225.jpg",
				},
				Ingredients: models.Ingredients{
					Values: []string{
						"1/4 cup (60ml) warm water (between 100-110°F, 38-43°C)",
						"2 and 1/4 teaspoons Platinum Yeast from Red Star (1 standard packet)*",
						"1/4 cup (50g) granulated sugar",
						"1/2 cup (120ml) whole milk, at room temperature (between 68–72°F, 20-22°C)",
						"1 large egg, at room temperature",
						"1 teaspoon salt",
						"14 Tablespoons (196g) unsalted butter, cold",
						"2 and 1/2 cups (313g) all-purpose flour (spooned &amp; leveled), plus more for generously flouring hands, " +
							"surface, and dough",
						"2/3 cup filling (see recipe notes for options &amp; cheese filling)",
						"1 large egg",
						"2 Tablespoons (30ml) milk",
						"1 cup (120g) confectioners’ sugar",
						"2 Tablespoons (30ml) milk or heavy cream",
						"1 teaspoon pure vanilla extract",
					},
				},
				Instructions: models.Instructions{
					Values: []string{
						"To help guarantee success, I recommend reading through the recipe, watching the video tutorial, and " +
							"reading the explanations below this recipe. (All answer many FAQs.) Do not use an electric mixer " +
							"for this dough. It&#8217;s best if the dough is folded together with a wooden spoon or rubber spatula " +
							"since it is so sticky. There is very minimal mixing required.",
						"Whisk the warm water, yeast, and 1 Tablespoon (6g) of sugar together in a large bowl. Cover and allow to " +
							"rest until foamy on top, about 5 minutes. If the surface doesn&#8217;t have bubbles on top or look " +
							"foamy after 15 minutes (it should if the yeast isn&#8217;t expired), start over with a fresh packet of" +
							" yeast. Whisk in remaining sugar, the milk, egg, and salt. Once these wet ingredients are mixed together, " +
							"lightly cover and set the bowl aside as you work on the next step.",
						"Cut the cold butter into 1/4 inch slices and add to a food processor or blender. Top with 2 and 1/2 cups " +
							"flour. Pulse the mixture 12-15 times, until butter is crumbled into pea-size bits. See photo below " +
							"for a visual. Using a food processor or blender is best for this dough. Keeping that in mind, if you " +
							"don&#8217;t have one, you can use a pastry cutter to work in the butter.",
						"Pour the flour mixture into the wet yeast mixture. Very gently fold everything together using a rubber " +
							"spatula or wooden spoon. Fold *just until* the dry ingredients are moistened. The butter must remain " +
							"in pieces and crumbles, which creates a flaky pastry. Turn the sticky dough out onto a large piece of " +
							"plastic wrap, parchment paper, aluminum foil, or into any container you can tightly cover.",
						"Wrap the dough/cover up tightly and refrigerate for at least 4 hours and up to 48 hours.",
						"Take the dough out of the refrigerator to begin the “rolling and folding” process. If the dough sat for " +
							"more than 4 hours, it may have slightly puffed up and that&#8217;s ok. (It will deflate as you shape " +
							"it, which is also ok.) Very generously flour a work surface. The dough is very sticky, so make sure you " +
							"have more flour nearby as you roll and fold. Using the palm of your hands, gently flatten the dough into " +
							"a small square. Using a rolling pin, roll out into a 15&#215;8 inch rectangle. When needed, flour " +
							"the work surface and dough as you are rolling. Fold the dough into thirds as if it were a business " +
							"letter. (See photos and video tutorial.) Turn it clockwise and roll it out into a 15 inch long rectangle " +
							"again. Then, fold into thirds again. Turn it clockwise. You’ll repeat rolling and folding 1 more time for " +
							"a total of 3 times.",
						"Wrap up/seal tightly and refrigerate for at least 1 hour and up to 24 hours. You can also freeze the dough " +
							"at this point. See freezing instructions.",
						"Line two large baking sheets with parchment paper or silicone baking mats. Rimmed baking sheets are best " +
							"because butter may leak from the dough as it bakes. If you don&#8217;t have rimmed baking sheets, when " +
							"it&#8217;s time to preheat the oven, place another baking sheet on the oven rack below to catch any butter " +
							"that may drip.",
						"Take the dough out of the refrigerator and cut it in half. Wrap 1 half up and keep refrigerated as you " +
							"work with the first half. (You can freeze half of the dough at this point, use the freezing instructions " +
							"below.)",
						"Cut the first half of dough into 8 even pieces. This will be about 1/4 cup of dough per pastry. Roll each " +
							"into balls. Flatten each into a 2.5 inch circle. Use your fingers to create a lip around the edges. See " +
							"photos and video tutorial if needed. Press the center down to flatten the center as much as you can so you " +
							"can fit the filling inside. (Center puffs up as it bakes.) Arrange pastries 3 inches apart on a lined " +
							"baking sheet. Repeat with second half of dough.",
						"Spoon 2 teaspoons of fruity filling or 1 Tablespoon of cheese filling inside each.",
						"Whisk the egg wash ingredients together. Brush on the edges of each shaped pastry.",
						"This step is optional, but I very strongly recommend it. Chill the shaped pastries in the refrigerator, " +
							"covered or uncovered, for at least 15 minutes and up to 1 hour. See recipe note. You can preheat the " +
							"oven as they finish up chilling.",
						"Preheat oven to 400°F (204°C).",
						"Bake for 19-22 minutes or until golden brown around the edges. Some butter may leak from the dough, " +
							"that&#8217;s completely normal and expected. Feel free to remove the baking sheets from the oven halfway " +
							"through baking and brush the dough with any of the leaking butter, then place back in the oven to finish " +
							"baking. (That&#8217;s what I do!)",
						"Remove baked pastries from the oven. Cool for at least 5 minutes before icing/serving.",
						"Whisk the icing ingredients together. If you want a thicker icing, whisk in more confectioners’ sugar. " +
							"If you want a thinner icing, whisk in more milk or cream. Drizzle over warm pastries and serve.",
						"Cover leftover iced or un-iced pastries and store at room temperature for 1 day or in the refrigerator " +
							"for up to 5 days. Or you can freeze them for up to 3 months. Thaw before serving. Before enjoying, feel " +
							"free to reheat leftover iced or un-iced pastries in the microwave for a few seconds until warmed.",
					},
				},
				Name:     "Breakfast Pastries with Shortcut Homemade Dough",
				PrepTime: "PT6H",
				Yield:    models.Yield{Value: 16},
				URL:      "https://sallysbakingaddiction.com/breakfast-pastries/",
			},
		},
		{
			name: "saveur.com",
			in:   "https://www.saveur.com/recipes/varenyky-pierogi-recipe/",
			want: models.RecipeSchema{
				AtContext: atContext,
				AtType:    models.SchemaType{Value: "Recipe"},
				Name:      "Larisa Frumkin’s Varenyky",
				CookTime:  "PT0D1H30M",
				Description: models.Description{
					Value: "These sweet dumplings, known as pierogi in Poland and varenyky in Ukraine, are a staple of many Slavic cuisines.",
				},
				DatePublished: "2022-04-05 17:23:56",
				Image: models.Image{
					Value: "https://www.saveur.com/uploads/2022/04/HR-Pierogi-Saveur-08-scaled.jpg?auto=webp",
				},
				Yield: models.Yield{Value: 4},
				URL:   "https://www.saveur.com/recipes/varenyky-pierogi-recipe/",
				Ingredients: models.Ingredients{
					Values: []string{
						"2 cups all-purpose flour, plus more for dusting",
						"1 tsp. kosher salt",
						"2 large eggs, separated",
						"1 tbsp. vegetable oil",
						"2 cups (1 lb.) farmer’s cheese",
						"1 large egg yolk",
						"3 tbsp. sugar (or substitute salt to taste for a savory version)",
						"4 tbsp. softened unsalted butter",
						"Sour cherry preserves, sour cream, or crème fraîche, to serve",
					},
				},
				Instructions: models.Instructions{
					Values: []string{
						"Make the dough: To a food processor, add the flour and the salt. With the motor running, add the 2 egg yolks, one-by-one, then drizzle in the oil through the feeder tube. With the motor still running, drizzle in 8–10 tablespoons of cool water, just until the dough begins to form a ball around the blade. Lightly flour a clean work surface, then transfer the dough out onto it and knead just until smooth, about 2 minutes. Cover with a clean kitchen towel and set aside to rest for 30 minutes.",
						"Meanwhile, make the filling: In a medium bowl, mix together the farmer’s cheese, egg yolk, and sugar. Set aside.",
						"Lightly flour a large rimmed baking sheet and set it by your work surface.",
						"Begin shaping the varenyky. Dust your work surface lightly with flour; divide the dough in half and shape into 2 balls. Keep one ball covered with the kitchen towel and, using a lightly floured rolling pin or a hand-crank pasta roller, roll the other ball into a thin sheet, about 1⁄16 -inch thick. Using a 3-inch round cookie cutter, punch out circles of the dough. Place a heaping teaspoon of filling in the center of each circle. Brush the edges of the circles lightly with egg white, then fold into a half moons, pressing the edges firmly together with either your fingers or with the tines of a fork to seal. Place the varenyky on the baking sheet about ½-inch apart and cover with a damp cloth. Roll out the second ball of dough, and repeat, then combine all of the leftover dough scraps to make a third batch.",
						"Fill a large pot two thirds of the way with water and salt generously. Set over medium high heat and bring to a boil. Carefully lower half the varenyky into the pot. Boil, stirring occasionally to prevent sticking, until the dumplings rise to the surface and the dough is cooked through, 6–7 minutes. Using a slotted spoon, transfer the varenyky to a deep bowl and add the butter, tossing gently with the spoon to melt. Keep warm while you cook the remaining dumplings. Divide the varenyky among 4 deep plates, top with sour cherry preserves and sour cream, and serve warm.",
					},
				},
				PrepTime: "PT0D0H0M",
			},
		},
		{
			name: "seriouseats.com",
			in:   "https://www.seriouseats.com/miyeok-guk-korean-seaweed-and-brisket-soup",
			want: models.RecipeSchema{
				AtContext:     atContext,
				AtType:        models.SchemaType{Value: "Recipe"},
				Category:      models.Category{Value: "Soups and Stews"},
				CookTime:      "PT180M",
				Cuisine:       models.Cuisine{Value: "Korean"},
				DateModified:  "2023-07-12T11:51:57.502-04:00",
				DatePublished: "2020-03-02T08:00:03.000-05:00",
				Description: models.Description{
					Value: "Tender seaweed and pieces of beef brisket come together in this warming, comforting, and nutritious " +
						"Korean soup.",
				},
				Image: models.Image{
					Value: "https://www.seriouseats.com/thmb/BkhWm33gH4ho1SSMFOLwfO8O6Ww=/1500x0/filters:no_upscale():max_bytes(150000):strip_icc()/__opt__aboutcom__coeus__resources__content_migration__serious_eats__seriouseats.com__2020__02__20200128-miyeok-guk-korean-seaweed-soup-vicky-wasik-7-21447f5620914e4b9e19912a78b7306c.jpg",
				},
				Ingredients: models.Ingredients{
					Values: []string{
						"1 ounce (30g) dried miyeok seaweed (also sold under the Japanese name wakame )",
						"3 whole medium cloves garlic plus 3 finely minced medium cloves garlic, divided",
						"One 1-inch piece fresh ginger (about 1/3 ounce; 10g), peeled",
						"1/2 of a medium white onion (about 3 ounces; 85g for the half onion)",
						"12 ounces (350g) beef brisket, washed in cold water",
						"2 tablespoons (30ml) Joseon ganjang (Korean soup soy sauce; see note), divided",
						"Kosher or sea salt",
					},
				},
				Instructions: models.Instructions{
					Values: []string{
						"In a medium bowl, cover seaweed with at least 3 inches cold water and let stand at room temperature until " +
							"fully softened and hydrated, about 2 hours.",
						"Meanwhile, in a Dutch oven or pot, combine whole garlic cloves, ginger, onion, and brisket with 1 1/2 " +
							"quarts (1 1/2L) cold water and bring to a boil over high heat. Lower heat to maintain a gentle simmer " +
							"and cook, covered, until brisket is tender and broth is slightly cloudy, about 2 hours. Using a slotted " +
							"spoon, remove and discard garlic cloves, ginger, and onion from broth.",
						"Transfer brisket to a work surface and allow to cool slightly, then slice across the grain into bite-size " +
							"pieces. Transfer brisket to a small bowl and toss well with 1 tablespoon soy sauce and remaining " +
							"3 cloves minced garlic. Set aside.",
						"Drain seaweed and squeeze well to remove excess water. Transfer to work surface and roughly chop into " +
							"bite-size pieces.",
						"Return broth to a simmer and add seaweed and seasoned brisket. If the proportion of liquid to solids is " +
							"too low for your taste, you can top up with water and return to a simmer. Add remaining 1 tablespoon " +
							"soy sauce and simmer until seaweed is tender, about 30 minutes. Season to taste with salt.",
						"Ladle soup into bowls and serve alongside hot rice and any banchan (side dishes) of your choosing.",
					},
				},
				Name: "Miyeok-Guk (Korean Seaweed and Brisket Soup)",
				NutritionSchema: models.NutritionSchema{
					Calories:       "173 kcal",
					Carbohydrates:  "2 g",
					Cholesterol:    "60 mg",
					Fat:            "11 g",
					Fiber:          "0 g",
					Protein:        "17 g",
					SaturatedFat:   "4 g",
					Servings:       "4",
					Sodium:         "421 mg",
					Sugar:          "0 g",
					UnsaturatedFat: "0 g",
				},
				PrepTime: "PT5M",
				Yield:    models.Yield{Value: 6},
				URL:      "https://www.seriouseats.com/miyeok-guk-korean-seaweed-and-brisket-soup",
			},
		},
		{
			name: "simplyquinoa.com",
			in:   "https://www.simplyquinoa.com/spicy-kimchi-quinoa-bowls/",
			want: models.RecipeSchema{
				AtContext:     atContext,
				AtType:        models.SchemaType{Value: "Recipe"},
				Category:      models.Category{Value: "Main Course"},
				CookTime:      "PT10M",
				Cuisine:       models.Cuisine{Value: "korean"},
				DatePublished: "2021-01-15T07:00:21+00:00",
				Description: models.Description{
					Value: "These spicy kimchi quinoa bowls are the perfect weeknight dinner. They&#039;re quick, easy, and " +
						"super healthy, packed with protein, fermented veggies, and greens!",
				},
				Keywords: models.Keywords{Values: "egg, kimchi, quinoa bowl"},
				Image: models.Image{
					Value: "https://www.simplyquinoa.com/wp-content/uploads/2015/06/spicy-kimchi-quinoa-bowls-3.jpg",
				},
				Ingredients: models.Ingredients{
					Values: []string{
						"2 teaspoons toasted sesame oil",
						"1/2 teaspoon freshly grated ginger",
						"1 teaspoon minced garlic",
						"2 cups cooked quinoa (cooled)",
						"1 cup kimchi (chopped)",
						"2 teaspoons kimchi \"juice\" (the liquid from the jar)",
						"2 teaspoons gluten-free tamari",
						"1 teaspoon hot sauce (optional)",
						"2 cups kale (finely chopped)",
						"2 eggs",
						"1/4 cup sliced green onions for garnish (optional)",
						"Fresh ground pepper for garnish (optional)",
					},
				},
				Instructions: models.Instructions{
					Values: []string{
						"Heat the oil in a large skillet over medium heat. Add ginger and garlic and saute for 30 - 60 seconds " +
							"until fragrant. Add the quinoa and kimchi and cook until hot, about 2 - 3 minutes. Stir in kimchi " +
							"juice, tamari and hot sauce if using. Turn to low and stir occasionally while you prepare the other " +
							"ingredients.",
						"In a separate skillet, cook the eggs on low until the whites have cooked through but the yolks are " +
							"still runny, about 3 - 5 minutes.",
						"Steam the kale in a separate pot for 30 - 60 seconds until soft.",
						"Assemble the bowls, dividing the kimchi-quinoa mixture and kale evenly between two dishes. Top with green " +
							"onions and fresh pepper if using.",
					},
				},
				Name: "Spicy Kimchi Quinoa Bowls",
				NutritionSchema: models.NutritionSchema{
					Calories:      "359 kcal",
					Carbohydrates: "46 g",
					Cholesterol:   "163 mg",
					Fat:           "12 g",
					Fiber:         "5 g",
					Protein:       "17 g",
					SaturatedFat:  "2 g",
					Servings:      "1",
					Sodium:        "489 mg",
					Sugar:         "1 g",
				},
				PrepTime: "PT3M",
				Yield:    models.Yield{Value: 2},
				URL:      "https://www.simplyquinoa.com/spicy-kimchi-quinoa-bowls/",
			},
		},
		{
			name: "simplyrecipes.com",
			in:   "https://www.simplyrecipes.com/recipes/chicken_tikka_masala/",
			want: models.RecipeSchema{
				AtContext:     atContext,
				AtType:        models.SchemaType{Value: "Recipe"},
				Category:      models.Category{Value: "Dinner"},
				CookTime:      "PT20M",
				Cuisine:       models.Cuisine{Value: "British"},
				DateModified:  "2023-09-29T18:15:43.573-04:00",
				DatePublished: "2017-02-27T04:30:56.000-05:00",
				Description: models.Description{
					Value: "This easy stovetop Chicken Tikka Masala tastes just like your favorite Indian take-out and is ready " +
						"in under an hour. Leftovers are even better the next day!",
				},
				Keywords: models.Keywords{
					Values: "Comfort Food, Quick and Easy, Restaurant Favorite, British, Indian, Gluten-Free, Dinner",
				},
				Image: models.Image{
					Value: "https://www.simplyrecipes.com/thmb/pYiHJojfyPYHFzhTQS8OU0GXUlE=/1500x0/filters:no_upscale():max_bytes(150000):strip_icc()/__opt__aboutcom__coeus__resources__content_migration__simply_recipes__uploads__2017__02__2017-02-27-ChickenTikkaMasala-18-2b30d704a54e4620a0f17fd085afeef5.jpg",
				},
				Ingredients: models.Ingredients{
					Values: []string{
						"For the chicken:",
						"1 1/4 pounds boneless skinless chicken breasts, thighs, or a mix",
						"6 tablespoons plain whole milk yogurt",
						"1/2 tablespoon grated ginger",
						"3 cloves of garlic, minced",
						"1 teaspoon cumin",
						"1 teaspoon paprika",
						"1 1/4 teaspoons salt",
						"For the tikka masala sauce:",
						"2 tablespoons canola oil, divided",
						"1 small onion, thinly sliced (about 5 ounces, or 1 1/2 cups sliced)",
						"2 teaspoons grated ginger",
						"4 cloves garlic, minced",
						"1 tablespoon ground coriander",
						"2 teaspoons paprika",
						"1 teaspoon garam masala",
						"1/2 teaspoon turmeric",
						"1/2 teaspoon freshly ground black pepper",
						"1 (14-ounce) can crushed fire-roasted tomatoes (regular crushed tomatoes work, too)",
						"6 tablespoons plain whole milk yogurt",
						"1/4 to 1/2 teaspoon cayenne pepper",
						"1/2 teaspoon salt",
						"Cooked rice , to serve",
						"Cilantro, to garnish",
					},
				},
				Instructions: models.Instructions{
					Values: []string{
						"Prepare the chicken: Trim chicken thighs of any extra fat. Chop into bite-sized pieces about 1/2 to 1-inch " +
							"wide. Place the chicken thigh pieces to a medium bowl. Add the yogurt, ginger, garlic, cumin, paprika " +
							"and salt. Using your hands, combine the chicken with the spices until the chicken is evenly coated.",
						"Marinate the chicken: Cover the bowl with plastic wrap and let the chicken marinate in the fridge for at " +
							"least 45 minutes or as long as overnight. (Marinating for 4 to 6 hours is perfect.)",
						"Cook the chicken: In a large skillet, heat 1 tablespoon of canola oil over medium-high heat. Add the chicken" +
							" thigh pieces and cook for about 6 to 7 minutes, until they’re cooked through. Transfer to a plate " +
							"and set aside.",
						"Toast the spices: Wipe down the pan you used to cook the chicken. Heat remaining canola oil over medium " +
							"heat. Add the onions and cook for 5 minutes, until softened, stirring often. Add the grated ginger," +
							" minced garlic, coriander, paprika, garam masala, turmeric, black pepper, salt, and cayenne. Let the " +
							"spices cook until fragrant, about 30 seconds to a minute.",
						"Make the sauce: Add the crushed tomatoes to the pan with the spices and let everything cook for 4 minutes," +
							" stirring often. Add the yogurt and stir to combine.",
						"Simmer the sauce: Reduce the heat to medium-low and let the sauce simmer for another 4 minutes. Add the c" +
							"hicken pieces to the pan and coat with sauce.",
						"Serve: Serve over cooked basmati rice and garnish with cilantro.",
					},
				},
				Name: "Chicken Tikka Masala",
				NutritionSchema: models.NutritionSchema{
					Calories:       "324 kcal",
					Carbohydrates:  "25 g",
					Cholesterol:    "84 mg",
					Fat:            "10 g",
					Fiber:          "3 g",
					Protein:        "34 g",
					SaturatedFat:   "2 g",
					Servings:       "4",
					Sodium:         "828 mg",
					Sugar:          "5 g",
					UnsaturatedFat: "0 g",
				},
				PrepTime: "PT15M",
				Yield:    models.Yield{Value: 6},
				URL:      "https://www.simplyrecipes.com/recipes/chicken_tikka_masala/",
			},
		},
		{
			name: "simplywhisked.com",
			in:   "https://www.simplywhisked.com/dill-pickle-pasta-salad/",
			want: models.RecipeSchema{
				AtContext:     atContext,
				AtType:        models.SchemaType{Value: "Recipe"},
				Category:      models.Category{Value: "Salads"},
				CookTime:      "PT10M",
				CookingMethod: models.CookingMethod{Value: "Stovetop"},
				Cuisine:       models.Cuisine{Value: "American"},
				DatePublished: "2022-03-02",
				Description: models.Description{
					Value: "Looking for something new to bring to your next potluck? This super flavorful dill pickle pasta salad " +
						"is a crowd pleaser, and it's so easy to make. It's loaded with crunchy dill pickles, savory bacon, " +
						"toasted cashews and topped with a creamy dill dressing.",
				},
				Keywords: models.Keywords{
					Values: "dill pickle pasta salad, pasta salad with pickles, dill pasta salad, pasta salad recipe, dairy free " +
						"dill pickle pasta salad, dairy free pasta salad, dairy free macaroni salad",
				},
				Image: models.Image{
					Value: "https://www.simplywhisked.com/wp-content/uploads/2022/01/Dairy-Free-Dill-Pickle-Pasta-Salad-3-225x225.jpg",
				},
				Ingredients: models.Ingredients{
					Values: []string{
						"1 pound pasta, cooked and cooled",
						"4 slices bacon, thinly sliced",
						"3/4 cup cashews, chopped",
						"1 cup chopped dill pickles",
						"2 stalks celery, thinly sliced",
						"3 green onions, thinly sliced",
						"2 tablespoons fresh dill, chopped",
						"2 cups mayonnaise",
						"1 tablespoon Dijon mustard",
						"4 tablespoons pickle juice",
						"2 tablespoons water",
						"Coarse salt &amp; black pepper, to taste",
					},
				},
				Instructions: models.Instructions{
					Values: []string{
						"If needed, cook pasta according to package directions for al dente in a large pot of salted, boiling " +
							"water. Drain pasta and rinse with cold water.",
						"In a medium mixing bowl, whisk the mayonnaise, Dijon mustard, pickle juice, and water until smooth. " +
							"Season salt &amp; pepper, to taste.",
						"Combine salad ingredients in a large mixing bowl, reserving about 1 teaspoon fresh dill. Add dressing " +
							"and stir until evenly coated.",
						"Before serving, adjust seasoning with salt &amp; pepper (to taste) and garnish with remaining dill.",
					},
				},
				Name: "Dill Pickle Pasta Salad",
				NutritionSchema: models.NutritionSchema{
					Calories:      "386 calories",
					Carbohydrates: "25.7 g",
					Cholesterol:   "16.2 mg",
					Fat:           "28.5 g",
					Fiber:         "1.6 g",
					Protein:       "7 g",
					SaturatedFat:  "5.1 g",
					Sodium:        "356.8 mg",
					Sugar:         "1.9 g",
					TransFat:      "0.1 g",
				},
				PrepTime: "PT15M",
				Yield:    models.Yield{Value: 16},
				URL:      "https://www.simplywhisked.com/dill-pickle-pasta-salad/",
			},
		},
		{
			name: "skinnytaste.com",
			in:   "https://www.skinnytaste.com/air-fryer-steak/",
			want: models.RecipeSchema{
				AtContext:     atContext,
				AtType:        models.SchemaType{Value: "Recipe"},
				Category:      models.Category{Value: "Dinner"},
				CookTime:      "PT10M",
				Cuisine:       models.Cuisine{Value: "American"},
				DatePublished: "2022-03-29T09:04:25+00:00",
				Description: models.Description{
					Value: "Make perfect Air Fryer Steak that is seared on the outside and juicy on the inside. Air frying " +
						"steak is quick and easy with no splatter or mess in the kitchen!",
				},
				Keywords: models.Keywords{
					Values: "Air Fryer Recipes, air fryer steak, sirloin",
				},
				Image: models.Image{
					Value: "https://www.skinnytaste.com/wp-content/uploads/2022/03/Air-Fryer-Steak-6.jpg",
				},
				Ingredients: models.Ingredients{
					Values: []string{
						"1 teaspoon garlic powder",
						"1/2 teaspoon sweet paprika",
						"1 teaspoon kosher salt",
						"1/4 teaspoon black pepper",
						"4 sirloin steaks (1 inch thick (1 1/2 lbs total))",
						"olive oil spray",
					},
				},
				Instructions: models.Instructions{
					Values: []string{
						"Combine the spices in a small bowl.",
						"Spray the steak with olive oil and coat both sides with the spices.",
						"Preheat the air fryer so the basket gets hot. For a 1-inch steak, air fry 400F 10 minutes " +
							"turning halfway, for medium rare, for medium, cook 12 minutes, flipping halfway. " +
							"See temp chart below, time may vary slightly with different air fryer models, " +
							"and the thickness of the steaks.",
						"Finish with a pinch of more salt and black pepper.",
						"Let it rest, tented with foil 5 minutes before slicing.",
					},
				},
				Name: "Air Fryer Steak",
				NutritionSchema: models.NutritionSchema{
					Calories:      "221 kcal",
					Carbohydrates: "0.5 g",
					Cholesterol:   "117.5 mg",
					Fat:           "7 g",
					Fiber:         "0.5 g",
					Protein:       "39.5 g",
					SaturatedFat:  "2 g",
					Servings:      "1",
					Sodium:        "391 mg",
				},
				PrepTime: "PT5M",
				Yield:    models.Yield{Value: 4},
				URL:      "https://www.skinnytaste.com/air-fryer-steak/",
			},
		},
		{
			name: "southernliving.com",
			in:   "https://www.southernliving.com/recipes/oven-roasted-corn-on-cob",
			want: models.RecipeSchema{
				AtContext:     atContext,
				AtType:        models.SchemaType{Value: "Recipe"},
				Name:          "Oven-Roasted Corn On The Cob",
				Category:      models.Category{Value: "Side Dish"},
				Cuisine:       models.Cuisine{Value: "American"},
				DateModified:  "2023-11-05T21:00:17.338-05:00",
				DatePublished: "2019-05-14T09:02:49.000-04:00",
				Description: models.Description{
					Value: "Great corn doesn&#39;t get much easier than our Oven-Roasted Corn on the Cob recipe. The trick? Flavored butter and foil. See how to bake corn on the cob in the oven.",
				},
				Yield: models.Yield{Value: 4},
				Image: models.Image{
					Value: "https://www.southernliving.com/thmb/-bpB7uavaEqLXMhmTD0mz3Fj9c0=/1500x0/filters:no_upscale():max_bytes(150000):strip_icc()/20220408_SL_OvenRoastedCornontheCobb_Beauty_1904-ed8011d403984f0aba111ec358359e02.jpg",
				},
				Ingredients: models.Ingredients{
					Values: []string{
						"1/4 cup unsalted butter, softened",
						"1 tablespoon chopped fresh flat-leaf parsley",
						"2 medium garlic cloves, minced (2 tsp.)",
						"1 teaspoon chopped fresh rosemary",
						"1 teaspoon chopped fresh thyme",
						"3/4 teaspoon kosher salt",
						"1/2 teaspoon black pepper",
						"4 ears fresh corn, husks removed",
					},
				},
				Instructions: models.Instructions{
					Values: []string{
						"Make butter mixture: Preheat oven to 425°F. Stir together butter, parsley, garlic, rosemary, thyme, salt, and pepper in a bowl until evenly combined.",
						"Spread butter on corn: Spread 1 tablespoon herb butter on each corn cob.",
						"Wrap corn in foil: Wrap each corn on the cob individually in aluminum foil.",
						"Roast corn in oven: Place foil-wrapped corn on a baking sheet. Bake in preheated oven until corn is soft, 20 to 25 minutes, turning once halfway through cook time. Remove corn from foil, and serve",
					},
				},
				URL: "https://www.southernliving.com/recipes/oven-roasted-corn-on-cob",
			},
		},
		{
			name: "spendwithpennies.com",
			in:   "https://www.spendwithpennies.com/split-pea-soup/",
			want: models.RecipeSchema{
				AtContext:     atContext,
				AtType:        models.SchemaType{Value: "Recipe"},
				Category:      models.Category{Value: "Dinner"},
				CookTime:      "PT130M",
				Cuisine:       models.Cuisine{Value: "American"},
				DatePublished: "2023-10-15T20:32:25+00:00",
				Description: models.Description{
					Value: "Split pea soup is the perfect way to use up leftover ham. Split peas and ham are simmered in a delicious broth to create a thick and hearty soup!",
				},
				Keywords: models.Keywords{
					Values: "best recipe, ham and pea soup, how to make, leftover ham, split pea soup",
				},
				Image: models.Image{
					Value: "https://www.spendwithpennies.com/wp-content/uploads/2023/10/1200-Split-Pea-Soup-SpendWithPennies.jpg",
				},
				Ingredients: models.Ingredients{
					Values: []string{
						"2 cups dried split peas (green or yellow (14 oz))",
						"1 meaty ham bone (or 2 cups diced leftover ham)",
						"4 cups chicken broth",
						"4 cups water (or additional broth if desired)",
						"2 teaspoons dried parsley",
						"1 bay leaf",
						"3 ribs celery (diced)",
						"2  carrots (diced)",
						"1 large onion (diced)",
						"½ teaspoon black pepper",
						"½ teaspoon dried thyme",
						"salt to taste",
					},
				},
				Instructions: models.Instructions{
					Values: []string{
						"Sort through the peas to ensure there is no debris. Rinse and drain well.",
						"In a large pot, combine peas, ham, water, broth, parsley, and bay leaf. Bring to a boil, reduce heat " +
							"to low, and simmer covered for 1 hour.",
						"Add in celery, carrots, onion, pepper, thyme, and salt. Cover and simmer 45 minutes more.",
						"Remove ham bone and chop the meat. Return the meat to the soup and cook uncovered until thickened and the peas have broken down and the soup has thickened, about 20 minutes more.",
						"Discard the bay leaf and season with salt and additional pepper to taste.",
					},
				},
				Name: "Split Pea Soup",
				NutritionSchema: models.NutritionSchema{
					Calories:      "365 kcal",
					Carbohydrates: "45 g",
					Cholesterol:   "29 mg",
					Fat:           "9 g",
					Fiber:         "18 g",
					Protein:       "27 g",
					SaturatedFat:  "3 g",
					Sodium:        "900 mg",
					Sugar:         "8 g",
				},
				PrepTime: "PT20M",
				Yield:    models.Yield{Value: 6},
				URL:      "https://www.spendwithpennies.com/split-pea-soup/",
			},
		},
		{
			name: "steamykitchen.com",
			in:   "https://steamykitchen.com/4474-korean-style-tacos-with-kogi-bbq-sauce.html",
			want: models.RecipeSchema{
				AtContext:     atContext,
				AtType:        models.SchemaType{Value: "Recipe"},
				Category:      models.Category{Value: "Main Course"},
				Cuisine:       models.Cuisine{Value: "American"},
				DatePublished: "2020-07-25T05:19:27+00:00",
				Description: models.Description{
					Value: "This is a great way to use your leftover pulled pork or roasted chicken. The BBQ Sauce from Kogi BBQ " +
						"was created by Chef Roy to be strong flavored enough to match the smokiness of BBQ’d pork or roasted " +
						"chicken. You can add use kimchi (spicy pickled Korean cabbage) to top the tacos, or make a quick " +
						"cucumber pickle like I have. The recipe for the quick cucumber pickle is below.",
				},
				Keywords: models.Keywords{Values: "korean bbq, taco"},
				Image: models.Image{
					Value: "https://steamykitchen.com/wp-content/uploads/2009/07/kogi-bbq-taco-151.jpg",
				},
				Ingredients: models.Ingredients{
					Values: []string{
						"1 pound cooked pulled pork or cooked shredded chicken",
						"12 corn or flour tortillas",
						"1/4 cup prepared store-bought Korean Kimchi ((optional))",
						"1 large English cucumber (or 2 Japanese cucumbers, sliced very thinly)",
						"2 tablespoons rice vinegar",
						"1/2 teaspoon sugar",
						"1/2 teaspoon finely minced fresh chili pepper (or more depending on your tastes)",
						"1 generous pinch of salt",
						"2 tablespoons Korean fermented hot pepper paste (gochujang)",
						"3 tablespoons sugar",
						"2 tablespoons soy sauce",
						"1 teaspoon rice wine vinegar",
						"2 teaspoons sesame oil",
					},
				},
				Instructions: models.Instructions{
					Values: []string{
						"Make the Quick Cucumber Pickle: Mix together all the Quick Pickle ingredients. You can make this a " +
							"few hours in advance and store in refrigerator. The longer it sits, the less “crunch” you’ll " +
							"have. I like making this cucumber pickle 1 hour prior, storing in refrigerator and serving it " +
							"cold on the tacos for texture and temperature contrast.",
						"Make the Koji BBQ Sauce: Whisk all BBQ sauce ingredients together until sugar has dissolved and mixture " +
							"is smooth. You can make this a few days in advance and store tightly covered in the refrigerator.",
						"Toss the Koji BBQ Sauce with your cooked pulled pork or shredded chicken. Warm the tortillas and serve" +
							" tacos with the Quick Cucumber Pickle.",
					},
				},
				Name: "Korean Style Tacos with Kogi BBQ Sauce Recipe",
				NutritionSchema: models.NutritionSchema{
					Calories:      "503 kcal",
					Carbohydrates: "48 g",
					Cholesterol:   "102 mg",
					Fat:           "20 g",
					Fiber:         "5 g",
					Protein:       "34 g",
					SaturatedFat:  "6 g",
					Servings:      "1",
					Sodium:        "722 mg",
					Sugar:         "11 g",
				},
				PrepTime: "PT60M",
				Yield:    models.Yield{Value: 4},
				URL:      "https://steamykitchen.com/4474-korean-style-tacos-with-kogi-bbq-sauce.html",
			},
		},
		{
			name: "streetkitchen.co",
			in:   "https://streetkitchen.co/recipe/thai-red-duck-curry/",
			want: models.RecipeSchema{
				AtContext: atContext,
				AtType:    models.SchemaType{Value: "Recipe"},
				Description: models.Description{
					Value: "This exquisite Thai Red Duck Curry is made with pineapple, cherry tomatoes and authentic red curry spices and coconut.",
				},
				Name:  "Thai Red Duck Curry",
				Yield: models.Yield{Value: 4},
				Image: models.Image{
					Value: "https://streetkitchen.co/wp-content/uploads/2022/10/Thai-Red-Duck-Curry-feature.jpg",
				},
				Ingredients: models.Ingredients{
					Values: []string{
						"2 tablespoons vegetable oil",
						"380g (approx. 2) fresh duck breast fillets, pat dry",
						"Salt to taste",
						"1 red onion, cut into thin wedges",
						"1 x 285g packet Street Kitchen Red Thai Curry Kit",
						"150g tinned pineapple slices in juice, drained, quartered",
						"1/2 x 250g punnet cherry tomatoes",
						"Thai basil, mint leaves and sliced red chilli for garnish",
						"Steamed jasmine rice and roti to serve",
					},
				},
				Instructions: models.Instructions{
					Values: []string{
						"Brush the oil over duck breasts and season with salt. Heat a non-stick frying pan over medium heat. When hot, add duck breasts skin-side down. Cook for 4-5 minutes or until the skin is golden and crisp. Turn and cook for 3-4 minutes or until just cooked through. Transfer to a plate. Cut into 1cm slices",
						"Discard excess oil, retaining 1 tbsp. Return pan to medium heat. Add spice pack and cook for 5 seconds. Add onion and cook for 3 minutes or until softened.\u00a0",
						"Stir in curry paste, coconut milk sachet, pineapple, tomatoes and 1/2 cup water. Stir until combined. Bring to the boil and simmer for 2 minutes. Nestle duck into sauce and simmer for 5 minutes or until duck is cooked through and sauce thickens. Spoon rice into serving bowls. Top with curry and garnish with basil and mint leaves and sliced red chilli. Serve with extra grilled roti.",
					},
				},
				URL: "https://streetkitchen.co/recipe/thai-red-duck-curry/",
			},
		},
		{
			name: "sunbasket.com",
			in:   "https://sunbasket.com/recipe/chicken-and-dumplings",
			want: models.RecipeSchema{
				AtContext: atContext,
				AtType:    models.SchemaType{Value: "Recipe"},
				Category:  models.Category{Value: ""},
				CookTime:  "PT35M",
				Description: models.Description{
					Value: "This is Sunbasket’s easy (and gluten-free!) spin on an American classic.",
				},
				Keywords: models.Keywords{Values: ""},
				Image: models.Image{
					Value: "https://cdn.sunbasket.com/c46a59d6-5745-4b86-9574-0e3e4ab4318b.jpg",
				},
				Ingredients: models.Ingredients{
					Values: []string{
						"1 cup milk",
						"1 teaspoon apple cider vinegar",
						"2 ounces organic cremini or other button mushrooms",
						"1 or 2 cloves organic peeled fresh garlic",
						"Sunbasket gluten-free dumpling mix (Cup4Cup gluten-free flour - sugar - kosher salt - baking powder - " +
							"baking soda)",
						"Chicken options:",
						"2 to 4 boneless skinless chicken thighs (about 10 ounces total)",
						"2 boneless skinless chicken breasts (about 6 ounces each)",
						"1 cup organic mirepoix (onions - carrots - celery)",
						"1 tablespoon tomato paste",
						"4 or 5 sprigs organic fresh flat-leaf parsley",
						"2 tablespoons Cup4Cup gluten-free flour",
						"1 cup chicken broth",
					},
				},
				Instructions: models.Instructions{
					Values: []string{
						"Prep the milk-vinegar mixture",
						"Prep the vegetables; make the dumpling dough",
						"Prep and brown the chicken",
						"Cook the vegetables",
						"Finish the chicken; cook the dumplings",
						"Serve",
					},
				},
				Name: "Chicken and dumplings",
				NutritionSchema: models.NutritionSchema{
					Calories:     "520",
					Cholesterol:  "135mg",
					Fat:          "21g",
					Fiber:        "2g",
					Protein:      "34g",
					SaturatedFat: "3.5g",
					Sodium:       "830mg",
					Sugar:        "7g",
				},
				Yield: models.Yield{Value: 2},
				URL:   "https://sunbasket.com/recipe/chicken-and-dumplings",
			},
		},
		{
			name: "sweetcsdesigns.com",
			in:   "https://sweetcsdesigns.com/roasted-tomato-marinara-sauce/",
			want: models.RecipeSchema{
				AtContext:     atContext,
				AtType:        models.SchemaType{Value: "Recipe"},
				Category:      models.Category{Value: "Pasta"},
				CookTime:      "PT45M",
				Cuisine:       models.Cuisine{Value: "italian"},
				DatePublished: "2022-03-09",
				Description: models.Description{
					Value: "Tired of using those bland jars of spaghetti sauce? This homemade roasted tomato marinara " +
						"sauce is packed with more flavor than those store-bought jars.",
				},
				Keywords: models.Keywords{
					Values: "pasta, sauce, tomato, italian, tomato sauce",
				},
				Image: models.Image{
					Value: "https://sweetcsdesigns.com/wp-content/uploads/2022/03/roasted-tomato-marinara-sauce-recipe-picture-720x720.jpg",
				},
				Ingredients: models.Ingredients{
					Values: []string{
						"2 pounds roma tomatoes",
						"1 onion, peeled and quartered",
						"2 tablespoons olive oil",
						"4 tablespoons balsamic vinegar, divided",
						"1 whole head garlic",
						"¼ cup fresh basil leaves, lightly packed",
						"Salt and Pepper",
					},
				},
				Instructions: models.Instructions{
					Values: []string{
						"Preheat your oven to 425 and line a large baking sheet with foil (it will help keep the tomatoes from " +
							"sticking to the pan).",
						"Add the tomatoes and onion to the pan and drizzle over 1.5 tablespoons olive oil along with the balsamic " +
							"vinegar. Sprinkle generously with salt and pepper and toss to combine.",
						"Slice the top off the head of garlic, making sure to leave it intact.",
						"Drizzle the remaining half tablespoon of olive oil over the garlic and wrap it tightly in foil.",
						"Place the garlic cut side up on the baking sheet.",
						"Roast for 35 minutes, or until the tomatoes are caramelized and the onions are golden brown.",
						"Remove the garlic from the foil and squeeze out the cloves-- be careful, it will be hot!",
						"Add the garlic cloves to a food processor, along with the roasted tomatoes, roasted onion, remaining two " +
							"tablespoons balsamic vinegar, and fresh basil.",
						"Puree to your desired consistency and season with salt and pepper to taste.",
						"Store the marinara sauce in an airtight container in the fridge for up to five days or in an airtight " +
							"container in the freezer for up to three months.",
					},
				},
				Name: "Roasted Tomato Marinara Sauce",
				NutritionSchema: models.NutritionSchema{
					Calories:       "68 calories",
					Carbohydrates:  "8 grams carbohydrates",
					Cholesterol:    "0 milligrams cholesterol",
					Fat:            "4 grams fat",
					Fiber:          "2 grams fiber",
					Protein:        "1 grams protein",
					SaturatedFat:   "1 grams saturated fat",
					Sodium:         "46 milligrams sodium",
					Sugar:          "5 grams sugar",
					TransFat:       "0 grams trans fat",
					UnsaturatedFat: "3 grams unsaturated fat",
				},
				PrepTime: "PT5M",
				Yield:    models.Yield{Value: 2},
				URL:      "https://sweetcsdesigns.com/roasted-tomato-marinara-sauce/",
			},
		},
		{
			name: "sweetpeasandsaffron.com",
			in:   "https://sweetpeasandsaffron.com/slow-cooker-cilantro-lime-chicken-tacos-freezer-slow-cooker/",
			want: models.RecipeSchema{
				AtContext:     atContext,
				AtType:        models.SchemaType{Value: "Recipe"},
				Category:      models.Category{Value: "Main Course"},
				CookTime:      "PT240M",
				Cuisine:       models.Cuisine{Value: "Mexican"},
				DatePublished: "2022-03-24T01:10:00+00:00",
				Description: models.Description{
					Value: "Cilantro lime chicken tacos are full of simple, bright flavors: cilantro, lime juice, garlic and a " +
						"touch of honey! No sauteeing required (just 15 min prep!), and easy to meal prep.",
				},
				Keywords: models.Keywords{
					Values: "cilantro lime chicken crockpot tacos, cilantro lime chicken tacos, crockpot tacos, meal prep tacos",
				},
				Image: models.Image{
					Value: "https://sweetpeasandsaffron.com/wp-content/uploads/2017/08/Slow-Cooker-Cilantro-Lime-Chicken-Tacos.jpg",
				},
				Ingredients: models.Ingredients{
					Values: []string{
						"4 chicken breasts (roughly 2 lbs; boneless skinless chicken thighs may also be used * see note 1)",
						"11.5 oz can of corn kernels (drained; 341 mL)",
						"15 oz can of black beans (drained &amp; rinsed; optional)",
						"1 red onion (sliced into strips)",
						"1/2 cup chicken stock",
						"2 cloves garlic",
						"1/4 teaspoon salt",
						"1/2 teaspoon cumin",
						"1/4 teaspoon ground coriander",
						"1 lime (zested)",
						"2 tablespoons honey (note 2)",
						"1/4 cup packed cilantro leaves",
						"Tortillas (2 small tortillas per person)",
						"shredded cabbage",
						"radish slices",
						"greek yogurt",
					},
				},
				Instructions: models.Instructions{
					Values: []string{
						"Combine - In the base of a 6-quart slow cooker, place the chicken breasts, corn and onion slices.",
						"Blend sauce - Using a stand or immersion blender, blend the sauce ingredients and pour over slow cooker " +
							"contents.",
						"Slow cook - Cover and cook on low for 4-5 hours, until chicken is cooked through.",
						"Serve - Shred chicken, then serve in tortillas topped with yogurt, shredded cabbage, and radish slices. " +
							"* see note 3",
					},
				},
				Name: "Cilantro Lime Chicken Crockpot Tacos",
				NutritionSchema: models.NutritionSchema{
					Calories:      "267 kcal",
					Carbohydrates: "28 g",
					Cholesterol:   "73 mg",
					Fat:           "4 g",
					Fiber:         "6 g",
					Protein:       "31 g",
					SaturatedFat:  "1 g",
					Servings:      "0.5 cup",
					Sodium:        "227 mg",
					Sugar:         "7 g",
				},
				PrepTime: "PT15M",
				Yield:    models.Yield{Value: 8},
				URL:      "https://sweetpeasandsaffron.com/slow-cooker-cilantro-lime-chicken-tacos-freezer-slow-cooker/",
			},
		},
		{
			name: "tasteofhome.com",
			in:   "https://www.tasteofhome.com/recipes/cast-iron-skillet-steak/",
			want: models.RecipeSchema{
				AtContext:     atContext,
				AtType:        models.SchemaType{Value: "Recipe"},
				Category:      models.Category{Value: "Dinner"},
				CookTime:      "PT5M",
				DateModified:  "2023-10-21",
				DatePublished: "2019-02-13",
				Description: models.Description{
					Value: `If you’ve never cooked steak at home before, it can be a little intimidating. That’s why I came up with this simple steak recipe that’s so easy, you could make it any day of the week. —<a href="https://www.tasteofhome.com/author/jschend/">James Schend</a>, <a href="https://www.dairyfreed.com/" target="_blank">Dairy Freed</a>`,
				},
				Keywords: models.Keywords{Values: ""},
				Image: models.Image{
					Value: "https://www.tasteofhome.com/wp-content/uploads/2019/02/Cast-Iron-Skillet-Steak_EXPS_CIMZ19_235746_B01_15_10b-6.jpg",
				},
				Ingredients: models.Ingredients{
					Values: []string{
						"1 beef New York strip or ribeye steak (1 pound), 1 inch thick",
						"3 teaspoons kosher salt, divided",
					},
				},
				Instructions: models.Instructions{
					Values: []string{
						"When making steak, you want to make sure it's well-seasoned. You don't need a lot of fancy flavors to make the meat taste amazing. In fact, we opt only for salt—just make sure that it's kosher. Salt with a smaller grain, such as table salt, breaks down faster and can give your steak a briny flavor. (This bamboo salt cellar is perfect for storing your kosher salt.) To season, start by removing the steak from the refrigerator and generously sprinkle two teaspoons of kosher salt on all sides of the filet. Let it stand for 45-60 minutes. This resting period gives the meat time to come up to room temperature, which helps the steak cook more evenly. It also gives the meat time to absorb some of the salt. If you've already mastered steak basics, try one of these steak rubs and marinades.",
						"The other key to a delicious steak is heat. And since that signature sear comes from a sizzling hot pan, a cast-iron skillet is the way to go. This hearty pan gets extremely hot and also retains heat for a long time, making it the perfect vessel for steak. You'll want to preheat your pan over high heat for 4-5 minutes, or until very hot. Then, pat your steak dry with paper towels and sprinkle the remaining teaspoon of salt in the bottom of the skillet. Now you're ready to sear! Love your skillet? You're going to be obsessed with all the amazing cast-iron accessories. We're big fans of these Lodge handle covers to protect your hands from the cast iron's super hot handle.",
						"Place the steak into the skillet and cook until it's easily moved. This takes between one and two minutes. Carefully flip the steak, placing it in a different section of the skillet. Cook for 30 seconds, and then begin moving the steak around, occasionally pressing slightly to ensure even contact with the skillet. A grill press is great for this. Moving the steak around the pan helps it cook faster and more evenly. Editor's Tip: This step will produce a lot of smoke, so make sure you're cooking in a well-ventilated space. It's also a good idea to turn your kitchen vent or fan on.",
						"Continue turning and flipping the steak until it's cooked to your desired degree of doneness. Let the steak rest for 10 minutes before cutting in. Have leftovers? Here's the right way to reheat steak and how to repurpose last night's dinner into one of these amazing leftover steak recipes.",
					},
				},
				Name: "Cast-Iron Skillet Steak",
				NutritionSchema: models.NutritionSchema{
					Calories:      "494 calories",
					Carbohydrates: "0 carbohydrate (0 sugars",
					Cholesterol:   "134mg cholesterol",
					Fat:           "36g fat (15g saturated fat)",
					Fiber:         "0 fiber)",
					Protein:       "40g protein. ",
					Sodium:        "2983mg sodium",
				},
				PrepTime: "PT5M",
				Yield:    models.Yield{Value: 2},
				URL:      "https://www.tasteofhome.com/recipes/cast-iron-skillet-steak/",
			},
		},
		{
			name: "tastesoflizzyt.com",
			in:   "https://www.tastesoflizzyt.com/easter-ham-pie/",
			want: models.RecipeSchema{
				AtContext:     atContext,
				AtType:        models.SchemaType{Value: "Recipe"},
				Category:      models.Category{Value: "Brunch"},
				CookTime:      "PT75M",
				Cuisine:       models.Cuisine{Value: "Italian"},
				DatePublished: "2022-04-04T04:28:00+00:00",
				Description: models.Description{
					Value: "A midwestern take on Italian Easter Pie, this Easter Ham Pie is perfect for Sunday brunch. " +
						"It&#039;s a great leftover ham recipe.",
				},
				Keywords: models.Keywords{
					Values: "breakfast and brunch, easter breakfast, leftover ham recipe",
				},
				Image: models.Image{
					Value: "https://www.tastesoflizzyt.com/wp-content/uploads/2022/03/Easter-Ham-Pie-15.jpg",
				},
				Ingredients: models.Ingredients{
					Values: []string{
						"double pie dough for 9-inch pie pan",
						"6 large eggs (beaten)",
						"2 cups diced ham (about ½ pound)",
						"15 ounces ricotta cheese",
						"2 cups mozzarella cheese",
						"½ cup Parmesan cheese",
						"1 teaspoon oregano",
						"½ teaspoon garlic powder",
						"Salt and pepper (for topping)",
					},
				},
				Instructions: models.Instructions{
					Values: []string{
						"Place one of the pie crusts in the bottom of a greased 9” pie plate. Flute the edges.",
						"Chill the pie crust for 30 minutes.",
						"Preheat the oven to 375ºF.",
						"Place parchment paper in the bottom of the pie crust, then fill with pie weights or dry beans.",
						"Blind bake the pie crust for 15 minutes or until the edges are starting to brown. Remove the crust " +
							"from the oven and allow it to cool while you prepare the filling.",
						"In a large bowl, whisk the eggs.",
						"Then add the ham, ricotta, mozzarella, parmesan, oregano and garlic powder. Mix well.",
						"Pour the mixture into the baked bottom crust.",
						"Top with the second crust and flute the edges to seal.",
						"Use a sharp knife to make 3 slits across the top of the pie crust.",
						"Sprinkle with sea salt and freshly ground pepper.",
						"Bake at 350ºF for one hour.",
						"Store any leftoveres in the refrigerator in an airtight container.",
					},
				},
				Name: "Easter Ham Pie",
				NutritionSchema: models.NutritionSchema{
					Calories:       "474 kcal",
					Carbohydrates:  "24 g",
					Cholesterol:    "192 mg",
					Fat:            "30 g",
					Fiber:          "1 g",
					Protein:        "26 g",
					SaturatedFat:   "14 g",
					Servings:       "1",
					Sodium:         "913 mg",
					Sugar:          "1 g",
					TransFat:       "1 g",
					UnsaturatedFat: "13 g",
				},
				PrepTime: "PT20M",
				Yield:    models.Yield{Value: 8},
				URL:      "https://www.tastesoflizzyt.com/easter-ham-pie/",
			},
		},
		{
			name: "tasty.co",
			in:   "https://tasty.co/recipe/honey-soy-glazed-salmon",
			want: models.RecipeSchema{
				AtContext:     atContext,
				AtType:        models.SchemaType{Value: "Recipe"},
				Category:      models.Category{Value: "Lunch"},
				Cuisine:       models.Cuisine{Value: "North American"},
				DateModified:  "2022-11-28T23:00:00",
				DatePublished: "2017-05-11T21:21:36",
				Description: models.Description{
					Value: "Two words: honey salmon! Sure, it takes a tiny bit of prep work, but once you marinate your " +
						"salmon, you won’t be able to go back. A simple mix of honey, soy sauce, garlic, and ginger coats " +
						"and flavors your fish for 30 minutes before you throw it on the pan until the outside is perfectly " +
						"crispy. Once that’s done, you heat up and reduce some extra marinade to make a thick, to-die-for " +
						"glaze to pour over your filet. Serve with your favorite veggies or rice and enjoy!",
				},
				Image: models.Image{
					Value: "https://img.buzzfeed.com/video-api-prod/assets/04ff8cfcc4b5428a8bcc6b03099d4492/Thumb_A_FB.jpg?" +
						"resize=1200:*",
				},
				Ingredients: models.Ingredients{
					Values: []string{
						"12 oz skinless salmon",
						"1 tablespoon olive oil",
						"4 cloves garlic, minced",
						"2 teaspoons ginger, minced",
						"½ teaspoon red pepper",
						"1 tablespoon olive oil",
						"⅓ cup less sodium soy sauce",
						"⅓ cup honey",
					},
				},
				Instructions: models.Instructions{
					Values: []string{
						"Place salmon in a sealable bag or medium bowl.",
						"In a small bowl or measuring cup, mix marinade ingredients.",
						"Pour half of the marinade on the salmon. Save the other half for later.",
						"Let the salmon marinate in the refrigerator for at least 30 minutes.",
						"In a medium pan, heat oil. Add salmon to the pan, but discard the used marinade. Cook salmon on one " +
							"side for about 2-3 minutes, then flip over and cook for an additional 1-2 minutes.",
						"Remove salmon from pan. Pour in remaining marinade and reduce.",
						"Serve the salmon with sauce and a side of veggies. We used broccoli.",
						"Enjoy!",
					},
				},
				Name: "Honey Soy-Glazed Salmon Recipe by Tasty",
				NutritionSchema: models.NutritionSchema{
					Calories:      "705 calories",
					Carbohydrates: "60 grams",
					Fat:           "35 grams",
					Fiber:         "0 grams",
					Protein:       "37 grams",
					Sugar:         "57 grams",
				},
				Yield: models.Yield{Value: 2},
				URL:   "https://tasty.co/recipe/honey-soy-glazed-salmon",
			},
		},
		{
			name: "tastykitchen.com",
			in:   "https://tastykitchen.com/recipes/main-courses/garlic-shrimp-scampi-with-angel-hair-pasta/",
			want: models.RecipeSchema{
				AtContext: atContext,
				AtType:    models.SchemaType{Value: "Recipe"},
				Name:      "Garlic Shrimp Scampi with Angel Hair Pasta",
				Category:  models.Category{Value: "Main Courses"},
				Description: models.Description{
					Value: "This shrimp scampi with angel hair pasta is a delicious dinner with plenty of cheese that you can " +
						"make in just 15 minutes! ",
				},
				Image: models.Image{
					Value: "https://tastykitchen.com/recipes/wp-content/uploads/sites/2/2020/09/SHRIMP-SCAMPI-WITH-ANGEL-HAIR-" +
						"PASTA-15-410x308.jpg",
				},
				PrepTime: "PT5M",
				CookTime: "PT10M",
				Ingredients: models.Ingredients{
					Values: []string{
						"3 Tablespoons Butter",
						"4 cloves Garlic, Minced",
						"1 pound Large Shrimp, Peeled & Deveined",
						"½ teaspoons Red Pepper Flakes (or To Taste)",
						"1 pinch Sea Salt",
						"1 piece Lemon, Zested",
						"⅓ cups Freshly Squeezed Lemon Juice",
						"6 cups Vegetable Or Fish Stock",
						"12 ounces, weight Angel Hair Pasta",
						"⅔ cups Freshly Grated Parmesan Cheese",
						"⅓ cups Finely Chopped Fresh Parsley",
						"Salt And Pepper, to taste",
					},
				},
				Instructions: models.Instructions{
					Values: []string{
						"Melt the butter in a large skillet over medium heat. Add the garlic and cook until fragrant, about " +
							"1 minute, stirring constantly. Add in the shrimp and sprinkle with chili flakes and sea salt.",
						"Cook the shrimp for about 1 to 2 minutes per side, or until pink and opaque. Set aside on a plate.",
						"To the same skillet, add the lemon zest and juice. Pour in 5 cups of broth. Keep the remaining 1 cup " +
							"on hand, in case you need it for later. Stir to combine then bring the liquid to a gentle simmer.",
						"Add the pasta, and using a pair of tongs, stir occasionally for a few minutes, or until the pasta starts " +
							"to soften and bend. Fully immerse the pasta into the liquid, and keep stirring frequently to avoid " +
							"sticking to the bottom of the pan.",
						"Cook the pasta according to the timing on the package directions. Angel hair pasta needs no more than 3 " +
							"minutes to cook to al dente.",
						"Once the pasta is cooked to al dente, return the shrimp to the pan and turn the heat off. Stir in the " +
							"Parmesan cheese and chopped parsley, and mix until the cheese has melted into the sauce. The " +
							"pasta will continue to absorb the sauce if not served immediately. So quickly season with salt " +
							"and pepper to your taste, and dig in!",
					},
				},
				Yield: models.Yield{Value: 4},
				URL:   "https://tastykitchen.com/recipes/main-courses/garlic-shrimp-scampi-with-angel-hair-pasta/",
			},
		},
		{
			name: "tesco.com",
			in:   "https://realfood.tesco.com/recipes/salted-honey-and-rosemary-lamb-with-roasties-and-rainbow-carrots.html",
			want: models.RecipeSchema{
				AtContext:     atContext,
				AtType:        models.SchemaType{Value: "Recipe"},
				Name:          "Salted honey and rosemary lamb with roasties and rainbow carrots",
				Category:      models.Category{Value: "Dinner"},
				CookTime:      "PT2H0M",
				Cuisine:       models.Cuisine{Value: "British"},
				DateModified:  "06/04/2022 10:26:45",
				DatePublished: "23/03/2022 14:55:51",
				Description: models.Description{
					Value: "This one-tray wonder is a great option for Sunday roast, Mother's Day or even Easter",
				},
				Keywords: models.Keywords{
					Values: "honey, lamb, easter, mother's Day, sunday roast, sunday lunch, roast dinner, veg, meat and two veg, potatoes, roast potatoes, roasted potatoes",
				},
				Image: models.Image{
					Value: "https://realfood.tesco.com/media/images/1400x919-OneTrayEasterRoast-7f95aa82-f903-4339-ab87-" +
						"6cffda6d26d8-0-1400x919.jpg",
				},
				Yield: models.Yield{Value: 6},
				NutritionSchema: models.NutritionSchema{
					Calories:      "855 calories",
					Carbohydrates: "64.2 grams carbohydrate",
					Cholesterol:   "",
					Fat:           "44 grams fat",
					Fiber:         "10.4 grams fibre",
					Protein:       "51.2 grams protein",
					SaturatedFat:  "16 grams saturated fat",
					Sugar:         "21 grams sugar",
				},
				Ingredients: models.Ingredients{
					Values: []string{
						"2kg lamb leg joint", "¼ tsp flaky sea salt, plus extra to season",
						"20g pack fresh rosemary, most leaves finely chopped, some whole sprigs",
						"12 garlic cloves, crushed", "7 tbsp olive oil",
						"1.5kg King Edward potatoes, peeled and cut into 4-5cm chunks",
						"2 x 450g packs Tesco Finest rainbow or Imperator carrots, scrubbed, halved lengthways if large",
						"50g clear honey, to taste",
						"20g fresh mint, leaves picked",
						"1 tbsp granulated sugar",
						"3 tbsp white wine vinegar",
					},
				},
				Instructions: models.Instructions{
					Values: []string{
						"Preheat the oven to gas 3, 160 ̊C, fan 140 ̊C. Use a small sharp knife to pierce holes about 4cm deep all " +
							"over the lamb, then season well with sea salt. In a small bowl, combine the chopped rosemary and " +
							"garlic with 6 tbsp olive oil.",
						"Rinse the chopped potatoes and tip into one side of a large, rimmed baking tray. Put the carrots in the " +
							"other half of the tray. Toss the carrots with 1 tbsp oil, then toss the potatoes in half the rosemary- " +
							"garlic oil and spread out in a single layer over half of the tray; add the whole rosemary sprigs. Put " +
							"the lamb in the centre of the tray and rub with the remaining rosemary-garlic oil.",
						"Roast for 1 hr 45 mins. Mix the honey with &frac14; tsp flaky sea salt and dab all over the lamb with a pastry " +
							"brush. Return to the oven, increase the temperature to gas 7, 220 ̊C, fan 200 ̊C, then roast for a " +
							"final 15 mins. Transfer the lamb to a warmed plate, then set aside to rest, covered loosely with " +
							"foil, for at least 20 mins. Keep the veg warm while the lamb rests.",
						"Meanwhile, make the mint sauce. Finely chop the mint leaves with the sugar (this helps to stop oxidisation), " +
							"then mix well in a bowl with 3 tbsp boiling water and the vinegar. Serve alongside the roast lamb, " +
							"potatoes and carrots.",
					},
				},
				PrepTime: "PT1H0M",
				URL:      "https://realfood.tesco.com/recipes/salted-honey-and-rosemary-lamb-with-roasties-and-rainbow-carrots.html",
			},
		},
		{
			name: "theclevercarrot.com",
			in:   "https://www.theclevercarrot.com/2021/10/homemade-sourdough-breadcrumbs/",
			want: models.RecipeSchema{
				AtContext:     atContext,
				AtType:        models.SchemaType{Value: "Recipe"},
				Category:      models.Category{Value: "Sourdough Bread"},
				CookTime:      "PT25M",
				CookingMethod: models.CookingMethod{Value: "Oven-Baked"},
				Cuisine:       models.Cuisine{Value: "American"},
				DatePublished: "2021-10-03",
				Description: models.Description{
					Value: "Transform leftover bread into delicious homemade breadcrumbs, with just 10 minutes hands on time and " +
						"minimal effort. I like using sourdough bread, but any bread will do. Breadcrumbs can be stored " +
						"in the freezer for up to 3-6 months.",
				},
				Keywords: models.Keywords{
					Values: "homemade, sourdough bread, breadcrumbs, Italian, seasoned bread crumbs",
				},
				Image: models.Image{
					Value: "https://www.theclevercarrot.com/wp-content/uploads/2021/09/SOURDOUGH-BREADCRUMBS-2-225x225.jpg",
				},
				Ingredients: models.Ingredients{
					Values: []string{
						"1 tsp dried garlic powder",
						"1 tsp dried onion powder",
						"1 tsp fine sea salt",
						"2 tsp dried oregano",
						"1 tbsp dried parsley",
						"1/4 cup (30 g) ground Pecorino Romano or Parmesan cheese",
					},
				},
				Instructions: models.Instructions{
					Values: []string{
						"Preheat your oven to 300˚ F (150˚ C). Grab (2x) rimmed baking sheets.",
						"Cut the bread into small cubes, including the crust, about 1-inch in size.",
						"Add the cubes to a food processor or high-powdered blender (you will need to work in batches). Process " +
							"until fine crumbs form. Note: if you cannot get the crumbs small enough at this stage, you " +
							"can process them again after baking.",
						"Divide the crumbs over the (2x) sheet pans in one even layer.",
						"Bake for 15-30 minutes, stirring once, or until the crumbs are crispy. Bake time will vary depending on " +
							"bread type and freshness.",
						"Remove the breadcrumbs from the oven. Allow to cool.",
						"At this point, you can either keep the breadcrumbs plain or season, Italian-style (my preference). Add the " +
							"dried garlic, onion, salt, oregano, parsley and cheese; mix well.",
						"Portion the cooled breadcrumbs into containers or a zip-top bag and freeze until ready to use.",
					},
				},
				Name:     "Homemade Sourdough Breadcrumbs",
				PrepTime: "PT10M",
				Yield:    models.Yield{Value: 4},
				URL:      "https://www.theclevercarrot.com/2021/10/homemade-sourdough-breadcrumbs/",
			},
		},
		{
			name: "thehappyfoodie.co.uk",
			in:   "https://thehappyfoodie.co.uk/recipes/leek-and-lentil-gratin/",
			want: models.RecipeSchema{
				AtContext:    atContext,
				AtType:       models.SchemaType{Value: "Recipe"},
				DateModified: "2022-02-07T16:00:36+00:00",
				Name:         "Leek and Puy Lentil Gratin with a Crunchy Feta Topping",
				Yield:        models.Yield{Value: 4},
				Description: models.Description{
					Value: "A creamy, cheesy, and deeply comforting dish, Rukmini Iyer's one-tin vegetarian gratin also happens " +
						"to be packed with nutritious Puy lentils.",
				},
				Image: models.Image{
					Value: "https://thehappyfoodie.co.uk/wp-content/uploads/2022/02/Rukmini-Iyer-Leek-and-Lentil-Gratin-scaled.jpg",
				},
				PrepTime: "PT10M",
				CookTime: "PT40M",
				Keywords: models.Keywords{
					Values: "Vegetarian, Feta, Leek, Lentil, One Pot, One-pot, Dinner, Main Course, Easy, Quick",
				},
				Ingredients: models.Ingredients{
					Values: []string{
						"30g butter", "3 cloves of garlic, crushed", "500g leeks, thinly sliced",
						"2 tsp sea salt", "freshly ground black pepper",
						"500g vac-packed cooked Puy lentils", "450ml crème fraîche",
						"125g feta cheese, crumbled", "50g panko or fresh white breadcrumbs",
						"1 tbsp olive oil",
					},
				},
				Instructions: models.Instructions{
					Values: []string{
						"Preheat the oven to 180°C fan/200°C/gas 6. Put the butter and garlic into a roasting tin and pop into " +
							"the oven to melt while you get on with slicing the leeks.",
						"Mix the sliced leeks with the melted garlic butter, season well with the sea salt and black pepper, then " +
							"return to the oven to roast for 20 minutes.",
						"After 20 minutes, stir through the Puy lentils, crème fraîche and another good scatter of sea salt, then " +
							"top with the feta cheese and breadcrumbs. Drizzle with the olive oil, then return to the oven " +
							"for a further 20–25 minutes, until golden brown on top.",
						"Serve the gratin hot, with a mustard or balsamic dressed green salad alongside.",
					},
				},
				URL: "https://thehappyfoodie.co.uk/recipes/leek-and-lentil-gratin/",
			},
		},
		{
			name: "thekitchenmagpie.com",
			in:   "https://www.thekitchenmagpie.com/blt-pasta-salad/",
			want: models.RecipeSchema{
				AtContext:     atContext,
				AtType:        models.SchemaType{Value: "Recipe"},
				Category:      models.Category{Value: "Salad"},
				Cuisine:       models.Cuisine{Value: "American"},
				DatePublished: "2021-06-16T04:00:00+00:00",
				Description: models.Description{
					Value: "Fantastic BLT pasta salad with bacon, lettuce, tomatoes, two types of cheese and tangy Ranch " +
						"dressing! The perfect side dish for BBQ&#039;s, or eat it as a meal!",
				},
				Keywords: models.Keywords{Values: "BLT pasta salad"},
				Image: models.Image{
					Value: "https://www.thekitchenmagpie.com/wp-content/uploads/images/2021/05/BLTpastasalad.jpg",
				},
				Ingredients: models.Ingredients{
					Values: []string{
						"one 454 gram box Farfalle pasta (bow-tie pasta cooked until al dente)",
						"1 pound bacon diced and cooked",
						"3 cups chopped romaine lettuce",
						"2 cups red cherry tomatoes (halved)",
						"1 cup orange or yellow cherry tomatoes (halved)",
						"1 cup cheddar cheese ( cut into small cubes then measure)",
						"1 cup pepper Jack cheese ( cut into small cubes then measure)",
						"1/4 cup red onion (sliced into very thin almost transparent rings)",
						"2 avocados  (pitted, peeled then diced)",
						"1 1/2 - 2 cups ranch dressing",
						"2 tablespoons fresh parsley chopped",
					},
				},
				Instructions: models.Instructions{
					Values: []string{
						"In a large serving bowl place the cooked pasta, then pour one cup of the dressing on top. Stir " +
							"to coat the pasta.",
						"Add in the cooked bacon, romaine lettuce, two types of tomatoes, the two types of cheese, the red " +
							"onion, and the avocados.",
						"Pour the remaining half a cup of dressing over, then gently mix until the other ingredients are " +
							"slightly coated. Taste test, and add more dressing if desired.",
						"Garnish with the fresh parsley and serve.",
					},
				},
				Name: "BLT Pasta Salad",
				NutritionSchema: models.NutritionSchema{
					Calories:      "653 kcal",
					Carbohydrates: "10 g",
					Cholesterol:   "78 mg",
					Fat:           "60 g",
					Fiber:         "3 g",
					Protein:       "20 g",
					SaturatedFat:  "15 g",
					Servings:      "1",
					Sodium:        "1533 mg",
					Sugar:         "3 g",
					TransFat:      "1 g",
				},
				PrepTime: "PT30M",
				Yield:    models.Yield{Value: 12},
				URL:      "https://www.thekitchenmagpie.com/blt-pasta-salad/",
			},
		},
		{
			name: "thenutritiouskitchen.co",
			in:   "http://thenutritiouskitchen.co/fluffy-paleo-blueberry-pancakes/",
			want: models.RecipeSchema{
				AtContext:     atContext,
				AtType:        models.SchemaType{Value: "Recipe"},
				Category:      models.Category{Value: "breakfast"},
				CookTime:      "PT15M",
				CookingMethod: models.CookingMethod{Value: "stove top"},
				Cuisine:       models.Cuisine{Value: "pancakes"},
				DatePublished: "2022-04-08",
				Description: models.Description{
					Value: "Delicious paleo blueberry pancakes made extra fluffy with simple, healthy ingredients all made in " +
						"one bowl! The perfect weekend breakfast or brunch recipe packed with refreshing blueberry flavors, " +
						"no added sugars and 100% dairy-free + gluten-free!",
				},
				Keywords: models.Keywords{
					Values: "paleo, pancakes, gluten-free, dairy-free",
				},
				Image: models.Image{
					Value: "https://thenutritiouskitchen.co/wp-content/uploads/2022/04/paleoblueberrypancakes-225x225.jpg",
				},
				Ingredients: models.Ingredients{
					Values: []string{
						"3/4 cup unsweetened apple sauce",
						"1/4 cup cashew butter",
						"1 large egg",
						"1 cup Bobs Red Mill super fine almond flour",
						"1 cup tapioca flour, or arrowroot flour",
						"2 teaspoons baking powder",
						"Sea salt + cinnamon to taste",
						"1/3 cup fresh blueberries",
						"Pure maple syrup",
						"Coconut cream (optional)",
						"Vegan butter or ghee",
					},
				},
				Instructions: models.Instructions{
					Values: []string{
						"In a large bowl whisk the apple sauce, cashew butter and egg until fully combined.",
						"Then add in the almond flour, tapioca flour, baking powder, sea salt, and cinnamon. Combine with a mixing " +
							"spoon until a thick batter forms. It will be thick but still moist!",
						"Heat ghee or butter in a pan over medium heat. Once the pan is hot, scoop the batter using a cookie scoop " +
							"or about 1/4 cup full of batter. I like to cook 3 at a time.",
						"Place some blueberries onto the pancakes while on pan and cook for about 3 minutes. Flip gently, and " +
							"continue cooking for about 2 minutes over medium-low heat.",
						"Repeat with remaining batter then serve with maple syrup, extra butter and toppings of choice. I love " +
							"coconut cream or cashew butter!",
					},
				},
				Name:     "Fluffy Paleo Blueberry Pancakes",
				PrepTime: "PT5M",
				Yield:    models.Yield{Value: 6},
				URL:      "https://thenutritiouskitchen.co/fluffy-paleo-blueberry-pancakes/",
			},
		},
		{
			name: "thepioneerwoman.com",
			in:   "https://www.thepioneerwoman.com/food-cooking/recipes/a8865/eggs-benedict/",
			want: models.RecipeSchema{
				AtContext:     atContext,
				AtType:        models.SchemaType{Value: "Recipe"},
				Category:      models.Category{Value: "brunch"},
				CookTime:      "PT10M",
				Cuisine:       models.Cuisine{Value: "American"},
				DateModified:  "2023-03-22T14:48:00Z",
				DatePublished: "2007-10-12T09:33:50Z",
				Description: models.Description{
					Value: "Ree Drummond shares her secrets to making perfect eggs Benedict. From flawless poached eggs to velvety Hollandaise sauce, it's the best brunch recipe.",
				},
				Keywords: models.Keywords{Values: "Recipes, Cooking, Food"},
				Image: models.Image{
					Value: "https://hips.hearstapps.com/thepioneerwoman/wp-content/uploads/2007/10/1546875357_506daa8f1c.jpg?crop=0.664xw:1xh;center,top&resize=1200:*",
				},
				Ingredients: models.Ingredients{
					Values: []string{
						"3 whole English muffins",
						"6 slices Canadian bacon",
						"2 sticks butter, plus more for the muffins",
						"6 whole eggs (plus 3 egg yolks)",
						"1 whole lemon, juiced",
						"Cayenne pepper, to taste",
						"Paprika, to garnish",
						"Chopped chives, to garnish",
					},
				},
				Instructions: models.Instructions{
					Values: []string{
						"Bring a pot of water to a boil. While the water's boiling, place the English muffin halves and an equal number of Canadian bacon slices on a cookie sheet. Lightly butter the English muffins and place the sheet under the broiler for just a few minutes, or until the English muffins are very lightly golden. Be careful not to dry out the Canadian bacon.",
						"Now if you do not have an egg poacher you can poach your eggs by doing the following: With a spoon, begin stirring the boiling water in a large, circular motion. When the tornado's really twisting, crack in an egg. The reason for the swirling is so the egg will wrap around itself as it cooks, keeping it together. Cook for about 2 1/2 to 3 minutes. Repeat with the remaining eggs.",
						"Melt 2 sticks of butter in a small saucepan until sizzling, but don't let it burn! Separate three eggs and place the yolks into a blender. Turn the blender on low to allow the yolks to combine, then begin pouring the very hot butter in a thin stream into the blender. The blender should remain on the whole time, and you should be careful to pour in the butter very slowly. Keep pouring butter until it’s all gone, then immediately begin squeezing lemon juice into the blender. If you are going to add cayenne pepper, this is the point at which you would do that.",
						"Place the English muffins on the plate, face up. Next, place a slice of Canadian bacon on each half. Place an egg on top of the bacon and then top with a generous helping of Hollandaise sauce. <em>Vegetarian variation:</em> you can omit the Canadian bacon altogether, or you can wilt fresh spinach and place it on the muffins for Eggs Florentine, which is divine in its own right. Top with more cayenne, or a sprinkle of paprika, and chopped chives if you like.",
					},
				},
				Name:     "Eggs Benedict",
				PrepTime: "PT15M",
				Yield:    models.Yield{Value: 3},
				URL:      "https://www.thepioneerwoman.com/food-cooking/recipes/a8865/eggs-benedict/",
			},
		},
		{
			name: "thespruceeats.com",
			in:   "https://www.thespruceeats.com/pasta-with-anchovies-and-breadcrumbs-recipe-5215384",
			want: models.RecipeSchema{
				AtContext:     atContext,
				AtType:        models.SchemaType{Value: "Recipe"},
				Category:      models.Category{Value: "Dinner"},
				CookTime:      "PT28M",
				Cuisine:       models.Cuisine{Value: "Italian"},
				DateModified:  "2023-02-20T21:36:14.876-05:00",
				DatePublished: "2022-01-12T16:19:41.204-05:00",
				Description: models.Description{
					Value: "Anchovies and pepper provide complex flavor in this simple Sicilian-style pasta dish. Crisp garlicky " +
						"breadcrumbs are the perfect finishing touch.",
				},
				Keywords: models.Keywords{Values: "anchovie pasta"},
				Image: models.Image{
					Value: "https://www.thespruceeats.com/thmb/_al2WJ0fr7Kt_LFIVpq-jQtimk0=/1500x0/filters:no_upscale():max_bytes(150000):strip_icc()/pasta-with-anchovies-and-breadcrumbs-recipe-5215384-Hero_01-a52a47010bd04ead814b972d518738ef.jpg",
				},
				Ingredients: models.Ingredients{
					Values: []string{
						"For the Breadcrumbs:",
						"4 slices day-old bread, about 7 ounces",
						"3 tablespoons olive oil",
						"5 cloves garlic, minced",
						"1/2 teaspoon kosher salt, or to taste",
						"1 teaspoon lemon zest, optional",
						"For the Pasta:",
						"12 ounces dry spaghetti or linguine",
						"1 1/2 tablespoons kosher salt for the cooking water, plus more, to taste",
						"3 tablespoons olive oil",
						"2 cloves garlic, cut in half lengthwise",
						"3 oil-packed anchovy fillets",
						"1/2 teaspoon crushed red pepper flakes, or more, to taste",
						"1 tablespoon lemon juice, optional",
						"1/4 teaspoon freshly ground black pepper, or to taste",
						"For Serving:",
						"3 tablespoons chopped Italian flat-leaf parsley",
						"1/4 cup freshly grated Parmesan cheese, optional",
					},
				},
				Instructions: models.Instructions{
					Values: []string{
						"Gather the ingredients.",
						"Bring a large pot of well salted water to a boil. While the water is heating, remove the crusts from the " +
							"bread, if you like, and tear or cut it into cubes.",
						"Put bread cubes, garlic, and 1/4 teaspoon of kosher salt in a food processor and pulse to make coarse " +
							"breadcrumbs. You should have about 2 cups of crumbs, a bit more if the crusts are not removed.",
						"Heat the 3 tablespoons of the oil in a large Dutch oven or other heavy-duty pot over medium heat. When " +
							"the oil shimmers, add the breadcrumbs. Cook, stirring frequently, until the crumbs are lightly brown " +
							"and crisp, about 5 to 12 minutes, depending on the moisture in the bread.",
						"Transfer the breadcrumbs to a bowl, toss with lemon zest, if using, and set aside.",
						"Meanwhile, cook the pasta according to al dente according to package instructions, reserving 1/2 cup of " +
							"the pasta water. Drain the pasta and set aside.",
						"Wipe out the pot used for the breadcrumbs. Add 3 tablespoons of the oil over medium heat until it shimmers. " +
							"Add the halved garlic cloves and cook until lightly brown, about 2 minutes.",
						"Remove and discard the garlic pieces. Add the anchovies and crushed red pepper to the garlic-flavored oil " +
							"and cook for 1 minute longer. Add the lemon juice, if using.",
						"Add the pasta to the pot and toss, cooking, until warmed through. Add some cooking water to loosen the " +
							"mixture as needed. Taste and adjust the seasonings with salt and more crushed red pepper flakes, if desired.",
						"Plate the pasta in wide, shallow pasta bowls and top with a generous amount of garlic breadcrumbs. Garnish " +
							"with chopped parsley and Parmesan cheese, if using. Serve with lemon wedges, if desired.",
					},
				},
				Name: "Pasta With Anchovies and Breadcrumbs Recipe",
				NutritionSchema: models.NutritionSchema{
					Calories:       "642 kcal",
					Carbohydrates:  "90 g",
					Cholesterol:    "3 mg",
					Fat:            "24 g",
					Fiber:          "4 g",
					Protein:        "17 g",
					SaturatedFat:   "3 g",
					Sodium:         "518 mg",
					Sugar:          "5 g",
					UnsaturatedFat: "0 g",
				},
				PrepTime: "PT15M",
				Yield:    models.Yield{Value: 4},
				URL:      "https://www.thespruceeats.com/pasta-with-anchovies-and-breadcrumbs-recipe-5215384",
			},
		},
		{
			name: "thevintagemixer.com",
			in:   "https://www.thevintagemixer.com/roasted-asparagus-grilled-cheese/",
			want: models.RecipeSchema{
				AtContext:     atContext,
				AtType:        models.SchemaType{Value: "Recipe"},
				Category:      models.Category{Value: "Main Course"},
				CookTime:      "PT5M",
				Cuisine:       models.Cuisine{Value: "American"},
				DatePublished: "2017-04-03T04:00:45+00:00",
				Description: models.Description{
					Value: "A seasonally fresh take on grilled cheese with asparagus.",
				},
				Image: models.Image{
					Value: "http://d6h7vs5ykbiug.cloudfront.net/wp-content/uploads/2017/04/asparagus-grilled-cheese-9.jpg",
				},
				Ingredients: models.Ingredients{
					Values: []string{
						"6 spears of asparagus",
						"1 teaspoon olive oil",
						"sea salt and freshly ground pepper",
						"2 ounces of white cheddar cheese*",
						"2 ounces of gruyere cheese*",
						"4 slices of sourdough bread",
						"grass fed butter",
					},
				},
				Instructions: models.Instructions{
					Values: []string{
						"Preheat oven to 400 degrees and line a baking sheet with foil.",
						"Remove the woody ends of the asparagus and toss the spears with olive oil, then sprinkle with the sea salt " +
							"and pepper. Spread the asparagus spears out onto the prepared baking sheet and roast for 8-10 " +
							"minutes or until slightly brown.",
						"Meanwhile, slice the cheese and butter the bread. Remove asparagus from oven.",
						"Heat up a large skillet over medium heat. Add two slices of the bread, butter side down, to the pan. " +
							"Add cheese then add 3 spears of asparagus to each sandwich. Top with the other bread slices, butter side up.",
						"Toast for 3-4 minutes then flip and toast for 2 minutes. Remove from heat and serve hot with tomato soup!",
					},
				},
				Name:     "Asparagus Grilled Cheese",
				PrepTime: "PT5M",
				Yield:    models.Yield{Value: 2},
				URL:      "https://www.thevintagemixer.com/roasted-asparagus-grilled-cheese/",
			},
		},
		{
			name: "thewoksoflife.com",
			in:   "https://thewoksoflife.com/fried-wontons/",
			want: models.RecipeSchema{
				AtContext:     atContext,
				AtType:        models.SchemaType{Value: "Recipe"},
				Category:      models.Category{Value: "Appetizers and Snacks"},
				CookTime:      "PT20M",
				Cuisine:       models.Cuisine{Value: "Chinese"},
				DatePublished: "2015-09-05T13:38:44+00:00",
				Description: models.Description{
					Value: "Fried wontons are a easy-to-make crispy, crunchy, delicious appetizer. Your guests will be talking " +
						"about these fried wontons long after the party's over!",
				},
				Keywords: models.Keywords{Values: "fried wontons"},
				Image: models.Image{
					Value: "https://thewoksoflife.com/wp-content/uploads/2015/09/fried-wontons-6-1.jpg",
				},
				Ingredients: models.Ingredients{
					Values: []string{
						"12 oz. ground pork ((340g))",
						"2 tablespoons finely chopped scallions",
						"1 teaspoon sesame oil",
						"1 tablespoon soy sauce",
						"1 tablespoon shaoxing wine ((or dry sherry))",
						"1/2 teaspoon sugar",
						"1 tablespoon peanut or canola oil",
						"2 tablespoons water ((plus more for sealing the wontons))",
						"1/8 teaspoon white pepper",
						"40-50 Wonton skins ((1 pack, medium thickness))",
						"2 tablespoons apricot jam",
						"1/2 teaspoon soy sauce",
						"1/2 teaspoon rice wine vinegar",
						"2 tablespoons honey",
						"2 tablespoons Sriracha",
						"1 ½ tablespoons light soy sauce",
						"1 tablespoon sugar ((dissolved in 1 tablespoon hot water))",
						"1 teaspoon Worcestershire sauce",
						"1/2 teaspoon toasted sesame seeds ((optional))",
					},
				},
				Instructions: models.Instructions{
					Values: []string{
						"Start by making the filling. Simply combine the ground pork, chopped scallions, sesame oil, soy sauce, " +
							"wine (or sherry), sugar, oil, water, and white pepper in a bowl. Whip everything together by hand " +
							"for 5 minutes or in a food processor for 1 minute. You want the pork to look a bit like a paste.",
						"To make the wontons, take a wrapper, and add about a teaspoon of filling. Overstuffed wontons will pop " +
							"open during the cooking process and make a mess. Use your finger to coat the edges with a little " +
							"water (this helps the two sides seal together).",
						"For shape #1:",
						"Fold the wrapper in half into a rectangle, and press the two sides together so you get a firm seal. Hold " +
							"the bottom two corners of the little rectangle you just made, and bring the two corners together, " +
							"pressing firmly to seal. (Use a little water to make sure it sticks.)",
						"Shape #2:",
						"Fold the wonton in half so you have a triangle shape. Bring together the two outer corners, and press to " +
							"seal (you can use a little water to make sure it sticks).",
						"Keep assembling until all the filling is gone (this recipe should make between 40 and 50 wontons). Place " +
							"the wontons on a baking sheet or plate lined with parchment paper to prevent sticking.",
						"At this point, you can cover the wontons with plastic wrap, put the baking sheet/plate into the freezer, " +
							"and transfer them to Ziploc bags once they’re frozen. They’ll keep for a couple months in the " +
							"freezer and be ready for the fryer whenever you’re ready.",
						"To conserve oil, use a small pot to fry the wontons. Fill it with 2 to 3 inches of oil, making sure the " +
							"pot is deep enough so the oil does not overflow when adding the wontons. Heat the oil to 350 degrees, " +
							"and fry in small batches, turning the wontons occasionally until they are golden brown.",
						"If you have a small spider strainer or slotted spoon, you can use it to keep the wontons submerged when " +
							"frying. This method will give you the most uniform golden brown look without the fuss of turning them. " +
							"Remove the fried wontons to a sheet pan lined with paper towels or a metal cooling rack to drain.",
						"To make one or all of the sauces, simply mix the respective ingredients in a small bowl, and you’re ready " +
							"to eat!",
					},
				},
				Name: "Fried Wontons",
				NutritionSchema: models.NutritionSchema{
					Calories:      "164 kcal",
					Carbohydrates: "15 g",
					Cholesterol:   "23 mg",
					Fat:           "8 g",
					Fiber:         "1 g",
					Protein:       "7 g",
					SaturatedFat:  "2 g",
					Servings:      "1",
					Sodium:        "243 mg",
					Sugar:         "1 g",
				},
				PrepTime: "PT90M",
				Yield:    models.Yield{Value: 12},
				URL:      "https://thewoksoflife.com/fried-wontons/",
			},
		},
		{
			name: "timesofindia.com",
			in:   "https://recipes.timesofindia.com/recipes/beetroot-cold-soup/rs90713582.cms",
			want: models.RecipeSchema{
				AtContext:     atContext,
				AtType:        models.SchemaType{Value: "Recipe"},
				Category:      models.Category{Value: "Appetizers"},
				Cuisine:       models.Cuisine{Value: "Vegetarian"},
				DatePublished: "2022-04-07T22:43:02+05:30",
				Description: models.Description{
					Value: "Yearning for a satiating and delicious meal, then try this easy and tasty cold soup made with " +
						"beetroot, curd, hard boiled eggs, coriander leaves and spices. To make this simple soup, just " +
						"follow us through some simple steps and make a sumptuous and enjoy it cold.",
				},
				Keywords: models.Keywords{
					Values: "Beetroot Cold Soup recipe, Vegetarian, cook Beetroot Cold Soup",
				},
				Image: models.Image{
					Value: "https://static.toiimg.com/thumb/90713582.cms?width=1200&height=900",
				},
				Ingredients: models.Ingredients{
					Values: []string{
						"1 Numbers egg",
						"1 cup yoghurt (curd)",
						"1 handfuls coriander leaves",
						"0 As required salt",
						"0 As required black pepper",
						"1/2 teaspoon cumin powder",
						"1 teaspoon oregano",
						"0 As required water",
						"1 Numbers beetroot",
					},
				},
				Instructions: models.Instructions{
					Values: []string{
						"To make this easy recipe, take a pan and add water, boil the egg with a dash of salt. Once done, remove " +
							"the shell and save for garnishing. In the meantime, wash and chop the beetroot and make a smooth blend.",
						"Next, whisk the curd and add beetroot blend along with spices, chopped coriander leaves and mix it well. " +
							"Garnish with oregano and egg and enjoy.",
					},
				},
				Name: "Beetroot Cold Soup Recipe",
				NutritionSchema: models.NutritionSchema{
					Calories: "189 cal",
					Servings: "1",
				},
				PrepTime: "PT10M",
				Yield:    models.Yield{Value: 2},
				URL:      "https://recipes.timesofindia.com/recipes/beetroot-cold-soup/rs90713582.cms",
			},
		},
		{
			name: "tine.no",
			in:   "https://www.tine.no/oppskrifter/middag-og-hovedretter/kylling-og-fjarkre/rask-kylling-tikka-masala",
			want: models.RecipeSchema{
				AtContext:     atContext,
				AtType:        models.SchemaType{Value: "Recipe"},
				Category:      models.Category{Value: "middag"},
				Cuisine:       models.Cuisine{Value: "indisk"},
				DateModified:  "2022-12-19T09:14:12.161Z",
				DatePublished: "2018-09-05T18:43:37.088Z",
				Description: models.Description{
					Value: "En god og rask oppskrift på en kylling tikka masala. Dette er en rett med små smakseksplosjoner som " +
						"sender tankene til India.",
				},
				Keywords: models.Keywords{Values: "kylling"},
				Image: models.Image{
					Value: "https://www.tine.no/_/recipeimage/w_1080,h_1080,c_fill,x_2880,y_1920,g_xy_center/recipeimage/w1r3ydbmyeqcngqpxatv.jpg",
				},
				Ingredients: models.Ingredients{
					Values: []string{
						"4 dl basmatiris",
						"400 g kyllingfileter",
						"1 ss TINE® Meierismør til steking",
						"1 stk paprika",
						"0.5 dl chili",
						"3 stk vårløk",
						"1 ts hvitløksfedd",
						"1 ss hakket, frisk ingefær",
						"0.5 dl hakket, frisk koriander",
						"2 ts garam masala",
						"3 dl TINE® Lett Crème Fraîche 18 %",
						"3 ss tomatpuré",
						"0.5 ts salt",
						"0.25 ts pepper",
						"0.5 dl slangeagurk",
						"3 dl TINE® Yoghurt Naturell Gresk Type",
						"0.5 dl frisk mynte",
						"1 ts hvitløksfedd",
						"0.5 ts salt",
						"0.25 ts pepper",
					},
				},
				Instructions: models.Instructions{
					Values: []string{
						"Kok ris etter anvisningen på pakken.",
						"Del kylling i biter. Brun kyllingen i smør i en stekepanne på middels varme.",
						"Rens og hakk paprika, chili, vårløk og hvitløk og ha det i stekepannen sammen med kyllingen. Rens og " +
							"finhakk ingefær og frisk koriander. Krydre med garam masala, koriander og ingefær.",
						"Hell i crème fraîche og tomatpuré, og la småkoke i 5 minutter. Smak til med salt og pepper.",
						"Riv agurk og bland den med yoghurt. Hakk mynte og hvitløk og bland det i. Smak til med salt og pepper.",
					},
				},
				Name:  "Rask kylling tikka masala",
				Yield: models.Yield{Value: 4},
				URL:   "https://www.tine.no/oppskrifter/middag-og-hovedretter/kylling-og-fjarkre/rask-kylling-tikka-masala",
			},
		},
		{
			name: "twopeasandtheirpod.com",
			in:   "https://www.twopeasandtheirpod.com/easy-chickpea-salad/",
			want: models.RecipeSchema{
				AtContext:     atContext,
				AtType:        models.SchemaType{Value: "Recipe"},
				Category:      models.Category{Value: "Salad"},
				Cuisine:       models.Cuisine{Value: "American"},
				DatePublished: "2022-03-28T06:03:00+00:00",
				Description: models.Description{
					Value: "This chickpea salad is a vegetarian version of a classic chicken salad, some refer to it as chickpea " +
						"chicken salad. It&#039;s made with basic ingredients, loaded with flavor, and perfect for picnics, " +
						"work lunches, or a simple, healthy lunch at home.",
				},
				Image: models.Image{
					Value: "https://www.twopeasandtheirpod.com/wp-content/uploads/2022/02/Chickpea-Salad-4.jpg",
				},
				Ingredients: models.Ingredients{
					Values: []string{
						"2 (15 oz) cans chickpeas, (drained and rinsed)",
						"3/4 cup diced celery",
						"1/2 cup diced dill pickles",
						"1/2 cup sliced green onion",
						"1/2 cup plain Greek yogurt",
						"1 tablespoon lemon juice",
						"1 to 2 tablespoons Dijon mustard",
						"2 teaspoons red wine vinegar",
						"2 tablespoons freshly chopped dill",
						"2 tablespoons freshly chopped parsley",
						"1/4 teaspoon garlic powder",
						"Kosher salt and black pepper, (to taste)",
					},
				},
				Instructions: models.Instructions{
					Values: []string{
						"Place the chickpeas on a clean dish towel or paper towel. Put another towel on top. Use your hands to " +
							"roll and rub the chickpeas for about 20-30 seconds. This will help the skins come off easier. " +
							"Remove the skins and discard. I try to remove most of the skins, but if you don’t get them all, " +
							"that is ok. And if you don&#39;t have time to remove the skins, just rinse and drain. The salad " +
							"will still be good. Removing the skins just makes the salad a little creamier and smooth.",
						"Place the chickpeas in a large bowl and mash with a fork or potato masher until most of the chickpeas " +
							"are smashed. Stir in the celery, onion, and pickles.",
						"In a small bowl, whisk together the Greek yogurt, lemon juice, mustard, red wine vinegar, dill, parsley, " +
							"garlic powder, salt, and pepper.",
						"Add the sauce to the chickpea mixture and stir until well combined. Taste and adjust ingredients, if " +
							"necessary.",
						"Serve in between two slices of bread to make a sandwich, in pita bread, in a lettuce wrap, in a tortilla, " +
							"with crackers or chips, or on top of a rice cake, or add to a bed of greens to make a salad! The " +
							"options are endless.",
					},
				},
				Keywords: models.Keywords{Values: "vegetarian"},
				Name:     "Chickpea Salad",
				NutritionSchema: models.NutritionSchema{
					Calories:       "215 kcal",
					Carbohydrates:  "32 g",
					Cholesterol:    "1 mg",
					Fat:            "5 g",
					Fiber:          "10 g",
					Protein:        "14 g",
					SaturatedFat:   "1 g",
					Servings:       "1",
					Sodium:         "803 mg",
					Sugar:          "2 g",
					TransFat:       "1 g",
					UnsaturatedFat: "3 g",
				},
				PrepTime: "PT10M",
				Yield:    models.Yield{Value: 4},
				URL:      "https://www.twopeasandtheirpod.com/easy-chickpea-salad/",
			},
		},
		{
			name: "valdemarsro.dk",
			in:   "https://www.valdemarsro.dk/butter_chicken/",
			want: models.RecipeSchema{
				AtContext:     atContext,
				AtType:        models.SchemaType{Value: "Recipe"},
				Name:          "Butter chicken",
				CookTime:      "PT30M",
				DateModified:  "2023-03-30T10:56:42+00:00",
				DatePublished: "2014-06-25T05:51:40+00:00",
				Description: models.Description{
					Value: "Krydret og lækker butter chicken med ris, indisk salat og godt brød til – det er virkelig en skøn ret.Selv om butter chicken måske umiddelbart tager lidt tid, så er det mest af alt tiden hvor kyllingen skal marinere, for selve arbejdstiden er hverdagsvenlig. Sæt evt kyllingen i marinade allerede aftenen før eller fra morgenstunden – det bliver den blot bedre af.Det er også en herlig weekendret og god til gæster, hvor man kan servere flere skål med lækkert indisk mad at samles om.Prøv også: min bedste opskrift på lækkert naan brød >>",
				},
				Image: models.Image{
					Value: "https://www.valdemarsro.dk/wp-content/2014/06/butterchicken.jpg",
				},
				PrepTime: "PT1H30M",
				Yield:    models.Yield{Value: 4},
				Ingredients: models.Ingredients{
					Values: []string{
						"100 g græsk yoghurt 10 %",
						"2 tsk chiliflager",
						"0,50 tsk stødt nellike",
						"2 tsk stødt spidskommen",
						"1 tsk stødt kardemomme",
						"2 tsk garam masala",
						"2 fed hvidløg, finthakkede",
						"1 tsk stødt gurkemeje",
						"1 spsk ingefær, friskrevet",
						"1 dåse hakkede tomater",
						"500 g kyllingebryst, skåret i tern",
						"1 løg, skåret i ringe",
						"0,50 dl piskefløde",
						"50 g smør",
						"2 spsk olivenolie",
						"flagesalt",
						"sort peber, friskkværnet",
						"3 dl basmati ris, kogt efter anvisning på emballagen",
						"2 håndfulde frisk koriander",
					},
				},
				Instructions: models.Instructions{
					Values: []string{
						"Blend yoghurt, chili, nellike, spidskommen, kardemomme, garam masala, gurkemeje, ingefær, hvidløg og " +
							"hakkede tomater sammen til en lind sauce.",
						"Hæld den over kyllingestykkerne, dæk dem til og lad dem trække i min. 30 minutter i køleskabet og gerne " +
							"natten over eller fra morgenstund til aftensmadstid.",
						"Smelt smør og olie i en sauterpande. Sautér løgene, til de bliver bløde og blanke.",
						"Tilsæt kylling, sauce og fløde.",
						"Lad det simre ved lav varme i 30-35 minutter, eller til kyllingen er mør. Server med ris og et drys frisk " +
							"koriander",
					},
				},
				URL: "https://www.valdemarsro.dk/butter_chicken/",
			},
		},
		{
			name: "vanillaandbean.com",
			in:   "https://vanillaandbean.com/carrot-cake-bread/",
			want: models.RecipeSchema{
				AtContext:     atContext,
				AtType:        models.SchemaType{Value: "Recipe"},
				Category:      models.Category{Value: "Brunch"},
				CookTime:      "PT60M",
				Cuisine:       models.Cuisine{Value: "American"},
				DatePublished: "2022-04-08T17:54:24+00:00",
				Description: models.Description{
					Value: "Carrot Cake Bread  is cake in carrot cake loaf form! It&#039;s just as easy to make as my moist Carrot " +
						"Cake Recipe with Pineapple, but because carrot cake quick bread is made in a loaf pan, it&#039;s a bit " +
						"more casual. Share with an orange spiked cream cheese frosting, or enjoy without. Either way, it&#039;s " +
						"perfect for all your occasions and general snacking!*Time above does not include two hours to cool the " +
						"cake completely before icing. See blog post for more tips.",
				},
				Keywords: models.Keywords{
					Values: "Carrot Cake Bread, Carrot Cake Loaf Recipe",
				},
				Image: models.Image{
					Value: "https://vanillaandbean.com/wp-content/uploads/2022/04/CarrotCakeBreadFinal-5.jpg",
				},
				Ingredients: models.Ingredients{
					Values: []string{
						"3/4 cup (80 grams) Pecans (or Walnuts)",
						"1 1/4 cup (175 grams) Unbleached All Purpose Flour (I use Bob&#039;s Red Mill)",
						"1 tablespoon Corn Starch",
						"1/2 teaspoon Ground Ginger",
						"1 teaspoon Ground Cinnamon",
						"1/2 teaspoon Fine Sea Salt",
						"1 1/4 teaspoon Baking Powder",
						"1/2 teaspoon Baking Soda",
						"1/3 cup (70 grams) Dark Brown Sugar (or light, packed)",
						"1/3 cup (75 grams) Cane Sugar",
						"3/4 cup (150 grams) Coconut Oil (melted and warm to touch, or Canola Oil)",
						"1 Large Orange (zested and 2 tablespoons of orange juice)",
						"2 Whole Eggs (room temperature* see note)",
						"2 teaspoons Vanilla Extract",
						"1/3 cup (75 grams) Unsweetened Whole Milk Yogurt or Greek Yogurt",
						"1 1/4 cups (160 grams) Finely Shredded Carrots (packed, about 3 medium carrots)",
						"5 tablespoons Unsalted Butter (room temperature)",
						"7 ounces (200 grams) Cream Cheese (room temperature)",
						"1/3 cup (40g) Powdered Sugar (sifted)",
						"1 Large Orange (zested)",
						"1/2 Vanilla Bean (scraped, or use 1/2 tsp vanilla bean paste or 1 teaspoon vanilla extract)",
						"Orange Juice (fresh squeezed, a few teaspoons for loosening the icing if needed.)",
					},
				},
				Instructions: models.Instructions{
					Values: []string{
						"Arrange an oven rack in the center of the oven. Preheat the oven to 350F (180C). Place the nuts on a small " +
							"sheet pan and put them in the oven while the oven is preheating. Toast the nuts for about 15 minutes, " +
							"or until the nuts are fragrant and starting to turn slightly darker. Chop fine when cool. Set aside.",
						"Grease a 9 inch by 5 inch (23 centimeters by 13 centimeters) loaf pan thoroughly on the all sides and b" +
							"ottom. Be sure to get the corners good! Line the pan with parchment paper, just one strip across the " +
							"bottom and sides will act as a handle to remove the bread once baked. Clip the sides of the parchment " +
							"to the pan if desired.",
						"In a medium mixing bowl, whisk the flour, corn starch, ginger, cinnamon, salt, baking powder, and baking " +
							"soda. Set aside.",
						"In a large mixing bowl, whisk the brown sugar, cane sugar and oil until ingredients are combined, about 30 " +
							"seconds. Add the eggs, orange zest and juice and vanilla. Whisk until the ingredients are emulsified, " +
							"about 20 seconds.",
						"Add the dry ingredients to the egg/sugar mixture. Using a silicone spatula, gently mix/fold until no flour " +
							"streaks remain. If using coconut oil, the batter will start to stiffen at this point because the coconut " +
							"oil is cooling (solidifying). Fold in the yogurt until no white streaks remain.Working quickly, fold in " +
							"the shredded carrots and nuts. If using canola oil, the batter should be thick but loose and almost pourable " +
							"(pictured in the photos). For coconut oil, the batter should be stiff and thick.",
						"Transfer the cake batter to the loaf pan and using an offset spatula,  smooth the batter into an even layer. " +
							"If using coconut oil, spread the batter while pressing it into the pan. Tap the pan on the counter to " +
							"disperse any air pockets. Remove the parchment clips if using.",
						"Bake the cake for 60-70 minutes. The loaf is ready when it&#039;s golden, a toothpick poked in the center " +
							"comes out clean, the cake slightly springs back under gentle pressure at center, and the edges of the " +
							"loaf are just starting to pull away from the sides of the pan. Allow cake to cool in the pan for 20 " +
							"minutes on a cooling rack. Lift cake out and transfer to a cooking rack. Cool completely at room " +
							"temperature before icing or covering (about 2 hours).",
						"Start with ingredients at room temperature. Place butter in mixer and beat on high with the paddle " +
							"attachment until light and fluffy, about three minutes. Scrape down your bowl several times to make sure " +
							"the butter is getting whipped. Add the cream cheese and beat another  minute. Scrape down the bowl.",
						"Sift in the powder sugar, add the vanilla, and orange zest then mix on low until incorporated, about 30 " +
							"seconds. Taste for sweet adjustment and add another tablespoon of powered sugar if desired. To loosen the " +
							"icing, you can add a few teaspoons of orange juice.",
						"Store the icing in a lidded container, in the refrigerator, if not icing the bread after it cools. Bring " +
							"to room temperature before slathering!",
						"Because cream cheese needs to be refrigerated, consider how/when you&#039;ll share the cake before icing it. " +
							"An iced cake needs to be consumed, refrigerated or frozen within three hours. An uniced cake can set " +
							"covered at room temperature for up to three days.Ice the Cake:Once the loaf cake is cool and just before " +
							"serving, spread the icing evenly over the top of the cake. Slice and share within three hours.Individual " +
							"Pieces:Once the loaf is cool, slice the loaf and share with a dollop or slather of cream cheese icing. " +
							"*Uniced carrot cake bread is lovely toasted or warmed, then slathered with cream cheese icing!",
						"Room Temperature: Store an iced loaf cake at room temperature (70F or less) for up to three hours. " +
							"Afterwards, store the cake in the fridge, covered for up to three days. Before serving, pull the cake " +
							"from the fridge and rest at room temperature for about 30 minutes to allow the fats to soften. Food safety " +
							"says cheese should set out for no more than three hours. An uniced loaf cake can be stored covered at room " +
							"temperature for up to three days.",
						"To Freeze: This cake freezes beautifully, iced or uniced. Simply allow the cake to cool completely, ice the " +
							"cake or not, then freeze individual pieces on a sheet pan. Once frozen wrap snugly in plastic wrap or " +
							"store in a lidded container. Store in freezer then thaw at room temperature for about two hours before " +
							"enjoying. " +
							"If iced, take the plastic wrap off before thawing so the icing doesn&#039;t stick.I&#039;ve only tested " +
							"freezing " +
							"the iced cake for up to two days. Uniced will freeze well for up to two weeks. Unwrap and thaw at room " +
							"temperature, " +
							"covered with a cake dome.",
					},
				},
				Name: "Carrot Cake Bread Recipe",
				NutritionSchema: models.NutritionSchema{
					Calories:       "476 kcal",
					Carbohydrates:  "38 g",
					Cholesterol:    "68 mg",
					Fat:            "35 g",
					Fiber:          "2 g",
					Protein:        "6 g",
					SaturatedFat:   "22 g",
					Servings:       "1",
					Sodium:         "263 mg",
					Sugar:          "22 g",
					TransFat:       "1 g",
					UnsaturatedFat: "11 g",
				},
				PrepTime: "PT30M",
				Yield:    models.Yield{Value: 10},
				URL:      "https://vanillaandbean.com/carrot-cake-bread/",
			},
		},
		{
			name: "vegolosi.it",
			in:   "https://www.vegolosi.it/ricette-vegane/pancake-vegani-senza-glutine-alla-quinoa-e-cocco/",
			want: models.RecipeSchema{
				AtContext:     atContext,
				AtType:        models.SchemaType{Value: "Recipe"},
				CookTime:      "PT15M",
				DatePublished: "2018-08-02",
				Description: models.Description{
					Value: "I pancake vegani senza glutine alla quinoa e cocco sono una deliziosa colazione che permette anche " +
						"a chi deve evitare il glutine di gustarsi dei morbidi e golosissimi pancake, completati in questo " +
						"caso dall'immancabile sciroppo d'acero, fragole fresche e cocco in scaglie.",
				},
				Image: models.Image{
					Value: "https://www.vegolosi.it/wp-content/uploads/2018/06/pancake-senza-glutine-cocco-quinoa_1592_650.jpg",
				},
				Ingredients: models.Ingredients{
					Values: []string{
						"120 g di farina di quinoa",
						"60 g di farina di cocco",
						"20 g di amido di mais",
						"30 g di zucchero di canna grezzo",
						"2 cucchiaini di lievito (cremor tartaro)",
						"2 cucchiai di farina di semi di lino",
						"½ cucchiaino di bicarbonato",
						"400 g di latte di soia",
						"30 g di olio di semi di girasole",
						"1 cucchiaino di aceto di mele",
						"Fragole",
						"Sciroppo d&#8217;acero",
					},
				},
				Instructions: models.Instructions{
					Values: []string{
						"Mescolate la farina di semi di lino con 6 cucchiai di acqua e lasciate riposare 10 minuti fino a che si " +
							"sarà formato un composto gelatinoso. In una ciotolina versate il latte di soia e l&#8217;aceto d" +
							"i mele e lasciate cagliare per 5 minuti. Riunite in una ciotola la farina di quinoa, la farina di " +
							"cocco, l&#8217;amido di mais, lo zucchero di canna, il lievito e il bicarbonato e mescolate. " +
							"Aggiungete agli ingredienti secchi il latte di soia, il composto di semi di lino e l&#8217;olio e " +
							"amalgamate bene il tutto fino ad ottenere un composto omogeneo e abbastanza denso.",
						"Scaldate una padella antiaderente e ungetela leggermente con un pezzo di carta assorbente imbevuto di " +
							"olio di semi. Versate 1-2 cucchiaiate di impasto per ciascun pancake e lasciate cuocere a fiamma " +
							"medio-bassa per 3-4 minuti per lato. Man mano che i vostri pancake saranno pronti disponeteli su un " +
							"piatto, e completate poi ciascuna porzione con sciroppo d&#8217;acero a piacere, fragole fresche e cocco " +
							"in scaglie.",
					},
				},
				Name:     "Pancake vegani senza glutine alla quinoa e cocco",
				PrepTime: "PT20M",
				Yield:    models.Yield{Value: 4},
				URL:      "https://www.vegolosi.it/ricette-vegane/pancake-vegani-senza-glutine-alla-quinoa-e-cocco/",
			},
		},
		{
			name: "vegrecipesofindia.com",
			in:   "https://www.vegrecipesofindia.com/paneer-butter-masala/",
			want: models.RecipeSchema{
				AtContext:     atContext,
				AtType:        models.SchemaType{Value: "Recipe"},
				Category:      models.Category{Value: "Main Course"},
				CookTime:      "PT30M",
				Cuisine:       models.Cuisine{Value: "North Indian"},
				DatePublished: "2022-10-31T23:35:31+00:00",
				Description: models.Description{
					Value: "Paneer Butter Masala Recipe is one of India’s most popular paneer preparation. This restaurant style recipe with soft paneer cubes dunked in a creamy, lightly spiced tomato sauce or gravy is a best one that I have been making for a long time. This rich dish is best served with roti or chapati, paratha, naan or rumali roti.",
				},
				Keywords: models.Keywords{Values: "Paneer Butter Masala"},
				Image: models.Image{
					Value: "https://www.vegrecipesofindia.com/wp-content/uploads/2020/01/paneer-butter-masala-5.jpg",
				},
				Ingredients: models.Ingredients{
					Values: []string{
						"18 to 20 cashews (- whole)",
						"⅓ cup hot water (- for soaking cashews)",
						"2 cups tomatoes (- diced, or 300 grams tomatoes or 4 to 5 medium sized, pureed)",
						"1 inch ginger (- peeled and roughly chopped)",
						"3 to 4 garlic cloves (- small to medium-sized, peeled)",
						"2 tablespoons Butter (or 1 tablespoon oil + 1 or 2 tablespoons butter)",
						"1 tej patta ((Indian bay leaf), optional)",
						"½ to 1 teaspoon kashmiri red chili powder (or deghi mirch or ¼ to ½ teaspoon cayenne pepper or paprika)",
						"1.5 cups water (or add as required)",
						"1 inch ginger (- peeled and julienned, reserve a few for garnish)",
						"1 or 2 green chili (- slit, reserve a few for garnish)",
						"200 to 250 grams Paneer ((Indian cottage cheese) - cubed or diced)",
						"1 teaspoon dry fenugreek leaves ((kasuri methi) - optional)",
						"½ to 1 teaspoon Garam Masala (or tandoori masala)",
						"2 to 3 tablespoons light cream (or half &amp; half or 1 to 2 tablespoons heavy cream - optional)",
						"¼ to 1 teaspoon sugar (- optional, add as required depending on the sourness of the tomatoes)",
						"salt (as required)",
						"1 to 2 tablespoons coriander leaves (- chopped, (cilantro) - optional)",
						"1 inch ginger (- peeled and julienned)",
						"1 tablespoon light cream (or 1 tablespoon heavy cream - optional)",
						"1 to 2 teaspoons Butter (- optional)",
					},
				},
				Instructions: models.Instructions{
					Values: []string{
						"Soak cashews in a hot water for 20 to 30 minutes. When the cashews are soaking, you can prep the other " +
							"ingredients like chopping tomatoes, preparing ginger-garlic paste, slicing paneer etc.",
						"Then drain and add the soaked cashews in a blender or mixer-grinder.",
						"Add 2 to 3 tablespoons water and blend to a smooth and fine paste without any tiny bits or chunks of cashews.",
						"In the same blender add the roughly chopped tomatoes. No need to blanch the tomatoes before blending.",
						"Blend to a smooth tomato puree. Set aside. Don’t add any water while blending the tomatoes.",
						"Melt butter in a pan on a low heat. Add tej patta and fry for 2 to 3 seconds or till the oil become fragrant.",
						"Add ginger-garlic paste and sauté for about 10 to 12 seconds till the raw aroma disappears.",
						"Add the tomato puree and stir well. Cook for 5 to 6 minutes stirring a few times.",
						"Next add kashmiri red chili powder and stir again. Continue to sauté till the oil starts to leave the sides of the tomato paste. The tomato paste will thicken considerably and will start coming together as one whole lump.",
						"Then add cashew paste and stir well. Sauté the cashew paste for a few minutes till the oil begins to leave " +
							"the sides of the masala paste.",
						"The cashew paste will begin to cook fast. Approx 3 to 4 minutes on a low heat. So keep stirring non-stop.",
						"Add water and mix very well. Simmer on a low to medium-low heat.",
						"The curry will come to a boil.&nbsp;",
						"After 2 to 3 minutes of boiling, add ginger julienne. Reserve a few for garnishing. The curry will also " +
							"begin to thicken.",
						"Add julienned ginger and green chillies, salt and sugar and simmer till the curry begins to thicken.",
						"After 3 to 4 minutes, add slit green chillies.also add salt as per taste and ½ to 1 teaspoon sugar (optional).",
						"You can vary the sugar quantity from ¼ tsp to 1 teaspoon or more depending on the sourness of the tomatoes. " +
							"Sugar is optional and you can skip it too. If you add cream, then you will need to add less sugar.",
						"Mix very well and simmer for a minute.",
						"After the gravy thickens to your desired consistency, then add the paneer cubes and stir gently.I keep the " +
							"gravy to a medium consistency.",
						"After that add crushed kasuri methi (dry fenugreek leaves), garam masala and cream. Gently mix and then " +
							"switch off the heat.",
						"Garnish the curry with coriander leaves and ginger julienne.",
						"You can even dot the gravy with some butter or drizzle some cream.",
						"Serve Paneer Butter Masala hot with plain naan, garlic naan, roti, paratha or steamed basmati or jeera rice " +
							"or even peas pulao.",
						"Side accompaniments can be an onion-cucumber salad or some pickle. Also serve some lemon wedges by the side.",
					},
				},
				Name: "Paneer Butter Masala Recipe (Restaurant Style)",
				NutritionSchema: models.NutritionSchema{
					Calories:      "307 kcal",
					Carbohydrates: "9 g",
					Cholesterol:   "66 mg",
					Fat:           "27 g",
					Fiber:         "2 g",
					Protein:       "9 g",
					SaturatedFat:  "15 g",
					Servings:      "1",
					Sodium:        "493 mg",
					Sugar:         "4 g",
				},
				PrepTime: "PT10M",
				Yield:    models.Yield{Value: 4},
				URL:      "https://www.vegrecipesofindia.com/paneer-butter-masala/",
			},
		},
		{
			name: "watchwhatueat.com",
			in:   "https://www.watchwhatueat.com/healthy-fried-brown-rice/",
			want: models.RecipeSchema{
				AtContext:     atContext,
				AtType:        models.SchemaType{Value: "Recipe"},
				Category:      models.Category{Value: "Main or Side"},
				CookTime:      "PT30M",
				Cuisine:       models.Cuisine{Value: "Asian Inspired"},
				DatePublished: "2019-04-26T13:31:54+00:00",
				Description: models.Description{
					Value: "Learn how to make healthy fried brown rice with fresh vegetables. Simple, delicious and wholesome " +
						"vegetable fried rice recipe that is better than takeout.",
				},
				Keywords: models.Keywords{Values: "fried rice, healthy fried rice"},
				Image: models.Image{
					Value: "https://www.watchwhatueat.com/wp-content/uploads/2019/04/Healthy-Fried-Brown-Rice-6.jpg",
				},
				Ingredients: models.Ingredients{
					Values: []string{
						"2 cup Jasmine brown rice",
						"3 1/2 cup water",
						"1/2 tbsp sesame oil",
						"3-4 medium garlic cloves (finely chopped)",
						"1/2\"  ginger (peeled and finely chopped)",
						"1/2 bell pepper diced (red and green each)",
						"1 large carrot (peeled and diced)",
						"1 cup green peas (fresh or frozen)",
						"3-4 springs of green onion ((scallions))",
						"1 1/2 tbsp soy sauce",
						"1 tbsp rice vinegar",
						"salt to taste",
					},
				},
				Instructions: models.Instructions{
					Values: []string{
						"In a medium pot add rice and water and bring mixture to boil. Then simmer on low heat with cover for 25 " +
							"mins. Let it cool down completely. Using a fork fluff the rice to separate the grains before using.",
						"In a large skillet or wok heat oil on medium to high heat.",
						"Add chopped garlic, ginger and the white portion of scallions (reserve green for garnishing). Cook it for " +
							"1-2 mins.",
						"Then add diced peppers, carrot and peas. Stir fry them for few minutes (see notes).",
						"Now add soy sauce, rice vinegar and mix well.",
						"Finally, add cold rice and mix well with vegetables. Season with salt if necessary.",
						"Garnish with green onions and serve warm.",
					},
				},
				Name: "Healthy Fried Brown Rice With Vegetables",
				NutritionSchema: models.NutritionSchema{
					Calories: "336 kcal",
					Servings: "1",
				},
				PrepTime: "PT15M",
				Yield:    models.Yield{Value: 5},
				URL:      "https://www.watchwhatueat.com/healthy-fried-brown-rice/",
			},
		},
		{
			name: "whatsgabycooking.com",
			in:   "https://whatsgabycooking.com/pea-prosciutto-spring-pizza/",
			want: models.RecipeSchema{
				AtContext:     atContext,
				AtType:        models.SchemaType{Value: "Recipe"},
				Category:      models.Category{Value: "Breakfast"},
				CookTime:      "PT10M",
				Cuisine:       models.Cuisine{Value: "Italian"},
				DatePublished: "2022-03-31T07:00:00+00:00",
				Description: models.Description{
					Value: "When in doubt of what to do with a bunch of fab farmers market produce, put it all on a pizza and " +
						"slap an egg on it, duh!",
				},
				Image: models.Image{
					Value: "https://whatsgabycooking.com/wp-content/uploads/2015/05/ALDI-Spring-Pea-Pizza-2-copy-2.jpg",
				},
				Ingredients: models.Ingredients{
					Values: []string{
						"1 recipe fresh pizza dough",
						"1/3 cup Basil Vinaigrette",
						"1 cup mozzarella (shredded or fresh mozzarella sliced )",
						"1/2 cup fresh peas (blanched)",
						"1/2 cup sugar snap peas (sliced thin )",
						"1 bunch asparagus (tips only, blanched)",
						"2 eggs",
						"4 ounces prosciutto",
						"Kosher salt and freshly cracked black pepper to taste",
						"Fresh Basil (torn into pieces)",
					},
				},
				Instructions: models.Instructions{
					Values: []string{
						"Pre-heat oven to 475 degrees F.",
						"Shape the dough into 2 medium-ish pizzas while on a clean floured surface. Let the dough sit for 5 " +
							"minutes and then re-form to make sure it's as big as you'd like. Place the pizza dough on a " +
							"lightly floured rimless baking sheet, or pizza peel.",
						"Spread the basil vinaigrette over the top of each pizza. Top with the mozzarella, scatter the peas and " +
							"asparagus on top of the cheese. Transfer to an oven and cook for about 5 minutes. Remove the " +
							"pizza and add the egg on top of each pizza and transfer them back into the oven to continue " +
							"to cook until the egg white is set and the yolk still runny .",
						"Remove from the oven, Add the prosciutto on top and garnish with basil.",
					},
				},
				Keywords: models.Keywords{Values: "homemade pizza, spring pizza"},
				Name:     "Pea Prosciutto Spring Pizza",
				NutritionSchema: models.NutritionSchema{
					Calories:       "610 kcal",
					Carbohydrates:  "56 g",
					Cholesterol:    "123 mg",
					Fat:            "33 g",
					Fiber:          "5 g",
					Protein:        "24 g",
					SaturatedFat:   "11 g",
					Servings:       "1",
					Sodium:         "1124 mg",
					Sugar:          "11 g",
					TransFat:       "0.04 g",
					UnsaturatedFat: "19 g",
				},
				PrepTime: "PT10M",
				Yield:    models.Yield{Value: 4},
				URL:      "https://whatsgabycooking.com/pea-prosciutto-spring-pizza/",
			},
		},
		{
			name: "wikibooks.org",
			in:   "https://en.wikibooks.org/wiki/Cookbook:Creamed_Spinach",
			want: models.RecipeSchema{
				AtContext: atContext,
				AtType:    models.SchemaType{Value: "Recipe"},
				Name:      "Creamed Spinach",
				Category:  models.Category{Value: "Sauce recipes"},
				Description: models.Description{
					Value: "Creamed spinach makes a nutritious sauce that goes well with fish and meat dishes. In Swedish cuisine it has traditionally been used with boiled potatoes and fish or with chipolata, the Swedish \"prince-sausage\". Creamed spinach can be done in two different ways: either with whole spinach or with chopped spinach.",
				},
				Image: models.Image{
					Value: "https://upload.wikimedia.org/wikipedia/commons/thumb/1/11/Chipolata.jpg/300px-Chipolata.jpg",
				},
				Instructions: models.Instructions{
					Values: []string{
						"Boil the fresh spinach leaves and the salt in the water for about 5 minutes. If you use frozen spinach, follow the instructions on the box.",
						"Drain away the water, and let the spinach dry for a minute or so.",
						"Put the spinach back into the pan and add the cream. Simmer for a few more minutes.",
						"Add salt and pepper to taste.",
						"Boil the fresh spinach leaves as in variation I, or gently thaw the frozen chopped spinach in a little bit of water in a pan.",
						"Mix the flour with the chopped spinach in the pan, and add the milk.",
						"Bring the spinach to the boil, and simmer gently for 3–5 minutes.",
						"Add salt and pepper to taste.",
					},
				},
				Ingredients: models.Ingredients{
					Values: []string{
						"500 g fresh spinach (or whole frozen spinach leaves)",
						"2 dl water",
						"½ dl cream or 1 tbsp butter",
						"½ tsp salt",
						"Black pepper",
						"400 g of chopped spinach, fresh or frozen",
						"1 ½ tbsp flour",
						"1 dl milk",
						"½ tsp salt",
						"Black pepper",
					},
				},
				Yield: models.Yield{Value: 4},
				URL:   "https://en.wikibooks.org/wiki/Cookbook:Creamed_Spinach",
			},
		},
		{
			name: "wikibooks.org_mobile",
			in:   "https://en.m.wikibooks.org/wiki/Cookbook:Creamed_Spinach",
			want: models.RecipeSchema{
				AtContext: atContext,
				AtType:    models.SchemaType{Value: "Recipe"},
				Name:      "Creamed Spinach",
				Category:  models.Category{Value: "Sauce recipes"},
				Description: models.Description{
					Value: "Creamed spinach makes a nutritious sauce that goes well with fish and meat dishes. In Swedish cuisine it has traditionally been used with boiled potatoes and fish or with chipolata, the Swedish \"prince-sausage\". Creamed spinach can be done in two different ways: either with whole spinach or with chopped spinach.",
				},
				Image: models.Image{
					Value: "https://upload.wikimedia.org/wikipedia/commons/thumb/1/11/Chipolata.jpg/300px-Chipolata.jpg",
				},
				Instructions: models.Instructions{
					Values: []string{
						"Boil the fresh spinach leaves and the salt in the water for about 5 minutes. If you use frozen spinach, follow the instructions on the box.",
						"Drain away the water, and let the spinach dry for a minute or so.",
						"Put the spinach back into the pan and add the cream. Simmer for a few more minutes.",
						"Add salt and pepper to taste.",
						"Boil the fresh spinach leaves as in variation I, or gently thaw the frozen chopped spinach in a little bit of water in a pan.",
						"Mix the flour with the chopped spinach in the pan, and add the milk.",
						"Bring the spinach to the boil, and simmer gently for 3–5 minutes.",
						"Add salt and pepper to taste.",
					},
				},
				Ingredients: models.Ingredients{
					Values: []string{
						"500 g fresh spinach (or whole frozen spinach leaves)",
						"2 dl water",
						"½ dl cream or 1 tbsp butter",
						"½ tsp salt",
						"Black pepper",
						"400 g of chopped spinach, fresh or frozen",
						"1 ½ tbsp flour",
						"1 dl milk",
						"½ tsp salt",
						"Black pepper",
					},
				},
				Yield: models.Yield{Value: 4},
				URL:   "https://en.m.wikibooks.org/wiki/Cookbook:Creamed_Spinach",
			},
		},
		{
			name: "woop.co.nz",
			in:   "https://woop.co.nz/thai-marinated-beef-sirlion-344-2-f.html",
			want: models.RecipeSchema{
				AtContext: atContext,
				AtType:    models.SchemaType{Value: "Recipe"},
				Name:      "Thai marinated beef sirlion",
				Description: models.Description{
					Value: "with crispy noodle salad",
				},
				Yield: models.Yield{Value: 2},
				Image: models.Image{
					Value: "https://woop.co.nz/media/catalog/product/cache/f4f005ad5960ef8c7b8a08a9a3fc244e/f/-/f-marinated-thai-beef-sirloin_mrypusp3a6h8fzas.jpg",
				},
				Ingredients: models.Ingredients{
					Values: []string{
						"1 pack of marinated beef sirloin steak",
						"1 pot of Thai dressing",
						"1 pack of crispy noodles",
						"1 sachet of roasted peanuts",
						"1 bag of baby leaves",
						"Cucumber",
						"1 tomato",
						"1 red onion",
						"1 bag of coriander",
					},
				},
				Instructions: models.Instructions{
					Values: []string{
						"1. TO PREPARE THE SALAD:   Using half the cucumber cut in half lengthways and scoop out the seeds with a " +
							"teaspoon then slice on the diagonal into ½ cm halfmoons. Dice the tomato into 1 cm cubes then " +
							"thinly slice ½ of the red onion.",
						"2. TO FINISH THE SALAD:   Place the cucumber, tomato and red onion into a serving bowl along with the crispy " +
							"noodles and baby leaves. Roughly chop the coriander leaves and stalk then add to the serving bowl. " +
							"Add half of the Thai dressing and season with salt and pepper and toss before serving.",
						"3. TO COOK THE BEEF SIRLOIN STEAK:   Remove the marinated beef sirloin steaks from their packaging and pat " +
							"dry with a paper towel and season with salt and pepper. Heat a drizzle of oil in a non-stick frying " +
							"pan over a medium-high heat. Once hot cook the beef for 2-3 mins each side for medium-rare – a little " +
							"longer for well done. Remove from the pan and allow to rest for a few mins before slicing thinly. BBQ " +
							"Instructions: Heat your BBQ up to a medium heat. Once hot cook beef steaks for 2-3 mins each side. " +
							"Remove from the BBQ and allow to rest for a few mins before slicing thinly.",
						"TO SERVE:   Divide salad between plates then top with sliced beef. Drizzle over remaining Thai dressing and " +
							"sprinkle with roasted peanuts.",
					},
				},
				Keywords: models.Keywords{Values: "Magento, Varien, E-commerce"},
				NutritionSchema: models.NutritionSchema{
					Calories:      "2536kj (606Kcal)",
					Carbohydrates: "43g",
					Protein:       "44g",
					Fat:           "28g",
				},
				URL: "https://woop.co.nz/thai-marinated-beef-sirlion-344-2-f.html",
			},
		},
		{
			name: "ye-mek.net",
			in:   "https://ye-mek.net/recipe/walnut-turkish-baklava-recipe",
			want: models.RecipeSchema{
				AtContext: atContext,
				AtType:    models.SchemaType{Value: "Recipe"},
				Name:      "Walnut Turkish Baklava Recipe",
				Image: models.Image{
					Value: "https://cdn.ye-mek.com/img/f/hazir-yufkadan-buzme-burma-baklava.jpg",
				},
				Ingredients: models.Ingredients{
					Values: []string{
						"11-12 sheets filo pastry",
						"170 g butter",
						"Crushed walnut",
						"For Syrup:",
						"2 cups water",
						"2 cups granulated sugar",
						"Quarter lemon juice",
					},
				},
				Instructions: models.Instructions{
					Values: []string{
						"Firstly, prepare your dessert syrup.",
						"For the syrup: Put water and sugar in a saucepan, stir until sugar is dissolved. Turn down the bottom after boiling syrup. Boil for 20 minutes. Add lemon juice into the syrup. Boil over low heat for 10 minutes.",
						"Remove syrup from heat, leave to cool. Then, melt the butter in a small pan.",
						"On the other hand, take a filo pastry. Put on the kitchen counter. Apply melted butter with a brush onto filo pastry. Sprinkle with crushed walnuts on to filo pastry. Slowly wrap in roll form. Then, follow the insructions at photo 9-10-11-12.  Put into a greased baking tray. Repeat the same process.",
						"Heat the remaining butter. Slowly pour butter onto baklava.",
						"Give a preheated 180 degree oven. Cook until golden brown (about 30 minutes).",
						"Remove cooked baklava from the oven. Leave to cool for 8-10 minutes. Then, cut the baklava with a knife or spatula.",
						"Finally, pour the cool syrup onto baklava. Rest for 20-25 minutes. Then, service.",
					},
				},
				URL: "https://ye-mek.net/recipe/walnut-turkish-baklava-recipe",
			},
		},
		{
			name: "zenbelly.com",
			in:   "https://www.zenbelly.com/short-ribs/",
			want: models.RecipeSchema{
				AtContext: atContext,
				AtType:    models.SchemaType{Value: "Recipe"},
				Name:      "pressure cooker honey balsamic short ribs",
				Description: models.Description{
					Value: "rich and decadent short ribs are cooked in a fraction of the time in the instant pot.",
				},
				Category:      models.Category{Value: "beef"},
				CookingMethod: models.CookingMethod{Value: "pressure cook"},
				Image: models.Image{
					Value: "https://www.zenbelly.com/wp-content/uploads/2020/08/short-ribs-1-scaled-225x225.jpg",
				},
				Ingredients: models.Ingredients{
					Values: []string{
						"3.5 pounds bone-in short ribs (try to get ones that are at least 1.5 inches thick)",
						"salt and pepper",
						"1 tablespoon extra virgin olive oil",
						"1 medium onion, sliced",
						"4-5 cloves garlic, smashed and roughly chopped",
						"1/2 cup balsamic vinegar",
						"1/4 cup tamari or coconut aminos",
						"2 tablespoons honey",
						"1 tablespoon dijon mustard",
						"1 cup chicken or beef broth or stock",
						"2 sprigs thyme",
						"1 lemon, zested and juiced",
						"fresh parsley, roughly chopped",
					},
				},
				Instructions: models.Instructions{
					Values: []string{
						"Liberally season the short ribs with salt and pepper &#8211; ideally a day in advance and refrigerated, " +
							"but if not, up to two hours at room temperature.",
						"Turn your Instant Pot on Sauté (high, if it&#8217;s an option). Once it reads HOT, add the oil. Brown " +
							"the short ribs in batches, until well browned on the meaty side (you don&#8217;t " +
							"need to brown them on all sides)",
						"Remove the short ribs to a plate and add the onions. Sauté for 5-6 minutes, until browned and softened.",
						"Add the garlic and saut\u00e9 for another minute, until fragrant.",
						"Add the balsamic vinegar, tamari, honey, mustard, broth, and thyme. Once it starts to bubble, add the " +
							"short ribs back in, meaty side down.",
						"Hit Cancel on the Instant Pot. Lock on the lid, making sure the seal is in place. Set to cook for 40 " +
							"minutes at high pressure, making sure the valve is set to sealing.",
						"When the time is up, release the pressure and remove the lid once it unlocks. Remove the ribs to a " +
							"platter and turn the Instant Pot back to Saut\u00e9. Reduce the sauce until it&#8217;s about 3 cups in volume.",
						"Pour the sauce over the ribs. Pour over the lemon juice and sprinkle with lemon zest and parsley.",
						"Serve over polenta, mashed potatoes, or whatever you&#8217;d like."},
				},
				Keywords:      models.Keywords{Values: "short ribs"},
				Yield:         models.Yield{Value: 4},
				PrepTime:      "PT20M",
				CookTime:      "PT1H10M",
				DatePublished: "2020-09-01",
				URL:           "https://www.zenbelly.com/short-ribs/",
			},
		},
		{

			name: "101cookbooks.com",
			in:   "https://www.101cookbooks.com/simple-bruschetta/",
			want: models.RecipeSchema{
				AtContext: atContext,
				AtType:    models.SchemaType{Value: "Recipe"},
				Name:      "Simple Bruschetta",
				Description: models.Description{
					Value: "Good tomatoes are the thing that matters most when it comes to making bruschetta - the classic " +
						"Italian antipasto. It is such a simple preparation that paying attention to the little details matters.",
				},
				Category: models.Category{Value: "Appetizer"},
				Cuisine:  models.Cuisine{Value: "Easy"},
				Image: models.Image{
					Value: "https://images.101cookbooks.com/bruschetta-recipe-h1.jpg?w=1200&auto=format",
				},
				Ingredients: models.Ingredients{
					Values: []string{
						"3 fresh tomatoes, ripe",
						"A small handful of basil leaves",
						"1 teaspoon good-tasting white wine vinegar (or red/balsamic), or to taste",
						"1/4 teaspoon sea salt, or to taste",
						"4 tablespoons extra-virgin olive oil, plus more for serving",
						"3 - 4 sourdough or country-style bread slices (at least 1/2-inch thick)",
						"2 cloves garlic, peeled",
					},
				},
				Instructions: models.Instructions{
					Values: []string{
						"Rinse and dry your tomatoes. Halve each of them, use a finger to remove the seeds, and cut out the " +
							"cores. Roughly cut the tomatoes into 1/2-inch pieces and place in a medium bowl. " +
							"Tear the basil into small pieces, and add that as well. Add 2 tablespoons of the olive oil, " +
							"a small splash of vinegar, and a pinch of salt. Gently toss, taste, adjust if necessary, and set aside.",
						"Heat a grill or oven to medium-high. When it’s ready, use the remaining 2 tablespoons of the olive oil to " +
							"brush across the slices of bread. Grill or bake until well-toasted and golden brown with a hint of char. " +
							"Flipping when the first side is done. Remove from grilled and when cool enough to handle, " +
							"rub both sides of each slice of bread with garlic.",
						"Cut each slice of bread in half if you like, and top each segment with the tomato mixture. And a " +
							"finishing drizzle of olive oil is always nice.",
					},
				},
				NutritionSchema: models.NutritionSchema{
					Calories:      "233 kcal",
					Carbohydrates: "19 g",
					Cholesterol:   "",
					Fat:           "15 g",
					Protein:       "6 g",
					Sodium:        "287 mg",
					Servings:      "1",
				},
				Keywords:      models.Keywords{Values: "simple"},
				Yield:         models.Yield{Value: 4},
				PrepTime:      "PT15M",
				CookTime:      "PT5M",
				DatePublished: "2022-06-29T14:40:39+00:00",
				URL:           "https://www.101cookbooks.com/simple-bruschetta/",
			},
		},
	}
	for _, tc := range testcases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			defer func() {
				err := recover()
				if err != nil {
					t.Fatalf("panic while testing %s: %s", tc.name, err)
				}
			}()

			// updateHTMLFile(t, tc.in)
			actual := testFile(t, tc.in)
			// actual := testHTTP(t, tc.in)

			if !cmp.Equal(actual, tc.want) {
				t.Logf(cmp.Diff(actual, tc.want))
				t.Fatal()
			}
		})
	}
}

func testFile(t *testing.T, in string) models.RecipeSchema {
	t.Helper()
	t.Parallel()

	host := getHost(in)
	_, fileName, _, _ := runtime.Caller(0)
	f, err := os.Open(filepath.Join(path.Dir(fileName), "testdata", host+".html"))
	if err != nil {
		t.Fatalf("%s open file: %s", in, err)
	}
	defer f.Close()

	doc, err := goquery.NewDocumentFromReader(f)
	if err != nil {
		t.Fatalf("%s could not parse HTML: %s", host, err)
	}

	actual, err := scrapeWebsite(doc, getHost(in))
	if err != nil {
		t.Fatal(err)
	}

	if actual.URL == "" {
		actual.URL = in
	}

	actual.AtContext = atContext
	return actual
}

/*func updateHTMLFile(t *testing.T, url string) {
	t.Helper()
	t.Parallel()

	res, err := http.Get(url)
	if err != nil {
		t.Log("could not fetch url")
		return
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		t.Logf("got status code %d", res.StatusCode)
		return
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		t.Log(err)
		return
	}

	host := getHost(url)
	_, fileName, _, _ := runtime.Caller(0)
	filePath := filepath.Join(path.Dir(fileName), "testdata", host+".html")
	err = os.WriteFile(filePath, body, os.ModePerm)
	if err != nil {
		t.Log(err)
		return
	}
}

// TODO: Change package name to scraper_test
func testHTTP(t *testing.T, in string) models.RecipeSchema {
	t.Helper()
	rs, err := Scrape(in, &mockFiles{})
	if err != nil {
		t.Error(err)
	}
	return rs
}

type mockFiles struct {
	exportHitCount      int
	extractRecipesFunc  func(fileHeaders []*multipart.FileHeader) models.Recipes
	ReadTempFileFunc    func(name string) ([]byte, error)
	uploadImageHitCount int
	uploadImageFunc     func(rc io.ReadCloser) (uuid.UUID, error)
}

func (m *mockFiles) ExportCookbook(cookbook models.Cookbook, fileType models.FileType) (string, error) {
	m.exportHitCount++
	return cookbook.Title + fileType.Ext(), nil
}

func (m *mockFiles) ExportRecipes(recipes models.Recipes, _ models.FileType) (string, error) {
	var s string
	for _, recipe := range recipes {
		s += recipe.Name + "-"
	}
	m.exportHitCount++
	return s, nil
}

func (m *mockFiles) ExtractRecipes(fileHeaders []*multipart.FileHeader) models.Recipes {
	if m.extractRecipesFunc != nil {
		return m.extractRecipesFunc(fileHeaders)
	}
	return models.Recipes{}
}

func (m *mockFiles) ReadTempFile(name string) ([]byte, error) {
	if m.ReadTempFileFunc != nil {
		return m.ReadTempFileFunc(name)
	}
	return []byte(name), nil
}

func (m *mockFiles) UploadImage(rc io.ReadCloser) (uuid.UUID, error) {
	if m.uploadImageFunc != nil {
		return m.uploadImageFunc(rc)
	}
	m.uploadImageHitCount++
	return uuid.New(), nil
}*/
