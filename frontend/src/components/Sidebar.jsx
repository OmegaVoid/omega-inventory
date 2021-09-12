import {Button, Col, Image, Link} from 'arwes';


import {FontAwesomeIcon} from '@fortawesome/react-fontawesome'

import React from "react";


class Sidebar extends React.Component {


  render() {

    return <Col noMargin noGutter s={1} style={{width: "5%", margin: 0}}>
      <div style={{padding: "0.2rem", paddingLeft: 0, display: "flex", height: "100%", flexBasis: "max-content"}}>
        <div style={{display: "flex", alignItems: "center", flexDirection: "column", gap: "1rem"}}>

          <Image animate resources='TerranEmpireOfOmegaGlowSprite.webp'/>
          <Button><FontAwesomeIcon icon={"home"}/></Button>
          <Button><FontAwesomeIcon icon={"list"}/></Button>
          <Button><FontAwesomeIcon icon={"sign-out-alt"}/></Button>
          <Link style={{marginTop: "auto"}} href="https://omegavoid.codes">
            <Button><FontAwesomeIcon icon={"arrow-alt-circle-left"}/></Button>
          </Link>
        </div>
      </div>
    </Col>;
  }
}

export default Sidebar;
