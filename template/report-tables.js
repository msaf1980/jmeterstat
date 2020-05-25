var aggHeads = ["Mean", "Min", "Max", "P90", "P95", "P99"]

function DescriptionTables(data) {
	var table = $('<table></table>').attr({ border: 1 });
	var tr = $('<tr></tr>');
	$('<th></th>').text("Test").appendTo(tr);
	$('<th></th>').text("Started").appendTo(tr);
	$('<th></th>').text("Ended").appendTo(tr);
	table.append(tr);
	var tr = $('<tr></tr>');
	$('<td></td>').text(data.Name).appendTo(tr);
	$('<td></td>').text(new Date(data.Started).toISOString()).appendTo(tr);
	$('<td></td>').text(new Date(data.Ended).toISOString()).appendTo(tr);
	table.append(tr);
	table.appendTo("body");
}

// Create queries tables from data
function Tables(data, aggHeads, maxErrors) {
	var stats = data.Stat;
	// Loop over label
	var labels = []
	for (key in stats) {
		labels.push(key)
	}
	labels.sort()

	for (i in labels) {
		// {
		// "Started": ..,
		// "Ended": ..,
		// "Stat": {
		// 	 "label1": {
		//       ..

		var label = labels[i]

		var values = labelQueriesTableRows(label, stats[label], aggHeads);
		var foots = values.first
		var rows = values.second
		var columns = values.third
		createQueriesTable(i, label, columns, foots, rows);

		var values = labelErrorsTableRows(label, stats[label], maxErrors);
		var foots = values.first
		var rows = values.second
		var columns = values.third		
		createErrorsTable(i, label, columns, foots, rows);
	}
}

function createQueriesTable(tableId, label, columns, foots, rows) {
		var table = $('<table class="display" style="width:100%"></table>').attr({ id: "tableQuery" + tableId, border: 1 });

		var thead = $('<thead>');
		var tr = $('<tr>');
		for (i in columns[0]) {
			var attr = ""
			for (key in columns[0][i]) {
				if (key != 'text') {
					if (attr == "") {
						attr += " "
					}
					attr += (key + "=" + columns[0][i][key]);
				}
			}
			var th = $('<th' + attr + '></th>')
			th.text(columns[0][i].text)
			th.appendTo(tr);
		}
		thead.append(tr)
		
		var tr = $('<tr>');
		for (i in columns[1]) {
			$('<th></th>').text(columns[1][i]).appendTo(tr);
		}
		thead.append(tr);

		table.append(thead);

		var tfoot = $('<tfoot>');
		var tr = $('<tr>');
		for (i in foots) {
			$('<th></th>').text(foots[i]).appendTo(tr);
		}
		tfoot.append(tr);
		table.append(tfoot);

		var tbody = $('<tbody>');
		for (i in rows) {
			var tr = $('<tr>');
			for (j in rows[i]) {
				$('<td></td>').text(rows[i][j]).appendTo(tr);
			}
			tbody.append(tr);
		}				
		table.append(tbody);

		table.appendTo("body");	

		$(table).DataTable( {
			//data: dataSet,
			//searching: false,
			/*
			paging: false,
			*/
		} );
}

function createErrorsTable(tableId, label, columns, foots, rows) {
	var table = $('<table class="display" style="width:100%"></table>').attr({ id: "tableErr" + tableId, border: 1 });

	var thead = $('<thead>');

	
	var tr = $('<tr>');
	for (i in columns) {
		$('<th></th>').text(columns[i]).appendTo(tr);
	}
	thead.append(tr);

	table.append(thead);

	var tfoot = $('<tfoot>');
	var tr = $('<tr>');
	for (i in foots) {
		$('<th></th>').text(foots[i]).appendTo(tr);
	}
	tfoot.append(tr);
	table.append(tfoot);

	var tbody = $('<tbody>');
	for (i in rows) {
		var tr = $('<tr>');
		for (j in rows[i]) {
			$('<td></td>').text(rows[i][j]).appendTo(tr);
		}
		tbody.append(tr);
	}				
	table.append(tbody);

	table.appendTo("body");	

	$(table).DataTable( {
		//data: dataSet,
		//searching: false,
		/*
		paging: false,
		*/
	} );
}
