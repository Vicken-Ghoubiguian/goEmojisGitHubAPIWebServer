$(function() {

    //
    $("[data-toggle='tooltip']").tooltip({
        
        html: true,
        tooltipClass: "main_tooltip-styling",
        show: {
            effect: "explode",
            delay: 250
        },
        hide: {
            effect: "explode",
            delay: 250
        }
    });
});