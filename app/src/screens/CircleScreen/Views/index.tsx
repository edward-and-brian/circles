import React, { PureComponent } from 'react';
import { View, Animated } from 'react-native';
import { Query } from 'react-apollo';
import gql from 'graphql-tag';
import CircleMessageWindow from '../../../components/circle/MessageWindow';
import CircleNavbar from '../../../components/shared/Navbars/CircleNavbar/';
import CircleFooter from '../../../components/circle/Footer';
import styles from './styles';

export interface Props {
  onPressBack(): void;
  footerHeight: Animated.Value;
  message: string;
  onMessageChange(newMessage: string): void;
}

const GET_EDDIE = gql`
  {
    Users {
      name
    }
  }
`;

const Dogs = () => (
  <Query query={GET_EDDIE}>
    {({ loading, error, data }) => {
      console.log({loading, error, data});

      return <View />;
    }}
  </Query>
);

class CircleScreenView extends PureComponent<Props> {
  constructor(props: Props) {
    super(props);

    this.PopulatedCircleNavbar = this.PopulatedCircleNavbar.bind(this);
  }

  PopulatedCircleNavbar() {
    return (
      <Query query={GET_EDDIE}>
        {({ loading, error, data }) => {
          console.log(loading, error, data);
          return (
            <CircleNavbar
              onPressBack={this.props.onPressBack}
              style={styles.headerContainer}
              circleName={data}
            />
          );
        }}
      </Query>
    );
  }

  render() {
    return (
      <View style={styles.container}>
        <Dogs />
        <CircleMessageWindow />
        <CircleFooter
          height={this.props.footerHeight}
          message={this.props.message}
          onMessageChange={this.props.onMessageChange}
        />
      </View>
    );
  }
}

export default CircleScreenView;
