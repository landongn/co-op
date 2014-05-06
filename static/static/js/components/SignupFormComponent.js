App.SignupFormComponent = Ember.Component.extend({
	password: '',
	username: '',
	email: '',
	passwordConfirm: '',

	signupAttempts: 0,

	focusIn: function (e) {
		$(e.target).siblings('span.tooltip').addClass('show-tooltip');
	},
	focusOut: function (e) {
		$(e.target).siblings('span.tooltip').removeClass('show-tooltip');
	},

	isValid: function () {
		if (this.get('username').length > 3 &&
			this.get('password').length > 4 &&
			this.get('email').length > 3 || this.get('email').match(/.+@.+\..+/i) !== null) {
			return true;
		} else {
			if (this.get('username').length === 0 || this.get('username').length < 3) {
				this.set('invalidUsername', true);
				this.set('failureUsername', 'Usernames should be at least 3 characters long. Try another one.');
			}

			if (this.get('password').length === 0 || this.get('password').length < 3) {
				this.set('invalidPassword', true);
				this.set('failurePassword', 'Passwords need to have at least 3 characters. Get creative!');
			}

			if (this.get('password') !== this.get('passwordConfirm')) {
				this.setProperties({
					invalidPassword: true,
					failurePassword: "Passwords must match"
				});
			}

			var em = this.get('email');
			if (em.match(/.+@.+\..+/i) === null) {
				this.set('invalidEmail', true);
				this.set('failureEmail', 'That isn\'t an email address.  C\'mon.');
			}

			return false;
		}
	}.property('signupAttempts', 'password', 'email', 'username', 'invalidUsername', 'invalidPassword', 'invalidEmail', 'passwordConfirm'),

	invalidate: function (failure) {
		if (failure.code === "name") {
			this.setProperties({
				invalidUsername: true,
				failureUsername: failure.msg
			});
		}

		if (failure.code === "email") {
			this.setProperties({
				invalidEmail: true,
				failureEmail: failure.msg
			});
		}

		if (failure.code === "error") {
			this.setProperties({
				invalidEmail: true,
				invalidUsername: true,
				failureUsername: failure.msg.username,
				failureEmail: failure.msg.email
			});
		}

		if (failure.code === "password") {
			this.setProperties({
				invalidPassword: true,
				failurePassword: "Passwords must match"
			});
		}
	},

	click: function (e) {
		if (e.target.nodeName === 'INPUT' && e.target.className !== "submit-signup-form") {
			this.setProperties({
				invalidUsername: false,
				invalidPassword: false,
				invalidEmail: false
			});
		}
	},

	keyPress: function () {
		this.setProperties({
			invalidUsername: false,
			invalidPassword: false,
			invalidEmail: false
		});
	},
	actions: {
		attemptRegistration: function () {
			this.incrementProperty('signupAttempts');
			if (this.get('isValid')) {
				$.post('/account/signup', {
					username: this.get("username"),
					password: this.get('password'),
					passwordConfirm: this.get('passwordConfirm'),
					email: this.get("email")
				}, function (resp) {
					if (resp.code === 200) {
						this.transitionTo('login');
					} else {
						this.invalidate(resp);
					}
				}.bind(this));
			}
		}
	}
});