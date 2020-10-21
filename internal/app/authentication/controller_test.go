package authentication_test

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
