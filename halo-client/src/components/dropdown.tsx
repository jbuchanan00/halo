
import type { ReactElement } from "react"
import type { Location } from "../types/location"

import './styles/dropdown.css'
import './styles/cityAutocomplete.css'



export default function Dropdown({locationList, itemClicked}: {locationList: Location[], itemClicked: (index: number) => void}): ReactElement {
    return (
        <div className="contained">
            {locationList ? locationList.map((value, ind) => 
                <li key={ind} className="dropdownItem" onClick={() => itemClicked(ind)}>
                    {value.name}{value.state ? `, ${value.state}` : null}
                </li>
            ) : null}
        </div>
    )
}