package apiserver

// func TestServer_AuthenticateUser(t *testing.T) {
// 	store := teststore.New()
// 	u := entity.TestUser(t)
// 	store.User().Create(u)

// 	testCases := []struct {
// 		desc         string
// 		cookieValue  map[interface{}]interface{}
// 		expectedCode int
// 	}{
// 		{
// 			desc: "when authenticated",
// 			cookieValue: map[interface{}]interface{}{
// 				"user_id": u.ID,
// 			},
// 			expectedCode: http.StatusOK,
// 		},
// 		{
// 			desc:         "when not authenticated",
// 			cookieValue:  nil,
// 			expectedCode: http.StatusUnauthorized,
// 		},
// 	}

// 	secretKey := []byte("secret")
// 	s := newServer(store, sessions.NewCookieStore(secretKey))
// 	sc := securecookie.New(secretKey, nil)
// 	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		w.WriteHeader(http.StatusOK)
// 	})

// 	for _, tC := range testCases {
// 		t.Run(tC.desc, func(t *testing.T) {
// 			rec := httptest.NewRecorder()
// 			req, _ := http.NewRequest(http.MethodGet, "/", nil)
// 			cookieStr, _ := sc.Encode(sessionName, tC.cookieValue)
// 			req.Header.Set("Cookie", fmt.Sprintf("%s=%s", sessionName, cookieStr))
// 			s.authenticateUser(handler).ServeHTTP(rec, req)

// 			assert.Equal(t, tC.expectedCode, rec.Code)
// 		})
// 	}
// }

// func TestServer_HandleUsersCreate(t *testing.T) {
// 	s := newServer(teststore.New(), sessions.NewCookieStore([]byte("secret")))
// 	testCases := []struct {
// 		desc         string
// 		payload      interface{}
// 		expectedCode int
// 	}{
// 		{
// 			desc: "when params valid",
// 			payload: map[string]string{
// 				"email":    "user@example.com",
// 				"password": "password",
// 			},
// 			expectedCode: http.StatusCreated,
// 		},
// 		{
// 			desc:         "when invalid request payload",
// 			payload:      "invalid",
// 			expectedCode: http.StatusBadRequest,
// 		},
// 		{
// 			desc: "when invalid params",
// 			payload: map[string]string{
// 				"email":    "invalid",
// 				"password": "1",
// 			},
// 			expectedCode: http.StatusUnprocessableEntity,
// 		},
// 	}
// 	for _, tC := range testCases {
// 		t.Run(tC.desc, func(t *testing.T) {
// 			rec := httptest.NewRecorder()
// 			b := &bytes.Buffer{}
// 			json.NewEncoder(b).Encode(tC.payload)
// 			req, _ := http.NewRequest(http.MethodPost, "/users", b)
// 			s.ServeHTTP(rec, req)
// 			assert.Equal(t, tC.expectedCode, rec.Code)
// 		})
// 	}
// }

// func TestServer_HandleSessionsCreate(t *testing.T) {
// 	u := entity.TestUser(t)
// 	store := teststore.New()

// 	s := newServer(store, sessions.NewCookieStore([]byte("secret")))

// 	store.User().Create(u)

// 	testCases := []struct {
// 		desc         string
// 		payload      interface{}
// 		expectedCode int
// 	}{
// 		{
// 			desc: "when user exists and email and password match",
// 			payload: map[string]string{
// 				"email":    u.Email,
// 				"password": u.Password,
// 			},
// 			expectedCode: http.StatusOK,
// 		},
// 		{
// 			desc:         "when payload is invalid",
// 			payload:      "invalid",
// 			expectedCode: http.StatusBadRequest,
// 		},
// 		{
// 			desc: "when user with given email does not exist",
// 			payload: map[string]string{
// 				"email":    "unknown@example.com",
// 				"password": u.Password,
// 			},
// 			expectedCode: http.StatusUnauthorized,
// 		},
// 		{
// 			desc: "when user with given email exists and passwords doesn't match",
// 			payload: map[string]string{
// 				"email":    u.Email,
// 				"password": "invalid-password",
// 			},
// 			expectedCode: http.StatusUnauthorized,
// 		},
// 	}
// 	for _, tC := range testCases {
// 		t.Run(tC.desc, func(t *testing.T) {
// 			rec := httptest.NewRecorder()
// 			b := &bytes.Buffer{}
// 			json.NewEncoder(b).Encode(tC.payload)
// 			req, _ := http.NewRequest(http.MethodPost, "/sessions", b)
// 			s.ServeHTTP(rec, req)
// 			assert.Equal(t, tC.expectedCode, rec.Code)
// 		})
// 	}
// }
