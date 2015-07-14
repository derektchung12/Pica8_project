//This is for the tree
$(document).ready(function () {
    $('label.tree-toggler').click(function () {
        $(this).parent().children('ul.tree').toggle(300);
    });

    $('#firstbutton').click(function () {
        $.get("http://localhost:8080/view/sumIp", function(data) {
  			alert(data);
		});
	});

	$('#ip1button').click(function () {
        $.get("http://localhost:8080/view/setIp1", function(data) {
  			alert(data);
		});
    });


	/*function myFunction() {
    var x = document.getElementById("frm1");
    var text = "";
    var i;
    for (i = 0; i < x.length ;i++) {
        text += x.elements[i].value + "<br>";
    }
    document.getElementById("demo").innerHTML = text;
}	*/





    $('.collapse').collapse()
});

