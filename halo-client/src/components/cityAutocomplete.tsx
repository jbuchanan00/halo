import { useEffect, useState } from "react";
import Dropdown from "./dropdown";

import './styles/cityAutocomplete.css'

import type { Location } from "../types/location";


export default function CitySearch(){
    const [searchText, setSearchText] = useState("")
    const [locations, setLocations] = useState<Location[]>([])
    const [searchError, setSearchError] = useState(false)

    useEffect(() => {

        if(searchText.length > 2){
            console.log(searchText)
            const delayDebounce = setTimeout(() => {
                fetch('http://localhost:8081/autofill', {
                    method: "POST",
                    body: JSON.stringify({
                        location: searchText.trim(), 
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
                    setSearchError(false)
                    const reader = res.body?.getReader()
                    const decoder = new TextDecoder()
                    
                    return reader?.read().then(function process( {done, value }): any{
                        if(done){
                            return
                        }
                        const text = decoder.decode(value, {stream: true})
                        if(text != 'null'){
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

    function handleItemClicked(index: number){
        if(locations?.length > 0){
            setSearchText(`${locations[index].name}${locations[index].state ? `, ${locations[index].state}` : ''}`)
        }
    }

    function onSubmit(){
        if(locations === null){
            setSearchError(true)
        }else{
            window.location.href=`/polvo?lon=${locations[0].coord.lon}&lat=${locations[0].coord.lat}`
        }
    }


    return (
        <div className="contained">
            <div className="input-with-button">
                <input id={`${searchError ? "errorState" : ""}`} onChange={e => setSearchText(e.target.value)} value={searchText}></input>
                <button type="submit" onClick={onSubmit}>GO</button>
            </div>
            <div className="dropdownContainer" hidden={locations?.length < 1}>
                <Dropdown locationList={locations} itemClicked={handleItemClicked}/>
            </div>
        </div>
    )
}

