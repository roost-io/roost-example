URL = "http://localhost"

function updateResult() {
    var xhttp = new XMLHttpRequest();
    xhttp.onreadystatechange = function() {
        if (this.readyState == 4 && this.status == 200) {
            data = JSON.parse(this.responseText)

            if(data.status.code == 204) {
                console.log(data.status.message)
                out = '<h1>'+ data.status.message +'</h1>';
                document.getElementById("error_msg").innerHTML = out;
                document.getElementById("main_content").style.display = "none";
                document.getElementById("error_msg").style.display = "block";
            }
            else if (data.status.code == 200) {
                
                data.results.sort((a, b) => {
                    if(a.total_votes === b.total_votes) {
                        // If two elements have same total_votes, then the one who has larger rating.average wins
                        return b.total_votes - a.total_votes
                    } else {
                        // If two elements have different total_votes, then the one who has larger total_votes wins
                        return b.total_votes - a.total_votes;
                    }
                });
            
                let total = 0;
                let out = '';
                data.results.forEach(item => {
                    total += item.total_votes
                });
            
                console.log("total_votes: " + total)
            
                data.results.forEach(item => {
                    percent = Math.round(((100 * item.total_votes) / total) / 2);
                    if(percent < 1) {
                        percent = 0
                    }
                    out += '<div class="result-card">'+
                            '    <div class="progressbar" style="width: '+ percent +'%;"></div>'+
                            '    <div class="logo">'+ item.name + '</div>  '+
                            '    <div class="percentage">'+(percent * 2)+' %</div>  '+
                            '</div>';
                });
                document.getElementById("candidates").innerHTML = out;
                document.getElementById("main_content").style.display = "block";
                document.getElementById("error_msg").style.display = "none";

            }
        }
    };
    xhttp.open("GET", URL, true);
    xhttp.send();
    
}

// setInterval( () => {
//     updateResult()
// } , 5000)

updateResult()

// var data = {
//     candidates: [
//        {
//            ID: 0,
//            name: "kind",
//            total_votes: 1000,
//        },
//        {
//            ID: 1,
//            name: "Roost",
//            total_votes: 200,
//        },
//        {
//             ID: 2,
//             name: "Docker-Desktop",
//             total_votes: 300,
//         },
//         {
//              ID: 3,
//              name: "K38",
//              total_votes: 200,
//          },
//          {
//             ID: 4,
//             name: "minikube",
//             total_votes: 300,
//         }
//    ],
//    status:{
//       code: 200,
//       message: "vote saved successfully"
//    }
// }