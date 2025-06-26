
import type { ReactElement } from "react"
import type { Location } from "../types/location"



export default function Dropdown({locationList}: {locationList: Location[]}): ReactElement {
    return (
        <div>
            {locationList ? locationList.map((value, ind) => 
                <li key={ind}>
                    {value.name}
                </li>
            ) : null}
        </div>
    )
}