import {AccessTokenResponseModel} from "./access-token.response.model";

export class LoginResponseModel {
  response_type: string;
  client_id: string;
  username: string;
  code: string;
  state: string;
  scope: string;
  confirm_scope: boolean;
  redirect_uri: string;
  access_token: AccessTokenResponseModel;
}

