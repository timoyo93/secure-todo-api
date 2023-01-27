import { Injectable } from "@angular/core";
import { HttpClient, HttpResponse } from "@angular/common/http";
import { BehaviorSubject, Observable } from "rxjs";
import { environment } from "../../environments/environment";

@Injectable({
  providedIn: "root",
})
export class AuthService {
  private backendUrl = environment.backendUrl + "/auth";

  isAuthenticated: BehaviorSubject<boolean> = new BehaviorSubject<boolean>(
    false
  );

  constructor(private http: HttpClient) {}

  loginUser(request: {
    username: string;
    password: string;
  }): Observable<HttpResponse<string>> {
    return this.http.post<string>(this.backendUrl + "/login", request, {
      withCredentials: true,
      observe: "response",
    });
  }

  registerUser(request: {
    username: string;
    password: string;
  }): Observable<HttpResponse<string>> {
    return this.http.post<string>(this.backendUrl + "/register", request, {
      withCredentials: true,
      observe: "response",
    });
  }

  logoutUser(): Observable<HttpResponse<string>> {
    return this.http.post<string>(
      this.backendUrl + "/logout",
      {},
      {
        withCredentials: true,
        observe: "response",
      }
    );
  }

  checkAuth(): Observable<HttpResponse<string>> {
    return this.http.get<string>(this.backendUrl, {
      withCredentials: true,
      observe: "response",
    });
  }
}
