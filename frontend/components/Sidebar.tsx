// @ts-ignore
import {Button, Col, Image, Link, Frame} from 'arwes';


import {FontAwesomeIcon} from '@fortawesome/react-fontawesome'

import React from "react";


class Sidebar extends React.Component {
  render() {
    return <Col animate noMargin noGutter s={1} style={{width: "5%", margin: 0}}>
      <Frame style={{padding: "0.2rem", paddingLeft: 0, display: "flex", height: "100%", flexBasis: "max-content"}}>
        <div style={{display: "flex", alignItems: "center", flexDirection: "column", gap: "1rem"}}>
          <Link animate href="http://localhost:3001/">
            <Image animate resources='images/TerranEmpireOfOmegaGlowSprite.webp'/>
          </Link>
          <Button animate><FontAwesomeIcon icon={"home"}/></Button>
          <Button animate><FontAwesomeIcon icon={"list"}/></Button>
          <Button animate><FontAwesomeIcon icon={"sign-out-alt"}/></Button>

        </div>
      </Frame>
    </Col>;
  }
}

export default Sidebar;
