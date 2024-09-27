// https://www.mongodb.com/docs/mongodb-vscode/playgrounds/

const database = "Thunbo";

// Select the database to use.
use('Thunbo');


db.Weather.aggregate([
	{
		$project: {
			date: {
				$dateToParts : {date: "$timestamp"}
			},
			temp : 1
		}
	},{
		$group: {
			_id: {
				date: {
					year: "$date.year",
					month: "$date.month",
					day: "$date.day"
			}
		},
		avgTmp: {$avg: "$temp"}
	}
	}
])