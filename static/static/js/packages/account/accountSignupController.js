App.AccountSignupController = Ember.Controller.extend({

	username: null,
	password: null,
	email: null,

	formIsValid: function () {
		if (this.get('username') && this.get('password') && this.get('email')) {
			if (this.get('username').length < 4) {
				return false;
			}
			return true;
		} else {
			return false;
		}
	}.property('username', 'password', 'email'),



	invalidate: function (data) {
		
	},
	
	actions: {
		attemptRegistration: function () {
			if (this.get('formIsValid')) {
				$.post('/account/signup', {
					username: this.get("username"),
					password: this.get('password'),
					email: this.get("email")
				}, function (resp) {
					if (resp.code === 200) {
						this.transitionTo('login');
					} else {
						this.invalidate(resp.failure);
					}
				}.bind(this));
			}
		}
	}

});