package scraper_test

import (
	"github.com/reaper47/recipya/internal/models"
	"testing"
)

func TestScraper_O(t *testing.T) {
	testcases := []testcase{
		{
			name: "ohsheglows.com",
			in:   "https://ohsheglows.com/2017/11/23/bread-free-stuffing-balls/",
			want: models.RecipeSchema{
				AtContext:     atContext,
				AtType:        &models.SchemaType{Value: "Recipe"},
				Category:      &models.Category{Value: "Vegan"},
				CookingMethod: &models.CookingMethod{},
				CookTime:      "PT23M",
				Cuisine:       &models.Cuisine{Value: "Canadian"},
				DatePublished: "2018-12-21 23:37:54",
				Description: &models.Description{
					Value: `My recipe tester Nicole likes to call these “bread-free stuffing balls," and I think I would have to agree! These festive bites have all the flavours of traditional stuffing, but they’re protein-packed, bite-sized, and gluten-free as well. This is a new and improved version of my popular Lentil Mushroom Walnut Balls recipe. I've streamlined the procedure and provided a make-ahead version in the Tips below. This recipe moves quickly using quite a few components, so my advice is to gather all of the ingredients and do as much prep as you can before you begin. If you aren't a cranberry sauce fan, my Vegan Mushroom Gravy is a nice option too!`,
				},
				Keywords: &models.Keywords{
					Values: "Vegan, Gluten-Free, Soy-Free, Budget Friendly, Freezer Friendly, Kid Friendly, Make-Ahead, Party Favourite",
				},
				Image: &models.Image{Value: anUploadedImage.String()},
				Ingredients: &models.Ingredients{
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
				Instructions: &models.Instructions{
					Values: []models.HowToItem{
						{Type: "HowToStep", Text: "Preheat the oven to 350°F (180°C) and line a baking sheet with parchment paper."},
						{Type: "HowToStep", Text: "Add the oil to a large pot and turn heat to medium. Finely chop the mushrooms until they’re roughly the size of peas. Add chopped mushrooms to the pot along with minced garlic and a pinch of salt. Stir until combined. Sauté for about 6 to 8 minutes, until the water from the mushrooms cooks off, reducing heat to low if necessary to prevent burning."},
						{Type: "HowToStep", Text: "Meanwhile, tear the kale into large pieces and place into a food processor. Pulse (do not process) the kale until finely chopped (pieces roughly the size of almonds), being careful not to overprocess it. Remove and place into a bowl for later."},
						{Type: "HowToStep", Text: "To the processor (no need to clean it out!), add the rolled oats. Process the oats until they’re finely chopped and resemble coarse flour, about 30 seconds."},
						{Type: "HowToStep", Text: "Add the drained lentils and walnuts to the processor bowl with the oat flour. Pulse the mixture, stopping to check on it every few pulses, until it’s coarsely chopped. Be sure not to overprocess it into a paste as you still want a lot of texture and crunchy walnut pieces. Set aside."},
						{Type: "HowToStep", Text: "To the pot with the mushrooms and garlic, add the herbs and sauté for 30 seconds until fragrant. Stir in the kale and chopped dried cranberries, then turn off the heat."},
						{Type: "HowToStep", Text: "Stir the flax and water together in a small cup (no need to let it sit)."},
						{Type: "HowToStep", Text: "Now add all of the food processor contents, vinegar, and flax mixture to the pot. Stir until thoroughly combined. The dough should be heavy and dense. Add salt and pepper to taste."},
						{Type: "HowToStep", Text: "With lightly wet hands, shape and roll about 14 to 15 balls, roughly 3 to 4 tablespoons of dough each. Place them on the prepared baking sheet about two inches apart."},
						{Type: "HowToStep", Text: "Bake for 22 to 24 minutes, until golden on the bottom and firm to touch. Remove and let cool for 5 minutes."},
						{Type: "HowToStep", Text: "While the Bread-Free Stuffing Balls are baking, make the Cranberry-Pear Sauce. Add the cranberries, pear, maple syrup, and salt to a medium pot. Bring to a low boil over high heat and then reduce to medium. Simmer uncovered for 10 to 20 minutes until thickened. Use a potato masher to mash up the pear near the end of cooking, if desired."},
						{Type: "HowToStep", Text: "Leftover balls can be refrigerated in an airtight container for a few days. To reheat, add oil to a skillet and fry over medium heat, tossing occasionally, until heated through."},
					},
				},
				Name: "Bread-Free Stuffing Balls",
				NutritionSchema: &models.NutritionSchema{
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
				PrepTime:     "PT30M",
				ThumbnailURL: &models.ThumbnailURL{},
				Tools:        &models.Tools{Values: []models.HowToItem{}},
				TotalTime:    "PT53M",
				Yield:        &models.Yield{Value: 14},
				URL:          "https://ohsheglows.com/2017/11/23/bread-free-stuffing-balls/",
				Video:        &models.Videos{},
			},
		},
		{
			name: "omnivorescookbook.com",
			in:   "https://omnivorescookbook.com/chinese-scallion-pancakes/",
			want: models.RecipeSchema{
				AtContext:     "https://schema.org",
				AtType:        &models.SchemaType{Value: "Recipe"},
				Category:      &models.Category{Value: "Appetizer"},
				CookTime:      "PT30M",
				Cuisine:       &models.Cuisine{Value: "Chinese"},
				DatePublished: "2020-08-25T09:07:00+00:00",
				Description: &models.Description{
					Value: "Super crispy and flaky on the outside and slightly chewy inside, my dim sum favorite, scallion pancakes, make a wonderful snack that you’ll love! {Vegan}",
				},
				Keywords: &models.Keywords{Values: "restaurant-style"},
				Image:    &models.Image{Value: anUploadedImage.String()},
				Ingredients: &models.Ingredients{
					Values: []string{
						"2 cups (300 g) all-purpose flour", "1/2 teaspoon salt",
						"1/2 cup boiling water", "1/4 cup cool water",
						"1/4 cup + 2 tablespoons (50 g) all-purpose flour",
						"1/4 cup peanut oil ((or your favorite oil like olive oil, melted coconut oil, melted butter, or melted chicken fat etc.))",
						"3/4 teaspoon salt",
						"8 to 10 green onions (, split down the middle and chopped (yield 1 cup))",
						"Vegetable oil for pan frying", "1 1/2 tablespoon soy sauce",
						"1 tablespoon Chinkiang vinegar ((or rice vinegar))", "1/4 teaspoon sugar",
					},
				},
				Instructions: &models.Instructions{
					Values: []models.HowToItem{
						{Type: "HowToStep", Text: "OPTION 1 - USING YOUR HANDS: Combine the flour and salt in a big bowl. Stir to mix well. Slowly drizzle in the hot water while mixing it with a pair of chopsticks (or a fork), until the water is fully absorbed. Slowly drizzle in the cool water, continuing to stir until many dough flakes form. Press the dough together, and try to combine the wet dough with the dry flour. Drizzle in a little extra water if there is any dry flour remaining. Knead until a firm ball is formed, about 5 minutes. Cover and let rest for 20 minutes."},
						{Type: "HowToStep", Text: "OPTION 2 - USING A MIXER: Combine the flour and salt in the mixer bowl with the dough hook attachment. Turn it to the mix setting and slowly drizzle in the hot water followed by the cool water. After a minute of mixing, drizzling in a little extra water if there is any dry flour remaining. Turn to setting 4 and knead until a ball of dough is formed, about 3 minutes. Cover and let rest for 20 minutes."},
						{Type: "HowToStep", Text: "While the dough is resting, combine the flour, oil, and salt in a small bowl. Mix until a smooth paste is formed."},
						{Type: "HowToStep", Text: "If making the dipping sauce, combine all the ingredients with 1 tablespoon of water in a small bowl and mix until the sugar is dissolved."},
						{Type: "HowToStep", Text: "When the dough is done resting, knead for another minute until it is smooth. Cut the dough into 6 even pieces, about 76 grams per piece. Form the dough into balls using your hands."},
						{Type: "HowToStep", Text: "Work on one dough ball at a time, covering the rest with plastic wrap to prevent drying out. Roll each dough ball into a thin rectangle with a rolling pin, aiming for around 6x10” (15x25 cm). Lift and turn the dough regularly as you roll it to prevent sticking. Spoon about 1 tablespoon of the filling onto the dough and spread it evenly with the back of a spoon, leaving about 1” (2.5 cm) on top and the left end without the filling. Sprinkle 2 heaping tablespoons of green onions onto the paste, concentrating most of it towards one side of the length and one side of the width, creating a loose L shape (see the blog post above for the step-by-step pictures)."},
						{Type: "HowToStep", Text: "Begin rolling the dough from the longer side of the dough where the green onions are concentrated, until you have one long tube. Lightly flatten the tube using your hand. Pinch the side with more green onion to seal it. Take the sealed side and gently pull and roll it towards the unsealed side, using your hand to smooth out the dough to push any large air bubbles out. Tuck the loose end under the rolled bun. Gently press down on the round disk with your hand to seal the pancake."},
						{Type: "HowToStep", Text: "Set the formed pancake aside and cover it with plastic wrap. Repeat steps 4 and 5 until each pancake is ready."},
						{Type: "HowToStep", Text: "Roll each prepared pancake into a circle, about 7-8” (17-20 cm) wide. Flip and move it as you do to prevent sticking. Don’t worry if air bubbles burst through or some green onions fall out."},
						{Type: "HowToStep", Text: "If you plan to store the pancakes and cook them later, place each pancake onto a piece of parchment paper and stack them. Transfer the pancakes to a large ziplock bag, squeeze out as much air as possible, and seal the bag. You can freeze the pancakes for up to 3 months."},
						{Type: "HowToStep", Text: "Heat a 9” (23 cm) cast iron pan (or a nonstick skillet) over medium-high heat and add enough oil to fully coat the bottom. Once the oil is hot, add a pancake. Use a pair of chopsticks (or a spatula) to swirl the pancake around to spread the oil and prevent sticking, for a few seconds. Cover the pan and turn the heat down to medium. Let the pancake cook, covered, for 1 minute. Remove the lid and flip the pancake, cover, and cook for another minute. Remove the lid. Use a spatula to press the pancake, to ensure even browning. Continue to cook, flipping regularly, until both sides are crisp and browned, about 3 minutes. Turn to medium-low heat if the pan gets too hot. Transfer the pancake to a cooling rack or a cutting board to cool. Repeat to cook all the pancakes you plan to serve."},
						{Type: "HowToStep", Text: "Once the pancakes are slightly cooled enough to handle, cut into 6 pieces and transfer to a serving platter. Serve hot with the dipping sauce as an appetizer."},
						{Type: "HowToStep", Text: "Let the frozen pancakes thaw for 10 minutes then proceed from step 8."},
						{Type: "HowToStep", Text: "Place the leftover pancakes in a ziplock bag by stacking them together. Store in the fridge up to 3 days. To reheat, add a pancake to a pan and heat over medium heat, cook, flip occasionally until the pancakes turn warm throughout."},
					},
				},
				Name: "Chinese Scallion Pancakes (葱油饼)",
				NutritionSchema: &models.NutritionSchema{
					Calories:      "202 kcal",
					Carbohydrates: "25.7 g",
					Fat:           "9.3 g",
					Fiber:         "1.1 g",
					Protein:       "3.6 g",
					SaturatedFat:  "1.6 g",
					Servings:      "1",
					Sodium:        "246 mg",
					Sugar:         "0.3 g",
				},
				PrepTime:  "PT10M",
				TotalTime: "PT60M",
				Yield:     &models.Yield{Value: 12},
				URL:       "https://omnivorescookbook.com/chinese-scallion-pancakes/",
			},
		},
		{
			name: "onceuponachef.com",
			in:   "https://www.onceuponachef.com/recipes/perfect-basmati-rice.html",
			want: models.RecipeSchema{
				AtContext:     atContext,
				AtType:        &models.SchemaType{Value: "Recipe"},
				Category:      &models.Category{Value: "Recipes"},
				CookingMethod: &models.CookingMethod{},
				CookTime:      "PT0M",
				Cuisine:       &models.Cuisine{Value: "Indian"},
				DatePublished: "2013-12-05T16:29:22-05:00",
				Description: &models.Description{
					Value: "This recipe makes tender and fluffy basmati rice every time.",
				},
				Keywords: &models.Keywords{Values: "All Seasons, Rice"},
				Image:    &models.Image{Value: anUploadedImage.String()},
				Ingredients: &models.Ingredients{
					Values: []string{
						"1 cup basmati rice (preferably imported from India or Pakistan)",
						"1¾ cups water",
						"1½ tablespoons unsalted butter",
						"½ teaspoon salt",
					},
				},
				Instructions: &models.Instructions{
					Values: []models.HowToItem{
						{Type: "HowToStep", Text: "Place the rice in a fine mesh strainer. Place under cold running water, swishing the rice with your hand, for 1 to 2 minutes to release excess starch. (Alternatively, place the rice in a medium bowl and add enough water to cover by 2 inches. Using your hands, gently swish the grains to release any excess starch. Carefully pour off the water, leaving the rice in the bowl. Repeat four times, or until the water runs almost clear. Use a fine mesh strainer to drain the rice.)"},
						{Type: "HowToStep", Text: "In a medium pot, bring the rice, water, butter, and salt to a boil. Cover the pot with a tight fitting lid, then turn the heat down to a simmer and cook for 15 to 20 minutes, until all of the water is absorbed and the rice is tender. If the rice is still too firm, add a few more tablespoons of water and continue cooking for a few minutes more. Remove the pan from the heat and allow it to sit covered for 5 minutes. Fluff the rice with a fork and serve."},
						{Type: "HowToStep", Text: "<strong>Freezer-Friendly Instructions:</strong> This rice can be frozen in an airtight container for up to 3 months. (Putting it in a flat layer in sealable plastic bags works well as it will take up less space in the freezer.) No need to thaw before reheating; remove it from the freezer and reheat in the microwave with 1 to 2 tablespoons of water."},
					},
				},
				Name: "Perfect Basmati Rice",
				NutritionSchema: &models.NutritionSchema{
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
				PrepTime:     "PT0M",
				ThumbnailURL: &models.ThumbnailURL{},
				Tools:        &models.Tools{Values: []models.HowToItem{}},
				TotalTime:    "PT25M",
				Yield:        &models.Yield{Value: 4},
				URL:          "https://www.onceuponachef.com/recipes/perfect-basmati-rice.html",
				Video:        &models.Videos{},
			},
		},
		{
			name: "ohmyveggies.com",
			in:   "https://ohmyveggies.com/korean-barbecue-jackfruit-sandwiches/",
			want: models.RecipeSchema{
				AtContext:     "https://schema.org",
				AtType:        &models.SchemaType{Value: "Recipe"},
				Category:      &models.Category{Value: "Main Course"},
				CookTime:      "PT20M",
				CookingMethod: nil,
				Cuisine:       &models.Cuisine{Value: "American"},
				DateCreated:   "",
				DateModified:  "",
				DatePublished: "2015-06-03T09:59:28+00:00",
				Description: &models.Description{
					Value: "Pulled jackfruit is simmered in spicy Korean barbecue sauce, topped with creamy tahini slaw and stuffed into buns in these flavor-packed pulled jackfruit sandwiches.",
				},
				Keywords: &models.Keywords{
					Values: "Jackfruit Sandwiches, Korean Barbecue Jackfruit Sandwiches, Korean Barbecue Jackfruit Sandwiches with Tahini Slaw, pulled jackfruit sandwiches",
				},
				Image: &models.Image{Value: anUploadedImage.String()},
				Ingredients: &models.Ingredients{
					Values: []string{
						"40 oz young green jackfruit (canned in brine or water)", "3/4 cup water",
						"1/2 cup soy sauce", "1/3 cup light brown sugar", "1/4 cup tomato paste",
						"2 tablespoons rice vinegar", "2 tablespoons sriracha sauce (or to taste)",
						"1 1/2 teaspoons fresh grated ginger",
						"1 teaspoon sesame oil",
						"1 tablespoon grape seed oil (or another neutral flavor vegetable oil)",
						"1 onion (sliced and separated into strips)",
						"2 garlic cloves (minced)",
						"1/4 cup tahini",
						"2 tablespoons rice vinegar",
						"2 teaspoons soy sauce",
						"2 teaspoons maple syrup (or agave syrup)",
						"2-4 tablespoons water",
						"3 cups shredded cabbage",
						"1 red bell pepper (cut into very thin strips)",
						"2 tablespoons finely chopped fresh cilantro (packed)",
						"2 scallions (chopped)",
						"4 burger buns or sandwich rolls (sliced open)",
					},
				},
				Instructions: &models.Instructions{
					Values: []models.HowToItem{
						{Text: "Drain and rinse jackfruit well. Blot pieces with a paper towel to remove excess water.", Type: "HowToStep"},
						{Text: "In small bowl, stir together water, soy sauce, brown sugar, tomato paste, rice vinegar, sriracha and ginger.", Type: "HowToStep"},
						{Text: "Place oil in a medium saucepan and place over medium heat. Add onion and garlic. Sauté 3 minutes. Add jackfruit. Continue to sauté, breaking up pieces of jackfruit with a fork or spatula, until onion is softened and jackfruit is lightly browned in spots, about 3 minutes more. Add soy sauce mixture and bring to a simmer. Lower heat and allow to simmer 15 minutes, until sauce is thick and jackfruit is tender, stirring frequently. While simmering, pull pieces of jackfruit apart using a fork and add more water if the mixture becomes too thick.", Type: "HowToStep"},
						{Text: "Stir tahini, rice vinegar, soy sauce and agave or maple syrup together in medium bowl. Thin with water until mixture is easy to stir, but still thick and creamy. Add cabbage, bell pepper, cilantro and scallions. Toss well to coat.", Type: "HowToStep"},
						{Text: "Divide Korean barbecue jackfruit among buns and top with creamy tahini slaw.", Type: "HowToStep"},
					},
				},
				Name: "Korean Barbecue Pulled Jackfruit Sandwiches",
				NutritionSchema: &models.NutritionSchema{
					Calories:       "421 kcal",
					Carbohydrates:  "63 g",
					Fat:            "14 g",
					Fiber:          "5 g",
					Protein:        "13 g",
					SaturatedFat:   "2 g",
					Servings:       "1",
					Sodium:         "2332 mg",
					Sugar:          "31 g",
					UnsaturatedFat: "11 g",
				},
				PrepTime:  "PT10M",
				TotalTime: "PT30M",
				Yield:     &models.Yield{Value: 4},
				URL:       "https://ohmyveggies.com/korean-barbecue-jackfruit-sandwiches/",
			},
		},
		{
			name: "owen-han.com",
			in:   "https://www.owen-han.com/recipes/cobb-chicken-sandwich",
			want: models.RecipeSchema{
				AtContext:     "https://schema.org",
				AtType:        &models.SchemaType{Value: "Recipe"},
				Category:      &models.Category{Value: "Chicken"},
				CookingMethod: &models.CookingMethod{},
				Cuisine:       &models.Cuisine{},
				Description: &models.Description{
					Value: "INGREDIENTS    6 tablespoons olive oil    3 tablespoons red wine vinegar    1 teaspoon Dijon mustard    1 teaspoon honey    1 teaspoon Italian seasoning    ¼ teaspoon kosher salt, plus more for seasoning    ¼ teaspoon freshly ground black pepper, plus more for seasoning    2 (4- or 5-ounce) thin-cut",
				},
				Keywords: &models.Keywords{},
				Image:    &models.Image{Value: anUploadedImage.String()},
				Ingredients: &models.Ingredients{
					Values: []string{
						"6 tablespoons olive oil", "3 tablespoons red wine vinegar",
						"1 teaspoon Dijon mustard", "1 teaspoon honey", "1 teaspoon Italian seasoning",
						"¼ teaspoon kosher salt, plus more for seasoning",
						"¼ teaspoon freshly ground black pepper, plus more for seasoning",
						"2 (4- or 5-ounce) thin-cut chicken breasts",
						"4 lettuce leaves, preferably Boston or Bibb", "2 hard-boiled eggs, sliced",
						"1 small avocado, sliced", "1 ripe vine or Roma tomato, sliced into rounds",
						"½ medium red onion, sliced into rings", "4 pieces of cooked bacon",
						"2 soft, oblong hoagie rolls, split", "4 tablespoons crumbled blue cheese",
						"2 tablespoons mayonnaise\u00a0",
						"1 tbsp mustard",
						"1 clove garlic minced",
						"1 tbsp chives",
					},
				},
				Instructions: &models.Instructions{
					Values: []models.HowToItem{
						{Type: "HowToStep", Text: "Combine the olive oil, vinegar, mustard, honey, Italian seasoning, salt, and pepper in a jar and shake until creamy."},
						{Type: "HowToStep", Text: "Place the chicken breasts in a bowl or zip-loc resealable bag, add half (about ⅓ cup) of the marinade, cover (or seal), and refrigerate for at least 30 minutes and up to 4 hours."},
						{Type: "HowToStep", Text: "Separately, in a small bowl, mix together the mayonnaise, blue cheese, mustard, garlic, and chives. Season to taste with salt and pepper."},
						{Type: "HowToStep", Text: "Heat a grill or grill pan over high heat. Remove the chicken from the marinade, shake off any excess, and season with more salt and pepper. Grill the chicken until done and grill marks form, 2–3 minutes per side."},
						{Type: "HowToStep", Text: "Brush both sides of the rolls with some of the remaining dressing and grill, face down, until toasted, 1–2 minutes. Spread the blue cheese mayonnaise on both sides of the rolls. Working with the bottom bun, layer on a piece of lettuce, chicken, bacon, some tomato slices, red onion, avocado slices, and a sliced hard-boiled egg. Drizzle some more dressing on the toppings. Close the sandwiches and enjoy!"},
					},
				},
				Name:            "Cobb Chicken Sandwich",
				NutritionSchema: &models.NutritionSchema{},
				ThumbnailURL:    &models.ThumbnailURL{},
				Tools:           &models.Tools{Values: []models.HowToItem{}},
				Yield:           &models.Yield{Value: 1},
				URL:             "https://www.owen-han.com/recipes/cobb-chicken-sandwich",
				Video:           &models.Videos{},
			},
		},
	}
	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			test(t, tc)
		})
	}
}
