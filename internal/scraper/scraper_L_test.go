package scraper_test

import (
	"testing"

	"github.com/reaper47/recipya/internal/models"
)

func TestScraper_L(t *testing.T) {
	testcases := []testcase{
		{
			name: "latelierderoxane.com",
			in:   "https://www.latelierderoxane.com/blog/recette-cake-marbre",
			want: models.RecipeSchema{
				AtContext:     "https://schema.org",
				AtType:        &models.SchemaType{Value: "Recipe"},
				Category:      &models.Category{Value: "uncategorized"},
				CookingMethod: &models.CookingMethod{},
				Cuisine:       &models.Cuisine{},
				CookTime:      "PT45M",
				DatePublished: "2023-11-12T07:54:40+01:00",
				Description: &models.Description{
					Value: "Une recette facile, rapide et adorée des enfants : le cake marbré moelleux au chocolat façon Savane. Un cake parfumé à la vanille et au chocolat. ",
				},
				Keywords: &models.Keywords{},
				Image:    &models.Image{Value: anUploadedImage.String()},
				Ingredients: &models.Ingredients{
					Values: []string{
						"3 œufs", "70 g de sucre", "70 g de beurre  fondu",
						"1 sachet de levure chimique", "250 g de farine", "150 g de lait",
						"150 g de chocolat noir fondu", "1 càc d'arôme ou poudre de vanille",
					},
				},
				Instructions: &models.Instructions{
					Values: []models.HowToItem{
						{Type: "HowToStep", Text: "Préchauffe le four à 165°."},
						{Type: "HowToStep", Text: "Commence par fouetter les œufs et le sucre, à l’aide de ton robot ou batteur électrique, pendant 10 minutes : ton mélange doit s’éclaircir et doubler de volume !"},
						{Type: "HowToStep", Text: "Ajoute le beurre fondu, la levure, la farine et fouette brièvement."},
						{Type: "HowToStep", Text: "Verse le lait et fouette jusqu’à l’obtention d’un mélange homogène."},
						{Type: "HowToStep", Text: "Sépare la préparation obtenue dans deux bols."},
						{Type: "HowToStep", Text: "Dans un des deux bols, ajoute l’arôme ou la poudre de vanille."},
						{Type: "HowToStep", Text: "Fais fondre ton chocolat, au bain-marie ou au micro-ondes et incorpore-le dans le second bol à l’aide d’une maryse."},
						{Type: "HowToStep", Text: "Récupère ton moule à cake et beurre-le généreusement."},
						{Type: "HowToStep", Text: "Verse, dans le fond de ton moule, la moitié de la pâte à la vanille puis la moitié de celle au chocolat."},
						{Type: "HowToStep", Text: "Répète l’opération une deuxième fois."},
						{Type: "HowToStep", Text: "Enfourne pendant 45 min."},
						{Type: "HowToStep", Text: "Tu peux vérifier la cuisson à l’aide d’un couteau, plante-le au centre de ton cake : ta lame doit ressortir sèche."},
						{Type: "HowToStep", Text: "À la sortie du four, laisse tiédir ton cake afin de faciliter son démoulage."},
						{Type: "HowToStep", Text: "À manger sans modération !"},
					},
				},
				NutritionSchema: &models.NutritionSchema{},
				Tools:           &models.Tools{Values: []models.HowToItem{}},
				PrepTime:        "PT20M",
				Yield:           &models.Yield{Value: 6},
				URL:             "https://www.latelierderoxane.com/blog/recette-cake-marbre",
			},
		},
		{
			name: "leanandgreenrecipes.net",
			in:   "https://leanandgreenrecipes.net/recipes/italian/spaghetti-squash-lasagna/",
			want: models.RecipeSchema{
				AtContext:     "https://schema.org",
				AtType:        &models.SchemaType{Value: "Recipe"},
				Category:      &models.Category{Value: `<a href="/recipes/category/main-course/" hreflang="en">Main course</a>`},
				CookTime:      "P0Y0M0DT0H64M0S",
				Cuisine:       &models.Cuisine{Value: `<a href="/recipes/italian/" hreflang="en">Italian</a>`},
				DatePublished: "2021-03-10T13:26:59-0600",
				Description: &models.Description{
					Value: "<p>If you're not familiar with spaghetti squash then it's time you get acquainted! This simple to prepare Spaghetti Squash Lasagna recipe will have you wondering why you haven't been eating spaghetti this way your entire life. Light, delicious and 100% on plan!</p>\n<p> </p>\n<p>Tip: If you do not like spicy heat you can reduce or omit the crushed red pepper flakes. This recipe call for several of the spices to be divided.</p>",
				},
				Image: &models.Image{Value: anUploadedImage.String()},
				Ingredients: &models.Ingredients{
					Values: []string{
						"1 medium Spaghetti Squash", "4 tsp Olive Oil", "1 tsp Salt",
						"1 tsp Black Pepper", "2 tsp Garlic", "1 lbs Lean Ground Turkey",
						"1(14.5oz.) can Diced Tomatoes", "1/2 tsp Basil", "1 tsp Whole Leaf Oregano",
						"1/2 cup Part-skim Ricotta", "1/2 cup 1% Cottage Cheese",
						"1 tsp Crushed Red Pepper Flakes", "1 cup Low-fat Mozzarela",
					},
				},
				Instructions: &models.Instructions{
					Values: []models.HowToItem{
						{Type: "HowToStep", Text: "Preheat oven to 400*F."},
						{Type: "HowToStep", Text: "Prepare the spaghetti squash. Cut in half"},
						{Type: "HowToStep", Text: "and remove seeds and pulp strands. Rub 1 teaspoon of olive oil into each half of squash and season each half with 1/4 teaspoon each of salt (optional) and pepper. Place each spaghetti squash half face down in a large baking dish and bake for 40 to 60 minutes"},
						{Type: "HowToStep", Text: "cook until the middle is tender and pulls apart easily."},
						{Type: "HowToStep", Text: "While the spaghetti squash is cooking in a large saucepan"},
						{Type: "HowToStep", Text: "saute garlic in remaining olive oil over a medium heat until fragrant. Add Turkey. Season with 1/4 teaspoon of each salt (optional) and pepper"},
						{Type: "HowToStep", Text: "and then cook until it has browned."},
						{Type: "HowToStep", Text: "Add tomatoes"},
						{Type: "HowToStep", Text: "onion powder"},
						{Type: "HowToStep", Text: "and 1/2 teaspoon each of basil and oregano. When the sauce starts to bubble"},
						{Type: "HowToStep", Text: "reduce heat to a simmer until it has thickened (about 3 to 4 minutes)."},
						{Type: "HowToStep", Text: "Combine the ricotta and cottage cheese into a medium bow. Season with crushed red pepper flakes (optional) and the remaining basil"},
						{Type: "HowToStep", Text: "oregano"},
						{Type: "HowToStep", Text: "salt (optional)"},
						{Type: "HowToStep", Text: "and pepper. Lightly mix until combined."},
						{Type: "HowToStep", Text: "When spaghetti squash is fully cooked"},
						{Type: "HowToStep", Text: "flip it in the baking dish so that it is now skin-side down. Lightly scrape flesh of squash with a fork to create spaghetti-like strands."},
						{Type: "HowToStep", Text: "Evenly divide the ricotta and cheese mixture between each squash half. Repeat with the meat sauce. Top each half with 1/2 cup of mozzarella cheese."},
						{Type: "HowToStep", Text: "Turn oven to broil"},
						{Type: "HowToStep", Text: "and cook for an additional 2 minutes"},
						{Type: "HowToStep", Text: "or until cheese is browned and bubbling. Watch that it does not burn due to different oven temperatures. Serve immediately"},
					},
				},
				Name: "Healthy Spaghetti Squash Lasagna Recipe",
				NutritionSchema: &models.NutritionSchema{
					Calories:      "337",
					Carbohydrates: "25",
					Fat:           "16",
					Protein:       "26",
				},
				PrepTime:  "P0Y0M0DT0H10M0S",
				TotalTime: "P0Y0M0DT0H74M0S",
				Yield:     &models.Yield{Value: 4},
				URL:       "https://leanandgreenrecipes.net/recipes/italian/spaghetti-squash-lasagna/",
			},
		},
		{
			name: "lecker.de",
			in:   "https://www.lecker.de/gemuesepfanne-mit-haehnchen-zuckerschoten-und-brokkoli-79685.html",
			want: models.RecipeSchema{
				AtContext:     "https://schema.org",
				AtType:        &models.SchemaType{Value: "Recipe"},
				Category:      &models.Category{Value: "Hauptgericht"},
				CookingMethod: &models.CookingMethod{},
				Cuisine:       &models.Cuisine{},
				CookTime:      "PT25M",
				DateModified:  "2023-02-23T13:12:02.767Z",
				Description: &models.Description{
					Value: "Unser beliebtes Rezept für Gemüsepfanne mit Hähnchen, Zuckerschoten und Brokkoli und mehr als 45.000 weitere kostenlose Rezepte auf LECKER.de.",
				},
				Keywords: &models.Keywords{
					Values: "Hähnchen,Geflügel,Fleisch,Zutaten,Mittagessen,Mahlzeit,Rezepte,Abendbrot,Low Carb,Gesundes Essen,Hauptgerichte,Menüs,Brokkoli,Kohl,Gemüse,Zuckerschoten,Gemüsepfanne,Pfannengerichte",
				},
				Image: &models.Image{Value: anUploadedImage.String()},
				Ingredients: &models.Ingredients{
					Values: []string{
						"1 Brokkoli", "150 g Zuckerschoten", "2 Lauchzwiebeln", "Salz",
						"4 kleine Hähnchenbrustfilets (à ca. 140 g)", "2 EL Sonnenblumenöl",
						"4 EL Sojasoße", "50 ml Gemüsebrühe", "6 Stiele Koriander",
						"1 EL Sesamsaat",
					},
				},
				Instructions: &models.Instructions{
					Values: []models.HowToItem{
						{Type: "HowToStep", Text: "Brokkoli putzen und in kleinen Röschen vom Strunk schneiden. Zuckerschoten putzen. Lauchzwiebeln putzen und in ca. 10 cm lange Stücke schneiden."},
						{Type: "HowToStep", Text: "Brokkoli, Zuckerschoten und Lauchzwiebeln in reichlich kochendem Salzwasser 2–3 Minuten garen. Herausnehmen und mit kaltem Wasser abspülen. Auf ein Sieb gießen und gut abtropfen lassen."},
						{Type: "HowToStep", Text: "Fleisch trocken tupfen und in Würfel (ca. 1,5 x 1,5 cm) schneiden. Öl in einer weiten Pfanne oder einem Wok erhitzen. Fleisch darin, unter Wenden, ca. 3 Minuten kräftig anbraten. Hitze reduzieren und weitere ca. 2 Minuten braten. Gemüse, Sojasoße und Brühe zufügen und 3–4 Minuten dünsten. Dabei gelegentlich wenden."},
						{Type: "HowToStep", Text: "Koriander waschen und trocken schütteln und, bis auf einige Blätter zum Garnieren, samt Stiel grob hacken. Sesam und Koriander zur Gemüsepfanne geben und untermischen. Mit Salz und Pfeffer abschmecken und mit Koriander garnieren."},
					},
				},
				Name: "Gemüsepfanne mit Hähnchen, Zuckerschoten und Brokkoli",
				NutritionSchema: &models.NutritionSchema{
					Calories:      "260",
					Carbohydrates: "7",
					Fat:           "7",
					Protein:       "38",
					Servings:      "1",
				},
				PrepTime:  "PT0M",
				Tools:     &models.Tools{Values: []models.HowToItem{}},
				TotalTime: "PT25M",
				Yield:     &models.Yield{Value: 4},
				URL:       "https://www.lecker.de/gemuesepfanne-mit-haehnchen-zuckerschoten-und-brokkoli-79685.html",
			},
		},
		{
			name: "lecremedelacrumb.com",
			in:   "https://www.lecremedelacrumb.com/instant-pot-pot-roast-potatoes/",
			want: models.RecipeSchema{
				AtContext:     atContext,
				AtType:        &models.SchemaType{Value: "Recipe"},
				Category:      &models.Category{Value: "Main Course"},
				CookTime:      "PT80M",
				Cuisine:       &models.Cuisine{Value: "American"},
				DatePublished: "2018-01-19T11:11:53+00:00",
				Description: &models.Description{
					Value: "Juicy and tender instant pot pot roast and potatoes with gravy makes the perfect family-friendly dinner. This easy one pot dinner recipe will please even the picky eaters!",
				},
				Keywords: &models.Keywords{
					Values: "instant pot pot roast, pot roast and potatoes",
				},
				Image: &models.Image{Value: anUploadedImage.String()},
				Ingredients: &models.Ingredients{
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
				Instructions: &models.Instructions{
					Values: []models.HowToItem{
						{Type: "HowToStep", Text: "Turn on your instant pot and set it to \"saute\". In a small bowl stir together salt, pepper, garlic powder, onion powder, and smoked paprika. Rub mixture all over the roast to coat all sides."},
						{Type: "HowToStep", Text: "Drizzle oil in instant pot, wait about 30 seconds, then use tongs to place roast in the pot. Do not move it for 3-4 minutes until well-seared and browned. Use tongs to turn the roast onto another side for 3-4 minutes, repeating until all sides are browned."},
						{Type: "HowToStep", Text: "Switch instant pot to \"pressure cook\" on high and set to 60-80 minutes (60 for a 3 pound roast, 80 for a 5 pound roast. see notes if using baby carrots). Add potatoes, onions, and carrots to pot (just arrange them around the roast) and pour beef broth and worcestershire sauce over everything. Place lid on the pot and turn to locked position. Make sure the vent is set to the sealed position."},
						{Type: "HowToStep", Text: "When the cooking time is up, do a natural release for 10 minutes (don't touch anything on the pot, just let it de-pressurize on it's own for 10 minutes). After 10 minutes, turn vent to the venting release position and allow all of the steam to vent and the float valve to drop down before removing the lid."},
						{Type: "HowToStep", Text: "Transfer the roast, potatoes, onions, and carrots to a platter and shred the roast with 2 forks into chunks. Use a handheld strainer to scoop out bits from the broth in the pot. Set instant pot to \"soup\" setting. Whisk together the water and corn starch. Once broth is boiling, stir in corn starch mixture until the gravy thickens. Add salt, pepper, and garlic powder to taste."},
						{Type: "HowToStep", Text: "Serve gravy poured over roast and veggies and garnish with fresh thyme or parsley if desired."},
					},
				},
				Name: "Instant Pot Pot Roast Recipe",
				NutritionSchema: &models.NutritionSchema{
					Calories:      "133",
					Carbohydrates: "23",
					Fat:           "3",
					Fiber:         "3",
					Protein:       "4",
					SaturatedFat:  "1",
					Servings:      "1",
					Sodium:        "1087",
					Sugar:         "5",
				},
				PrepTime:  "PT20M",
				TotalTime: "PT100M",
				Yield:     &models.Yield{Value: 6},
				URL:       "https://www.lecremedelacrumb.com/instant-pot-pot-roast-potatoes/",
			},
		},
		/*{
			name: "lekkerensimpel.com",
			in:   "https://www.lekkerensimpel.com/gougeres/",
			want: models.RecipeSchema{
				AtContext: atContext,
				Type:    &models.SchemaType{Value: "Recipe"},
				Name:      "Gougères",
				Category:  &models.Category{Value: "Snacks"},
				Yield:     &models.Yield{Value: 4},
				URL:       "https://www.lekkerensimpel.com/gougeres/",
				PrepTime:  "PT20M",
				CookTime:  "PT25M",
				Description: &models.Description{
					Value: "Vandaag een receptje uit de Franse keuken, namelijk deze gougères. Gougères zijn een soort hartige kaassoesjes, erg lekker! We hadden een tijdje geleden een stuk gruyère kaas gekocht bij de kaasboer, meer uit nood want parmezaanse had hij op dat moment even niet. Inmiddels lag de kaas al een tijdje in de koelkast en moesten we er toch echt wat mee gaan doen. Iemand tipte ons dat we echt eens gougères moesten maken en eerlijk gezegd hadden we er nooit eerder van gehoord. Een kleine speurtocht bracht ons uiteindelijk bij een recept van ‘The Guardian – How to make the perfect gougères‘. We zijn ermee aan de slag gegaan en zie hier het resultaat! \n\nNog meer van dit soort lekkere snacks en borrelhapjes vind je in onze categorie tapas recepten en tussen de high-tea recepten.",
				},
				Ingredients: &models.Ingredients{
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
				Instructions: &models.Instructions{
					Values: []models.HowToItem{
						"Verwarm de oven voor op 200 graden. Klop vervolgens 2 eieren los in een beker. Doe het water, de boter en een snuf zout in een pan en laat de boter al roerend smelten. Zet het ‘vuur’ laag en doe de bloem erbij. Zelf doen we de bloem eerst door een zeef zodat er geen kleine klontjes meer inzitten. Roer de bloem door het botermengsel totdat er een soort deeg ontstaat. Haal de pan van het vuur en mix het deeg, bij voorkeur met een mixer, een minuut of 3-4. Voeg dan de helft van het losgeklopte ei toe, even goed mengen en dan kan de andere helft erbij. Mix daarna nog de nootmuskaat en geraspte gruyère door het deeg. Bekleed een bakplaat met bakpapier. Schep met twee lepels kleine bolletjes deeg op de bakplaat of gebruik hiervoor een spuitzak. Smeer de bovenkant in met een beetje losgeklopt ei, bestrooi met nog geraspte gruyère kaas en dan kan de bakplaat de oven in voor 20-25 minuten. Eet smakelijk!",
						"Bewaar dit recept op Pinterest !",
					},
				},
				DatePublished: "2021-09-28T04:00:00+00:00",
				DateModified:  "2021-09-21T08:22:19+00:00",
			},
		},*/
		{
			name: "leukerecepten.nl",
			in:   "https://www.leukerecepten.nl/recepten/pita-tandoori",
			want: models.RecipeSchema{
				AtContext:     "https://schema.org",
				AtType:        &models.SchemaType{Value: "Recipe"},
				Category:      &models.Category{Value: "Hoofdgerechten"},
				CookingMethod: &models.CookingMethod{},
				Cuisine:       &models.Cuisine{},
				CookTime:      "PT0M",
				DatePublished: "2024-03-17T22:00:52+01:00",
				Description: &models.Description{
					Value: "Dit recept voor pita tandoori is ideaal als je weinig tijd hebt en toch iets lekkers wil maken. Met kip, een kruidige tandoori marinade en verfrissende yoghurtsaus in een knapperig pitabroodje.",
				},
				Keywords: &models.Keywords{},
				Image:    &models.Image{Value: anUploadedImage.String()},
				Ingredients: &models.Ingredients{
					Values: []string{
						"250 gr kip (vega)", "5 grote pitabroodjes", "1 komkommer",
						"1 eetlepel tandoori kruiden (pasta)", "50 gr yoghurt", "olie om te bakken",
						"1 Bosje verse munt", "120 gr yoghurt", "1 Snuf peper", "1 limoen",
					},
				},
				Instructions: &models.Instructions{
					Values: []models.HowToItem{
						{Type: "HowToStep", Text: "Meng de stukjes kip met de tandoori kruiden (dit kunnen droge kruiden zijn of een kruidenpasta) en de yoghurt in een kom. Dek af en laat ondertussen in de koelkast staan."},
						{Type: "HowToStep", Text: "Ga verder met het afbakken van de pitabroodjes. Snijd de komkommer in plakjes."},
						{Type: "HowToStep", Text: "Maak de saus: hak de munt fijn en rasp de limoen en pers het sap uit. Meng de munt en limoen met de yoghurt voor de saus en breng op smaak met een snufje peper."},
						{Type: "HowToStep", Text: "Verhit een beetje olie in een koekenpan en bak de stukjes kip tandoori gaar."},
						{Type: "HowToStep", Text: "Neem een pitabroodje en vul deze met plakjes komkommer, de kip tandoori en wat van de frisse yoghurtsaus."},
					},
				},
				Name:            "Pita tandoori",
				NutritionSchema: &models.NutritionSchema{},
				PrepTime:        "PT0H30M",
				Tools:           &models.Tools{Values: []models.HowToItem{}},
				TotalTime:       "PT0H30M",
				Yield:           &models.Yield{Value: 5},
				URL:             "https://www.leukerecepten.nl/recepten/pita-tandoori",
			},
		},
		{
			name: "lifestyleofafoodie.com",
			in:   "https://lifestyleofafoodie.com/chick-fil-a-peppermint-milkshake/",
			want: models.RecipeSchema{
				AtContext:     "https://schema.org",
				AtType:        &models.SchemaType{Value: "Recipe"},
				Category:      &models.Category{Value: "Drinks"},
				Cuisine:       &models.Cuisine{Value: "American"},
				DatePublished: "2023-11-18T18:26:00+00:00",
				Description: &models.Description{
					Value: "These Chick-Fil-A peppermint milkshakes are an easy and delicious way to enjoy your favorite holiday drink year round!",
				},
				Keywords: &models.Keywords{Values: "chick fil a peppermint milkshake, peppermint milkshake"},
				Image:    &models.Image{Value: anUploadedImage.String()},
				Ingredients: &models.Ingredients{
					Values: []string{
						"3 cups vanilla ice cream", "1/4 cup milk (whole or 2% for creamier results)",
						"1/4 teaspoon peppermint extract (adjust to taste)",
						"1/3 cup crushed peppermint candies (plus extra for garnish)",
						"1/4 cup chocolate chips",
						"2-3 drops of red food coloring (optional, for a pinkish color)",
						"Whipped cream (for topping)", "Maraschino cherries (for garnish)",
					},
				},
				Instructions: &models.Instructions{
					Values: []models.HowToItem{
						{Type: "HowToStep", Text: "In a blender, combine the vanilla ice cream, milk, peppermint extract, crushed peppermint candies (roughly crushed if using a high-speed blender), chocolate chips, and red food coloring (if using)."},
						{Type: "HowToStep", Text: "Blend the mixture until it&#39;s smooth and all the ingredients are well combined. You can adjust the milk or ice cream to achieve your desired milkshake consistency."},
						{Type: "HowToStep", Text: "Taste the milkshake and adjust the amount of peppermint extract and crushed candies to your liking. Some people prefer a stronger peppermint flavor, while others like it milder."},
						{Type: "HowToStep", Text: "Pour the milkshakes into your serving glasses, top them with whipped cream and a maraschino cherry and enjoy."},
					},
				},
				Name: "Chick Fil A Peppermint Milkshake",
				NutritionSchema: &models.NutritionSchema{
					Calories:       "649",
					Carbohydrates:  "90",
					Cholesterol:    "88",
					Fat:            "28",
					Fiber:          "1",
					Protein:        "8",
					SaturatedFat:   "17",
					Servings:       "1",
					Sodium:         "171",
					Sugar:          "76",
					UnsaturatedFat: "7",
				},
				PrepTime:  "PT1M",
				TotalTime: "PT1M",
				Yield:     &models.Yield{Value: 2},
				URL:       "https://lifestyleofafoodie.com/chick-fil-a-peppermint-milkshake/",
			},
		},
		{
			name: "lidl-kochen.de",
			in:   "https://www.lidl-kochen.de/rezeptwelt/schweinemedaillons-mit-ofenkartoffeln-butterbohnen-und-rosmarinbroeseln-147914",
			want: models.RecipeSchema{
				AtContext: atContext,
				AtType:    &models.SchemaType{Value: "Recipe"},
				Category:  &models.Category{Value: "Mittagessen, Abendessen"},
				Cuisine:   &models.Cuisine{Value: "Deutschland"},
				Description: &models.Description{
					Value: "Rezept für Schweinemedaillons mit Ofenkartoffeln, Butterbohnen und Rosmarinbröseln » Über 562x nachgekocht » 40min Zubereitung » 10 Zutaten » 558 kcal/Portion",
				},
				Keywords: &models.Keywords{
					Values: "Bohnen, Buschbohnen, Kartoffeln, Schwein, Schweinelende, Schweinefilet, einfach, lecker, leicht, Mittagessen, Abendessen, NährwertKompass 7-8, Deutschland, Gäste, Familie, Hauptspeise, Fleisch, Gemüse, Schweinefleisch, Einfaches Mittagessen, Einfaches Abendessen, Einfache Familienrezepte, Abendessen für Gäste, Familien Mittagessen",
				},
				Image: &models.Image{Value: anUploadedImage.String()},
				Ingredients: &models.Ingredients{
					Values: []string{
						"Salz Prise", "Schweinefilet 600 g", "Kartoffeln, vorw. festk. 1 kg",
						"Rosmarin, frisch 10 g", "Olivenöl 5 EL", "Buschbohnen 400 g",
						"Schalotten 2 St.", "Pfeffer, schwarz Prise", "Paniermehl 3 EL", "Butter 3 EL",
					},
				},
				Instructions: &models.Instructions{
					Values: []models.HowToItem{
						{Type: "HowToStep", Text: "Ofen auf 220 °C (Umluft) vorheizen. In einem Topf ca. 1 l Salzwasser zugedeckt aufkochen. Schweinefilets waschen, trocken tupfen, in Medaillons schneiden und zum Temperieren beiseitelegen. Kartoffeln gründlich waschen und längs vierteln. Rosmarin waschen, trocken schütteln, Nadeln von den Stielen streifen und sehr fein hacken."},
						{Type: "HowToStep", Text: "Kartoffeln in einer Schüssel mit zwei Drittel vom Rosmarin, 1 EL Olivenöl und Salz vermengen. Mit der Hautseite nach unten auf ein Backblech legen und im Ofen ca. 25 Min. rösten."},
						{Type: "HowToStep", Text: "Inzwischen Bohnen waschen, trocken tupfen, Enden dünn abschneiden und im kochenden Salzwasser ca. 4 Min. kochen. Bohnen anschließend in ein Sieb abgießen, unter kaltem Wasser abspülen und abtropfen lassen."},
						{Type: "HowToStep", Text: "Derweil Schalotten halbieren, schälen und fein würfeln. In einer Pfanne 1 EL Olivenöl auf mittlerer bis hoher Stufe erhitzen, Medaillons salzen und im Öl von jeder Seite ca. 3 Min. anbraten. Fleisch pfeffern und in Alufolie gewickelt bis zum Anrichten durchziehen lassen. Pfanne nicht säubern."},
						{Type: "HowToStep", Text: "2 EL Olivenöl und Paniermehl zum Bratensatz in die Pfanne geben und Brösel auf mittlerer Stufe ca. 2 Min. goldgelb rösten. Restlichen Rosmarin und Salz zugeben, kurz mitrösten und auf einem Teller beiseitestellen. Pfanne säubern."},
						{Type: "HowToStep", Text: "1 EL Olivenöl in der gesäuberten Pfanne auf mittlerer Stufe erhitzen und Schalotten weitere ca. 2 Min. mitbraten. Abgetropfte Bohnen und 3 EL Butter zugeben, durchmischen, ca. 1 Min. erwärmen und mit Salz und Pfeffer abschmecken. Kartoffeln aus dem Ofen nehmen und mit Schweinemedaillons, Bohnen und Bröseln auf Tellern anrichten und servieren.\r\n\r\nGuten Appetit!"},
					},
				},
				Name: "Schweinemedaillons mit Ofenkartoffeln, Butterbohnen und Rosmarinbröseln",
				NutritionSchema: &models.NutritionSchema{
					Calories: "558",
				},
				PrepTime: "PT40M",
				Tools: &models.Tools{
					Values: []models.HowToItem{
						{Type: "HowToTool", Text: "große Schüssel", Quantity: 1},
						{Type: "HowToTool", Text: "mittlerer Topf", Quantity: 1},
						{Type: "HowToTool", Text: "mittlere Pfanne", Quantity: 1},
						{Type: "HowToTool", Text: "Backblech mit Backpapier", Quantity: 1},
					},
				},
				Yield: &models.Yield{Value: 4},
				URL:   "https://www.lidl-kochen.de/rezeptwelt/schweinemedaillons-mit-ofenkartoffeln-butterbohnen-und-rosmarinbroeseln-147914",
			},
		},
		{
			name: "littlespicejar.com",
			in:   "https://littlespicejar.com/starbucks-pumpkin-loaf/",
			want: models.RecipeSchema{
				AtContext:     atContext,
				AtType:        &models.SchemaType{Value: "Recipe"},
				Category:      &models.Category{Value: "Bread & Baking"},
				CookingMethod: &models.CookingMethod{},
				CookTime:      "PT55M",
				Cuisine:       &models.Cuisine{Value: "American"},
				DatePublished: "2021-11-09",
				Description: &models.Description{
					Value: "Learn how to make an easy delicious copycat Starbucks Pumpkin Loaf right at home! This pumpkin bread is studded with roasted pepitas and loaded with spices and so much pumpkin goodness!",
				},
				Keywords: &models.Keywords{Values: ""},
				Image:    &models.Image{Value: anUploadedImage.String()},
				Ingredients: &models.Ingredients{
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
				Instructions: &models.Instructions{
					Values: []models.HowToItem{
						{Type: "HowToStep", Text: "PREP: Position a rack in the center of the oven and preheat the oven to 350°F. Spray two 8 ½ x 4 ½ (or 9x5) bread pans with cooking spray, you can also line with parchment if you’d like; set aside for now."},
						{Type: "HowToStep", Text: "DRY INGREDIENTS: Add the dry ingredients: flour, baking soda baking powder, all the spices, and salt to a medium bowl. Whisk to combine; set aside for now."},
						{Type: "HowToStep", Text: "WET INGREDIENTS: Add the granulated sugar, brown sugar, and oil to a large bowl. Whisk to combine, then add the pumpkin puree, eggs, vanilla, and orange zest and combine to whisk until all the eggs have been incorporated into the wet batter. Don't be alarmed if the batter splits or curdles! It's totally fine!"},
						{Type: "HowToStep", Text: "BREAD BATTER: Add the dry ingredients into the wet ingredients in two batches, stirring just long enough so each batch of flour is incorporated. Do not over-mix or you’ll end up with dry bread!"},
						{Type: "HowToStep", Text: "BAKE: Divide the batter into the to pans, taking care to only fill each pan about ¾ of the way full. The bread will rise significantly! Smooth out the batter then sprinkle with the pepitas. Bake the bread for 52-62 minutes or until a toothpick inserted in the center of the loaf comes out clean. Cool the pans on a wire baking rack for at least 10 minutes before removing from the pan and allowing the bread to cool further."},
					},
				},
				Name:            "The Best Starbucks Pumpkin Loaf Recipe (Copycat)",
				NutritionSchema: &models.NutritionSchema{},
				PrepTime:        "PT15M",
				Tools:           &models.Tools{Values: []models.HowToItem{}},
				TotalTime:       "PT1H10M",
				Yield:           &models.Yield{Value: 2},
				URL:             "https://littlespicejar.com/starbucks-pumpkin-loaf/",
			},
		},
		{
			name: "livelytable.com",
			in:   "https://livelytable.com/bbq-ribs-on-the-charcoal-grill/",
			want: models.RecipeSchema{
				AtContext:     atContext,
				AtType:        &models.SchemaType{Value: "Recipe"},
				Category:      &models.Category{Value: "main dish"},
				CookTime:      "PT2H30M",
				CookingMethod: &models.CookingMethod{Value: "grilled"},
				Cuisine:       &models.Cuisine{Value: "BBQ"},
				DatePublished: "2019-07-25",
				Description: &models.Description{
					Value: "Nothing says summer like grilled BBQ ribs! These baby back ribs on the charcoal grill are simple, delicious, and sure to please a crowd! (gluten-free, dairy-free, nut-free)",
				},
				Keywords: &models.Keywords{Values: "BBQ ribs, ribs on the charcoal grill"},
				Image:    &models.Image{Value: anUploadedImage.String()},
				Ingredients: &models.Ingredients{
					Values: []string{
						"1 rack baby back pork ribs",
						"1/3 cup BBQ spice rub",
						"water",
						"BBQ sauce of choice (optional)",
					},
				},
				Instructions: &models.Instructions{
					Values: []models.HowToItem{
						{Type: "HowToStep", Text: "Prepare fire in the charcoal grill. Remove the grates, place a pile of charcoal on one side of the grill only. On the other side, place a small foil pan filled with water. Start the fire and return the grates to the grill. Let the grill get to a low temperature (about 275°F.) You may also add pieces of wood to the charcoal for a more smoky flavor."},
						{Type: "HowToStep", Text: "While the fire is heating, prepare ribs. Turn ribs over so that the bone side is facing up. Remove the membrane along the back by sliding a dull knife (such as a butter knife) under the membrane along the last bone until you get under the membrane. Hold on tight, and pull it until the whole thing is removed from the rack of ribs."},
						{Type: "HowToStep", Text: "Rub ribs all over with spice rub. Once fire is ready, place the ribs on indirect heat - the side of the grill that has the foil pan. Cover and cook about 2 hours, watching to make sure the fire is maintained at a steady low temperature, adding charcoal as needed, and rotating the rack of ribs roughly every 30 minutes so that different edges of the rack are turned toward the hot side."},
						{Type: "HowToStep", Text: "After 1 1/2 to 2 hours, remove ribs and wrap in foil. Return to the grill for another 30 minutes or so."},
						{Type: "HowToStep", Text: "When ribs are done, you can either remove them from the foil and place back on the grill, meat side down, for a little char, or place them meat side up and brush with barbecue sauce in layers, waiting about 5 minutes between layers. Or simply remove them from the grill to a cutting board, slice, and serve!"},
					},
				},
				Name: "BBQ Ribs on the Charcoal Grill",
				NutritionSchema: &models.NutritionSchema{
					Calories:      "416",
					Carbohydrates: "8.9",
					Cholesterol:   "122.9",
					Fat:           "26.3",
					Fiber:         "0.8",
					Protein:       "36.1",
					SaturatedFat:  "9.1",
					Servings:      "2",
					Sodium:        "512.8",
					Sugar:         "5.1",
					TransFat:      "0.2",
				},
				PrepTime:  "PT10M",
				Tools:     &models.Tools{Values: []models.HowToItem{}},
				TotalTime: "PT2H40M",
				Yield:     &models.Yield{Value: 1},
				URL:       "https://livelytable.com/bbq-ribs-on-the-charcoal-grill/",
			},
		},
		{
			name: "livingthegreenlife.com",
			in:   "https://livingthegreenlife.com/recepten/vegan-tikka-masala-met-rijst/",
			want: models.RecipeSchema{
				AtContext:     atContext,
				AtType:        &models.SchemaType{Value: "Recipe"},
				Category:      &models.Category{Value: "Diner"},
				CookTime:      "PT30M",
				DateModified:  "2021-04-20T07:30:51+02:00",
				DatePublished: "2020-03-30T08:06:37+02:00",
				Description: &models.Description{
					Value: "Tikka masala is een mix van geroosterde kruiden en specerijen. Zowel India als Engeland beweren deze heerlijke kruidenmix uitgevonden te hebben. Wat voor ons het belangrijkste is, is dat we een vegan variant hebben gemaakt. Mega simpel en minstens net zo lekker als het origineel!",
				},
				Keywords: &models.Keywords{Values: "Living the Green Life"},
				Image:    &models.Image{Value: anUploadedImage.String()},
				Ingredients: &models.Ingredients{
					Values: []string{
						"0,75 st bloemkool",
						"3 el plantaardige olie",
						"3 tl garam masala",
						"200 g zilvervliesrijst",
						"1 st ui",
						"3 teentje(s) knoflook",
						"1 st  gemberwortel (1 x 1 cm)",
						"1 tl kurkuma",
						"1 tl gemalen komijn",
						"1 tl gerookte-paprikapoeder",
						"1 blik(ken) tomatenblokjes",
						"1 snuf(jes) zout",
						"1 blik(ken) linzen",
						"2 el kokosyoghurt",
						"1 el ahornsiroop",
						"1 handje(s) cashewnoten",
					},
				},
				Instructions: &models.Instructions{
					Values: []models.HowToItem{
						{Type: "HowToStep", Text: "Verwarm de oven voor op 180 ˚C en bekleed een bakplaat met bakpapier."},
						{Type: "HowToStep", Text: "Snijd de bloemkool in kleine roosjes. Leg deze op de met bakpapier beklede bakplaat en schep ze om met 2 eetlepels olie en garam masala. Rooster de roosjes 25-30 minuten in de voorverwarmde oven."},
						{Type: "HowToStep", Text: "Kook intussen de rijst volgens de instructies op de verpakking."},
						{Type: "HowToStep", Text: "Snijd de ui, de knoflook en de gember fijn. Verwarm 1 eetlepel olie in de pan en fruit hierin de ui, knoflook en gember. Voeg de overige kruiden en de tomatenblokjes toe. Laat 15 minuten sudderen op laag vuur en breng op smaak met zout."},
						{Type: "HowToStep", Text: "Voeg de geroosterde bloemkool en de linzen toe aan de saus en laat het nog 10 minuten sudderen."},
						{Type: "HowToStep", Text: "Verdeel de rijst over 2 diepe borden en schep de saus ernaast. Maak swirls van kokosyoghurt en ahornsiroop, en voeg eventueel wat gehakte cashewnoten toe voor een extra bite."},
					},
				},
				Name:            "Vegan tikka masala met rijst",
				NutritionSchema: &models.NutritionSchema{},
				Yield:           &models.Yield{Value: 1},
				URL:             "https://livingthegreenlife.com/recepten/vegan-tikka-masala-met-rijst/",
			},
		},
		{
			name: "lovingitvegan.com",
			in:   "https://lovingitvegan.com/vegan-buffalo-chicken-dip/",
			want: models.RecipeSchema{
				AtContext:     atContext,
				AtType:        &models.SchemaType{Value: "Recipe"},
				Category:      &models.Category{Value: "Appetizer"},
				CookTime:      "PT20M",
				Cuisine:       &models.Cuisine{Value: "American"},
				DatePublished: "2022-01-21T14:31:28+00:00",
				Description: &models.Description{
					Value: "This baked vegan buffalo chicken dip is rich, creamy and so cheesy. It&#39;s packed with spicy flavor and makes the perfect crowd pleasing party dip.",
				},
				Keywords: &models.Keywords{
					Values: "vegan buffalo chicken dip, vegan buffalo dip",
				},
				Image: &models.Image{Value: anUploadedImage.String()},
				Ingredients: &models.Ingredients{
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
				Instructions: &models.Instructions{
					Values: []models.HowToItem{
						{Type: "HowToStep", Text: "Soak the cashews. Place the cashews into a bowl. Pour boiling hot water from the kettle over the top of the cashews to submerge them. Leave the cashews to soak for 1 hour and then drain and rinse."},
						{Type: "HowToStep", Text: "Preheat the oven to 375°F (190°C)."},
						{Type: "HowToStep", Text: "Add the soaked cashews, lemon juice, coconut cream, distilled white vinegar, salt, onion powder, vegan chicken spice, vegan buffalo sauce and nutritional yeast to the blender and blend until smooth."},
						{Type: "HowToStep", Text: "Transfer the blended mix to a mixing bowl."},
						{Type: "HowToStep", Text: "Add chopped artichoke hearts and chopped spring onions and gently fold them in."},
						{Type: "HowToStep", Text: "Transfer to an oven safe 9-inch round dish and smooth down."},
						{Type: "HowToStep", Text: "Bake for 20 minutes until lightly browned on top."},
						{Type: "HowToStep", Text: "Serve topped with chopped spring onions with tortilla chips, crackers, breads or veggies for dipping."},
					},
				},
				Name: "Vegan Buffalo Chicken Dip",
				NutritionSchema: &models.NutritionSchema{
					Calories:       "214",
					Carbohydrates:  "13",
					Fat:            "16",
					Fiber:          "3",
					Protein:        "8",
					SaturatedFat:   "7",
					Servings:       "1",
					Sodium:         "938",
					Sugar:          "2",
					UnsaturatedFat: "8",
				},
				PrepTime:  "PT10M",
				TotalTime: "PT90M",
				Yield:     &models.Yield{Value: 8},
				URL:       "https://lovingitvegan.com/vegan-buffalo-chicken-dip/",
			},
		},
	}
	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			test(t, tc)
		})
	}
}
