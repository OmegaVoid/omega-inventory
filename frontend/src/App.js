import logo from './logo.svg';

import {ThemeProvider, createTheme, Arwes, Button, Image, Frame, Project, Words, Content, Heading} from 'arwes';

import './App.css';

import Tree, {FilteringContainer, renderers} from 'react-virtualized-tree'


import {FontAwesomeIcon} from '@fortawesome/react-fontawesome'
import {library} from '@fortawesome/fontawesome-svg-core'
import {
  faFolder,
  faFolderOpen,
  faFile,
  faFolderPlus,
  faFolderMinus,
  faPlusSquare,
  faMinusSquare,
  faEdit,
  faSync
} from '@fortawesome/free-solid-svg-icons'
import React, {Component} from "react";
import tree from './tree';

import cx from 'classnames';

const {Expandable} = renderers;
library.add(faFolder, faFolderOpen, faFile, faFolderPlus, faFolderMinus, faPlusSquare, faMinusSquare, faEdit, faSync)


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
  return new Array(minNumOfNodes).fill(deepness).map((si, i) => {
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

            <Tree nodes={this.state.nodes} onChange={this.handleChange}>
              {({style, node, ...rest}) => (
                  <div style={style}>
                    <Expandable node={node}  {...rest}>
                      <Button>{node.name}</Button>
                    </Expandable>
                  </div>
              )}
            </Tree>
          </Arwes>
        </ThemeProvider>
    )
  }
}


export default App;
