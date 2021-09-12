import type { NextPage } from 'next'
import Sidebar from "../components/Sidebar";
// @ts-ignore
import {Arwes, Col, createTheme, Frame, Row, ThemeProvider, createResponsive, Button} from 'arwes';
import React from "react";

const Home: NextPage = () => {
  return (
      <ThemeProvider theme={createTheme()}>
          <Arwes>
            <Row  animate noMargin noGutter
                  style={{display: "flex", height: "100%", justifyItems: "flex-start", flexBasis: "max-content"}}>
              <Sidebar/>
              <Col animate s={11} style={{margin: 0}}>
                <Frame animate style={{height: "100%", width: "100%", margin: 0}}>
                </Frame>
              </Col>
            </Row>



          </Arwes>
      </ThemeProvider>
  )
}


export default Home
