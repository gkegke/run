'use strict'

function populateRecent() {

    for (let i=1; i <= 5; i++) {

        let rec = localStorage.getItem("rec:" + i);

        if (rec === null) {
            localStorage.setItem("rec:" + i, "")
        } else {
            $("#recent").append("<div class=rec>" + rec + "</div>")
        }

    }

}

function updateRecentStorage(data) {

    let i = 4;

    while (i > 0) {

        let rec = localStorage.getItem("rec:" + i);

        if (rec === "" || i === 5) {
            i = i - 1;
            continue
        } else {
            localStorage.setItem("rec:" + (i + 1), rec)
        }

        i = i - 1;

    }

    localStorage.setItem("rec:" + 1, data)

}

async function success() {

    var i = 5; 

    while ($("#current").width() < $(window).width()) {


        var temp = $("#current").width();
        $("#current").width(temp + (40 * 5));

        await new Promise(r => setTimeout(r, 200));

        i = i - 1;
    }

    $("#current").width(0);

    $("#query").val("");

}
  
function update(data) {

    $("#recent .rec:eq(0)").before("<div class=rec>" + data + "</div>")

    var l = $(".rec").length

    if (l > 5) {

        $(".rec").eq(l - 1).remove();

    }
                
    updateRecentStorage(data);

    $("#recent").show(500);
    $(".com").show(500);

}
              
$(document).ready(function() {


    /* Runner code */
    populateRecent();

    // filter out commands
    $("#query").on("input", function() {

        $("#errors").hide();

        var q = $(this).val();

        $(".com").each(function(i, v) {

            if ($(this).text().includes(q)) {

                $(this).show(500);

            } else {

                $(this).hide(500);

            }

        })

        if (q.length === 0) {
            $("#recent").show(500);
        } else {
            $("#recent").hide(500);
        }

    });

    // attempt run command
    $("#queryForm").on("submit", function(e) {

        console.log("submitting...");

        var command = $("#query").val();

        $.ajax({
            type : "POST",
            url : "/command/",
            data : $("#queryForm").serialize(),
        }).done(function(data) {
                
            if (data.success === true) {

                $("#errors").text("");
                $("#errors").hide();
                update(command);
                success();

            } else {

                $("#errors").text(data.error);
                $("#errors").show(300);

            }

       });
 
        e.preventDefault();
    });


    /* Kanban-esque todo shit */

    $(".kb").on("change input", function(e) {

        var id = $(this).attr("id");
        var text = $(this).html();

        localStorage.setItem(id, text);

        var test = localStorage.getItem(id);

    })


});
            
// load kanban data stored in localstorage
(function () {

    for (let i=1; i < 5; i++) {

        const id = "kb" + String(i);
        let el = document.getElementById(id);
        let text = localStorage.getItem(id);

        el.innerHTML = text;
    }

})()
