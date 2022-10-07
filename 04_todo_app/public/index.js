function removeFromDb(item_id){
    fetch(`/api/todo?item=${item_id}`, {method: "Delete"}).then(res =>{
        if (res.status == 200){
            window.location.pathname = "/"
        }
    })
}

function updateDb(item_id) {
    let input = document.getElementById(item_id)
    let newitem = input.value
    const data={
        "olditem": item_id,
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
