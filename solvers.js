var solverInfo = {
	"Conduction": {
		"FDM": {
			"1D": ["Dirichlet", "Neumann", "Robin"],
			"2D": ["Dirichlet", "Neumann", "Robin"]
		},
		"FVM": {
			"1D": ["Dirichlet", "Neumann", "Robin"],
			"2D": ["Dirichlet", "Neumann", "Robin"]
		}
	},
	"Radiation": {
        "FDM": {
			"1D": ["Dirichlet", "Neumann", "Robin"],
			"2D": ["Dirichlet", "Neumann", "Robin"]
		},
		"FVM": {
			"1D": ["Dirichlet", "Neumann", "Robin"],
			"2D": ["Dirichlet", "Neumann", "Robin"]
		}
	}
}


window.onload = function () {
	
	//Get html elements
	var heatSel = document.getElementById("heatSel");
	var finiteSel = document.getElementById("finiteSel");	
	var dimnSel = document.getElementById("dimnSel");
	var bcSel = document.getElementById("bcSel");
	
	//Load countries
	for (var country in solverInfo) {
		heatSel.options[heatSel.options.length] = new Option(country, country);
	}
	
	//County Changed
	heatSel.onchange = function () {
		 
		 finiteSel.length = 1; // remove all options bar first
		 dimnSel.length = 1; // remove all options bar first
		 bcSel.length = 1; // remove all options bar first
		 
		 if (this.selectedIndex < 1)
			 return; // done
		 
		 for (var state in solverInfo[this.value]) {
			 finiteSel.options[finiteSel.options.length] = new Option(state, state);
		 }
	}
	
	//State Changed
	finiteSel.onchange = function () {		 
		 
		 dimnSel.length = 1; // remove all options bar first
		 bcSel.length = 1; // remove all options bar first
		 
		 if (this.selectedIndex < 1)
			 return; // done
		 
		 for (var city in solverInfo[heatSel.value][this.value]) {
			 dimnSel.options[dimnSel.options.length] = new Option(city, city);
		 }
	}
	
	//City Changed
	dimnSel.onchange = function () {
		bcSel.length = 1; // remove all options bar first
		
		if (this.selectedIndex < 1)
			return; // done
		
		var zips = solverInfo[heatSel.value][finiteSel.value][this.value];
		for (var i = 0; i < zips.length; i++) {
			bcSel.options[bcSel.options.length] = new Option(zips[i], zips[i]);
		}
	}	
}
