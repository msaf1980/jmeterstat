// labelQueriesTableRows prepare table data for label from input data 
// return 
// first: foots[columns],
// second: rows[2][columns],
// third: array[queries][columns]
function labelQueriesTableRows(label, data, aggHeads) {
	let foots = []
	let rows = [];
	let columns1 = [];
	let columns2 = [];

	columns1.push({text:label});
	columns2.push("Request");

	// Summary
	var sumStat = data["SumStat"];
	// Executions: Count, Errors %
	columns1.push({text: "Executions", "colspan": 2});
	columns2.push("Samples", "Errors %");
	foots.push("Summary", sumStat["Count"], Math.round(10000 * ((sumStat["Count"] - sumStat["Success"]) / sumStat["Count"])) / 100);
	// Responce Times
	columns1.push({text: "Responce Times (ms)", "colspan": aggHeads.length})
	columns2.push.apply(columns2, aggHeads)
	for (let head of aggHeads) {
		foots.push(Math.round(100 *sumStat["Elapsed"][head]) / 100);
	}
	// Sent
	columns1.push({text: "Sent (bytes)", "colspan": aggHeads.length})
	columns2.push.apply(columns2, aggHeads)
	for (let head of aggHeads) {
		foots.push(Math.round(100 *sumStat["SentBytes"][head]) / 100);
	}
	// Received
	columns1.push({text: "Received (bytes)", "colspan": aggHeads.length})
	columns2.push.apply(columns2, aggHeads)
	for (let head of aggHeads) {
		foots.push(Math.round(100 *sumStat["Bytes"][head]) / 100);
	}

	for (url in data["Stat"]) {
		let row = [];
		var stat = data["Stat"][url];
		// Executions
		// Executions: Count, Errors %
		row.push(url, stat["Count"], Math.round(10000 * ((stat["Count"] - stat["Success"]) / stat["Count"])) / 100);
		// Responce Times
		for (let head of aggHeads) {
			row.push(Math.round(100 *stat["Elapsed"][head]) / 100);
		}
		// Sent
		for (let head of aggHeads) {
			row.push(Math.round(100 *stat["SentBytes"][head]) / 100);
		}
		// Received
		for (let head of aggHeads) {
			row.push(Math.round(100 *stat["Bytes"][head]) / 100);
		}
		rows.push(row);
	}

	var columns = [columns1, columns2];
	return {
		first: foots,
		second: rows,
		third: columns
	};
}

function queriesTables(data) {
	let aggHeads = ["Mean", "Min", "Max", "P90", "P95", "P99"]

	var body = d3.select("body");
	body.append("h3").text("Benchmark report (" + new Date(data.Started).toISOString() + " "
												+ new Date(data.Ended).toISOString() + ")")

	var stats = data.Stat;
	for (key in stats) {
		var table = body.append("table");
		table.append("th").text("Label");
		table.append("th").text(key).attr("colspan", "24");

		// Main Header
		var rowHead = table.append("tr")

		rowHead.append("th").text("").attr("width", "1")
		rowHead.append("th").text("Executions").attr("colspan", "2")
		rowHead.append("th").text("").attr("width", "1")
		rowHead.append("th").text("Responce Times (ms)").attr("colspan", "6")
		rowHead.append("th").text("").attr("width", "1")
		rowHead.append("th").text("Sent (bytes)").attr("colspan", "6")
		rowHead.append("th").text("").attr("width", "1")
		rowHead.append("th").text("Received (bytes)").attr("colspan", "6")
		rowHead.append("th").text("").attr("width", "1")

		var rowHead2 = table.append("tr")
		rowHead2.append("th").text("Requests")

		// Executions
		rowHead2.append("th").text("Count")
		rowHead2.append("th").text("Errors %")
		rowHead2.append("th").text("").attr("width", "1")
		// Responce Times
		for (let head of aggHeads) {
			rowHead2.append("th").text(head)
		}
		rowHead2.append("th").text("").attr("width", "1")
		// Sent
		for (let head of aggHeads) {
			rowHead2.append("th").text(head)
		}
		rowHead2.append("th").text("").attr("width", "1")
		// Received
		for (let head of aggHeads) {
			rowHead2.append("th").text(head)
		}
		rowHead2.append("th").text("").attr("width", "1")


		// Summary
		var sumHead = table.append("tr")
		var allStats = stats[key]
		var sumStat = allStats["SumStat"]
		sumHead.append("td").text("Summary")
		// Executions
		sumHead.append("td").text(sumStat["Count"])
		var KO = Math.round(10000 * ((sumStat["Count"] - sumStat["Success"]) / sumStat["Count"])) / 100
		sumHead.append("td").text(KO)
		sumHead.append("td").text("").attr("width", "1")
		// Responce Times
		for (let head of aggHeads) {
			sumHead.append("td").text(Math.round(100 *sumStat["Elapsed"][head]) / 100)
		}
		sumHead.append("td").text("").attr("width", "1")
		// Sent
		for (let head of aggHeads) {
			sumHead.append("td").text(Math.round(100 *sumStat["SentBytes"][head]) / 100)
		}
		sumHead.append("td").text("").attr("width", "1")
		// Received
		for (let head of aggHeads) {
			sumHead.append("td").text(Math.round(100 *sumStat["Bytes"][head]) / 100)
		}
		sumHead.append("td").text("").attr("width", "1")

		for (url in allStats["Stat"]) {
			var stat = allStats["Stat"][url]
			var head = table.append("tr")
			head.append("td").text(url).attr("class", "a")
			// Executions
			head.append("td").text(stat["Count"])
			var KO = Math.round(10000 * ((stat["Count"] - stat["Success"]) / stat["Count"])) / 100
			head.append("td").text(KO)
			head.append("td").text("").attr("width", "1")
			// Responce Times
			for (let h of aggHeads) {
				head.append("td").text(Math.round(100 *stat["Elapsed"][h]) / 100)
			}
			head.append("td").text("").attr("width", "1")
			// Sent
			for (let h of aggHeads) {
				head.append("td").text(Math.round(100 *stat["SentBytes"][h]) / 100)
			}
			head.append("td").text("").attr("width", "1")
			// Received
			for (let h of aggHeads) {
				head.append("td").text(Math.round(100 *stat["Bytes"][h]) / 100)
			}
			head.append("td").text("").attr("width", "1")
		}
	}
}

function compareErrorsDesc(a, b) {
	if (a.Value > b.Value)
		return -1;
	else if (a.Value < b.Value)
		return 1;
	else
		return 0;
}

// labelErrorsTableRows prepare table data for label from input data 
// return only maxErros columns
// first: foots[columns],
// second: rows[columns],
// third: array[queries][columns]
function labelErrorsTableRows(label, data, maxErros) {
	let foots = []
	let rows = [];
	let columns = [];

	columns.push(label);
	columns.push("Samples");
	for (i = 0; i < maxErros; i++) {
		columns.push("Errors");
		columns.push("Error");
	}

	// Summary
	var sumStat = data["SumStat"];

	let errors = [];
	errorCodes = sumStat["ErrorCodes"];
	for (key in errorCodes) {
		errors.push({ Name: key, Value: errorCodes[key]});
	}
	errors.sort(compareErrorsDesc);

	foots.push("Summary")
	foots.push(sumStat["Count"])
	for (i = 0; i < maxErros; i++) {
		if (i < errors.length) {
			foots.push(errors[i].Value);
			foots.push(errors[i].Name);
		} else {
			foots.push("");
			foots.push("");
		}
	}

	for (url in data["Stat"]) {
		let row = [];
		var stat = data["Stat"][url];
		errorCodes = stat["ErrorCodes"];

		let errors = [];
		errorCodes = stat["ErrorCodes"];
		for (key in errorCodes) {
			errors.push({ Name: key, Value: errorCodes[key]});
		}
		errors.sort(compareErrorsDesc);

		row.push(url);
		row.push(stat["Count"]);
		for (i = 0; i < maxErros; i++) {
			if (i < errors.length) {
				row.push(errors[i].Value);
				row.push(errors[i].Name);
			} else {
				row.push("");
				row.push("");
			}
		}
		rows.push(row);
	}

	return {
		first: foots,
		second: rows,
		third: columns
	};
}
