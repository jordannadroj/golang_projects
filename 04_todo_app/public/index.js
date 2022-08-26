function removeFromDb(item){
    fetch(`/api/todo?item=${item}`, {method: "Delete"}).then(res =>{
        if (res.status == 200){
            window.location.pathname = "/"
        }
    })
}

function updateDb(item) {
    let input = document.getElementById(item)
    let newitem = input.value
    const data={
        "olditem": item,
        "newitem": newitem
    };

    fetch(`/api/todo`, {method: "PUT",
        headers: {
            'Accept': 'application/json',
            'Content-Type': 'application/json'
        },
        body: JSON.stringify(data)})
        .then((response) => response.json())
        .then((data) => {
            console.log('Success:', data);
        })
        .then(res =>{
        if (res.status == 200){
            window.location.pathname = "/"
        }
    })
}
