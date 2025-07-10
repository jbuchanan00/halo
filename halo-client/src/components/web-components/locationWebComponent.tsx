import React from "react";
import ReactDOM from 'react-dom/client'
import reactToWebComponent from 'react-to-webcomponent'
import Input from '../cityAutocomplete'

const WebComponent = reactToWebComponent(Input, React, ReactDOM)

customElements.define('autofill-location', WebComponent)