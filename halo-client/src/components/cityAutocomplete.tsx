import { useEffect, useState } from "react";
import Dropdown from "./dropdown";

import type { Location } from "../types/location";


export default function CitySearch(){
    const [searchText, setSearchText] = useState("")
    const [locations, setLocations] = useState<Location[]>([])

    useEffect(() => {

        console.log(locations)
        if(searchText.length > 2){
            const delayDebounce = setTimeout(() => {
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
                    setLocations([])
                    const reader = res.body?.getReader()
                    const decoder = new TextDecoder()
                    
                    return reader?.read().then(function process( {done, value }): any{
                        if(done){
                            console.log("Stream complete")
                            return
                        }
                        const text = decoder.decode(value, {stream: true})
                        if(text != 'null'){
                            console.log("GOING HERE", text)
                            setLocations(JSON.parse(text))
                        }
                        return reader.read().then(process)
                    })
                })
            }, 300)

            return () => clearTimeout(delayDebounce)
        }else{
            setLocations([])
        }
    }, [searchText])


    return (
        <div>
            <input onChange={e => setSearchText(e.target.value)}></input>
            <Dropdown locationList={locations}/>
        </div>
    )
}

