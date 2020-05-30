import React from 'react'
import './styles.css'
import { Link } from 'react-router-dom';

class Auth extends React.Component {

    constructor(props:any){
        super(props);
        this.state = {
            isRegister: false
        }
    }

    changeView(){
        const { isRegister } : any = this.state;
        this.setState({
            isRegister: !isRegister
        })
    }

    registerForm(){

        return (
            <div className="form">
                <div className="row">
                    <div className="col-sm-6">
                        <div className="form-group">
                            <label>Email</label>
                            <input type="text" className="form-control" placeholder="example@mail.com"/>
                        </div>
                        <div className="form-group">
                            <label>Password</label>
                            <input type="password" className="form-control" placeholder="**************"/>
                        </div>
                    </div>
                    <div className="col-sm-6">
                        <div className="form-group">
                            <label>Nombre</label>
                            <input type="text" className="form-control" placeholder="yefferson"/>
                        </div>
                        <div className="form-group">
                            <label>Confirm password</label>
                            <input type="password" className="form-control" placeholder="**************"/>
                        </div>
                    </div>
                </div>
                <button className="btn btn-secondary">Register</button>
                <br></br>
                <label>You already have an account? <Link to="/login" onClick={() => this.changeView()}>Login</Link> </label>
            </div>
        );

    }

    loginForm(){

        return(
            <div className="form">
                <div className="form-group">
                    <label>User Name</label>
                    <input type="text" className="form-control" placeholder="example@mail.com"/>
                </div>
                <div className="form-group">
                    <label>Password</label>
                    <input type="password" className="form-control" placeholder="**************"/>
                </div>
                <button className="btn btn-black">Login</button>
                <button className="btn btn-secondary" onClick={() => this.changeView()}>Register</button>
                <br></br>
                <label>Forget your password? <Link to="/">Recover</Link> </label>
            </div>
        );

    }

    render(){

        const { isRegister } : any = this.state;

        return(

            <div className="convert-top">
                <div className="sidenav">
                    <div className="login-main-text">
                        <h2>Los Yefferson</h2>
                        <p>Login or Register.</p>
                    </div>
                </div>
                <div className="main">
                    <div className="col-sm-12">
                        <h1 className="font-MMJ">Aqui va <br></br> imagen</h1>
                        <div className="login-form">
                            {
                                !isRegister?
                                    this.loginForm()
                                :
                                    this.registerForm()
                            }
                        </div>
                    </div>
                </div>
            </div>

        );

    }

} 

export default Auth;