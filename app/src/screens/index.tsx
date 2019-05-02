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
import { ApolloProvider } from 'react-apollo';
import ApolloClient from 'apollo-boost';

const AppContainer = createAppContainer(Routes);

const client = new ApolloClient({
  uri: 'http://localhost:3000/query',
});

class App extends Component {
  render() {
    return (
      <ApolloProvider client={client}>
        <AppContainer
          ref={navigatorRef => {
            navigatorRef = navigatorRef as NavigationContainerComponent;
            NavigationService.setTopLevelNavigator(navigatorRef);
          }}
        />
      </ApolloProvider>
    );
  }
}

export default App;
