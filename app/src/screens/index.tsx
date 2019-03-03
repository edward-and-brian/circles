/**
 * Sample React Native App
 * https://github.com/facebook/react-native
 *
 * Generated with the TypeScript template
 * https://github.com/emin93/react-native-template-typescript
 *
 * @format
 */

import React, { Component } from 'react';
import { createAppContainer } from 'react-navigation';
import Routes from '../Routes';

const AppContainer = createAppContainer(Routes);

interface Props {}

class App extends Component<Props> {
  render() {
    return <AppContainer />;
  }
}

export default App;
