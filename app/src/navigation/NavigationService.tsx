import {
  NavigationActions,
  NavigationParams,
  NavigationContainerComponent,
} from 'react-navigation';

let navigator: NavigationContainerComponent;

const setTopLevelNavigator = (navigatorRef: NavigationContainerComponent) => {
  navigator = navigatorRef as NavigationContainerComponent;
};

const navigate = (routeName: string, params?: NavigationParams) => {
  navigator.dispatch(
    NavigationActions.navigate({
      routeName,
      params,
    }),
  );
};

const back = () => {
  navigator.dispatch(NavigationActions.back());
};

export default {
  navigate,
  back,
  setTopLevelNavigator,
};
