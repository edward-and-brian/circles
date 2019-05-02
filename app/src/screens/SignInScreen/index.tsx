import React, { PureComponent } from 'react';
import { Query } from 'react-apollo';
import { View } from 'react-native';
import gql from 'graphql-tag';
import NavigationService from '../../navigation/NavigationService';
import SignInScreenView from './Views';

export interface Props {}
interface State {
  username: string;
  phone: string;
  performQuery: boolean;
}

const LOGIN = gql`
  {
    users {
      name
      id
      phoneNumber
      displayName
    }
  }
`;

interface dataType {
  users: {
    name: string;
    id: string;
    phoneNumber: string;
    displayName: string;
  }[];
}

interface paramTypes {}

class SignInScreen extends PureComponent<Props, State> {
  constructor(props: Props) {
    super(props);

    this.state = {
      username: '',
      phone: '',
      performQuery: false,
    };

    this.onChangeUsername = this.onChangeUsername.bind(this);
    this.onChangePhone = this.onChangePhone.bind(this);
    this.login = this.login.bind(this);
  }

  onPressBack() {
    NavigationService.back();
  }

  onChangeUsername(newUsername: string) {
    this.setState({ username: newUsername });
  }

  onChangePhone(newPhone: string) {
    this.setState({ phone: newPhone });
  }

  async login() {
    this.setState({ performQuery: true });
  }

  render() {
    const {
      username: displayName,
      phone: phoneNumber,
      performQuery,
    } = this.state;

    return (
      <Query<dataType, paramTypes>
        query={LOGIN}
        // variables={{ displayName, phoneNumber }}
        skip={!performQuery}
      >
        {({ loading, error, data }) => {
          if (loading) return <View />;
          if (error) return <View />;

          return (
            <SignInScreenView
              username={this.state.username}
              phone={this.state.phone}
              onChangeUsername={this.onChangeUsername}
              onChangePhone={this.onChangePhone}
              login={this.login}
            />
          );
        }}
        )
      </Query>
    );
  }
}

export default SignInScreen;
