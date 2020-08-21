URL = "http://localhost"

var data = {
    results: [
       {
           id: 0,
           name: "Roost",
           total_votes: 1000,
       },
       {
           id: 1,
           name: "Docker",
           total_votes: 200,
       },
       {
            id: 2,
            name: "Minikube",
            total_votes: 300,
        },
        {
             id: 3,
             name: "K3d",
             total_votes: 200,
         },
         {
            id: 4,
            name: "kind",
            total_votes: 300,
        }
   ],
   status:{
      code: 200,
      message: "vote saved successfully"
   }
}

function updateResult() {
    // var xhttp = new XMLHttpRequest();
    // xhttp.onreadystatechange = function() {
    //     if (this.readyState == 4 && this.status == 200) {
            // data = JSON.parse(this.responseText)

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
                    src = "";
                    if(percent < 1) {
                        percent = 0
                    }
                    
                    if(item.id == 0) {
                        src = "./assets/roost.png"
                    }
                    else if(item.id == 1) {
                        src = "./assets/docker.png"
                    }
                    else if(item.id == 2) {
                        src = "./assets/minikube.png"
                    }
                    else if(item.id == 3) {
                        src = "./assets/k3d.svg"
                    }
                    else if(item.id == 4) {
                        src = "./assets/kind.png"
                    }
                    console.log("src: ", src)

                    out +=  '<div class="card">'+
                            '    <div class="cardBackgroundContainer">'+
                            '        <div class="cardBackground"></div>'+
                            '        <div class="cardBackgroundImage">'+
                            '            <img'+
                            '            src= '+ src +
                            '            width="150px"'+
                            '            height="150px"'+
                            '            class="image"'+
                            '            />'+
                            '        </div>'+
                            '    </div>'+
                            '    <div class="cardContent">'+
                            '        '+ item.name +
                            '        <div class="progressbar_back">'+
                            '           <div class="progressbar_front" style="width: '+ percent +'%;"></div>'+
                                + percent +
                            '        %</div>'+
                            '    </div>+'+
                            '</div>';
                        
                    // out += '<div class="result-card">'+
                    //         '    <div class="progressbar" style="width: '+ percent +'%;"></div>'+
                    //         '    <div class="logo">'+ item.name + '</div>  '+
                    //         '    <div class="percentage">'+(percent * 2)+' %</div>  '+
                    //         '</div>';
                });
                document.getElementById("candidates").innerHTML = out;
                document.getElementById("main_content").style.display = "block";
                document.getElementById("error_msg").style.display = "none";

            }
        }
//     };
//     xhttp.open("GET", URL, true);
//     xhttp.send();
    
// }

// setInterval( () => {
//     updateResult()
// } , 5000)

updateResult()

// 0: roost
// 1: docker
// 2: minikube
// 3: k3d
// 4: kind 