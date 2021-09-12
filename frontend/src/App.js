import Sidebar from "./components/Sidebar";

import {Arwes, Col, createTheme, Frame, Row, ThemeProvider} from 'arwes';

import './App.css';

import {library} from '@fortawesome/fontawesome-svg-core'
import {
  faArrowAltCircleLeft,
  faEdit,
  faFile,
  faFolder,
  faFolderMinus,
  faFolderOpen,
  faFolderPlus,
  faHome,
  faList,
  faMinusSquare,
  faPlusSquare,
  faSignOutAlt,
  faSync
} from '@fortawesome/free-solid-svg-icons'
import React, {Component} from "react";

library.add(faFolder, faFolderOpen, faFile, faFolderPlus, faFolderMinus,
    faPlusSquare, faMinusSquare, faEdit, faSync, faHome, faList, faSignOutAlt, faArrowAltCircleLeft)


let ids = {};
const getUniqueId = () => {
  const candidateId = Math.round(Math.random() * 1000000000);

  if (ids[candidateId]) {
    return getUniqueId();
  }

  ids[candidateId] = true;

  return candidateId;
};
const constructTree = (maxDeepness, maxNumberOfChildren, minNumOfNodes, deepness = 1) => {
  return new Array(minNumOfNodes).fill(deepness).map(() => {
    const id = getUniqueId();
    const numberOfChildren = deepness === maxDeepness ? 0 : Math.round(Math.random() * maxNumberOfChildren);

    return {
      id,
      name: `Leaf ${id}`,
      children: numberOfChildren ? constructTree(maxDeepness, maxNumberOfChildren, numberOfChildren, deepness + 1) : [],
      state: {
        expanded: numberOfChildren ? Boolean(Math.round(Math.random())) : false,
        favorite: Boolean(Math.round(Math.random())),
        deletable: Boolean(Math.round(Math.random())),
      },
    };
  });
};


const MAX_DEEPNESS = 3;
const MAX_NUMBER_OF_CHILDREN = 4;
const MIN_NUMBER_OF_PARENTS = 5;

const Nodes = constructTree(MAX_DEEPNESS, MAX_NUMBER_OF_CHILDREN, MIN_NUMBER_OF_PARENTS);

class App extends Component {
  state = {
    nodes: Nodes,
  };

  handleChange = nodes => {
    this.setState({nodes});
  };

  render() {
    return (
        <ThemeProvider theme={createTheme()}>
          <Arwes>

            <Row noMargin noGutter
                 style={{display: "flex", height: "100%", justifyItems: "flex-start", flexBasis: "max-content"}}>
              <Sidebar/>
              <Col s={11} style={{margin: 0}}>
                <Frame style={{height: "100%", width: "100%", margin: 0}}>

                </Frame>
              </Col>
            </Row>


            {/*<Tree nodes={this.state.nodes} onChange={this.handleChange}>*/}
            {/*  {({style, node, ...rest}) => (*/}
            {/*      <div style={style}>*/}
            {/*        <Expandable node={node}  {...rest}>*/}
            {/*          <Button>{node.name}</Button>*/}
            {/*        </Expandable>*/}
            {/*      </div>*/}
            {/*  )}*/}
            {/*</Tree>*/}
          </Arwes>
        </ThemeProvider>
    )
  }
}


export default App;
