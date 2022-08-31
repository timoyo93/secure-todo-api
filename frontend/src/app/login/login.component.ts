import { Component, OnInit } from '@angular/core';
import { AuthService } from '../services/auth.service';
import { catchError, tap } from 'rxjs';
import { HttpErrorResponse } from '@angular/common/http';
import { Router } from '@angular/router';

@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.scss'],
})
export class LoginComponent implements OnInit {
  formValid: boolean = true;
  errorMessage: string = '';
  loading: boolean = false;

  constructor(private authService: AuthService, private router: Router) {}

  ngOnInit() {}

  loginUserSubmit(data: LoginFormData): void {
    this.loading = true;
    this.formValid = true;
    this.authService.loginUser(data).subscribe({
      next: () => {
        this.loading = false;
        this.authService.isAuthenticated.next(true);
        this.router.navigate(['/']).then();
      },
      error: (err) => {
        this.loading = false;
        this.formValid = false;
        const error = err as HttpErrorResponse;
        switch (error.status) {
          case 400:
            this.errorMessage =
              'Username or passwore are wrong, please try again';
            break;
          case 404:
            this.errorMessage =
              'You are not registered yet. Please register first!';
            break;
          default:
            this.errorMessage = 'An unexpected error occured';
            break;
        }
      },
    });
  }
}

interface LoginFormData {
  username: string;
  password: string;
}
