package layout

type Layout struct {
	Tables []Table
}

type Table struct {
	MaxOccupants int
}

var EdTechLayout = Layout{
	Tables: []Table{
		Table{
			MaxOccupants: 5,
		},
		Table{
			MaxOccupants: 4,
		},
		Table{
			MaxOccupants: 4,
		},
		Table{
			MaxOccupants: 4,
		},
		Table{
			MaxOccupants: 6,
		},
		Table{
			MaxOccupants: 4,
		},
		Table{
			MaxOccupants: 6,
		},
	},
}
