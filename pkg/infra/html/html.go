package html

func NavBody() string {

	return `
	Navigate:
	<div class="container">
	<ul>
		<li><a href="all">All properties</a></li>
		<li><a href="add_form">Add a property</a></li>
	</ul>
	</div>
	`
}

func PropertyFormBody() string {
	return `
	Add a property:
	<div class="container">
	<form action="/homes/html/add">
	<label for="name">Name:</label><br>
	<input type="text" id="name" name="name"><br>

	<label for="postcode">Postcode:</label><br>
	<input type="text" id="postcode" name="postcode"><br>

	<label for="street">Street:</label><br>
	<input type="text" id="street" name="street"><br>

	<label for="type">Type:</label><br>
	<input type="text" id="type" name="type"><br>

	<label for="bedrooms">Bedrooms:</label><br>
	<input type="number" id="bedrooms" name="bedrooms"><br>

	<label for="inhabited">Inhabited:</label><br>
	<input type="checkbox" id="inhabited" name="inhabited"><br>

	<label for="safe">Safe:</label><br>
	<input type="checkbox" id="safe" name="safe"><br>

	<label for="owner-name">Owner Name:</label><br>
	<input type="text" id="owner-name" name="owner-name"><br>

	<label for="owner_contact">Owner Contact Info:</label><br>
	<input type="text" id="owner_contact" name="owner_contact"><br>

	<label for="last_checked">Last Checked:</label><br>
	<input type="datetime-local" id="last_checked" name="last_checked"><br>

	<input type="submit" value="Submit">
	</form>
	</div>
	`
}
