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
import {
  createAppContainer,
  NavigationContainerComponent,
} from 'react-navigation';

import NavigationService from '../navigation/NavigationService';
import Routes from '../navigation/Routes';

const AppContainer = createAppContainer(Routes);

class App extends Component {
  render() {
    return (
      <AppContainer
        ref={navigatorRef => {
          navigatorRef = navigatorRef as NavigationContainerComponent;
          NavigationService.setTopLevelNavigator(navigatorRef);
        }}
      />
    );
  }
}

export default App;
