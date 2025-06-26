import { useEffect, useState } from "react";


export default function CitySearch(){
    const [searchText, setSearchText] = useState("")

    useEffect(() => {

        console.log(searchText)
        if(searchText.length > 2){
            fetch('http://localhost:8081/autofill', {
                method: "POST",
                body: JSON.stringify({
                    location: searchText, 
                    baseLoc: {
                        "id": 0, 
                        "name": "", 
                        "state": "", 
                        "country": "",
                        "coords": {
                            "latitude": 0.00, 
                            "longitude": 0.00
                        } 
                    }
                })
            }).then((res) => {
                const reader = res.body?.getReader()
                const decoder = new TextDecoder()

                return reader?.read().then(function process( {done, value }): any{
                    if(done){
                        console.log("Stream complete")
                        return
                    }
                    const text = decoder.decode(value, {stream: true})
                    console.log("Chunk: ", text)
                    return reader.read().then(process)
                })
            })
        }
    }, [searchText])


    return (
        <div>
            <input onChange={e => setSearchText(e.target.value)}></input>
        </div>
    )
}

