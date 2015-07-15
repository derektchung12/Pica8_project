//This is for the tree
$(document).ready(function () {
    $('label.tree-toggler').click(function () {
        $(this).parent().children('ul.tree').toggle(300);
    });

    


    $('#firstbutton').click(function () {
        var ip1 = document.getElementById("ipAddressForm").value;
        var ip2 = document.getElementById("ipAddressForm2").value;
        var sum = null;

        

        var table = document.getElementById("SiteTable")
        var row = table.insertRow(1);
        var cell1 = row.insertCell(0);
        var cell2 = row.insertCell(1);
        var cell3 = row.insertCell(2);
        
        
        
        //cell3.innerHTML = sum;
        $.post("http://localhost:8080/view/setIp1", {ip1value: ip1}, function(data){
            //alert(data)
            cell1.innerHTML = data;
        });
        //alert(ip1)

        $.post("http://localhost:8080/view/setIp2", {ip1value: ip2}, function(data){
            //alert(data)
            cell2.innerHTML = data;
        });
        $.get("http://localhost:8080/view/sumIp", function(data) {
    			//document.getElementById("sum").innerHTML = data;
            cell3.innerHTML = data;
    	});

        //alert(sum)

        

    });

/*
<script>
function myFunction() {
    var table = document.getElementById("myTable");
    var row = table.insertRow(table.rows.length);
    var cell1 = row.insertCell(0);
    var cell2 = row.insertCell(1);
    cell1.innerHTML = "NEW CELL1";
    cell2.innerHTML = "NEW CELL2";
}
</script>
*/



	// $('#ip1button').click(function () {
 //        /*$.get("http://localhost:8080/view/setIp1", function(data) {
 //  			alert(data);
	// 	});*/

 //        var ip1 = document.getElementById("ipAddressForm").value;
 //        //alert(ip1)

 //        $.post("http://localhost:8080/view/setIp1", {ip1value: ip1}, function(data){
 //            alert(data)
 //        });
 //    });

 //    $('#ip2button').click(function () {
 //        var ip2 = document.getElementById("ipAddressForm2").value;
 //        //alert(ip1)

 //        $.post("http://localhost:8080/view/setIp2", {ip1value: ip2}, function(data){
 //            alert(data)
 //        });
 //    });

	/*function myFunction() {
    var x = document.getElementById("frm1");
    var text = "";
    var i;
    for (i = 0; i < x.length ;i++) {
        text += x.elements[i].value + "<br>";
    }
    document.getElementById("demo").innerHTML = text;
}	*/





    //$('.collapse').collapse()
});

