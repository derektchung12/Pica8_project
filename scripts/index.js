//This is for the tree
$(document).ready(function () {
    $('label.tree-toggler').click(function () {
        $(this).parent().children('ul.tree').toggle(300);
    });

    $('button').click(function () {
        $.get("http://localhost:8080/view/TestPage", function(data) {
  			alert(data);
	});    });


    $('.collapse').collapse()
});

